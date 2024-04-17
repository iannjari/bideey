package repo

import (
	"bideey/config"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AbstractRepository interface {
	Save(entity any) (*any, error)
	GetById(id *uuid.UUID) (interface{}, error)
	Delete(id uuid.UUID) (err error)
	Update(entity interface{})
	// TODO: have a query/pagination impl
	Query() (*[]interface{}, error)
}

type Repository struct {
}

var db *gorm.DB

func NewRepository() *Repository {
	db = config.Database
	return &Repository{}
}

func (*Repository) Save(entity any) (*any, error) {

	var err error

	tx := db.Begin()

	err = db.Create(&entity).Error

	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("could not create entity. error: " + err.Error())
	}
	tx.Commit()
	return &entity, nil
}

func (*Repository) GetById(id *uuid.UUID) (entity *interface{}, error error) {
	db.First(&entity, id.String)
	return
}

func (*Repository) Delete(id uuid.UUID) (err error) {
	var entity interface{}
	db.Delete(&entity)
	return
}

// TODO: have a query/pagination impl
func (*Repository) Query() (entity *interface{}, err error) {
	db.Find(&entity)
	return
}
