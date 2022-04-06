package controllers

import (
	"core/infra/errors"
	"core/infra/logger"
	"core/svc"
	"core/utils/msgutil"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type system struct {
	svc svc.ISystem
}

// NewSystemController will initialize the controllers
func NewSystemController(sysSvc svc.ISystem) *system {
	return &system{
		svc: sysSvc,
	}
}

// Root will let you see what you can slash 🐲
func (sh *system) Root(c *gin.Context) {
	c.JSON(http.StatusOK, msgutil.NewRestResp("app backend! let's play!!", nil))
}

// Health will let you know the heart beats ❤️
func (sys *system) Health(c *gin.Context) {
	resp, err := sys.svc.GetHealth()
	if err != nil {
		logger.Error(fmt.Sprintf("%+v", resp), err)
		c.JSON(http.StatusInternalServerError, errors.ErrSomethingWentWrong)
	}
	c.JSON(http.StatusOK, msgutil.NewRestResp("", resp))
}
