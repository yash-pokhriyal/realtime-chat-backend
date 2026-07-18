package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yash-pokhriyal/realtime-chat-backend/internal/config"
	"github.com/yash-pokhriyal/realtime-chat-backend/internal/repository"
	"github.com/yash-pokhriyal/realtime-chat-backend/internal/handlers"

	"github.com/yash-pokhriyal/realtime-chat-backend/internal/database"
	"github.com/yash-pokhriyal/realtime-chat-backend/internal/middleware"
	"github.com/yash-pokhriyal/realtime-chat-backend/internal/websocket"
)

func main(){

	cfg,err:=config.Load()
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("PORT:", cfg.Port)
	log.Println("DB_USER:", cfg.DBUser)
	log.Println("DB_NAME:", cfg.DBName)

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

   //  Initialize Repository
	repo := repository.NewUserRepository(db)

	//  Initialize Handler
	userHandler := handlers.NewUserHandler(repo,cfg)


	router := gin.Default()

	// Initialize WebSocket Hub
	hub := websocket.NewHub()

	go hub.Run()

	// Initialize WebSocket Handler
	wsHandler := websocket.NewHandler(hub)

	router.GET("/health",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"status":"success",
			"message":"Server is running",
		})
	})

	// Register API
	router.POST("/register", userHandler.Register)

	router.POST("/login", userHandler.Login)

	router.GET("/ws", wsHandler.HandleConnections)

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))

	protected.GET("/profile", func(c *gin.Context) {

	userID, _ := c.Get("userID")

	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome! You are authenticated.",
		"userID":  userID,
	})
	})


	log.Printf("Server started on port %s",cfg.Port)
	err = router.Run(":"+cfg.Port)
	if err!=nil{
		log.Fatal(err)
	}

	
}


// Anonymous function kya hota hai?
// Naam ke bina function, jise callback ki tarah pass kiya ja sakta hai.
// Gin Context kya hai?
// Ek object jo current HTTP request aur response se related saari information aur helper methods provide karta hai.
// *gin.Context pointer kyun hai?
// Taaki bada context object copy na ho aur handlers same request/response object par kaam kar sake.
// gin.H kya hai?
// map[string]interface{} ka shortcut.