package models

import "time"

type Negocio struct {
	ID          int
	Nombre      string
	Descripsion string
	Direccion   string
	telefono    string
	correo      string
	imagen      string
	Latitude    float64
	Longitude   float64
	Facebook    string
	Twitter     string
	Instagram   string
	Website     string
	Creado      time.Time
	Actualizado time.Time
}
