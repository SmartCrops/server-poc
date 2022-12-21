package models

import (
	"time"

	"gorm.io/gorm"
)

type SensorData struct {
	ID                        uint
	SoilHumidity              float64
	Temperature               float64
	LightIntesity             float64
	CreatedAt                 time.Time     `gorm:"index:,composite:index,priority:2"`
	DataCollectorSerialNumber string        `gorm:"index:,composite:index,priority:1"`
	DataCollector             DataCollector `gorm:"foreignKey:DataCollectorSerialNumber;references:SerialNumber"`
}

// Create SensorData
func (sensorData SensorData) Save(db *gorm.DB) error {
	sensorData.CreatedAt = time.Now()
	return db.Create(&sensorData).Error
}

// Get Latest SensorData by DataCollector Serial Number
func (sensorData *SensorData) GetLatest(db *gorm.DB, dataCollectorSerialNumber string) error {
	return db.Order("created_at desc").Where("data_collector_serial_number == ?", dataCollectorSerialNumber).First(sensorData).Error
}
