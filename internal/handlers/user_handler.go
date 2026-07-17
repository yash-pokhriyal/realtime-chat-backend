package handlers

import (
	"net/http"
    "log"
	"github.com/gin-gonic/gin"
	"github.com/yash-pokhriyal/realtime-chat-backend/internal/models"
	"github.com/yash-pokhriyal/realtime-chat-backend/internal/repository"
	"github.com/yash-pokhriyal/realtime-chat-backend/internal/utils"
	"errors"

    "gorm.io/gorm"
)

type UserHandler struct {
	Repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		Repo: repo,
	}
}

func (h *UserHandler) Register(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

    existingUser, err := h.Repo.GetByEmail(user.Email)

	if err == nil && existingUser != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Email already registered",
		})
		return
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database error",
		})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user.Password = hashedPassword

    err = h.Repo.Create(&user)
	if err != nil {
		log.Println("Create User Error:", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	  return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user": gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	    },
	})
}


// ShouldBindJSON()
//         │
//         ▼
// GetByEmail()
//         │
//         ├── User exists
//         │       │
//         │       ▼
//         │   409 Conflict
//         │
//         └── User doesn't exist
//                 │
//                 ▼
// HashPassword()
//                 │
//                 ▼
// Create()
//                 │
//                 ▼
// Response




// Ab samajhte hain ye code
// 1.
// existingUser, err := h.Repo.GetByEmail(user.Email)

// Ye database me check karega:

// SELECT * FROM users WHERE email='yash@gmail.com';
// 2.
// if err == nil && existingUser != nil

// Matlab:

// Query successful hui.
// User mil gaya.

// To:

// return "Email already registered"
// 3.
// errors.Is(err, gorm.ErrRecordNotFound)

// Ye bahut important hai.

// Agar email nahi mili, GORM ye error deta hai:

// gorm.ErrRecordNotFound

// Aur ye error normal hai.

// Matlab:

// User nahi mila
// ↓

// Registration allow karo.
// 4.

// Lekin agar koi aur error aaye:

// PostgreSQL band ho gaya
// Network issue
// Query fail ho gayi

// Tab:

// Database error

// Return karenge.

// Ek interview question ⭐

// Q: ErrRecordNotFound ko error kyu nahi maan rahe?

// Answer:

// Kyuki registration ke time user ka na milna expected behavior hai.

// Hum actually yahi expect karte hain ki email pehle se database me na ho.

// Isliye:

// Record Not Found

// = Success case for registration.