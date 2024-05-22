package models

import (
	"database/sql"
	"fmt"
	"negosioscol/src/db"
	"time"
)

type Servisio struct {
	IDProducto  int64     `json:"Id_Producto"`
	Nombre      string    `json:"Nombre"`
	Descripcion string    `json:"Descripcion"`
	Unidad      int64     `json:"Unidad"`
	Creado      time.Time `json:"Creado"`
	Actualizado time.Time `json:"Actualizado"`
}

func CrearServisio(nombre string, descripcion string, unidad int64) (*int64, *ErrorStatusCode) {
	fmt.Println("crearServis")
	db, err := db.ConnectDB()
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("CALL InsertarServisio($1, $2, $3)")
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(nombre, descripcion, unidad)
	if err != nil {
		return nil, Error500(err.Error())
	}

	return nil, nil
}

func ActualizarServisio(idProducto int64, nombre string, descripcion string, unidad int64) (*int64, *ErrorStatusCode) {
	db, err := db.ConnectDB()
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("CALL ActualizarServisio($1, $2, $3, $4)")
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(idProducto, nombre, descripcion, unidad)
	if err != nil {
		return nil, Error500(err.Error())
	}

	return nil, nil
}

func EliminarServisio(idProducto int64) (*int64, *ErrorStatusCode) {
	fmt.Println("EliminarServisio")
	db, err := db.ConnectDB()
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT EliminarServisio($1)")
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(idProducto)
	if err != nil {
		return nil, Error500(err.Error())
	}

	return nil, nil
}

func ListarServisios() ([]Servisio, *ErrorStatusCode) {
	fmt.Println("ListarServisios")
	db, err := db.ConnectDB()
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM ListarServisio()")
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer rows.Close()

	var servisios []Servisio
	for rows.Next() {
		var s Servisio
		if err := rows.Scan(&s.IDProducto, &s.Nombre, &s.Descripcion, &s.Unidad, &s.Creado, &s.Actualizado); err != nil {
			return nil, Error500(err.Error())
		}
		servisios = append(servisios, s)
	}

	if err := rows.Err(); err != nil {
		return nil, Error500(err.Error())
	}

	return servisios, nil
}

func ObtenerServisio(idProducto int64) (*Servisio, *ErrorStatusCode) {
	db, err := db.ConnectDB()
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer db.Close()

	row := db.QueryRow("SELECT * FROM ObtenerServisio($1)", idProducto)
	var s Servisio
	err = row.Scan(&s.IDProducto, &s.Nombre, &s.Descripcion, &s.Unidad, &s.Creado, &s.Actualizado)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, Error404("Registro no encontrado")
		}
		return nil, Error500(err.Error())
	}

	return &s, nil
}
