package models

import "gorm.io/gorm"

type User struct {
	ID            uint           `json:"id"`
	Username      string         `json:"username"`
	PasswordHash  string         `json:""` // Never serialize the password hash
	Installations []Installation `json:"installations"`
}

// Create or update User
func (u User) Save(db *gorm.DB) error {
	return db.Create(&u).Error
}

// Get User with its Installations by UserID
func (u *User) GetByID(db *gorm.DB, userID uint) error {
	return db.Model(&User{}).Where("id == ?", userID).First(u).Error
}

// Find user with a given username
func (u *User) FindByUsername(db *gorm.DB, username string) error {
	query := &User{Username: username}
	return db.Where(query).First(u).Error
}

// Find user by credentials (can be used for login)
func (u *User) FindByCreds(db *gorm.DB, username, passwordHash string) error {
	query := &User{Username: username, PasswordHash: passwordHash}
	return db.Where(query).First(u).Error
}
