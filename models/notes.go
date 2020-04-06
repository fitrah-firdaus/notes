package models

import "github.com/jinzhu/gorm"

type Notes struct {
	gorm.Model
	Title   string `json:"title" validate:"required" gorm:"unique;not null"`
	Content string `json:"content" validate:"required"`
	Author  string `json:"author"`
}
