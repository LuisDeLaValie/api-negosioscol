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

	api := router.Group("/api")
	{

		usuarios := api.Group("/usuarios")
		{

			// Ruta para obtener un usuario por su ID (GET)
			usuarios.GET("/:id", handlers.GetUsuarioPorID)

			// Ruta para crear un nuevo usuario (POST)
			usuarios.POST("/", handlers.CrearUsuario)

			// Ruta para actualizar un usuario existente (PUT)
			usuarios.PUT("/:id", handlers.ActualizarUsuario)
			usuarios.PATCH("/:id", handlers.RemplanzarUsuario)

			// Ruta para eliminar un usuario (DELETE)
			usuarios.DELETE("/:id", handlers.EliminarUsuario)
		}

		servicios := api.Group("/servicios")
		{
			// Ruta para obtener un usuario por su ID (GET)
			servicios.GET("/", handlers.GetServicios)
			servicios.GET("/:id", handlers.GetServicioPorID)

			// Ruta para crear un nuevo usuario (POST)
			servicios.POST("", handlers.CrearServicio)

			// Ruta para actualizar un usuario existente (PUT)
			servicios.PUT("/:id", handlers.ActualizarServicio)
			servicios.PATCH("/:id", handlers.RemplanzarServicio)

			// Ruta para eliminar un usuario (DELETE)
			servicios.DELETE("/:id", handlers.EliminarServicio)
		}

	}

	// Run the application
	router.Run(":8080")
}
