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
func (r *PersonRepository)FindWithFilters(name, gender, nationality string, page, limit int) ([]model.Person, error) {
	var people []model.Person
	query := r.db.Model(&model.Person{})

	if name != "" {
		query = query.Where("name = ?", name)
	}
	if gender != "" {
		query = query.Where("gender = ?", gender)
	}
	if nationality != "" {
		query = query.Where("? = ANY(nationalities)", nationality)
	}
	
	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Find(&people).Error
	return people, err
}

func (r *PersonRepository) DeleteByID(id uint) error {
	return r.db.Delete(&model.Person{}, id).Error
}

func (r *PersonRepository) UpdateBasicInfo(id uint, name, surname string, patronymic *string) error {
	return r.db.Model(&model.Person{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name":       name,
		"surname":    surname,
		"patronymic": patronymic,
	}).Error
}