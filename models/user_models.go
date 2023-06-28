package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
	Role     string
}

func GetUserByUsername(db *gorm.DB, username string) (*User, error) {
	user := new(User)
	result := db.Where("username = ?", username).First(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
