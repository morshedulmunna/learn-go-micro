package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/morshedulmunna/auth-service/internal/model"
)

type AuthHandler struct {
	// Add any dependencies here
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var creds model.Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Implement actual authentication logic
	token := generateToken(creds.Email)

	c.JSON(http.StatusOK, model.TokenResponse{
		AccessToken:  token,
		RefreshToken: "refresh_token",
		TokenType:    "Bearer",
		ExpiresIn:    3600,
	})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var creds model.Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Implement actual registration logic

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func generateToken(email string) string {
	// This is a simple example. In production, use proper secret management
	secret := []byte("your-secret-key")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString(secret)
	return tokenString
}
