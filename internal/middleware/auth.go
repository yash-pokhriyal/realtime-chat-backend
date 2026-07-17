package middleware
import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header missing",
			})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization format",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token claims",
		})
		c.Abort()
		return
	}

	id, ok := claims["user_id"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid user ID in token",
		})
		c.Abort()
		return
	}

	userID := uint(id)

	c.Set("userID", userID)



		c.Next()
	}
}

// 1.
// c.GetHeader("Authorization")

// Reads:

// Authorization: Bearer eyJhbGc...
// 2.
// strings.HasPrefix()

// Checks:

// Bearer ....

// Likha hai ya nahi.

// 3.
// TrimPrefix()

// Removes

// Bearer

// Aur sirf JWT token bachta hai.

// 4.
// jwt.Parse()

// Verify karta hai

// Signature
// Expiry
// Secret

// Sab sahi hua?

// token.Valid == true
// 5.
// c.Abort()

// Middleware yahin request ko rok deta hai.

// Agar ye na likho to invalid user bhi next handler tak pahunch jayega.

// 6.
// c.Next()

// Valid token hai.

// Next handler execute karo.

// Flow
// Login

// ↓

// JWT Token

// ↓

// Authorization Header

// ↓

// Middleware

// ↓

// Protected Route



// Q: c.Abort() aur c.Next() me difference?

// Answer:

// c.Abort() → Request chain yahin stop ho jati hai. Baaki handlers execute nahi hote.
// c.Next() → Current middleware ke baad next middleware ya route handler execute hota hai.



// Client
//    │
//    ▼
// Authorization: Bearer <token>
//    │
//    ▼
// JWT Middleware
//    │
//    ├── Invalid Token → 401
//    │
//    └── Valid Token
//            │
//            ▼
//       Next Handler