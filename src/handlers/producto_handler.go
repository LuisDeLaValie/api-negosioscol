package handlers

import (
	"fmt"
	"negosioscol/src/models"
	"negosioscol/src/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary Show an account
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} model.HTTPError
// @Router /accounts/{id} [get]
func ObtenerUltimoProducto(c *gin.Context) {
	servi, resE := models.ObtenerUltimoProducto()
	if resE != nil {
		c.JSON(resE.Code, resE)
		return
	}

	c.JSON(200, servi)
}

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

	negocio := struct {
		Nombre      string `form:"Nombre"`
		Descripcion string `form:"Descripcion"`
		Unidad      int64  `form:"Unidad"`
		Negocio     int64  `form:"Negocio"`
		Precio      int64  `form:"Precio"`
	}{}

	if err := c.ShouldBind(&negocio); err != nil {
		errcode := models.Error400("Error al procesar los datos: " + err.Error())
		c.JSON(errcode.Code, errcode)
		return
	}

	imagen, resE := utils.UploadToS3(c, "Imagen")
	if resE != nil {
		c.JSON(resE.Code, resE)
		return
	}

	fmt.Println(negocio)
	lastID, err := models.CrearProducto(negocio.Nombre, negocio.Descripcion, *imagen, negocio.Unidad, negocio.Negocio, negocio.Precio)
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

	negocio := struct {
		Nombre      string `form:"Nombre"`
		Descripcion string `form:"Descripcion"`
		Unidad      int64  `form:"Unidad"`
		Precio      int64  `form:"Precio"`
	}{}

	if err := c.ShouldBind(&negocio); err != nil {
		errcode := models.Error400("Error al procesar los datos: " + err.Error())
		c.JSON(errcode.Code, errcode)
		return
	}

	fmt.Print(negocio)
	var imagen *string
	var resE *models.ErrorStatusCode
	if _, _, err = c.Request.FormFile("Imagen"); err == nil {
		imagen, resE = utils.UploadToS3(c, "Imagen")
		if resE != nil {
			c.JSON(resE.Code, resE)
			return
		}
	}
	fmt.Println(imagen)

	resE = models.EditarProducto(idd, negocio.Nombre, negocio.Descripcion, imagen, negocio.Unidad, negocio.Precio)
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
