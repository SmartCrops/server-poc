package waterplanner

import (
	"server-poc/pkg/models"
	"server-poc/pkg/mqtt"
	"server-poc/pkg/services/datacollector"

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
	// - Make decision
	// - Send command over MQTT
}
