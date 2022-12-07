package waterplanner

import (
	"log"
	"server-poc/pkg/models"
	"server-poc/pkg/mqtt"
	"server-poc/pkg/services/datacollector"
	"server-poc/pkg/services/pumpcontroller"

	"gorm.io/gorm"
)

const (
	hours         = 6
	humDiffFactor = 0.1
	rainFactor    = 0.1
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

	var err error

	var sensor *models.Sensor
	err = sensor.GetByID(s.db, data.ID)
	if err != nil {
		log.Println("Error getting sensor from database:", err)
		return
	}

	var pump *models.Pump
	err = pump.GetByID(s.db, sensor.PumpID)
	if err != nil {
		log.Println("Error getting pump from database:", err)
		return
	}

	var tank *models.Tank
	err = tank.GetByID(s.db, pump.TankID)
	if err != nil {
		log.Println("Error getting tank from database:", err)
		return
	}

	var installation *models.Installation
	err = installation.GetByID(s.db, tank.InstallationID)
	if err != nil {
		log.Println("Error getting installation from database:", err)
		return
	}

	// rainVolume, err := GetAccumulatedRainVolume(hours, installation.Lat, installation.Lon)
	// if err != nil {
	// 	log.Println("Error calling API:", err)
	// 	return
	// }

	// - Get location of sensor
	// - Call weather api for that location
	// - Make decision
	// - Send command over MQTT
}
