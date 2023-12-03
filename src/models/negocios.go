package models

import "time"



type Restaurant struct {
    ID          int
    Nombre      string
    Descripsion string
    Direccion   string
    telefono    string
    correo      string
    Latitude    float64
    Longitude   float64
    Facebook    string
    Twitter     string
    Instagram   string
    Website     string
    Creado      time.Time
    Actualizado time.Time
}
