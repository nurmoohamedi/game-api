package models

import (
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Title string `json:"title"`
	Characters []Character `json:"characters,omitempty"`
}

type Character struct {
	gorm.Model
	Name string `json:"name"`
	Class string `json:"class"`
	GameID uint `json:"game_id"`
}

type User struct {
	gorm.Model
	Username     string `json:"username" gorm:"unique"`
	PasswordHash string `json:"-"`
}
