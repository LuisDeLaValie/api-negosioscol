package handlers

import (
	"fmt"
	"negosioscol/src/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Obtener un Servicio por su ID
func GetServicioPorID(c *gin.Context) {
	id := c.Param("id")
	idd, err := strconv.Atoi(id)
	if err != nil {

		errcode := models.Error400("El id no es valido.")
		c.JSON(errcode.Code, errcode)
		return
	}

	user, resE := models.ObtenerServisio(int64(idd))
	if resE != nil {
		c.JSON(resE.Code, resE)

		return
	}

	c.JSON(200, user)
}

// Listar los serviso
func GetServicios(c *gin.Context) {
	fmt.Println("getServicios")
	user, resE := models.ListarServisios()
	if resE != nil {
		c.JSON(resE.Code, resE)

		return
	}

	c.JSON(200, user)
}

// Crear un nuevo Servicio
func CrearServicio(c *gin.Context) {

	defer func() {
		err := recover()
		if err != nil {
			errcode := models.Error500("Ocurrió un problema para procesar la solicitud:\n %v", err)
			c.JSON(errcode.Code, errcode)
		}
	}()

	var Servicio map[string]interface{}
	if err := c.ShouldBindJSON(&Servicio); err != nil {
		errcode := models.Error500("Ocurrió un problema para procesar la solicitud" + err.Error())
		c.JSON(errcode.Code, errcode)

		return
	}

	fmt.Println(Servicio)
	// Aquí puedes usar los datos del Servicio
	nombre := Servicio["Nombre"].(string)
	descripcion := Servicio["Descripcion"].(string)
	unidad := int64(Servicio["Unidad"].(float64))

	if nombre == "" || descripcion == "" || unidad == 0 {
		errcode := models.Error400("Faltan datos.")
		c.JSON(errcode.Code, errcode)

		return
	}

	_, err := models.CrearServisio(nombre, descripcion, unidad)
	if err != nil {
		c.JSON(err.Code, err)

		return
	}

	c.JSON(201, gin.H{})
}

// Actualizar un Servicio existente por su ID
func ActualizarServicio(c *gin.Context) {
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

	var Servicio map[string]interface{}
	if err := c.ShouldBindJSON(&Servicio); err != nil {
		c.JSON(500, gin.H{
			"errno":             500,
			"error":             "internal_error",
			"error_description": "Ocurrió un problema para procesar la solicitud" + err.Error(),
		})
		return
	}

	// Aquí puedes usar los datos del Servicio
	nombre := Servicio["Nombre"].(string)
	descripcion := Servicio["Descripcion"].(string)
	unidad := int64(Servicio["Unidad"].(float64))

	if nombre == "" || descripcion == "" || unidad == 0 {
		errcode := models.Error400("Faltan datos.")
		c.JSON(errcode.Code, errcode)

		return
	}

	_, resE := models.ActualizarServisio(int64(idd), nombre, descripcion, unidad)
	if resE != nil {
		c.JSON(resE.Code, resE)

		return
	}

	user, resE := models.ObtenerServisio(int64(idd))
	if resE != nil {
		c.JSON(resE.Code, resE)

		return
	}

	c.JSON(200, user)
}

// Eliminar un Servicio por su ID
func EliminarServicio(c *gin.Context) {
	id := c.Param("id")
	idd, err := strconv.Atoi(id)
	if err != nil {

		errcode := models.Error400("El id no es valido.")
		c.JSON(errcode.Code, errcode)
		return
	}

	_, resE := models.EliminarServisio(int64(idd))
	if resE != nil {
		c.JSON(resE.Code, resE)

		return
	}

	c.JSON(201, gin.H{})
}

func RemplanzarServicio(c *gin.Context) {
	fmt.Println("remplanzarServicio")
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

	var Servicio map[string]interface{}
	if err := c.ShouldBindJSON(&Servicio); err != nil {
		c.JSON(500, gin.H{
			"errno":             500,
			"error":             "internal_error",
			"error_description": "Ocurrió un problema para procesar la solicitud" + err.Error(),
		})
		return
	}

	// Aquí puedes usar los datos del Servicio
	nombre := Servicio["Nombre"].(string)
	descripcion := Servicio["Descripcion"].(string)
	unidad := int64(Servicio["Unidad"].(float64))

	if nombre == "" || descripcion == "" || unidad == 0 {
		errcode := models.Error400("Faltan datos.")
		c.JSON(errcode.Code, errcode)

		return
	}

	if _, resE := models.EliminarServisio(int64(idd)); resE != nil {
		c.JSON(resE.Code, resE)

		return
	}

	_, resE := models.CrearServisio(nombre, descripcion, unidad)
	if resE != nil {
		c.JSON(resE.Code, resE)

		return
	}

	// user, resE := models.ObtenerServicio(*idC)
	// if resE != nil {
	// 	c.JSON(resE.Code, resE)

	// 	return
	// }

	c.JSON(201, gin.H{})
}
