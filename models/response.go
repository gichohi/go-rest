package models

import (
	"github.com/jinzhu/gorm"
)

type Response struct {
	gorm.Model
	Code  int	 `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
}
