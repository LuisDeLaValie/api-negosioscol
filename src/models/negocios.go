package models

import (
	"fmt"
	"log"
	"negosioscol/src/db"
	"time"
)

type Negocio struct {
	IDNegocio   int64     `json:"Id_Negocio"`
	Nombre      string    `json:"Nombre"`
	Descripsion string    `json:"Descripsion"`
	Direccion   string    `json:"Direccion"`
	Telefono    string    `json:"Telefono"`
	Correo      string    `json:"Correo"`
	Imagen      string    `json:"Imagen"`
	Latitude    float64   `json:"Latitude"`
	Longitude   float64   `json:"Longitude"`
	Facebook    string    `json:"Facebook,omitempty"`
	Twitter     string    `json:"Twitter,omitempty"`
	Instagram   string    `json:"Instagram,omitempty"`
	Website     string    `json:"Website,omitempty"`
	Creado      time.Time `json:"Creado"`
	Actualizado time.Time `json:"Actualizado"`
}

func CrearNegocio(nombre string, descripcion string, direccion string, telefono string, correo string, imagen string, latitud float64, longitud float64, facebook *string, twitter *string, instagra *string, website *string) (*int64, *ErrorStatusCode) {

	db, err := db.ConnectDB()
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("CALL registrarnegocio($1, $2, $3, $4, $5, $6, $7 $8, $9 $20, $11, $12);")
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(nombre, descripcion, direccion, telefono, correo, imagen, latitud, longitud, facebook, twitter, instagra, website)
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
func EditarNegocio(id int, nombre string, descripcion string, direccion string, telefono string, correo string, imagen string, latitud float64, longitud float64, facebook *string, twitter *string, instagra *string, website *string) *ErrorStatusCode {

	db, err := db.ConnectDB()
	if err != nil {
		return Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("CALL actualizarnegocio($1, $2, $3, $4, $5, $6, $7 $8, $9 $20, $11, $12, $13);")
	if err != nil {
		return Error500(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, nombre, descripcion, direccion, telefono, correo, imagen, latitud, longitud, facebook, twitter, instagra, website)
	if err != nil {
		return Error500(err.Error())
	}

	return nil

}
func EliminarNegocio(id int) *ErrorStatusCode {

	db, err := db.ConnectDB()
	if err != nil {
		return Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT eliminarnegocio($1);")
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
		return Error500("no se elimino el negocio")
	}
	return nil

}
func ObtenerNegocio(id int64) (*Negocio, *ErrorStatusCode) {

	db, err := db.ConnectDB()
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("select * FROM obtenernegocio($1);")
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer stmt.Close()

	resul, err := stmt.Query(id)
	if err != nil {
		return nil, Error500(err.Error())
	}

	var negocio Negocio
	if resul.Next() {
		err = resul.Scan(
			&negocio.IDNegocio,
			&negocio.Nombre,
			&negocio.Descripsion,
			&negocio.Direccion,
			&negocio.Telefono,
			&negocio.Correo,
			&negocio.Imagen,
			&negocio.Latitude,
			&negocio.Longitude,
			&negocio.Facebook,
			&negocio.Twitter,
			&negocio.Instagram,
			&negocio.Website,
			&negocio.Creado,
			&negocio.Actualizado,
		)
		if err != nil {
			return nil, Error500(err.Error())
		}
	} else {
		return nil, Error404(fmt.Sprintf("No se encontro el negocio %d", id))
	}

	return &negocio, nil
}

func ObtenerServicioNegocio(id int64) (*[]Servisio, *ErrorStatusCode) {

	db, err := db.ConnectDB()
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("select s.idservicio ,s.nombre ,s.descripcion ,s.imagen ,s.unidad ,s.creado ,s.actualizado from servisio s where s.idnegocio =$1;")
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer stmt.Close()

	resul, err := stmt.Query(id)
	if err != nil {
		return nil, Error500(err.Error())
	}

	var resultados []Servisio
	for resul.Next() {
		var buscar Servisio
		err := resul.Scan(
			&buscar.IDServicio,
			&buscar.Nombre,
			&buscar.Descripcion,
			&buscar.Imagen,
			&buscar.Unidad,
			&buscar.Creado,
			&buscar.Actualizado,
		)

		if err != nil {
			return nil, Error500(err.Error())
		}
		resultados = append(resultados, buscar)
	}

	return &resultados, nil
}

func ObtenerProductoNegocio(id int64) (*[]Producto, *ErrorStatusCode) {

	db, err := db.ConnectDB()
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("select p.idproducto ,p.nombre ,p.descripsion ,p.imagen ,p.unidad ,p.creado ,p.actualizado  from producto p  where p.idnegocio =$1;")
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer stmt.Close()

	resul, err := stmt.Query(id)
	if err != nil {
		return nil, Error500(err.Error())
	}

	var resultados []Producto
	for resul.Next() {
		var buscar Producto
		err := resul.Scan(
			&buscar.IDProducto,
			&buscar.Nombre,
			&buscar.Descripsion,
			&buscar.Imagen,
			&buscar.Unidad,
			&buscar.Creado,
			&buscar.Actualizado,
		)

		if err != nil {
			return nil, Error500(err.Error())
		}
		resultados = append(resultados, buscar)
	}

	return &resultados, nil
}

func ObtenerUltimosNegocios() (*[]Negocio, *ErrorStatusCode) {

	db, err := db.ConnectDB()
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM negocio ORDER BY Creado DESC LIMIT 10;")
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer stmt.Close()

	resul, err := stmt.Query()
	if err != nil {
		return nil, Error500(err.Error())
	}

	var resultados []Negocio
	for resul.Next() {
		var buscar Negocio
		err := resul.Scan(
			&buscar.IDNegocio,
			&buscar.Nombre,
			&buscar.Descripsion,
			&buscar.Direccion,
			&buscar.Telefono,
			&buscar.Correo,
			&buscar.Imagen,
			&buscar.Latitude,
			&buscar.Longitude,
			&buscar.Facebook,
			&buscar.Twitter,
			&buscar.Instagram,
			&buscar.Website,
			&buscar.Creado,
			&buscar.Actualizado,
		)

		if err != nil {
			return nil, Error500(err.Error())
		}
		resultados = append(resultados, buscar)
	}

	return &resultados, nil
}
