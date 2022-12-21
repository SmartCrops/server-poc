package models

import "gorm.io/gorm"

type PumpController struct {
	SerialNumber string `gorm:"primaryKey"`
	FieldID      uint
	Field        Field
	Pumps        []Pump
}

// Create or update PumpController
func (pumpController PumpController) Save(db *gorm.DB) error {
	return db.Save(&pumpController).Error
}

// Get PumpController with its Field and Pumps
func (pumpController *PumpController) GetBySerialNumber(db *gorm.DB, serialNumber string) error {
	return db.Model(&PumpController{}).Preload("Field").Preload("Pumps").First(pumpController, "serial_number == ?", serialNumber).Error
}
