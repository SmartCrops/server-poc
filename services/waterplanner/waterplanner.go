package waterplanner

import (
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

type WaterPlanningData struct {
	Lat                        float64
	Lon                        float64
	OptimalHumidity            float64
	PumpControllerSerialNumber string
	PumpGpio                   uint8
	SoilHumidityAvg            float64
}

func GetWaterPlanningData(db *gorm.DB, dataCollectorSerialNumber string) (data WaterPlanningData, err error) {
	data = WaterPlanningData{}
	err = nil

	// Get DataCollector
	collector := models.DataCollector{}
	err = collector.GetBySerialNumber(db, dataCollectorSerialNumber)
	if err != nil {
		return
	}

	// Get data from the parent Pump
	pump := models.Pump{}
	err = pump.GetBySerialNumber(db, collector.PumpSerialNumber)
	if err != nil {
		return
	}
	data.PumpGpio = pump.Gpio

	// Get data from the parent PumpController
	controller := models.PumpController{}
	err = controller.GetBySerialNumber(db, pump.PumpControllerSerialNumber)
	if err != nil {
		return
	}
	data.PumpControllerSerialNumber = controller.SerialNumber

	// Get data from the parent Installation
	installation := controller.Installation
	data.Lat = installation.Lat
	data.Lon = installation.Lon
	data.OptimalHumidity = installation.OptimalHumidity

	// Get average soil humidity of latest SensorData of every DataCollector
	// assigned to the same Pump as a DataCollector identified by the given dataCollectorSerialNumber
	soilHumidityAvg := 0.0
	count := 0.0
	for _, dataCollector := range pump.DataCollectors {
		sensorData := models.SensorData{}
		err = sensorData.GetLatest(db, dataCollector.SerialNumber)
		if err != nil {
			return
		}
		soilHumidityAvg += sensorData.SoilHumidity
		count++
	}
	soilHumidityAvg /= count
	data.SoilHumidityAvg = soilHumidityAvg

	return
}

/* ------------------------- service implementation ------------------------- */
type service struct {
	db                    *gorm.DB
	mqttClient            mqtt.Client
	datacollectorService  datacollector.Service
	pumpControllerService pumpcontroller.Service
}

func (s *service) handleData(sensorData models.SensorData) {
	// Get required waterplanning data from db
	data, err := GetWaterPlanningData(s.db, sensorData.DataCollectorSerialNumber)
	if err != nil {
		return
	}
	// Decide if and for how long to water the crops
	shouldWater := data.makeDecision()
	if !shouldWater {
		return
	}
	// Create and send pump controller command
	durationS := data.determineWateringDuration()
	command := pumpcontroller.PumpControllerCommand{
		PumpGpio:  data.PumpGpio,
		DurationS: durationS,
	}
	s.pumpControllerService.Send(command, data.PumpControllerSerialNumber)
}

/* ----------------------------- making decision ---------------------------- */
func (q WaterPlanningData) makeDecision() bool {
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

func (data WaterPlanningData) determineWateringDuration() uint16 {
	optimalAverageHumDiff := data.OptimalHumidity - data.SoilHumidityAvg
	return 5 * uint16(optimalAverageHumDiff)
}
