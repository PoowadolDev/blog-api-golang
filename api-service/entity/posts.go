package entity

import (
	"gorm.io/gorm"
)

type Posts struct {
	gorm.Model
	Title   string
	Content string
}
