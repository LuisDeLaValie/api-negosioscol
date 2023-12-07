package handlers

import "github.com/gin-gonic/gin"

func GetLocal(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World! asdasd",
	})
}
