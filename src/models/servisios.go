package models

import (
	"fmt"
	"log"
	"negosioscol/src/db"
	"time"
)

type Servisio struct {
	IDServicio  int64     `json:"id_servicio"`
	Nombre      string    `json:"nombre"`
	Descripcion string    `json:"descripcion"`
	Imagen      string    `json:"imagen"`
	Unidad      int64     `json:"unidad"`
	IDNegocio   int64     `json:"id_Negocio"`
	Creado      time.Time `json:"creado"`
	Actualizado time.Time `json:"actualizado"`
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
			&servisio.IDNegocio,
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

func ObtenerUltimosServisio() (*[]Servisio, *ErrorStatusCode) {

	db, err := db.ConnectDB()
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT idservicio, nombre, descripcion, unidad, imagen, idnegocio, creado, actualizado FROM Servisio ORDER BY Creado DESC LIMIT 10;")
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer stmt.Close()

	resul, err := stmt.Query()
	if err != nil {
		return nil, Error500(err.Error())
	}

	var servisio []Servisio
	for resul.Next() {
		var buscar Servisio
		err := resul.Scan(
			&buscar.IDServicio,
			&buscar.Nombre,
			&buscar.Descripcion,
			&buscar.Unidad,
			&buscar.Imagen,
			&buscar.IDNegocio,
			&buscar.Creado,
			&buscar.Actualizado,
		)

		if err != nil {
			return nil, Error500(err.Error())
		}
		servisio = append(servisio, buscar)
	}

	return &servisio, nil
}
