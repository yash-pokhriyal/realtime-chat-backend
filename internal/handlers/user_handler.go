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
	"github.com/yash-pokhriyal/realtime-chat-backend/internal/config"
)

type UserHandler struct {
	Repo *repository.UserRepository
	Config *config.Config
}

func NewUserHandler(repo *repository.UserRepository, cfg *config.Config) *UserHandler {
	return &UserHandler{
		Repo: repo,
		Config: cfg,
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

func (h *UserHandler) Login(c *gin.Context) {

	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	user, err := h.Repo.GetByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	if !utils.CheckPassword(req.Password, user.Password) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "Invalid email or password",
	})
	return
    }

	token, err := utils.GenerateToken(user.ID, h.Config.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
	"message": "Login successful",
	"token":   token,
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



// Aaj ka ek interview question

// Q: Registration me password hash kiya, lekin Login me original password kaise verify karoge jab hash ko reverse nahi kar sakte?

// Answer:

// Hum hash ko reverse nahi karte. User jo password login ke time deta hai, usko bcrypt.CompareHashAndPassword() se stored hash ke against compare karte hain. Agar match hota hai to login successful hota hai.

// Client
//    │
//    ▼
// ShouldBindJSON()
//    │
//    ▼
// GetByEmail()
//    │
//    ├── Not Found
//    │      ▼
//    │   401
//    │
//    ▼
// CheckPassword()
//    │
//    ├── Wrong Password
//    │      ▼
//    │   401
//    │
//    ▼
// Success
//    │
//    ▼
// 200 OK



// 📖 Aaj ka interview concept

// Q: Login me ye kyu nahi bataya ki "Email not found" ya "Wrong password"?

// A: Security reasons. Agar alag-alag messages doge, attacker valid emails enumerate kar sakta hai. Isliye dono cases me same response dete hain:

// {
//   "error": "Invalid email or password"
// }

// Isse account enumeration attack se protection milti hai.



