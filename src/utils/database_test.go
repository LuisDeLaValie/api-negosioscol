package utils_test

import (
	"negosioscol/src/utils"
	"testing"
)

// Escribe TestXXXX en donde XXXX es el nombre de la función original
func TestConnectDB(t *testing.T) { // Recibir struct de tipo testing.T
	resultado := utils.ConnectDB()
	if resultado != nil {
		t.Log("todo salio bien")

	} else {
		t.Errorf("Error al iniciar la db")
	}
}
