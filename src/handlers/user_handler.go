package handlers

import (
	"negosioscol/src/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Obtener un usuario por su ID
func GetUsuarioPorID(c *gin.Context) {
	id := c.Param("id")
	idd, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"errno":             400,
			"error":             "bad_request",
			"error_description": "El id no es valido.",
		})
		return
	}

	user, err := models.ObtenerUsuario(idd)
	if err != nil {
		c.JSON(404, gin.H{
			"errno":             404,
			"error":             "not_found",
			"error_description": "No se encontró el Uasuario.",
		})
		return
	}

	c.JSON(200, user)
}

// Crear un nuevo usuario
func CrearUsuario(c *gin.Context) {

	var usuario map[string]interface{}
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(500, gin.H{
			"errno":             500,
			"error":             "internal_error",
			"error_description": "Ocurrió un problema para procesar la solicitud" + err.Error(),
		})
		return
	}

	// Aquí puedes usar los datos del usuario
	nombre := usuario["Nombre"].(string)
	apellidos := usuario["Apellidos"].(string)
	cumpleanos := usuario["Cumpleanos"].(string)
	imagen := usuario["Imagen"].(string)

	if nombre == "" || apellidos == "" || cumpleanos == "" || imagen == "" {
		c.JSON(400, gin.H{
			"errno":             400,
			"error":             "bad_request",
			"error_description": "Faltan datos.",
		})
		return
	}

	err := models.CrearUsuario(nombre, apellidos, cumpleanos, imagen)
	if err != nil {
		c.JSON(500, gin.H{
			"errno":             500,
			"error":             "internal_error",
			"error_description": "Ocurrió un problema para procesar la solicitud",
		})
		return
	}

	c.JSON(201, gin.H{})
}

// Actualizar un usuario existente por su ID
func ActualizarUsuario(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "ActualizarUsuario:" + id,
	})
}

// Eliminar un usuario por su ID
func EliminarUsuario(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "EliminarUsuario:" + id,
	})
}
