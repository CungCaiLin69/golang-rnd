package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SampleController(c *gin.Context) {
	payload, _ := c.Get("payload")
	userAccount, _ := c.Get("UserAccount")
	matrix, _ := c.Get("UserMatrix")
	c.JSON(http.StatusOK, gin.H{"payload": payload, "uac": userAccount, "matrix": matrix})
}
