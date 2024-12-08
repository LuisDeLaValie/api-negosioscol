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

	var negocio map[string]interface{}

	if err := c.ShouldBindJSON(&negocio); err != nil {
		errcode := models.Error500("Ocurrió 000 un problema para procesar la solicitud: " + err.Error())
		c.JSON(errcode.Code, errcode)
		return
	}

	// Validación de datos obligatorios
	nombre, okNombre := negocio["nombre"].(string)
	password, okPassword := negocio["password"].(string)
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

	if !okNombre || !okPassword || !okDescripcion || !okDireccion || !okTelefono || !okCorreo || !okImagen || !okLatitude || !okLongitude {
		errcode := models.Error400("Faltan datos obligatorios o tienen el formato incorrecto.")
		c.JSON(errcode.Code, errcode)
		return
	}

	// Llamada al modelo
	lastID, err := models.CrearNegocio(nombre, password, descripcion, direccion, telefono, correo, imagen, latitude, longitude, &facebook, &twitter, &instagram, &website)
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
		Nombre      string `json:"nombre"`
		Password    string `json:"password,omitempty"`
		Descripcion string `json:"descripcion"`
		Direccion   string `json:"direccion"`
		Telefono    string `json:"telefono"`
		Correo      string `json:"correo"`
		// Imagen      string  `json:"imagen"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Facebook  string  `json:"facebook,omitempty"`
		Twitter   string  `json:"twitter,omitempty"`
		Instagram string  `json:"instagram,omitempty"`
		Website   string  `json:"website,omitempty"`
	}{}

	if err := c.ShouldBindJSON(&negocio); err != nil {
		errcode := models.Error400(err.Error())
		c.JSON(errcode.Code, errcode)
		return
	}

	imagen, resE := utils.SubirImagen(c, "imagen")
	if resE != nil {
		c.JSON(resE.Code, resE)

		return
	}

	resE = models.EditarNegocio(idd, negocio.Nombre, negocio.Password, negocio.Direccion, negocio.Direccion, negocio.Telefono, negocio.Correo, *imagen, negocio.Latitude, negocio.Longitude, &negocio.Facebook, &negocio.Twitter, &negocio.Instagram, &negocio.Website)
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
