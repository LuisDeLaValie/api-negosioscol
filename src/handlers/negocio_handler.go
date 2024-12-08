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

	negocio := struct {
		Nombre      string `form:"nombre"`
		Password    string `form:"password,omitempty"`
		Descripcion string `form:"descripcion"`
		Direccion   string `form:"direccion"`
		Telefono    string `form:"telefono"`
		Correo      string `form:"correo"`
		// Imagen      string  `form:"imagen"`
		Latitude  float64 `form:"latitude"`
		Longitude float64 `form:"longitude"`
		Facebook  string  `form:"facebook,omitempty"`
		Twitter   string  `form:"twitter,omitempty"`
		Instagram string  `form:"instagram,omitempty"`
		Website   string  `form:"website,omitempty"`
	}{}

	if err := c.ShouldBind(&negocio); err != nil {
		errcode := models.Error400("Error al procesar los datos: " + err.Error())
		c.JSON(errcode.Code, errcode)
		return
	}

	fmt.Println(negocio)
	imagen, resE := utils.UploadToS3(c, "imagen")
	if resE != nil {
		c.JSON(resE.Code, resE)
		return
	}

	// Llamada al modelo
	lastID, err := models.CrearNegocio(negocio.Nombre, negocio.Password, negocio.Descripcion, negocio.Direccion, negocio.Telefono, negocio.Correo, *imagen, negocio.Latitude, negocio.Longitude, &negocio.Facebook, &negocio.Twitter, &negocio.Instagram, &negocio.Website)
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

	negocio := struct {
		Nombre      string `form:"nombre"`
		Password    string `form:"password,omitempty"`
		Descripcion string `form:"descripcion"`
		Direccion   string `form:"direccion"`
		Telefono    string `form:"telefono"`
		Correo      string `form:"correo"`
		// Imagen      string  `form:"imagen"`
		Latitude  float64 `form:"latitude"`
		Longitude float64 `form:"longitude"`
		Facebook  string  `form:"facebook,omitempty"`
		Twitter   string  `form:"twitter,omitempty"`
		Instagram string  `form:"instagram,omitempty"`
		Website   string  `form:"website,omitempty"`
	}{}

	if err := c.ShouldBind(&negocio); err != nil {
		errcode := models.Error400("Error al procesar los datos: " + err.Error())
		c.JSON(errcode.Code, errcode)
		return
	}

	var imagen *string
	var resE *models.ErrorStatusCode
	if _, _, err = c.Request.FormFile("imagen"); err == nil {
		imagen, resE = utils.UploadToS3(c, "imagen")
		if resE != nil {
			c.JSON(resE.Code, resE)
			return
		}
	}

	resE = models.EditarNegocio(idd, negocio.Nombre, negocio.Password, negocio.Direccion, negocio.Direccion, negocio.Telefono, negocio.Correo, imagen, negocio.Latitude, negocio.Longitude, &negocio.Facebook, &negocio.Twitter, &negocio.Instagram, &negocio.Website)
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

func LonginNegocio(c *gin.Context) {
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
	correo := usuario["Correo"].(string)
	password := usuario["Password"].(string)

	if correo == "" || password == "" {
		errcode := models.Error400("Faltan datos.")
		c.JSON(errcode.Code, errcode)

		return
	}

	user, resE := models.LonginNegocio(correo, password)
	if resE != nil {
		c.JSON(resE.Code, resE)
		return
	}

	c.JSON(200, user)
}
