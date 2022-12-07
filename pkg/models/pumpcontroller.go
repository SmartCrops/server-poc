package models

import "gorm.io/gorm"

type PumpController struct {
	gorm.Model
	SerialNumber   string `gorm:"unique"`
	InstallationID uint
	Pumps          []Pump
}

func (pumpController *PumpController) GetByID(db *gorm.DB, id uint) error {
	return db.First(pumpController, id).Error
}

func (pumpController PumpController) Save(db *gorm.DB) error {
	return db.Save(pumpController).Error
}
