package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email,omitempty"`
	FirstName  string   `json:"firstname,omitempty"`
	LastName   string `json:"lastname,omitempty"`
	Password   string `json:"password,omitempty"`
	Token  string	 `json:"token,omitempty"`
	Roles	[]Role `json:"roles,omitempty"`
}
