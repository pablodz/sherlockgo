package models

import "gorm.io/gorm"

type Query struct {
	gorm.Model `json:"model"`

	Site     Sites  `json:"site"`
	Username string `json:"username"`
}
