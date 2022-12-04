package waterplanner

import (
	"server-poc/pkg/datacollector"
	"server-poc/pkg/mqtt"
	"server-poc/pkg/sensordata"

	"gorm.io/gorm"
)

func Start(db *gorm.DB, mqttClient mqtt.Client, datacollectorService datacollector.Service) {
	s := service{db, mqttClient, datacollectorService}
	s.datacollectorService.ListenForNewData(s.handleData)
}

type service struct {
	db                   *gorm.DB
	mqttClient           mqtt.Client
	datacollectorService datacollector.Service
}

func (s *service) handleData(data sensordata.SensorData) {
	// - Get sensor location from the database
	// - Call weather api for that location
	w, err := getWeather(lat, lng)
	if err != nil {
		return
	}
	// - Make decision
	shouldWater := !w.willItRainIn24h()
	// - Send command over MQTT
	if shouldWater {
		err = s.mqttClient.Pub("controllers/xxx", 1, false, "water")
		if err != nil {
			return
		}
	}
}
