package model

import "github.com/jinzhu/gorm"

type (
	TodoModel struct {
		gorm.Model
		Title     string `json:"title"`
		Completed int    `json:"completed", gorm:"default:0"`
	}

	TodoModelGetData struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
)
