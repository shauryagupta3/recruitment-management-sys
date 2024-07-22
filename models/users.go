package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID              int64   `json:"id" gorm:"primarykey"`
	Name            string  `json:"name"`
	Email           string  `json:"email" gorm:"unique"`
	Address         bool    `json:"address"`
	UserType        string  `json:"user_type"`
	PasswordHash    string  `json:"password_hash"`
	ProfileHeadline string  `json:"profile_headline"`
}
