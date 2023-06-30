package http

import (
	"net/http"
	"github.com/AvinFajarF/internal/usecase"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user struct {
		Username     string `json:"username"`
		Password     string `json:"password"`
		Alamat       string `json:"alamat"`
		Email        string `json:"email"`
		NomorTelepon int    `json:"nomor_telepon"`
	}

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashedPasswordString := string(hash)

	result, err := h.userUsecase.CreateUser(user.Username, hashedPasswordString, user.Alamat, user.Email, user.NomorTelepon)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   result,
	})

}
