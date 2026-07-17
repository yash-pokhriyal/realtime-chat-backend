package database

import (
	// "fmt"
	"log"

	"github.com/yash-pokhriyal/realtime-chat-backend/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/yash-pokhriyal/realtime-chat-backend/internal/models"
)


func Connect(cfg *config.Config) (*gorm.DB, error) {

	dsn := "postgres://er.yashpokhriyal@localhost:5432/realtime_chat?sslmode=disable"
	log.Println("DSN:", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	log.Println("PostgreSQL connected successfully")

	return db, nil
}

// func Connect(cfg *config.Config) (*gorm.DB, error)
// cfg *config.Config → Configuration object leta hai.
// * pointer hai, taaki object copy na ho.
// *gorm.DB → Connected database object return karta hai.
// error → Connection fail ho to error return karega.


// DSN (Data Source Name)

// Database connection string.

// Example:

// host=localhost user=postgres password=postgres dbname=realtime_chat port=5432

// Isse PostgreSQL ko pata chalta hai kis database se connect karna hai.

// gorm.Open()

// Database connection initialize karta hai aur *gorm.DB object return karta hai.

// postgres.Open(dsn)

// GORM ko batata hai ki PostgreSQL driver use karna hai.

// &gorm.Config{}

// GORM ki configuration object.

// Abhi default settings use ho rahi hain.

// return db, nil

// Database successfully connect hua.

// db → Connected database object
// nil → Koi error nahi.
// db, err := database.Connect(cfg)

// Database se connect karta hai.

// Agar error aaye to:

// if err != nil {
//     log.Fatal(err)
// }

// Application band ho jayegi.

// _ = db

// Temporary compiler ko batata hai ki variable intentionally abhi use nahi ho raha.

// Baad me isi db se queries chalengi.



// DSN? → Database connection string.
// Why *Config? → Copy avoid karne aur efficiency ke liye.
// Why *gorm.DB? → Connected DB instance ko share karne ke liye.
// fmt.Sprintf vs fmt.Println? → Sprintf string return karta hai, Println console par print karta hai.
// gorm.Open? → Database connection initialize karta hai.
// postgres.Open? → PostgreSQL driver select karta hai.
// &gorm.Config{}? → GORM ki configuration settings.
// return db, nil? → Database object return, error nahi.


// AutoMigrate kya karta hai?

// Ye GORM ko bolta hai:

// "Mere User struct ko dekh aur uske hisaab se database me table bana de."

// Flow:

// User Struct
//       ↓
// AutoMigrate()
//       ↓
// PostgreSQL
//       ↓
// users table create