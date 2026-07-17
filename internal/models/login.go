package models

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}



// Register:

// Client

// ↓

// Create User

// ↓

// Database

// Login:

// Client

// ↓

// Find User

// ↓

// Verify Password

// ↓

// Return Success