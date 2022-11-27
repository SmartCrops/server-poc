package waterplanner

import (
	"server-poc/pkg/models"
	"server-poc/pkg/mqtt"
	"server-poc/pkg/services/datacollector"
	"server-poc/pkg/services/pumpcontroller"

	"gorm.io/gorm"
)

func Start(
	db *gorm.DB,
	mqttClient mqtt.Client,
	datacollectorService datacollector.Service,
	pumpcontrollerService pumpcontroller.Service,
) {
	s := service{db, mqttClient, datacollectorService, pumpcontrollerService}
	s.datacollectorService.ListenForNewData(s.handleData)
}

type service struct {
	db                    *gorm.DB
	mqttClient            mqtt.Client
	datacollectorService  datacollector.Service
	pumpcontrollerService pumpcontroller.Service
}

func (s *service) handleData(data models.SensorData) {
	// - Get sensor location from the database
	// - Call weather api for that location
	// - Make decision
	// - Send command over MQTT
}
