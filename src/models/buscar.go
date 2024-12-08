package models

import (
	"negosioscol/src/db"
)

type Buscar struct {
	IDNegosio   int64   `json:"id_Negocio"`
	IDProducto  *int64  `json:"id_Producto,omitempty"`
	IDServicio  *int64  `json:"id_servicio,omitempty"`
	Negocio     string  `json:"negocio"`
	Nombre      *string `json:"nombre,omitempty"`
	Descripsion string  `json:"descripsion"`
	Imagen      string  `json:"imagen"`
}

func BuscarServicioProducto(buscar string) (*[]Buscar, *ErrorStatusCode) {

	db, err := db.ConnectDB()
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer db.Close()

	query := `SELECT * from buscarelementos($1);`
	rows, err := db.Query(query, buscar)
	if err != nil {
		return nil, Error500(err.Error())
	}
	defer rows.Close()

	var resultados []Buscar
	for rows.Next() {
		var buscar Buscar
		err := rows.Scan(
			&buscar.IDNegosio,
			&buscar.IDProducto,
			&buscar.IDServicio,
			&buscar.Negocio,
			&buscar.Nombre,
			&buscar.Descripsion,
			&buscar.Imagen,
		)

		if err != nil {
			return nil, Error500(err.Error())
		}
		resultados = append(resultados, buscar)
	}

	return &resultados, nil

}
