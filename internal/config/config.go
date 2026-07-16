package config
import (
	"os"
    "errors"
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

// Config ka kaam hai application ki settings ko manage karna.

// .env

// ↓

// config.Load()

// ↓

// Config Struct

// ↓

// Poori application

func Load()(*Config,error){
	err:=godotenv.Load()
	if err!=nil{
		return nil,err
	}

	cfg := &Config{
		Port:          os.Getenv("PORT"),
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
		RedisAddr:     os.Getenv("REDIS_ADDR"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
	}

	if cfg.Port == ""{
		return nil,errors.New("PORT is required")
	}
	if cfg.JWTSecret == ""{
		return nil,errors.New("JWT_SECRET is required")
	}
	return cfg,nil
}