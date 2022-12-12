package models

import "gorm.io/gorm"

type Pump struct {
	SerialNumber               string `gorm:"primaryKey"`
	PumpControllerSerialNumber string
	PumpController             PumpController `gorm:"foreignKey:PumpControllerSerialNumber;references:SerialNumber"`
	Gpio                       uint8
	DataCollectors             []DataCollector
}

// Create or update Pump
func (pump Pump) Save(db *gorm.DB) error {
	return db.Save(&pump).Error
}

// Get Pump with its DataCollectors and its PumpController
func (pump *Pump) GetBySerialNumber(db *gorm.DB, serialNumber string) error {
	return db.Model(&Pump{}).Preload("PumpController", "DataCollectors").Where("serial_number == ?", serialNumber).First(pump).Error
}
