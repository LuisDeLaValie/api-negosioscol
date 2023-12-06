package main

import (
	"negosioscol/src/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World! asdasd",
		})
	})

	// Ruta para obtener un usuario por su ID (GET)
	router.GET("/usuarios/:id", handlers.GetUsuarioPorID)

	// Ruta para crear un nuevo usuario (POST)
	router.POST("/usuarios", handlers.CrearUsuario)

	// Ruta para actualizar un usuario existente (PUT)
	router.PUT("/usuarios/:id", handlers.ActualizarUsuario)
	router.PATCH("/usuarios/:id", handlers.RemplanzarUsuario)

	// Ruta para eliminar un usuario (DELETE)
	router.DELETE("/usuarios/:id", handlers.EliminarUsuario)

	// Run the application
	router.Run(":8080")
}
