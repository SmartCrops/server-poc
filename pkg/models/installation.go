package models

import "gorm.io/gorm"

type Installation struct {
	gorm.Model
	Name   string
	Lat    float64
	Lon    float64
	UserID uint
	Tanks  []Tank
}

func (installation Installation) Save(db *gorm.DB) error {
	return db.Save(installation).Error
}
