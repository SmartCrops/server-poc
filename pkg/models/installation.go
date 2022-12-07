package models

import "gorm.io/gorm"

type Installation struct {
	gorm.Model
	Name            string
	Lat             float64
	Lon             float64
	UserID          uint
	PumpControllers []PumpController
}

func (installation *Installation) GetByID(db *gorm.DB, id uint) error {
	return db.First(installation, id).Error
}

func (installation Installation) Save(db *gorm.DB) error {
	return db.Save(installation).Error
}
