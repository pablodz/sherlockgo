package models

import "gorm.io/gorm"

type Sites struct {
	gorm.Model `json:"model"`

	IDSite            int    `json:"id_site"`
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
