package handlers

import (
	"negosioscol/src/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Obtener un servisio por su ID
func GetProductoPorID(c *gin.Context) {
	id := c.Param("id")
	idd, err := strconv.Atoi(id)
	if err != nil {
		errcode := models.Error400("El id no es valido.")
		c.JSON(errcode.Code, errcode)
		return
	}

	servi, resE := models.ObtenerProducto(int64(idd))
	if resE != nil {
		c.JSON(resE.Code, resE)
		return
	}

	c.JSON(200, servi)
}

// Crear un nuevo servisio
func CrearProducto(c *gin.Context) {

	defer func() {
		err := recover()
		if err != nil {
			errcode := models.Error500("Ocurrió un problema para procesar la solicitud:\n %v", err)
			c.JSON(errcode.Code, errcode)
		}
	}()

	var servisio map[string]interface{}
	if err := c.ShouldBindJSON(&servisio); err != nil {
		errcode := models.Error500("Ocurrió un problema para procesar la solicitud" + err.Error())
		c.JSON(errcode.Code, errcode)

		return
	}

	// Aquí puedes usar los datos del servisio
	nombre := servisio["Nombre"].(string)
	descripcion := servisio["Descripcion"].(string)
	imagen := servisio["Imagen"].(string)
	unidad := servisio["Unidad"].(float64)
	negocio := servisio["Negocio"].(float64)

	if nombre == "" || descripcion == "" || imagen == "" || unidad == 0 || negocio == 0 {
		errcode := models.Error400("Faltan datos.")
		c.JSON(errcode.Code, errcode)

		return
	}

	lastID, err := models.CrearProducto(nombre, descripcion, imagen, int64(unidad), int64(negocio))
	if err != nil {
		c.JSON(err.Code, err)

		return
	}

	servi, resE := models.ObtenerProducto(*lastID)
	if resE != nil {
		c.JSON(resE.Code, resE)
		return
	}

	c.JSON(200, servi)
}

// Actualizar un servisio existente por su ID
func ActualizarProducto(c *gin.Context) {
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

	var servisio map[string]interface{}
	if err := c.ShouldBindJSON(&servisio); err != nil {
		c.JSON(500, models.Error500("Ocurrió un problema para procesar la solicitud"+err.Error()))
		return
	}

	// Aquí puedes usar los datos del servisio
	nombre := servisio["Nombre"].(string)
	descripcion := servisio["Descripcion"].(string)
	imagen := servisio["Imagen"].(string)
	unidad := servisio["Unidad"].(float64)

	if nombre == "" || descripcion == "" || imagen == "" || unidad == 0 {
		errcode := models.Error400("Faltan datos.")
		c.JSON(errcode.Code, errcode)

		return
	}

	resE := models.EditarProducto(idd, nombre, descripcion, imagen, int64(int64(unidad)))
	if resE != nil {
		c.JSON(resE.Code, resE)

		return
	}

	servi, resE := models.ObtenerProducto(int64(idd))
	if resE != nil {
		c.JSON(resE.Code, resE)
		return
	}

	c.JSON(200, servi)
}

// Eliminar un servisio por su ID
func EliminarProducto(c *gin.Context) {
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

	resE := models.EliminarProducto(idd)
	if resE != nil {
		c.JSON(resE.Code, resE)

		return
	}

	c.JSON(200, gin.H{
		"message": "EliminarProducto:" + id,
	})

}

// func RemplanzarProducto(c *gin.Context) {
// 	defer func() {
// 		err := recover()
// 		if err != nil {
// 			errcode := models.Error500("Ocurrió un problema para procesar la solicitud:\n %v", err)
// 			c.JSON(errcode.Code, errcode)
// 		}
// 	}()

// 	id := c.Param("id")
// 	idd, err := strconv.Atoi(id)
// 	if err != nil {
// 		c.JSON(400, models.Error400("El id no es valido."))
// 		return
// 	}

// 	var servisio map[string]interface{}
// 	if err := c.ShouldBindJSON(&servisio); err != nil {
// 		c.JSON(500, models.Error500("Ocurrió un problema para procesar la solicitud"+err.Error()))
// 		return
// 	}

// 	// Aquí puedes usar los datos del servisio
// 	nombre := servisio["Nombre"].(string)
// 	descripcion := servisio["Descripcion"].(string)
// 	imagen := servisio["Imagen"].(string)
// 	unidad := servisio["Unidad"].(float64)

// 	if nombre == "" || descripcion == "" || imagen == "" || unidad == 0 {
// 		errcode := models.Error400("Faltan datos.")
// 		c.JSON(errcode.Code, errcode)

// 		return
// 	}

// 	if resE := models.EliminarProducto(idd); resE != nil {
// 		c.JSON(resE.Code, resE)

// 		return
// 	}

// 	lastID, resE := models.CrearProducto(nombre, descripcion, imagen, int64(int64(unidad)))
// 	if resE != nil {
// 		c.JSON(resE.Code, resE)

// 		return
// 	}

// 	servi, resE := models.ObtenerProducto(*lastID)
// 	if resE != nil {
// 		c.JSON(resE.Code, resE)
// 		return
// 	}

// 	c.JSON(200, servi)
// }
