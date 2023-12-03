package handlers

import "github.com/gin-gonic/gin"

// Obtener todos los usuarios
func GetUsuarios(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetUsuarios:",
	})
}

// Obtener un usuario por su ID
func GetUsuarioPorID(c *gin.Context) {
    id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "GetUsuarioPorID:"+id,
	})
}

// Crear un nuevo usuario
func CrearUsuario(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "CrearUsuario:",
	})
}

// Actualizar un usuario existente por su ID
func ActualizarUsuario(c *gin.Context) {
    id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "ActualizarUsuario:"+id,
	})
}

// Eliminar un usuario por su ID
func EliminarUsuario(c *gin.Context) {
    id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "EliminarUsuario:"+id,
	})
}
