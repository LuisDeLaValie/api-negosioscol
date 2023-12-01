
package main


import (
	// "github.com/gorilla/mux"
	// "log"
	// "net/http"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, World! asdasd",
        })
    })

    // Run the application
    router.Run(":8080")
}