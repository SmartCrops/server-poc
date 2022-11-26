package models

import "gorm.io/gorm"

type Pump struct {
	gorm.Model
	Gpio    uint
	TankID  uint
	Sensors []Sensor
}
