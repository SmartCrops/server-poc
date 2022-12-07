package models

import (
	"gorm.io/gorm"
)

type SensorData struct {
	gorm.Model
	SoilHumidity              float64
	Temperature               float64
	LightIntesity             float64
	DataCollectorSerialNumber string
	DataCollector             DataCollector `gorm:"foreignKey:DataCollectorSerialNumber;references:SerialNumber"`
}

func GetByDataCollectorSerialNumber(db *gorm.DB, DataCollectorSerialNumber string) ([]SensorData, error) {
	var data []SensorData
	err := db.Where(&SensorData{DataCollectorSerialNumber: DataCollectorSerialNumber}).Find(&data).Error
	return data, err
}

func (data SensorData) Save(db *gorm.DB) error {
	return db.Save(&data).Error
}
