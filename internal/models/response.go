package models

import "gorm.io/gorm"

type UsernameResponse struct {
	Found            bool   `json:"found"`
	URI              string `json:"uri"`
	MethodValidation string `json:"method_validation"`
	Username         string `json:"username"`
	ResponseStatus   int    `json:"response_status"`
	SiteName         string `json:"site_name"`
	IDSite           int    `json:"id_site"`
	IDDownload       string `json:"id_download"`
}

type UsernameResponseDatabase struct {
	gorm.Model `json:"model"`

	UsernameResponse UsernameResponse `json:"username_response"`
}
