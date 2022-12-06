package models

import "gorm.io/gorm"

type Sensor struct {
	gorm.Model
	PumpID      uint
	SensorDatas []SensorData
}

func (sensor *Sensor) GetByID(db *gorm.DB, id uint) error {
	return db.First(sensor, id).Error
}

func (sensor Sensor) Save(db *gorm.DB) error {
	return db.Save(sensor).Error
}
