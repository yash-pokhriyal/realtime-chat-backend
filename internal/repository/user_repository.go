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


// Repository ka kaam sirf database se baat karna hai.

// Ye decide karti hai:

// User save karna
// User find karna
// User update karna
// User delete karna

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User

	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}


// internal/repository

// Kyun banaya?

// 👉 Database ka saara code ek hi jagah rakhne ke liye.

// user_repository.go

// Kyun?

// 👉 User se related database operations alag file me rakhne ke liye.

// type UserRepository struct
// type UserRepository struct {
//     DB *gorm.DB
// }

// Kyun?

// 👉 Repository ko database access chahiye, isliye DB store kiya.

// DB *gorm.DB

// Kyun pointer?

// 👉 Naya database object copy na ho, wahi existing connection use ho.

// NewUserRepository(db *gorm.DB)
// repo := repository.NewUserRepository(db)

// Kyun?

// 👉 Repository create karne aur database inject karne ke liye.

// Ye constructor hai.

// return &UserRepository{}

// Kyun pointer return kiya?

// 👉 Object copy na ho, ek hi repository instance use ho.

// func (r *UserRepository)
// func (r *UserRepository) Create(...)

// Ye (r *UserRepository) kya hai?

// 👉 Ye receiver hai.

// Matlab Create() function UserRepository ka method hai.

// Jaise:

// repo.Create(user)
// Create(user *models.User)

// Pointer kyun?

// 👉 User object copy na ho.

// Aur GORM insert ke baad ID bhi isi object me fill karta hai.

// r.DB.Create(user)

// Kyun?

// 👉 User ko database me insert karne ke liye.

// .Error

// Kyun?

// 👉 GORM operation fail hua ya nahi, uska error return karta hai.

// GetByEmail(email string)

// Kyun?

// 👉 Login ke time email se user find karna hoga.

// var user models.User

// Kyun?

// 👉 Database ka result store karne ke liye.

// Where("email = ?", email)

// Kyun ??

// 👉 SQL Injection se bachne ke liye.

// ❌ Wrong:

// "email = '" + email + "'"

// ✅ Right:

// Where("email = ?", email)
// First(&user)

// Kyun?

// 👉 Pehla matching record lana.

// return &user

// Pointer kyun?

// 👉 User struct copy na ho.

// Repository Pattern kyun use kiya?

// ❌ Agar na use karte:

// handler1 -> db.Create()

// handler2 -> db.First()

// handler3 -> db.Delete()

// Har jagah database code.

// Database change hua to poora project badlega.

// ✅ Repository use karne par:

// Handler
//     ↓
// Repository
//     ↓
// Database

// Handler ko sirf pata hai:

// repo.Create(user)

// Database ke implementation se koi lena-dena nahi.

// Interview One-Liners
// Repository? → Database access ko alag layer me rakhne ka pattern.
// Constructor? → Object initialize karne ke liye.
// Dependency Injection? → Dependency (db) bahar se pass karna.
// Pointer (*gorm.DB)? → Copy avoid karna aur same DB instance use karna.
// Receiver (func (r *UserRepository)) ? → Function ko struct ka method banana.
// Where("email = ?", email)? → Parameterized query, SQL injection se protection.
// First()? → Pehla matching record fetch karta hai.