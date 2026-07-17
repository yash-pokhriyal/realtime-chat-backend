package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID uint, secret string) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}


// 1
// jwt.MapClaims

// Token ke andar data store hota hai.

// Jaise

// user_id = 5

// expiry = tomorrow
// 2
// exp

// Token kab expire hoga.

// Hum

// 24*time.Hour

// rakh rahe hain.

// 3
// jwt.NewWithClaims()

// JWT object banata hai.

// 4
// SignedString()

// Sabse important.

// Ye token ko

// Header
// Payload
// Signature

// me convert karta hai.

// Result kuch aisa dikhega

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
// Interview

// JWT me kya store karna chahiye?

// ✅

// user_id

// role

// expiry

// ❌

// password

// email

// OTP

// sensitive data

// JWT encrypted nahi hota, sirf signed hota hai.