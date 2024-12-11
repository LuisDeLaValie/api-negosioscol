package handlers

import (
	"fmt"
	"negosioscol/src/db"
	"negosioscol/src/models"
	"time"

	"github.com/gin-gonic/gin"
)

type DescargasTotales struct {
	Mes            string `json:"mes"`
	TotalDescargas int64  `json:"descargas"`
}

type BusquedasTotales struct {
	Mes            string `json:"mes"`
	TotalBusquedas int64  `json:"busquedas"`
}
type Totales struct {
	TotalDescargas                 int64 `json:"descargas"`
	TotalNegocios                  int64 `json:"negocios"`
	NegociosConProductosOServicios int64 `json:"negociosPS"`
}
type TotalesDias struct {
	Fecha          time.Time `json:"fecha"`
	TotalNegocios  int64     `json:"negocios"`
	TotalServicios int64     `json:"servicios"`
	TotalProductos int64     `json:"productos"`
	TotalBusquedas int64     `json:"busquedas"`
}

type Estadisticas struct {
	Mes            string `json:"mes"`
	TotalProductos int64  `json:"productos"`
	TotalServicios int64  `json:"servicios"`
}

type Grafiac struct {
	TotalesDias      []TotalesDias      `json:"totalesdias"`
	Totales          Totales            `json:"totales"`
	BusquedasTotales []BusquedasTotales `json:"busquedas"`
	DescargasTotales []DescargasTotales `json:"descargas"`
	Estadisticas     []Estadisticas     `json:"estadisticas"`
}

// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary Show an account
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} model.HTTPError
// @Router /accounts/{id} [get]
func Grafias(c *gin.Context) {
	var grafins Grafiac
	fmt.Println("obtenerEstadisticasMensuales")
	estadisticas, err := obtenerEstadisticasMensuales()
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	fmt.Println("obtenerDescargasPorMes")
	descargasTotales, err := obtenerDescargasPorMes()
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	fmt.Println("obtenerBusquedasPorMes")
	busquedas, err := obtenerBusquedasPorMes()
	if err != nil {
		c.JSON(err.Code, err)
		return

	}
	fmt.Println("obtenerTotales")
	totales, err := obtenerTotales()
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	fmt.Println("obtenerTotalesPorDia")
	totalesdias, err := obtenerTotalesPorDia()
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	grafins.Estadisticas = *estadisticas
	grafins.DescargasTotales = *descargasTotales
	grafins.BusquedasTotales = *busquedas
	grafins.Totales = *totales
	grafins.TotalesDias = *totalesdias

	c.JSON(200, grafins)

}

func obtenerDescargasPorMes() (*[]DescargasTotales, *models.ErrorStatusCode) {

	db, err := db.ConnectDB()
	if err != nil {
		erro := models.Error500(err.Error())
		return nil, erro
	}
	defer db.Close()

	// Llamada al	 SP ObtenerTotales()
	rows, err := db.Query("SELECT * FROM SP_ObtenerDescargasPorMes()")
	if err != nil {
		erro := models.Error500(err.Error())
		return nil, erro
	}
	defer rows.Close()

	// Variables para almacenar los resultados
	var descargasTotales []DescargasTotales

	// Leer los resultados del SP
	for rows.Next() {
		var descargas DescargasTotales
		err := rows.Scan(
			&descargas.Mes,
			&descargas.TotalDescargas,
		)
		if err != nil {
			erro := models.Error500(err.Error())
			return nil, erro
		}
		descargasTotales = append(descargasTotales, descargas)

	}

	return &descargasTotales, nil

}

func obtenerBusquedasPorMes() (*[]BusquedasTotales, *models.ErrorStatusCode) {

	db, err := db.ConnectDB()
	if err != nil {
		erro := models.Error500(err.Error())
		return nil, erro
	}
	defer db.Close()

	// Llamada al	 SP ObtenerTotales()
	rows, err := db.Query("SELECT * FROM SP_ObtenerBusquedasPorMes()")
	if err != nil {
		erro := models.Error500(err.Error())
		return nil, erro
	}
	defer rows.Close()

	// Variables para almacenar los resultados
	var busquedasTotales []BusquedasTotales

	// Leer los resultados del SP
	for rows.Next() {
		var busquedas BusquedasTotales
		err := rows.Scan(
			&busquedas.Mes,
			&busquedas.TotalBusquedas,
		)
		if err != nil {
			erro := models.Error500(err.Error())
			return nil, erro
		}
		busquedasTotales = append(busquedasTotales, busquedas)

	}

	return &busquedasTotales, nil

}

func obtenerTotales() (*Totales, *models.ErrorStatusCode) {

	db, err := db.ConnectDB()
	if err != nil {
		erro := models.Error500(err.Error())
		return nil, erro
	}
	defer db.Close()

	// Llamada al	 SP ObtenerTotales()
	rows, err := db.Query("SELECT * FROM SP_ObtenerTotales()")
	if err != nil {
		erro := models.Error500(err.Error())
		return nil, erro
	}
	defer rows.Close()

	// Variables para almacenar los resultados
	var totales Totales

	// Leer los resultados del SP
	for rows.Next() {
		err := rows.Scan(
			&totales.TotalDescargas,
			&totales.TotalNegocios,
			&totales.NegociosConProductosOServicios,
		)
		if err != nil {
			erro := models.Error500(err.Error())
			return nil, erro
		}

	}

	return &totales, nil

}

func obtenerTotalesPorDia() (*[]TotalesDias, *models.ErrorStatusCode) {

	db, err := db.ConnectDB()
	if err != nil {
		erro := models.Error500(err.Error())
		return nil, erro
	}
	defer db.Close()

	// Llamada al	 SP ObtenerTotales()
	rows, err := db.Query("SELECT * FROM SP_ObtenerTotalesPorDia()")
	if err != nil {
		erro := models.Error500(err.Error())
		return nil, erro
	}
	defer rows.Close()

	// Variables para almacenar los resultados
	var totalesDias []TotalesDias

	// Leer los resultados del SP
	for rows.Next() {
		var totaleDia TotalesDias
		err := rows.Scan(
			&totaleDia.Fecha,
			&totaleDia.TotalNegocios,
			&totaleDia.TotalServicios,
			&totaleDia.TotalProductos,
			&totaleDia.TotalBusquedas,
		)
		if err != nil {
			erro := models.Error500(err.Error())
			return nil, erro
		}
		totalesDias = append(totalesDias, totaleDia)

	}

	return &totalesDias, nil

}

func obtenerEstadisticasMensuales() (*[]Estadisticas, *models.ErrorStatusCode) {

	db, err := db.ConnectDB()
	if err != nil {
		erro := models.Error500(err.Error())
		return nil, erro
	}
	defer db.Close()

	// Llamada al	 SP ObtenerTotales()
	rows, err := db.Query("SELECT * FROM SP_ObtenerEstadisticasMensuales()")
	if err != nil {
		erro := models.Error500(err.Error())
		return nil, erro
	}
	defer rows.Close()

	// Variables para almacenar los resultados
	var estadisticas []Estadisticas

	// Leer los resultados del SP
	for rows.Next() {
		var estadistica Estadisticas
		err := rows.Scan(
			&estadistica.Mes,
			&estadistica.TotalProductos,
			&estadistica.TotalServicios,
		)
		if err != nil {
			erro := models.Error500(err.Error())
			return nil, erro
		}
		estadisticas = append(estadisticas, estadistica)

	}

	return &estadisticas, nil

}
