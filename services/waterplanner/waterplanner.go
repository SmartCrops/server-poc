package waterplanner

import (
	"fmt"
	"server-poc/pkg/models"
	"server-poc/pkg/mqtt"
	"server-poc/services/datacollector"
	"server-poc/services/pumpcontroller"

	"gorm.io/gorm"
)

/* ------------------------------- public API ------------------------------- */
func Start(
	db *gorm.DB,
	mqttClient mqtt.Client,
	datacollectorService datacollector.Service,
	pumpcontrollerService pumpcontroller.Service,
) {
	s := service{db, mqttClient, datacollectorService, pumpcontrollerService}
	s.datacollectorService.ListenForNewData(s.handleData)
}

/* ------------------------- implementation details ------------------------- */
type service struct {
	db                    *gorm.DB
	mqttClient            mqtt.Client
	datacollectorService  datacollector.Service
	pumpControllerService pumpcontroller.Service
}

func (s *service) handleData(data models.SensorData) {
	// Get required waterplanning data from db
	q, err := getQueryData(data.DataCollectorSerialNumber, s.db)
	if err != nil {
		return
	}
	// Decide if and for how long to water the crops
	shouldWater := q.makeDecision()
	if !shouldWater {
		return
	}

	// Create and send pump controller command
	durationS := q.determineWateringDuration()
	command := pumpcontroller.PumpControllerCommand{
		PumpGpio:  q.PumpGpio,
		DurationS: durationS,
	}
	s.pumpControllerService.Send(command, q.PumpControllerSerialNumber)
}

/* ----------------------------- making decision ---------------------------- */
func (q waterPlanningQuery) makeDecision() bool {
	// Check if crops need more water
	if q.SoilHumidityAvg > q.OptimalHumidity {
		return false
	}
	// Check if it will rain
	w, err := getWeather(q.Lat, q.Lon)
	if err != nil {
		return false
	}
	shouldWater := !w.itWillRainIn24h()
	if shouldWater {
		return true
	} else {
		return false
	}
}

func (q waterPlanningQuery) determineWateringDuration() uint16 {
	optimalAverageHumDiff := q.OptimalHumidity - q.SoilHumidityAvg
	return 5 * uint16(optimalAverageHumDiff)
}

/* --------------------- fetching required data from db --------------------- */
type waterPlanningQuery struct {
	Lat                        float64
	Lon                        float64
	OptimalHumidity            float64
	PumpControllerSerialNumber string
	PumpGpio                   uint8
	SoilHumidityAvg            float64
}

func getQueryData(dataCollectorSerialNumber string, db *gorm.DB) (waterPlanningQuery, error) {
	result := waterPlanningQuery{}
	err := db.Raw(getSqlQueryString(dataCollectorSerialNumber)).Scan(&result).Error
	if err != nil {
		return result, err
	}
	return result, nil
}

func getSqlQueryString(dataCollectorSerialNumber string) string {
	return fmt.Sprintf(`WITH water_planning_query AS (
		SELECT
			installations.lat AS lat,
			installations.lon AS lon,
			installations.optimal_humidity AS optimal_humidity,
			pump_controllers.serial_number AS pump_controller_serial_number,
			pumps.gpio AS pump_gpio,
			AVG(sensor_data.soil_humidity) AS soil_humidity_avg,
			row_number() OVER(PARTITION BY sensor_data.data_collector_serial_number ORDER BY sensor_data.created_at DESC) AS sensor_data_latest_record_time_index
		FROM sensor_data
		JOIN data_collectors 
		ON sensor_data.data_collector_serial_number=data_collectors.serial_number
		JOIN pumps
		ON data_collectors.pump_id=pumps.id
		JOIN pump_controllers
		ON pumps.pump_controller_id=pump_controllers.id
		JOIN installations
		ON pump_controllers.installation_id=installations.id
		WHERE sensor_data.data_collector_serial_number=\"%s\" -- sensor_data.data_collector_serial_number is an input for this command
	)
	SELECT
		lat,
		lon,
		optimal_humidity,
		pump_controller_serial_number,
		pump_gpio,
		soil_humidity_avg
	FROM water_planning_query
	WHERE sensor_data_latest_record_time_index = 1;`, dataCollectorSerialNumber)
}
