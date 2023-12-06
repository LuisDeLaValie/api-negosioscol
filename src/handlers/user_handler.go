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

		errcode := models.Error400("El id no es valido.")
		c.JSON(errcode.Code, errcode)
		return
	}

	user, resE := models.ObtenerUsuario(int64(idd))
	if resE != nil {
		c.JSON(resE.Code, resE)

		return
	}

	c.JSON(200, user)
}

// Crear un nuevo usuario
func CrearUsuario(c *gin.Context) {

	var usuario map[string]interface{}
	if err := c.ShouldBindJSON(&usuario); err != nil {
		errcode := models.Error500("Ocurrió un problema para procesar la solicitud" + err.Error())
		c.JSON(errcode.Code, errcode)

		return
	}

	// Aquí puedes usar los datos del usuario
	nombre := usuario["Nombre"].(string)
	apellidos := usuario["Apellidos"].(string)
	cumpleanos := usuario["Cumpleanos"].(string)
	imagen := usuario["Imagen"].(string)

	if nombre == "" || apellidos == "" || cumpleanos == "" || imagen == "" {
		errcode := models.Error400("Faltan datos.")
		c.JSON(errcode.Code, errcode)

		return
	}

	_, err := models.CrearUsuario(nombre, apellidos, cumpleanos, imagen)
	if err != nil {
		c.JSON(err.Code, err)

		return
	}

	c.JSON(201, gin.H{})
}

// Actualizar un usuario existente por su ID
func ActualizarUsuario(c *gin.Context) {
	id := c.Param("id")
	idd, err := strconv.Atoi(id)
	if err != nil {
		errcode := models.Error400("El id no es valido.")
		c.JSON(errcode.Code, errcode)
		return
	}

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
		errcode := models.Error400("Faltan datos.")
		c.JSON(errcode.Code, errcode)

		return
	}

	resE := models.EditarUsuario(idd, nombre, apellidos, cumpleanos, imagen)
	if resE != nil {
		c.JSON(resE.Code, resE)

		return
	}

	user, resE := models.ObtenerUsuario(int64(idd))
	if resE != nil {
		c.JSON(resE.Code, resE)

		return
	}

	c.JSON(200, user)
}

// Eliminar un usuario por su ID
func EliminarUsuario(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "EliminarUsuario:" + id,
	})
}

func RemplanzarUsuario(c *gin.Context) {
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

	if resE := models.EliminarUsuario(idd); resE != nil {
		c.JSON(resE.Code, resE)

		return
	}

	_, resE := models.CrearUsuario(nombre, apellidos, cumpleanos, imagen)
	if resE != nil {
		c.JSON(resE.Code, resE)

		return
	}

	// user, resE := models.ObtenerUsuario(*idC)
	// if resE != nil {
	// 	c.JSON(resE.Code, resE)

	// 	return
	// }

	c.JSON(201, gin.H{})
}
