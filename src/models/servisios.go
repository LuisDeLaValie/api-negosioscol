package models

import (
	"fmt"
	"log"
	"negosioscol/src/db"
	"time"
)

type Servisio struct {
	IDServicio  int64     `json:"Id_servicio"`
	Nombre      string    `json:"Nombre"`
	Descripcion string    `json:"Descripcion"`
	Imagen      string    `json:"Imagen"`
	Unidad      int64     `json:"Unidad"`
	Creado      time.Time `json:"Creado"`
	Actualizado time.Time `json:"Actualizado"`
}

func CrearServisio(nombre string, descripcion string, imagen string, unidad int64, negocio int64) (*int64, *ErrorStatusCode) {

	db, err := db.ConnectDB()
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("CALL registrarservisio($1, $2, $3, $4,$5);")
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(nombre, descripcion, imagen, unidad, negocio)
	if err != nil {
		return nil, Error500(err.Error())
	}

	// Obtener el último ID insertado
	var lastID int64
	err = db.QueryRow("SELECT lastval()").Scan(&lastID)
	if err != nil {
		log.Fatal(err)
	}

	return &lastID, nil

}
func EditarServisio(id int, nombre string, descripcion string, imagen string, unidad int64) *ErrorStatusCode {

	db, err := db.ConnectDB()
	if err != nil {
		return Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("CALL actualizarservisio($1, $2, $3, $4, $5);")
	if err != nil {
		return Error500(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, nombre, descripcion, imagen, unidad)
	if err != nil {
		return Error500(err.Error())
	}

	return nil

}
func EliminarServisio(id int) *ErrorStatusCode {

	db, err := db.ConnectDB()
	if err != nil {
		return Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT eliminarservisio($1);")
	if err != nil {
		return Error500(err.Error())
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return Error500(err.Error())
	}

	delet, err := result.RowsAffected()
	if err != nil {
		return Error500(err.Error())
	} else if delet == 0 {
		return Error500("no se elimino el servisio")
	}
	return nil

}
func ObtenerServisio(id int64) (*Servisio, *ErrorStatusCode) {

	db, err := db.ConnectDB()
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("select * FROM obtenerservisio($1);")
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer stmt.Close()

	resul, err := stmt.Query(id)
	if err != nil {
		return nil, Error500(err.Error())
	}

	var servisio Servisio
	if resul.Next() {
		err = resul.Scan(
			&servisio.IDServicio,
			&servisio.Nombre,
			&servisio.Descripcion,
			&servisio.Imagen,
			&servisio.Unidad,
			&servisio.Creado,
			&servisio.Actualizado,
		)
		if err != nil {
			return nil, Error500(err.Error())
		}
	} else {
		return nil, Error404(fmt.Sprintf("No se encontro el servisio %d", id))
	}

	return &servisio, nil
}
