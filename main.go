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
			// servicios.GET("/", handlers.Get)
			servicios.GET("/:id", handlers.GetServisioPorID)

			// Ruta para crear un nuevo usuario (POST)
			servicios.POST("", handlers.CrearServisio)

			// Ruta para actualizar un usuario existente (PUT)
			servicios.PUT("/:id", handlers.ActualizarServisio)
			servicios.PATCH("/:id", handlers.RemplanzarServisio)

			// Ruta para eliminar un usuario (DELETE)
			servicios.DELETE("/:id", handlers.EliminarServisio)
		}

		producto := api.Group("/producto")
		{
			// Ruta para obtener un usuario por su ID (GET)
			// producto.GET("/", handlers.Get)
			producto.GET("/:id", handlers.GetProductoPorID)

			// Ruta para crear un nuevo usuario (POST)
			producto.POST("", handlers.CrearProducto)

			// Ruta para actualizar un usuario existente (PUT)
			producto.PUT("/:id", handlers.ActualizarProducto)
			producto.PATCH("/:id", handlers.RemplanzarProducto)

			// Ruta para eliminar un usuario (DELETE)
			producto.DELETE("/:id", handlers.EliminarProducto)
		}

	}

	// Run the application
	router.Run(":8080")
}
