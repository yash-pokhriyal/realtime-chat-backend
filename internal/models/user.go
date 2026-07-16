package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-"`
}

// Interview One-Liners
// gorm.Model

// Built-in struct jo ID, timestamps aur soft delete fields provide karta hai.

// gorm:"unique"

// Duplicate values ko database level par prevent karta hai.

// gorm:"not null"

// Field ko empty save hone se rokta hai.

// json:"-"

// Field ko API response se hide karta hai.

// AutoMigrate()

// Struct ke basis par database schema create/update karta hai.