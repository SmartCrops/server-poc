package models

import "gorm.io/gorm"

type Pump struct {
	gorm.Model
	Gpio    uint8
	TankID  uint
	Sensors []Sensor
}
