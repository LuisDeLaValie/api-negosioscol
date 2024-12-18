package main

import (
	"fmt"
	"negosioscol/src/handlers"
	"negosioscol/src/models"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World! asdasd",
		})
	})

	// Configuración del middleware CORS
	config := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(config))

	api := router.Group("/api")
	{
		api.GET("/buscar", func(c *gin.Context) {

			defer func() {
				err := recover()
				if err != nil {
					errcode := models.Error500("Ocurrió un problema para procesar la solicitud:\n %v", err)
					c.JSON(errcode.Code, errcode)
				}
			}()

			buscar := c.Query("buscar")

			servi, resE := models.BuscarServicioProducto(buscar)
			if resE != nil {
				c.JSON(resE.Code, resE)
				return
			}

			c.JSON(200, servi)
		})
		api.POST("/login", handlers.LonginNegocio)
		api.GET("/graficas", handlers.Grafias)

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
			// servicios.PATCH("/:id", handlers.RemplanzarServisio)

			// Ruta para eliminar un usuario (DELETE)
			servicios.DELETE("/:id", handlers.EliminarServisio)

			servicios.GET("/lasts", handlers.ObtenerUltimosServisio)
		}

		negocio := api.Group("/negocios")
		{
			// Ruta para obtener un usuario por su ID (GET)
			// negocio.GET("/", handlers.Get)
			negocio.GET("/:id", handlers.GetNegocioPorID)
			negocio.GET("/servisios/:id", handlers.GetServicioNegocioPorID)
			negocio.GET("/productos/:id", handlers.GetProductoNegocioPorID)

			// Ruta para crear un nuevo usuario (POST)
			negocio.POST("", handlers.CrearNegocio)

			// Ruta para actualizar un usuario existente (PUT)
			negocio.PUT("/:id", handlers.ActualizarNegocio)
			// negocio.PATCH("/:id", handlers.RemplanzarNegocio)

			// Ruta para eliminar un usuario (DELETE)
			negocio.DELETE("/:id", handlers.EliminarNegocio)

			negocio.GET("/lasts", handlers.ObtenerUltimosNegocios)
		}

		producto := api.Group("/productos")
		{
			// Ruta para obtener un usuario por su ID (GET)
			// producto.GET("/", handlers.Get)
			producto.GET("/:id", handlers.GetProductoPorID)

			// Ruta para crear un nuevo usuario (POST)
			producto.POST("", handlers.CrearProducto)

			// Ruta para actualizar un usuario existente (PUT)
			producto.PUT("/:id", handlers.ActualizarProducto)
			// producto.PATCH("/:id", handlers.RemplanzarProducto)

			// Ruta para eliminar un usuario (DELETE)
			producto.DELETE("/:id", handlers.EliminarProducto)

			producto.GET("/lasts", handlers.ObtenerUltimoProducto)
		}

	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Run the application
	router.Run(fmt.Sprintf(":%s", port))
}
