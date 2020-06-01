package models

import "github.com/jinzhu/gorm"

type UserResponse struct {
	gorm.Model
	Email    string `json:"email,omitempty"`
	FirstName  string   `json:"firstname,omitempty"`
	LastName   string `json:"lastname,omitempty"`
	Token  string	 `json:"token,omitempty"`
}
