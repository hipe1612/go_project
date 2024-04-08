package repository

import (
	"crud-person/internal/config/database"
	"crud-person/internal/model"
	"gorm.io/gorm"
)

type IPersonRepository interface {
	CreatePerson(data *model.NamesCreation) error
	GetPerson(id int) (model.Names, error)
	UpdatePerson(id int, data *model.NamesUpdate) error
	DeletePerson(id int) error
	ListPerson() ([]model.Names, error)
}

var personRepository IPersonRepository

func GetPersonRepository() IPersonRepository {
	if personRepository == nil {
		personRepository = &PersonRepository{DB: database.Instance}
		return personRepository
	}
	return personRepository
}

type PersonRepository struct {
	DB *gorm.DB
}

func (r PersonRepository) CreatePerson(data *model.NamesCreation) error {
	err := r.DB.Create(data).Error
	return err
}

func (r PersonRepository) GetPerson(id int) (model.Names, error) {
	var person model.Names
	err := database.Instance.First(&person, "id = ?", id).Error
	return person, err
}

func (r PersonRepository) UpdatePerson(id int, data *model.NamesUpdate) error {
	err := r.DB.Where("id = ?", id).Updates(data).Error
	return err
	//if err := r.DB.Model(&model.Names{}).Where("id = ?", id).Updates(data).Error; err != nil {
	//	return &common.Error{Code: "500", Message: err.Error()}
	//}
	//return nil
}

func (r PersonRepository) DeletePerson(id int) error {
	err := r.DB.Delete(&model.Names{}, id).Error
	return err
	//if err := r.DB.Delete(&model.Names{}, id).Error; err != nil {
	//	return &common.Error{Code: "500", Message: err.Error()}
	//}
	//return nil
}

func (r PersonRepository) ListPerson() ([]model.Names, error) {
	var result []model.Names
	err := r.DB.Find(&result).Error
	return result, err
}
