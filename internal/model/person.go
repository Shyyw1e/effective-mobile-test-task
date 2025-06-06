package model

import (
	"github.com/lib/pq"
)

type Person struct {
	ID           uint            `json:"id" gorm:"primaryKey"`
	Name         string          `json:"name"`
	Surname      string          `json:"surname"`
	Patronymic   *string         `json:"patronymic,omitempty"`
	Age          int             `json:"age"`
	Gender       string          `json:"gender"`
	Nationalities pq.StringArray `gorm:"type:text[]" json:"nationalities" swaggertype:"array,string"`
}


