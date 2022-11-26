package models

import "gorm.io/gorm"

type Installation struct {
	gorm.Model
	Name   string `gorm:"unique"`
	Lat    float64
	Long   float64
	UserID uint
	Tanks  []Tank
}

func (installation Installation) Save(db *gorm.DB) error {
	return db.Save(installation).Error
}
