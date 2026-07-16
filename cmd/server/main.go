package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yash-pokhriyal/realtime-chat-backend/internal/config"
)

func main(){

	cfg,err:=config.Load()
	if err!=nil{
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/health",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"status":"success",
			"message":"Server is running",
		})
	})

	log.Printf("Server started on port %s",cfg.Port)
	err = router.Run(":"+cfg.Port)
	if err!=nil{
		log.Fatal(err)
	}

	
}

