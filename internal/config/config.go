package config
import (
	"os"

	"github.com/joho/godotenv"
)

// os → environment variables padhne ke liye.
// godotenv → .env file ko load karke environment variables me daalne ke liye.

// Production me usually .env file nahi hoti. Cloud (AWS, GCP, Docker, Kubernetes) environment variables directly provide karta hai.

type Config struct {
	Port           string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	JWTSecret      string
	RedisAddr      string
	RedisPassword  string
}

// 💡 Concept: Hum saari configuration ek struct me rakhte hain. Isse 8 alag variables pass karne ki jagah ek hi Config object pass hota hai.

