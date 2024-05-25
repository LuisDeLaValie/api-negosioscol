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

	defer func() {
		err := recover()
		if err != nil {
			errcode := models.Error500("Ocurrió un problema para procesar la solicitud:\n %v", err)
			c.JSON(errcode.Code, errcode)
		}
	}()

	var usuario map[string]interface{}
	if err := c.ShouldBindJSON(&usuario); err != nil {
		errcode := models.Error500("Ocurrió un problema para procesar la solicitud" + err.Error())
		c.JSON(errcode.Code, errcode)

		return
	}

	// Aquí puedes usar los datos del usuario
	nombre := usuario["Nombre"].(string)
	apellidos := usuario["Apellidos"].(string)
	correo := usuario["Correo"].(string)
	password := usuario["Password"].(string)
	cumpleanos := usuario["Cumpleanos"].(string)
	imagen := usuario["Imagen"].(string)

	if nombre == "" || apellidos == "" || correo == "" || password == "" || cumpleanos == "" || imagen == "" {
		errcode := models.Error400("Faltan datos.")
		c.JSON(errcode.Code, errcode)

		return
	}

	lastID, err := models.CrearUsuario(nombre, apellidos, correo, password, cumpleanos, imagen)
	if err != nil {
		c.JSON(err.Code, err)

		return
	}

	user, resE := models.ObtenerUsuario(*lastID)
	if resE != nil {
		c.JSON(resE.Code, resE)
		return
	}

	c.JSON(200, user)
}

// Actualizar un usuario existente por su ID
func ActualizarUsuario(c *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			errcode := models.Error500("Ocurrió un problema para procesar la solicitud:\n %v", err)
			c.JSON(errcode.Code, errcode)
		}
	}()

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
	correo := usuario["Correo"].(string)
	password := usuario["Password"].(string)
	cumpleanos := usuario["Cumpleanos"].(string)
	imagen := usuario["Imagen"].(string)

	if nombre == "" || apellidos == "" || correo == "" || password == "" || cumpleanos == "" || imagen == "" {
		errcode := models.Error400("Faltan datos.")
		c.JSON(errcode.Code, errcode)

		return
	}

	resE := models.EditarUsuario(idd, nombre, apellidos, correo, password, cumpleanos, imagen)
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
	defer func() {
		err := recover()
		if err != nil {
			errcode := models.Error500("Ocurrió un problema para procesar la solicitud:\n %v", err)
			c.JSON(errcode.Code, errcode)
		}
	}()

	id := c.Param("id")
	idd, err := strconv.Atoi(id)
	if err != nil {
		errcode := models.Error400("El id no es valido.")
		c.JSON(errcode.Code, errcode)
		return
	}

	resE := models.EliminarUsuario(idd)
	if resE != nil {
		c.JSON(resE.Code, resE)

		return
	}

	c.JSON(200, gin.H{
		"message": "EliminarUsuario:" + id,
	})

}

func RemplanzarUsuario(c *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			errcode := models.Error500("Ocurrió un problema para procesar la solicitud:\n %v", err)
			c.JSON(errcode.Code, errcode)
		}
	}()

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
	correo := usuario["Correo"].(string)
	password := usuario["Password"].(string)
	cumpleanos := usuario["Cumpleanos"].(string)
	imagen := usuario["Imagen"].(string)

	if nombre == "" || apellidos == "" || correo == "" || password == "" || cumpleanos == "" || imagen == "" {
		errcode := models.Error400("Faltan datos.")
		c.JSON(errcode.Code, errcode)

		return
	}

	if resE := models.EliminarUsuario(idd); resE != nil {
		c.JSON(resE.Code, resE)

		return
	}

	lastID, resE := models.CrearUsuario(nombre, apellidos, correo, password, cumpleanos, imagen)
	if resE != nil {
		c.JSON(resE.Code, resE)

		return
	}

	user, resE := models.ObtenerUsuario(*lastID)
	if resE != nil {
		c.JSON(resE.Code, resE)
		return
	}

	c.JSON(200, user)
}
