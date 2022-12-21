package models

import "gorm.io/gorm"

type User struct {
	ID            uint
	Username      string
	PasswordHash  string
	Installations []Installation
}

// Create or update User
func (user User) Save(db *gorm.DB) error {
	return db.Create(&user).Error
}

// Get User with its Installations by UserID
func (user *User) GetByID(db *gorm.DB, userId uint) error {
	return db.Model(&User{}).Where("id == ?", userId).First(user).Error
}
