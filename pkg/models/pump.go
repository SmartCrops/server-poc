package models

import "gorm.io/gorm"

type Pump struct {
	gorm.Model
	SerialNumber     string `gorm:"unique"`
	PumpControllerID uint
	Gpio             uint8
	DataCollectors   []DataCollector
}

func (pump *Pump) GetByID(db *gorm.DB, id uint) error {
	return db.First(pump, id).Error
}

func (pump Pump) Save(db *gorm.DB) error {
	return db.Save(pump).Error
}
