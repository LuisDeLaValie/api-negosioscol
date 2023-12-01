package handlers

import "github.com/gin-gonic/gin"


func getLocal(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World! asdasd",
	})
}