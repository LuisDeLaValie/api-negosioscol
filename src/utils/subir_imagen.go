package utils

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"negosioscol/src/models"
	"net/url"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SubirImagen(c *gin.Context, filed string) (*string, *models.ErrorStatusCode) {
	file, err := c.FormFile(filed)
	var imagenPath string

	if err != nil {
		return nil, models.Error400("No se enontro el archivo %s\n%s", filed, err.Error())
	}

	// Antes de guardar la imagen
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		os.Mkdir("uploads", os.ModePerm) // Crea el directorio si no existe
	}

	// Guardar la imagen en un directorio específico

	// Genera un nombre único basado en la marca de tiempo
	uniqueFilename := time.Now().Format("20060102150405") + "_" + file.Filename
	imagenPath = "uploads/" + uniqueFilename
	if err := c.SaveUploadedFile(file, imagenPath); err != nil {
		errcode := models.Error500("Error al guardar la imagen: " + err.Error())
		return nil, errcode
	}

	return &imagenPath, nil

}

// UploadToS3 sube un archivo a S3 y devuelve la URL pública
func UploadToS3(c *gin.Context, filed string) (*string, *models.ErrorStatusCode) {
	// Obtener el archivo desde el formulario
	file, header, err := c.Request.FormFile(filed)
	if err != nil {
		return nil, models.Error400("No se encontró el archivo %s\n%s", filed, err.Error())
	}
	defer file.Close()

	// Convertir el archivo en un buffer
	buffer := bytes.NewBuffer(nil)
	_, err = buffer.ReadFrom(file)
	if err != nil {
		return nil, models.Error500("Error leyendo el archivo: %s", err)
	}

	// Cargar variables de entorno
	err = godotenv.Load()
	if err != nil {
		log.Printf("Advertencia: no se pudo cargar el archivo .env, usando variables de entorno del sistema.")
	}

	bucketName := os.Getenv("S3_BUCKET_NAME")
	if bucketName == "" {
		return nil, models.Error500("El nombre del bucket (S3_BUCKET_NAME) no está configurado")
	}

	// Cargar configuración de AWS
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, models.Error500("Error cargando configuración de AWS: %s", err)
	}

	// Generar un nombre único para el archivo
	namefile := fmt.Sprintf("%d-%s", time.Now().Unix(), header.Filename)
	key := fmt.Sprintf("uploads/%s", namefile)

	// Crear cliente S3
	client := s3.NewFromConfig(cfg)

	// Subir el archivo a S3
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:        aws.String(bucketName),
		Key:           aws.String(key),
		Body:          bytes.NewReader(buffer.Bytes()),
		ContentLength: aws.Int64(header.Size),
		ContentType:   aws.String(header.Header.Get("Content-Type")), // Usa el tipo MIME del archivo subido
		// ACL:           types.ObjectCannedACLPublicRead,
	})
	if err != nil {
		return nil, models.Error500("Error subiendo archivo a S3: %s", err)
	}

	// Generar la URL pública del archivo
	// fileURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, key)
	encodedName := url.QueryEscape(namefile) // Codifica el nombre
	fileURL := fmt.Sprintf("https://%s.s3.us-east-2.amazonaws.com/uploads/%s", bucketName, encodedName)
	return &fileURL, nil
}
