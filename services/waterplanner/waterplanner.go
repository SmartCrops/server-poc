package waterplanner

import (
	"server-poc/pkg/models"
	"server-poc/pkg/mqtt"
	"server-poc/services/datacollector"

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

func (s *service) handleData(data models.SensorData) {
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
