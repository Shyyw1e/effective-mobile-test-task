package repository

import (
	"github.com/Shyyw1e/effective-mobile-test-task/internal/model"
    "gorm.io/gorm"
)

type PersonRepository struct {
	db *gorm.DB
}

func NewPersonRepository(db *gorm.DB) *PersonRepository {
	return &PersonRepository{db: db}
}

func (r *PersonRepository) Save(person *model.Person) error {
	return r.db.Create(person).Error
}