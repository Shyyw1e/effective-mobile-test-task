package model

import (
	"time"

	"gorm.io/gorm"
	"github.com/lib/pq"
)

type Person struct {
	ID 				uint 				`gorm:"primaryKey" json:"id"`
	Name			string				`json:"name"`
	Surname 		string 				`json:"surname"`
	Patronymic		*string				`json:"patronymic, omitempty"`
	Age				int					`json:"age"`
	Gender			string				`json:"gender"`
	Nationalities	pq.StringArray		`gorm:"type:text[]" json:"nationalities"`
	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt		gorm.DeletedAt		`gorm:"index"`	
}

