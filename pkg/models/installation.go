package models

import "gorm.io/gorm"

type Installation struct {
	ID     uint
	Name   string
	UserID uint    // owner
	Users  []*User `gorm:"many2many:user_installations;"` // managing users
}

// Create Installation owned by User with userId
func (installation Installation) Save(db *gorm.DB, userId uint) error {
	return db.Create(&installation).Association("Users").Append(&User{ID: userId})
}

// Add User to Installation
func (installation *Installation) AddUser(db *gorm.DB, userId uint) error {
	return db.Model(installation).Association("Users").Append(&User{ID: userId})
}

// Remove User from Installation
func (installation *Installation) RemoveUser(db *gorm.DB, userId uint) error {
	return db.Model(installation).Association("Users").Delete(&User{ID: userId})
}

// Get by InstallationID
func (installation *Installation) GetByID(db *gorm.DB, installationId uint) error {
	return db.Model(&Installation{}).Where("id == ?", installationId).First(installation).Error
}
