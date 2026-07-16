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


// "We use GORM because it reduces boilerplate code, maps Go structs to database tables, provides CRUD operations, supports auto migrations, and speeds up backend development. For applications that need maximum SQL control or performance, raw SQL or libraries like pgx can be a better choice."


// "Humne GORM use kiya taaki hume basic CRUD operations ke liye baar-baar raw SQL na likhni pade, Go structs ko database tables se easily map kar saken, aur development fast aur maintainable rahe."


// "Humne GORM use kiya taaki repetitive SQL na likhni pade, Go structs ko database tables se map kar saken, aur backend development faster aur cleaner ho."