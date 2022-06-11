package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Sites struct {
	gorm.Model `json:"model,omitempty"`

	Sitename          string `json:"_key" gorm:"unique"`
	ErrorType         string `json:"errorType"`
	ErrorMessage      string `json:"errorMsg"`
	URLDomain         string `json:"urlMain"`
	URLFormat         string `json:"url"`
	UsernameClaimed   string `json:"username_claimed"`
	UsernameUnclaimed string `json:"username_unclaimed"`
}

type SitesJSON struct {
	Sitename          string      `json:"_key" gorm:"unique"`
	ErrorType         string      `json:"errorType"`
	ErrorMessage      interface{} `json:"errorMsg"`
	URLDomain         string      `json:"urlMain"`
	URLFormat         string      `json:"url"`
	UsernameClaimed   string      `json:"username_claimed"`
	UsernameUnclaimed string      `json:"username_unclaimed"`
}
