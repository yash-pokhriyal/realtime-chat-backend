package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword converts plain password into hashed password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword compares plain password with hashed password.
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


// Agar kal Login API me password verify karna hua to dobara bcrypt code nahi likhenge.

// Flow:

// Register
//       ↓
// HashPassword()

// Login
//       ↓
// CheckPassword()

// Ek hi jagah password logic rahega.

// bcrypt ka concept

// Abhi:

// 123456

// Database me save ho raha tha.

// Ab hoga:

// $2a$10$uYf....

// Ye one-way hash hai.

// Isse original password wapas nahi nikala ja sakta.

// GenerateFromPassword()
// bcrypt.GenerateFromPassword(...)

// Ye plaintext leta hai.

// 123456

// ↓

// Hash

// ↓

// $2a$10$...
// CompareHashAndPassword()

// Login me:

// User likhega

// 123456

// Database me

// $2a$10$...

// bcrypt compare karega.