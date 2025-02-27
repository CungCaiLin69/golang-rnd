package handlers

import (
	"golang-rnd/controllers"
	"golang-rnd/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginHandler generates and returns JWT
func LoginHandler(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Dummy validation (replace with DB check)
	if req.Username != "sendy" || req.Password != "sendy123" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := controllers.GenerateJWT(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// DataHandler handles protected data access
func DataHandler(c *gin.Context) {
	var req models.DataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{
		"message": "Authorized access",
		"user":    username,
		"data":    req.Message,
	})
}
