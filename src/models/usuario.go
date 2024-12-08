package models

import (
	"fmt"
	"log"
	"negosioscol/src/db"
	"time"
)

type Usuario struct {
	IDUsuario   int64     `json:"id_Usuario"`
	Nombre      string    `json:"nombre"`
	Apellidos   string    `json:"apellidos"`
	Correo      string    `json:"correo"`
	Password    string    `json:"password"`
	Creado      string    `json:"creado"`
	Actualizado time.Time `json:"actualizado"`
	Cumpleanos  time.Time `json:"cumpleanos"`
	Imagen      string    `json:"imagen"`
}

func CrearUsuario(nombres string, apellidos string, correo string, password string, cumple string, imagen string) (*int64, *ErrorStatusCode) {

	db, err := db.ConnectDB()
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("CALL registrarusuario($1, $2, $3, $4, $5, $6);")
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(nombres, apellidos, correo, password, cumple, imagen)
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
func EditarUsuario(id int, nombres string, apellidos string, correo string, password string, cumple string, imagen string) *ErrorStatusCode {

	db, err := db.ConnectDB()
	if err != nil {
		return Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("CALL actualizarusuario($1, $2, $3, $4, $5, $6, $7);")
	if err != nil {
		return Error500(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, nombres, apellidos, correo, password, cumple, imagen)
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

	stmt, err := db.Prepare("SELECT eliminarusuario($1);")
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

	stmt, err := db.Prepare("select * FROM obtenerusuario($1);")
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
			&usuario.Correo,
			&usuario.Password,
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

func Login(use string, pass string) (*Usuario, *ErrorStatusCode) {

	db, err := db.ConnectDB()
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM usuario u WHERE u.correo = $1 AND u.password = $2")
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer stmt.Close()

	resul, err := stmt.Query(use, pass)
	if err != nil {
		return nil, Error500(err.Error())
	}

	var usuario Usuario
	if resul.Next() {
		err = resul.Scan(
			&usuario.IDUsuario,
			&usuario.Nombre,
			&usuario.Apellidos,
			&usuario.Correo,
			&usuario.Password,
			&usuario.Creado,
			&usuario.Actualizado,
			&usuario.Cumpleanos,
			&usuario.Imagen,
		)
		if err != nil {
			return nil, Error500(err.Error())
		}
	} else {
		return nil, Error401("Correo electronico o contraseña incorectos")
	}

	return &usuario, nil
}
