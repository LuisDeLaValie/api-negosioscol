package utils

import (
	"negosioscol/src/models"

	"github.com/gin-gonic/gin"
)

func SubirImagen(c *gin.Context, filed string) (*string, *models.ErrorStatusCode) {
	file, err := c.FormFile(filed)
	var imagenPath string

	if err != nil {
		return nil, models.Error400("No se enontro el archivo %s\n%s", filed, err.Error())
	}

	// Guardar la imagen en un directorio específico
	imagenPath = "./uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, imagenPath); err != nil {
		errcode := models.Error500("Error al guardar la imagen: " + err.Error())
		return nil, errcode
	}

	return &imagenPath, nil

}
