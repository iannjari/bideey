package repo

import (
	"bideey/config"
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AbstractRepository[T any] interface {
	Save(entity *T) (*T, error)
	GetById(id *uuid.UUID) (*T, error)
	Delete(id *uuid.UUID) (err error)
	// TODO: have a query/pagination impl
	Query() (*[]T, error)
}

type Repository[T any] struct {
}

var db *gorm.DB

func NewRepository[T any]() *Repository[T] {
	log.Println(".....................")
	sqlDB := config.NewPostgresDB()
	db = sqlDB.Database
	log.Println(db)
	return &Repository[T]{}
}

func (*Repository[T]) Save(entity *T) (*T, error) {

	var err error

	tx := db.Begin()

	err = db.Create(&entity).Error

	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("could not create entity. error: " + err.Error())
	}
	tx.Commit()
	return entity, nil
}

func (*Repository[T]) GetById(id *uuid.UUID) (entity *T, error error) {
	var value T
	if result := db.Where("id = ?", id).First(&value); result.Error != nil {
		return
	}
	return &value, nil
}

func (*Repository[T]) Delete(id *uuid.UUID) (err error) {
	var entity T
	db.Where("id = ?", id).Delete(&entity)
	return
}

// TODO: have a query/pagination impl
func (*Repository[T]) Query() (entities *[]T, err error) {
	db.Find(&entities)
	return
}
