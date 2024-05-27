package models

import (
	"fmt"
	"log"
	"negosioscol/src/db"
	"time"
)

type Producto struct {
	IDProducto  int64     `json:"Id_Producto"`
	Nombre      string    `json:"Nombre"`
	Descripsion string    `json:"Descripsion"`
	Imagen      string    `json:"Imagen"`
	Unidad      int64     `json:"Unidad"`
	Creado      time.Time `json:"Creado"`
	Actualizado time.Time `json:"Actualizado"`
}

func CrearProducto(nombre string, descripcion string, imagen string, unidad int64, negocio int64) (*int64, *ErrorStatusCode) {

	db, err := db.ConnectDB()
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("CALL registrarproducto($1, $2, $3, $4, $5);")
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
func EditarProducto(id int, nombre string, descripcion string, imagen string, unidad int64) *ErrorStatusCode {

	db, err := db.ConnectDB()
	if err != nil {
		return Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("CALL actualizarproducto($1, $2, $3, $4, $5);")
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
func EliminarProducto(id int) *ErrorStatusCode {

	db, err := db.ConnectDB()
	if err != nil {
		return Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT eliminarproducto($1);")
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
		return Error500("no se elimino el producto")
	}
	return nil

}
func ObtenerProducto(id int64) (*Producto, *ErrorStatusCode) {

	db, err := db.ConnectDB()
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("select * FROM obtenerproducto($1);")
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer stmt.Close()

	resul, err := stmt.Query(id)
	if err != nil {
		return nil, Error500(err.Error())
	}

	var producto Producto
	if resul.Next() {
		err = resul.Scan(
			&producto.IDProducto,
			&producto.Nombre,
			&producto.Descripsion,
			&producto.Imagen,
			&producto.Unidad,
			&producto.Creado,
			&producto.Actualizado,
		)
		if err != nil {
			return nil, Error500(err.Error())
		}
	} else {
		return nil, Error404(fmt.Sprintf("No se encontro el producto %d", id))
	}

	return &producto, nil
}
