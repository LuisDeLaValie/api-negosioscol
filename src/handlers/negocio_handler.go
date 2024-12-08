package handlers

import (
	"fmt"
	"negosioscol/src/models"
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
func ObtenerUltimosNegocios(c *gin.Context) {

	servi, resE := models.ObtenerUltimosNegocios()
	if resE != nil {
		c.JSON(resE.Code, resE)
		return
	}

	c.JSON(200, servi)

}

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

	var negocio map[string]interface{}

	if err := c.ShouldBindJSON(&negocio); err != nil {
		errcode := models.Error500("Ocurrió 000 un problema para procesar la solicitud: " + err.Error())
		c.JSON(errcode.Code, errcode)
		return
	}

	// Validación de datos obligatorios
	nombre, okNombre := negocio["nombre"].(string)
	descripcion, okDescripcion := negocio["descripcion"].(string)
	direccion, okDireccion := negocio["direccion"].(string)
	telefono, okTelefono := negocio["telefono"].(string)
	correo, okCorreo := negocio["correo"].(string)
	imagen, okImagen := negocio["imagen"].(string)
	latitude, okLatitude := negocio["latitude"].(float64)
	longitude, okLongitude := negocio["longitude"].(float64)

	// Validación de redes sociales (opcionales)
	facebook, _ := negocio["Facebook"].(string)
	twitter, _ := negocio["Twitter"].(string)
	instagram, _ := negocio["Instagram"].(string)
	website, _ := negocio["Website"].(string)

	fmt.Println(negocio)

	fmt.Printf("nombre: %s - %t\n", nombre, okNombre)
	fmt.Printf("descripcion: %s - %t\n", descripcion, okDescripcion)
	fmt.Printf("direccion: %s - %t\n", direccion, okDireccion)
	fmt.Printf("telefono: %s - %t\n", telefono, okTelefono)
	fmt.Printf("correo: %s - %t\n", correo, okCorreo)
	fmt.Printf("imagen: %s - %t\n", imagen, okImagen)
	fmt.Printf("latitude: %f - %t\n", latitude, okLatitude)
	fmt.Printf("longitude: %f - %t\n", longitude, okLongitude)

	if !okNombre || !okDescripcion || !okDireccion || !okTelefono || !okCorreo || !okImagen || !okLatitude || !okLongitude {
		errcode := models.Error400("Faltan datos obligatorios o tienen el formato incorrecto.")
		c.JSON(errcode.Code, errcode)
		return
	}

	// Llamada al modelo
	lastID, err := models.CrearNegocio(nombre, descripcion, direccion, telefono, correo, imagen, latitude, longitude, &facebook, &twitter, &instagram, &website)
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
