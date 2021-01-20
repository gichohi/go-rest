package models

import "github.com/jinzhu/gorm"

type LoginRequest struct {
	gorm.Model
	Password  string   `json:"password,omitempty"`
	UserName   string `json:"username,omitempty"`
}
