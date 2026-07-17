package repository

import (
	"github.com/yash-pokhriyal/realtime-chat-backend/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User

	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// 1.
// type UserRepository struct {
// 	DB *gorm.DB
// }
// Kyu?

// Ye repository ke paas database connection store karta hai.

// Without this:

// db.Create(...)

// chal hi nahi sakta.

// 2.
// DB *gorm.DB
// Pointer kyu?

// Database connection bahut bada object hota hai.

// Har function me copy nahi karna.

// Sab same connection use kare.

// 3.
// func NewUserRepository(db *gorm.DB) *UserRepository
// Ye kya hai?

// Ye constructor hai.

// Jaise car banane ke liye factory hoti hai.

// Waise repository banane ke liye constructor.

// Use:

// repo := repository.NewUserRepository(db)
// 4.
// return &UserRepository{
// 	DB: db,
// }
// & kyu?

// Pointer return kar rahe hain.

// Object copy nahi hoga.

// Ek hi repository instance use hoga.

// 5.
// func (r *UserRepository)

// Ye receiver hai.

// Matlab

// repo.Create(...)

// allowed hoga.

// Agar receiver nahi hota to

// Create(...)

// ek normal function hota.

// 6.
// Create(user *models.User)

// Pointer kyu?

// Struct copy nahi hota.
// GORM insert ke baad ID fill karta hai usi object me.

// Example:

// user := models.User{Name: "Yash"}

// repo.Create(&user)

// fmt.Println(user.ID)

// ID automatically aa jayegi.

// 7.
// r.DB.Create(user)

// GORM internally SQL banata hai.

// Tum likhte ho

// r.DB.Create(user)

// Andar ye chalta hai:

// INSERT INTO users (...)
// VALUES (...)
// 8.
// .Error

// Har database operation fail bhi ho sakta hai.

// Jaise:

// Duplicate email
// Database band
// Network issue

// Isliye error return karte hain.

// 9.
// Where("email = ?", email)

// ? placeholder hai.

// Ye safe hai.

// ❌ Galat:

// Where("email = '" + email + "'")

// Isse SQL Injection ka risk hota hai.

// ✅ Sahi:

// Where("email = ?", email)
// 10.
// First(&user)

// Pehla matching record laata hai.

// SQL:

// SELECT * FROM users
// WHERE email = ...
// LIMIT 1;
// 11.
// return &user

// Pointer return.

// Copy nahi hogi.

// 🔥 Sabse Important Concept
// Repository Pattern
// Without Repository
// Handler

// ↓

// db.Create()

// ↓

// db.First()

// ↓

// db.Delete()

// Har jagah database code.

// With Repository
// Handler

// ↓

// repo.Create()

// ↓

// Repository

// ↓

// GORM

// ↓

// Database

// Agar kal GORM hata kar pgx use karna ho, to sirf repository badlegi. Handlers ko touch nahi karna padega.

// Flow
// Browser

// ↓

// Handler

// ↓

// Repository

// ↓

// GORM

// ↓

// PostgreSQL