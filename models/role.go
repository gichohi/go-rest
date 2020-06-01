package models

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	roleid  int	 `json:"roleid,omitempty"`
	name   string `json:"name,omitempty"`
	description   string `json:"description,omitempty"`
}
