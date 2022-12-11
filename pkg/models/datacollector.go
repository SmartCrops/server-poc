package models

import "gorm.io/gorm"

type DataCollector struct {
	gorm.Model
	PumpID       uint
	SerialNumber string       `gorm:"unique"`
	SensorData   []SensorData `gorm:"foreignKey:DataCollectorSerialNumber;references:SerialNumber"`
}

func (dataCollector *DataCollector) GetByID(db *gorm.DB, id uint) error {
	return db.First(dataCollector, id).Error
}

func (dataCollector DataCollector) Save(db *gorm.DB) error {
	return db.Save(&dataCollector).Error
}
