package models

import (
	"fmt"
	"log"
	"negosioscol/src/db"
	"time"
)

type Producto struct {
	IDProducto  int64  `json:"id_Producto"`
	IDNegocio   int64  `json:"id_Negocio"`
	Nombre      string `json:"nombre"`
	Descripsion string `json:"descripsion"`
	Imagen      string `json:"imagen"`
	Unidad      int64  `json:"unidad"`
	Precio      int64  `json:"precio"`
	// precio      int64     `json:"precio"`
	Creado      time.Time `json:"creado"`
	Actualizado time.Time `json:"actualizado"`
}

func CrearProducto(nombre string, descripcion string, imagen string, unidad int64, negocio int64, precio int64) (*int64, *ErrorStatusCode) {

	db, err := db.ConnectDB()
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("CALL registrarproducto($1, $2, $3, $4, $5, $6);")
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(nombre, descripcion, imagen, unidad, negocio, precio)
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
func EditarProducto(id int, nombre string, descripcion string, imagen *string, unidad int64, precio int64) *ErrorStatusCode {

	db, err := db.ConnectDB()
	if err != nil {
		return Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("CALL actualizarproducto($1, $2, $3, $4, $5, $6);")
	if err != nil {
		return Error500("Error parse SQL: " + err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, nombre, descripcion, imagen, unidad, precio)
	if err != nil {
		return Error500("Error exejutar SQL: " + err.Error())
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
		return nil, Error500(err.Error() + "Consulta")
	}
	defer stmt.Close()

	resul, err := stmt.Query(id)
	if err != nil {
		return nil, Error500(err.Error() + "query")
	}

	var producto Producto
	if resul.Next() {
		err = resul.Scan(
			&producto.IDProducto,
			&producto.Nombre,
			&producto.Descripsion,
			&producto.Imagen,
			&producto.Unidad,
			&producto.IDNegocio,
			&producto.Precio,
			&producto.Creado,
			&producto.Actualizado,
		)
		if err != nil {
			return nil, Error500(err.Error() + "parce")
		}
	} else {
		return nil, Error404(fmt.Sprintf("No se encontro el producto %d", id))
	}

	return &producto, nil
}
func ObtenerUltimoProducto() (*[]Producto, *ErrorStatusCode) {

	db, err := db.ConnectDB()
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT idproducto, nombre, descripsion, imagen, unidad, idnegocio, precio, creado, actualizado FROM Producto ORDER BY Creado DESC LIMIT 10;")
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer stmt.Close()

	resul, err := stmt.Query()
	if err != nil {
		return nil, Error500(err.Error())
	}

	var producto []Producto
	for resul.Next() {
		var buscar Producto
		err := resul.Scan(
			&buscar.IDProducto,
			&buscar.Nombre,
			&buscar.Descripsion,
			&buscar.Imagen,
			&buscar.Unidad,
			&buscar.IDNegocio,
			&buscar.Precio,
			&buscar.Creado,
			&buscar.Actualizado,
		)

		if err != nil {
			return nil, Error500(err.Error())
		}
		producto = append(producto, buscar)
	}

	return &producto, nil
}
