package models

import "gorm.io/gorm"

type DataCollector struct {
	SerialNumber     string `gorm:"primaryKey"`
	PumpSerialNumber string
	Pump             Pump `gorm:"foreignKey:PumpSerialNumber;references:SerialNumber"`
}

// Create or update DataCollector
func (dataCollector DataCollector) Save(db *gorm.DB) error {
	return db.Save(&dataCollector).Error
}

// Get DataCollector with its Pump
func (dataCollector *DataCollector) GetBySerialNumber(db *gorm.DB, serialNumber string) error {
	return db.Model(&dataCollector).Preload("Pump").Where("serial_number == ?", serialNumber).First(dataCollector).Error
}
