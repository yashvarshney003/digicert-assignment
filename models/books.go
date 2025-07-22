package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `gorm:"size:100"`
	Description string `gorm:"size:500"`
	Author      string `gorm:"size:100"`
}
