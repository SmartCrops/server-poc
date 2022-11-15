package sensordata

import (
	"fmt"
	db "server-poc/pkg/db"

	"gorm.io/gorm"
)

type SensorData struct {
	gorm.Model
	Temperature   float64
	Pressure      float64
	LightIntesity float64
	SensorID      uint32
}

func GetAll() ([]SensorData, error) {
	data := make([]SensorData, 0)
	err := db.DB.Find(&data).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all sensor data: %w", err)
	}
	return data, nil
}
