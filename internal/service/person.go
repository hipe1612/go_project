package service

import (
	"crud-person/internal/common"
	"crud-person/internal/model"
	"crud-person/internal/repository"
	"github.com/gin-gonic/gin"
	"strconv"
)

type IPersonService interface {
	CreatePerson(g *gin.Context) *common.Error
	GetPerson(g *gin.Context) (*model.Names, *common.Error)
	UpdatePerson(g *gin.Context) (*model.NamesUpdate, *common.Error)
	DeletePerson(g *gin.Context) *common.Error
	ListPerson(g *gin.Context) ([]model.Names, *common.Error)
}

type PersonService struct {
	repository repository.IPersonRepository
}

func NewPersonService() IPersonService {
	return &PersonService{repository: repository.GetPersonRepository()}
}

func (c PersonService) CreatePerson(g *gin.Context) *common.Error {
	var data model.NamesCreation
	if err := g.ShouldBindJSON(&data); err != nil {
		return &common.Error{Code: "400", Message: "Invalid request body"}
	}

	if err := c.repository.CreatePerson(&data); err != nil {
		return nil
	}

	return nil
}

func (c PersonService) GetPerson(g *gin.Context) (*model.Names, *common.Error) {
	id := g.Param("id")
	personID, err := strconv.Atoi(id)
	if err != nil {
		return nil, &common.Error{Code: "400", Message: "Invalid ID format"}
	}

	person, err := c.repository.GetPerson(personID)
	if err != nil {
		commonErr := &common.Error{Code: "500", Message: err.Error()}
		return nil, commonErr
	}

	return &person, nil
}

func (c PersonService) UpdatePerson(g *gin.Context) (*model.NamesUpdate, *common.Error) {
	id := g.Param("id")
	personID, err := strconv.Atoi(id)
	if err != nil {
		return nil, &common.Error{Code: "400", Message: err.Error()}
	}
	var data model.NamesUpdate
	if err := g.ShouldBindJSON(&data); err != nil {
		return nil, &common.Error{Code: "400", Message: err.Error()}
	}
	if err := c.repository.UpdatePerson(personID, &data); err != nil {
		return nil, &common.Error{Code: "400", Message: err.Error()}
	}
	return &data, nil
}

func (c PersonService) DeletePerson(g *gin.Context) *common.Error {
	id := g.Param("id")
	personID, err := strconv.Atoi(id)
	if err != nil {
		return &common.Error{Code: "400", Message: "Invalid ID format"}
	}
	if err := c.repository.DeletePerson(personID); err != nil {
		return &common.Error{Code: "400", Message: err.Error()}
	}
	return nil
}

func (c PersonService) ListPerson(g *gin.Context) ([]model.Names, *common.Error) {
	persons, err := c.repository.ListPerson()
	if err != nil {
		return nil, &common.Error{Code: "400", Message: err.Error()}
	}
	return persons, nil
}
