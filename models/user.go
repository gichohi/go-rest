package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);unique_index"`
	FirstName  string   `json:"firstname,omitempty"`
	LastName   string `json:"lastname,omitempty"`
	Password   string `json:"password,omitempty"`
	Token  string	 `json:"token,omitempty"`
}
