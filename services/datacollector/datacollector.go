package datacollector

import (
	"encoding/json"
	"log"

	"server-poc/pkg/models"
	"server-poc/pkg/mqtt"

	"gorm.io/gorm"
)

type NotificationHandler func(models.SensorData)
type Service interface {
	ListenForNewData(NotificationHandler)
}

// Internal structure for keeping all data and dependencies.
type service struct {
	db         *gorm.DB
	mqttClient mqtt.Client
	observers  []NotificationHandler
}

func Start(mqttClient mqtt.Client, db *gorm.DB) (Service, error) {
	s := service{
		db:         db,
		mqttClient: mqttClient,
	}
	if err := mqttClient.Sub("sensors/#", 1, s.handleData); err != nil {
		return nil, err
	}
	return &s, nil
}

func (s *service) ListenForNewData(handler NotificationHandler) {
	// Add new observer to the list
	s.observers = append(s.observers, handler)
}

func (s *service) handleData(msg []byte) {
	type MessagePayload struct {
		SoilHumidity              float64 `json:"soilHumidity"`
		Temperature               float64 `json:"temperature"`
		LightIntensity            float64 `json:"lightIntensity"`
		DataCollectorSerialNumber string  `json:"serialNumber"`
	}

	// Decode the payload
	var payload MessagePayload
	if err := json.Unmarshal(msg, &payload); err != nil {
		log.Println("Received malformed sensordata:", err)
		return
	}

	// Add data to the database

	data := models.SensorData{
		SoilHumidity:              payload.SoilHumidity,
		Temperature:               payload.Temperature,
		LightIntesity:             payload.LightIntensity,
		DataCollectorSerialNumber: payload.DataCollectorSerialNumber,
	}

	if err := data.Save(s.db); err != nil {
		log.Println("Failed to save sensor data to the database:", err)
		return
	}

	// Notify all observers
	for _, observer := range s.observers {
		go observer(data)
	}
}
