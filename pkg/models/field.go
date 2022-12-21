package models

import "gorm.io/gorm"

type Field struct {
	ID              uint
	Name            string
	Lat             float64
	Lon             float64
	OptimalHumidity float64
	InstallationID  uint
	PumpControllers []PumpController
}

// Create or update Installation
func (field Field) Save(db *gorm.DB) error {
	return db.Save(&field).Error
}

// Get Installation with its PumpControllers by FieldID
func (field *Field) GetByID(db *gorm.DB, fieldId uint) error {
	return db.Model(&Field{}).Preload("PumpControllers").Find(field).Error
}
