package http

import (
	"net/http"
	"os"
	"time"

	"github.com/AvinFajarF/internal/entity"
	"github.com/AvinFajarF/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
		ID           string `json:"id"`
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

func (h *UserHandler) Login(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	result, err := h.userUsecase.Login(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	key := []byte(os.Getenv("SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": result.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(key)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"massage": "error membuat token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.Header("Authorization", tokenString)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  tokenString,
		"data":   result,
	})

}
