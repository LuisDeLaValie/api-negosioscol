
package main


import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Crear un enrutador utilizando gorilla/mux
	r := mux.NewRouter()

	// Define tus rutas aquí
	// Ejemplo:
	r.HandleFunc("/hello", handlerFunc).Methods("GET")

	// Iniciar el servidor utilizando el enrutador
	log.Print("Servidor Web en linea")
	log.Print("Coriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("¡Hola desde gorilla/mux! que haces"))
}