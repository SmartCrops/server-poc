package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username      string
	PasswordHash  string
	Installations []Installation
}

func (user User) Save(db *gorm.DB) error {
	return db.Create(&user).Error
}
