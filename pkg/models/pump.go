package models

import "gorm.io/gorm"

type Pump struct {
	gorm.Model
	Gpio    uint8
	TankID  uint
	Sensors []Sensor
}

func (pump *Pump) GetByID(db *gorm.DB, id uint) error {
	return db.First(pump, id).Error
}

func (pump Pump) Save(db *gorm.DB) error {
	return db.Save(pump).Error
}
