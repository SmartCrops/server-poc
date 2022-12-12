package models

import "gorm.io/gorm"

type Installation struct {
	ID              uint
	Name            string
	Lat             float64
	Lon             float64
	OptimalHumidity float64
	UserID          uint
	PumpControllers []PumpController
}

// Create or update Installation
func (installation Installation) Save(db *gorm.DB) error {
	return db.Save(&installation).Error
}

// Get Installation with its PumpControllers by InstallationID
func (installation *Installation) GetByID(db *gorm.DB, installationId uint) error {
	return db.Model(&Installation{}).Preload("PumpControllers").Find(installation).Error
}
