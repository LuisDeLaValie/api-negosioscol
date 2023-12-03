package models

import "time"


type Usuario struct {
    ID          int
    Nombre      string
    Apellidos   string
    Creado      time.Time
    Actualizado time.Time
    Cumpleanos  time.Time
    Imagen      string
}