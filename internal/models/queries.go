package models

import (
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Query struct {
	gorm.Model `json:"model"`

	Site     Sites  `json:"site"`
	Username string `json:"username"`
}
