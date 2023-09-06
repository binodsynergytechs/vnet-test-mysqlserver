package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	// Add other user fields here
}
