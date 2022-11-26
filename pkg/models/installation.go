package models

import "gorm.io/gorm"

type Installation struct {
	gorm.Model
	lat    float64
	long   float64
	userID uint
	Tanks  []Tank
}
