package models

import "gorm.io/gorm"

type Sensor struct {
	gorm.Model
	PumpID      uint
	SensorDatas []SensorData
}
