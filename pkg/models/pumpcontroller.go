package models

import "gorm.io/gorm"

type PumpController struct {
	SerialNumber   string `gorm:"primaryKey"`
	InstallationID uint
	Installation   Installation
	Pumps          []Pump
}

// Create or update PumpController
func (pumpController PumpController) Save(db *gorm.DB) error {
	return db.Save(&pumpController).Error
}

// Get PumpController with its pumps
func (pumpController *PumpController) GetBySerialNumber(db *gorm.DB, serialNumber string) error {
	return db.Model(&PumpController{}).Preload("Installation").Preload("Pumps").First(pumpController, "serial_number == ?", serialNumber).Error
}
