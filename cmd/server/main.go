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


// Anonymous function kya hota hai?
// Naam ke bina function, jise callback ki tarah pass kiya ja sakta hai.
// Gin Context kya hai?
// Ek object jo current HTTP request aur response se related saari information aur helper methods provide karta hai.
// *gin.Context pointer kyun hai?
// Taaki bada context object copy na ho aur handlers same request/response object par kaam kar sake.
// gin.H kya hai?
// map[string]interface{} ka shortcut.