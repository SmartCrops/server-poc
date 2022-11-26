package models

import "gorm.io/gorm"

type Tank struct {
	gorm.Model
	InstallationID int
	Pumps          []Pump
}
