package handlers

import (
	"net/http"
    "log"
	"github.com/gin-gonic/gin"
	"github.com/yash-pokhriyal/realtime-chat-backend/internal/models"
	"github.com/yash-pokhriyal/realtime-chat-backend/internal/repository"
	"github.com/yash-pokhriyal/realtime-chat-backend/internal/utils"
)

type UserHandler struct {
	Repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		Repo: repo,
	}
}

func (h *UserHandler) Register(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user.Password = hashedPassword

    err = h.Repo.Create(&user)
	if err != nil {
		log.Println("Create User Error:", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	  return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}
