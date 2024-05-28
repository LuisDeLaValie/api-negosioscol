package handlers

import (
	"negosioscol/src/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Obtener un servisio por su ID
func GetNegocioPorID(c *gin.Context) {
	id := c.Param("id")
	idd, err := strconv.Atoi(id)
	if err != nil {
		errcode := models.Error400("El id no es valido.")
		c.JSON(errcode.Code, errcode)
		return
	}

	servi, resE := models.ObtenerNegocio(int64(idd))
	if resE != nil {
		c.JSON(resE.Code, resE)
		return
	}

	c.JSON(200, servi)
}

// Crear un nuevo servisio
func CrearNegocio(c *gin.Context) {

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

	negocio := struct {
		Nombre      string  `json:"Nombre"`
		Descripcion string  `json:"Descripcion"`
		Direccion   string  `json:"Direccion"`
		Telefono    string  `json:"Telefono"`
		Correo      string  `json:"Correo"`
		Imagen      string  `json:"Imagen"`
		Latitude    float64 `json:"Latitude"`
		Longitude   float64 `json:"Longitude"`
		Facebook    string  `json:"Facebook,omitempty"`
		Twitter     string  `json:"Twitter,omitempty"`
		Instagram   string  `json:"Instagram,omitempty"`
		Website     string  `json:"Website,omitempty"`
	}{}

	if err := c.ShouldBind(&negocio); err != nil {
		errcode := models.Error400(err.Error())
		c.JSON(errcode.Code, errcode)
		return
	}

	lastID, err := models.CrearNegocio(negocio.Nombre, negocio.Direccion, negocio.Direccion, negocio.Telefono, negocio.Correo, negocio.Imagen, negocio.Latitude, negocio.Longitude, &negocio.Facebook, &negocio.Twitter, &negocio.Instagram, &negocio.Website)
	if err != nil {
		c.JSON(err.Code, err)

		return
	}

	servi, resE := models.ObtenerNegocio(*lastID)
	if resE != nil {
		c.JSON(resE.Code, resE)
		return
	}

	c.JSON(200, servi)
}

// Actualizar un servisio existente por su ID
func ActualizarNegocio(c *gin.Context) {
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

	negocio := struct {
		Nombre      string  `json:"Nombre"`
		Descripcion string  `json:"Descripcion"`
		Direccion   string  `json:"Direccion"`
		Telefono    string  `json:"Telefono"`
		Correo      string  `json:"Correo"`
		Imagen      string  `json:"Imagen"`
		Latitude    float64 `json:"Latitude"`
		Longitude   float64 `json:"Longitude"`
		Facebook    string  `json:"Facebook,omitempty"`
		Twitter     string  `json:"Twitter,omitempty"`
		Instagram   string  `json:"Instagram,omitempty"`
		Website     string  `json:"Website,omitempty"`
	}{}

	if err := c.ShouldBind(&negocio); err != nil {
		errcode := models.Error400(err.Error())
		c.JSON(errcode.Code, errcode)
		return
	}

	resE := models.EditarNegocio(idd, negocio.Nombre, negocio.Direccion, negocio.Direccion, negocio.Telefono, negocio.Correo, negocio.Imagen, negocio.Latitude, negocio.Longitude, &negocio.Facebook, &negocio.Twitter, &negocio.Instagram, &negocio.Website)
	if resE != nil {
		c.JSON(resE.Code, resE)

		return
	}

	servi, resE := models.ObtenerNegocio(int64(idd))
	if resE != nil {
		c.JSON(resE.Code, resE)
		return
	}

	c.JSON(200, servi)
}

// Eliminar un servisio por su ID
func EliminarNegocio(c *gin.Context) {
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

	resE := models.EliminarNegocio(idd)
	if resE != nil {
		c.JSON(resE.Code, resE)

		return
	}

	c.JSON(200, gin.H{
		"message": "EliminarNegocio:" + id,
	})

}

// Obtener un servisio por su ID
func GetServicioNegocioPorID(c *gin.Context) {
	id := c.Param("id")
	idd, err := strconv.Atoi(id)
	if err != nil {
		errcode := models.Error400("El id no es valido.")
		c.JSON(errcode.Code, errcode)
		return
	}

	servi, resE := models.ObtenerServicioNegocio(int64(idd))
	if resE != nil {
		c.JSON(resE.Code, resE)
		return
	}

	c.JSON(200, servi)
}

func GetProductoNegocioPorID(c *gin.Context) {
	id := c.Param("id")
	idd, err := strconv.Atoi(id)
	if err != nil {
		errcode := models.Error400("El id no es valido.")
		c.JSON(errcode.Code, errcode)
		return
	}

	servi, resE := models.ObtenerProductoNegocio(int64(idd))
	if resE != nil {
		c.JSON(resE.Code, resE)
		return
	}

	c.JSON(200, servi)
}

// func RemplanzarNegocio(c *gin.Context) {
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

// 	if resE := models.EliminarNegocio(idd); resE != nil {
// 		c.JSON(resE.Code, resE)

// 		return
// 	}

// 	lastID, resE := models.CrearNegocio(nombre, descripcion, imagen, int64(int64(unidad)))
// 	if resE != nil {
// 		c.JSON(resE.Code, resE)

// 		return
// 	}

// 	servi, resE := models.ObtenerNegocio(*lastID)
// 	if resE != nil {
// 		c.JSON(resE.Code, resE)
// 		return
// 	}

// 	c.JSON(200, servi)
// }
