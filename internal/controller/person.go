package controller

import (
	"crud-person/internal/service"
	"github.com/gin-gonic/gin"
)

type IPersonController interface {
	CreatePerson() gin.HandlerFunc
	GetPerson() gin.HandlerFunc
	UpdatePerson() gin.HandlerFunc
	DeletePerson() gin.HandlerFunc
	ListPerson() gin.HandlerFunc
}

type PersonController struct {
	BaseController
	service service.IPersonService
}

func (c PersonController) CreatePerson() gin.HandlerFunc {
	return func(g *gin.Context) {
		err := c.service.CreatePerson(g)
		if err != nil {
			c.Error(g, err, nil)
			return
		}
		c.Success(g)
	}
}

func (c PersonController) GetPerson() gin.HandlerFunc {
	return func(g *gin.Context) {
		person, err := c.service.GetPerson(g)
		if err != nil {
			c.Error(g, err, nil)
			return
		}
		c.Success(g, person)
	}
}

func (c PersonController) UpdatePerson() gin.HandlerFunc {
	return func(g *gin.Context) {
		person, err := c.service.UpdatePerson(g)
		if err != nil {
			c.Error(g, err, nil)
			return
		}
		c.Success(g, person)
	}
}

func (c PersonController) DeletePerson() gin.HandlerFunc {
	return func(g *gin.Context) {
		err := c.service.DeletePerson(g)
		if err != nil {
			c.Error(g, err, nil)
			return
		}
		c.Success(g)
	}
}

func (c PersonController) ListPerson() gin.HandlerFunc {
	return func(g *gin.Context) {
		person, err := c.service.ListPerson(g)
		if err != nil {
			c.Error(g, err, nil)
			return
		}
		c.Success(g, person)
	}
}

func NewPersonController() IPersonController {
	return &PersonController{service: service.NewPersonService()}
}
