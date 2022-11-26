package models

import "gorm.io/gorm"

type Tank struct {
	gorm.Model
	Pumps []Pump
}
