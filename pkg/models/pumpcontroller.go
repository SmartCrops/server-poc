package models

import "gorm.io/gorm"

type PumpController struct {
	SerialNumber   string `gorm:"primaryKey"`
	InstallationID uint
	Pumps          []Pump
}

// Create or update PumpController
func (pumpController PumpController) Save(db *gorm.DB) error {
	return db.Save(&pumpController).Error
}

// Get PumpController with its pumps
func (pumpController *PumpController) GetBySerialNumber(db *gorm.DB, serialNumber string) error {
	return db.Model(&PumpController{}).Preload("Pumps").First(pumpController, serialNumber).Error
}
