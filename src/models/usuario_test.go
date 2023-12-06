package models_test

import (
	"negosioscol/src/models"
	"testing"
)

func TestCrearUsuario(t *testing.T) {
	_, err := models.CrearUsuario("EMilio 1", "ZAnches", "2023-12-02 22:32:17", "https://bestprofilepictures.com/wp-content/uploads/2021/08/Amazing-Profile-Picture.jpg")
	if err != nil {
		t.Error(err)
	}

	_, err = models.CrearUsuario("EMilio 2", "ZAnches", "2023-12-02", "https://bestprofilepictures.com/wp-content/uploads/2021/08/Amazing-Profile-Picture.jpg")
	if err != nil {
		t.Error(err)
	}

	_, err = models.CrearUsuario("EMilio 3", "ZAnches", "02/12/2023", "https://bestprofilepictures.com/wp-content/uploads/2021/08/Amazing-Profile-Picture.jpg")
	if err != nil {
		t.Error(err)
	}
	_, err = models.CrearUsuario("EMilio 4", "ZAnches", "02/12/2023", "https://bestprofilepictures.com/wp-content/uploads/2021/08/Amazing-Profile-Picture.jpghttps://bestprofilepictures.com/wp-content/uploads/2021/08/Amazing-Profile-Picture.jpghttps://bestprofilepictures.com/wp-content/uploads/2021/08/Amazing-Profile-Picture.jpg")
	if err != nil {
		t.Log(err)
	} else {
		t.Error("no se genero un error")
	}
}
func TestEditarUsuario(t *testing.T) {
	err := models.EditarUsuario(1, "Prueba de usuario 1", "NuevosApellidos", "1990-01-01", "https://bestprofilepictures.com/wp-content/uploads/2021/08/Amazing-Profile-Picture.jpg")
	if err != nil {
		t.Error(err)
	}

	err = models.EditarUsuario(2, "Prueba de usuario 2", "NuevosApellidos", "1990-01-01", "https://bestprofilepictures.com/wp-content/uploads/2021/08/Amazing-Profile-Picture.jpg")
	if err != nil {
		t.Error(err)
	}

	err = models.EditarUsuario(3, "Prueba de usuario 3", "NuevosApellidos", "1990-01-01", "https://bestprofilepictures.com/wp-content/uploads/2021/08/Amazing-Profile-Picture.jpg")
	if err != nil {
		t.Error(err)
	}
	err = models.EditarUsuario(4, "Prueba de usuario 4", "NuevosApellidos", "1990-01-01", "https://bestprofilepictures.com/wp-content/uploads/2021/08/Amazing-Profile-Picture.jpg")
	if err != nil {
		t.Log(err)
	} else {
		t.Errorf("no se genero un error [%v]", err)
	}

}

func TestEliminarUsuario(t *testing.T) {
	err := models.EliminarUsuario(3)
	if err != nil {
		t.Error(err)
	}
	err = models.EliminarUsuario(4)
	if err != nil {
		t.Log(err)
	} else {
		t.Errorf("no se genero un error [%v]", err)
	}
}
func TestObtenerUsuario(t *testing.T) {
	user, err := models.ObtenerUsuario(1)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(user)
	}
	_, err = models.ObtenerUsuario(4)
	if err != nil {
		t.Log(err)
	} else {
		t.Errorf("no se genero un error [%v]", err)
	}
}
