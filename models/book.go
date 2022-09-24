package models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title" form:"title"`
	Author string `json:"author" form:"author"`
	Genres string `json:"genre" form:"genre"`
	Year   int    `json:"year" form:"year"`
}
