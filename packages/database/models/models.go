package models

import (
	"gorm.io/gorm"
)

type Film struct {
	gorm.Model
	Title    string
	Director string
	GenreId  uint
	Starred  bool
}

type Genre struct {
	gorm.Model
	Description string
}
