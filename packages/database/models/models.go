package models

import (
	"gorm.io/gorm"
)

type Film struct {
	gorm.Model
	Title    string
	Director string
	GenreId  uint64
	Starred  bool
}

type Genre struct {
	gorm.Model
	Description string
}
