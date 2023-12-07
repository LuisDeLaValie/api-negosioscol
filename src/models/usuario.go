package models

import (
	"fmt"
	"negosioscol/src/db"
	"time"
)

type Usuario struct {
	IDUsuario   int64     `json:"Id_Usuario"`
	Nombre      string    `json:"Nombre"`
	Apellidos   string    `json:"Apellidos"`
	Creado      string    `json:"Creado"`
	Actualizado time.Time `json:"Actualizado"`
	Cumpleanos  time.Time `json:"Cumpleanos"`
	Imagen      string    `json:"Imagen"`
}

func CrearUsuario(nombres string, apellidos string, cumple string, imagen string) (*int64, *ErrorStatusCode) {

	db, err := db.ConnectDB()
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("CALL RegistrarUsuario($1, $2, $3, $4)")
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(nombres, apellidos, cumple, imagen)
	if err != nil {
		return nil, Error500(err.Error())
	}

	return nil, nil

}
func EditarUsuario(id int, nombres string, apellidos string, cumple string, imagen string) *ErrorStatusCode {

	db, err := db.ConnectDB()
	if err != nil {
		return Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("CALL ActualizarUsuario($1, $2, $3, $4, $5)")
	if err != nil {
		return Error500(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, nombres, apellidos, cumple, imagen)
	if err != nil {
		return Error500(err.Error())
	}

	return nil

}
func EliminarUsuario(id int) *ErrorStatusCode {

	db, err := db.ConnectDB()
	if err != nil {
		return Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT EliminarUsuario($1);")
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
		return Error500("no se elimino el usuario")
	}
	return nil

}
func ObtenerUsuario(id int64) (*Usuario, *ErrorStatusCode) {

	db, err := db.ConnectDB()
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM ObtenerUsuario($1);")
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer stmt.Close()

	resul, err := stmt.Query(id)
	if err != nil {
		return nil, Error500(err.Error())
	}

	var usuario Usuario
	if resul.Next() {
		err = resul.Scan(
			&usuario.IDUsuario,
			&usuario.Nombre,
			&usuario.Apellidos,
			&usuario.Creado,
			&usuario.Actualizado,
			&usuario.Cumpleanos,
			&usuario.Imagen,
		)
		if err != nil {
			return nil, Error500(err.Error())
		}
	} else {
		return nil, Error404(fmt.Sprintf("No se encontro el usuario %d", id))
	}

	return &usuario, nil
}
