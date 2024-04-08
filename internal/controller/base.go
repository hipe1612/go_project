package controller

import (
	"crud-person/internal/common"
	"crud-person/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
}

func (controller *BaseController) Success(c *gin.Context, v ...interface{}) {
	c.JSON(http.StatusOK, model.ResponseWrapper{
		Code:    common.Success.Code,
		Message: common.Success.Message,
		Data:    v,
	})
}

func (controller *BaseController) Error(c *gin.Context, err *common.Error, v ...interface{}) {
	c.JSON(http.StatusOK, model.ResponseWrapper{
		Code:    err.Code,
		Message: err.Message,
		Data:    v,
	})
}
