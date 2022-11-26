package models

import "gorm.io/gorm"

type Tank struct {
	gorm.Model
	InstallationID uint
	Pumps          []Pump
}

func (tank Tank) Save(db *gorm.DB) error {
	return db.Save(tank).Error
}
