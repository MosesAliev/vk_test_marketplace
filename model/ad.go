package model

import "gorm.io/gorm"

type Ad struct {
	gorm.Model
	Title       string
	Description string
	Image       string
	Price       int
	UserLogin   string `json:"-"`
	IsYours     bool   `gorm:"-" json:"is_yours,omitempty"`
}
