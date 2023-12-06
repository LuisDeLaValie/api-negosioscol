package utils

import (
	"negosioscol/src/models"

	"github.com/gin-gonic/gin"
)

func GetParam(params gin.Params, param string) (*string, *models.ErrorStatusCode) {
	val, exis := params.Get(param)
	if !exis {
		return nil, models.Error400("falta el parametro %s", param)
	}
	return &val, nil
}
