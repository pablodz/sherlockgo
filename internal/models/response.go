package models

import "gorm.io/gorm"

type UsernameResponse struct {
	Found            string `json:"found" csv:"found" xml:"found"`
	URI              string `json:"uri" csv:"uri" xml:"uri"`
	MethodValidation string `json:"method_validation" csv:"method_validation" xml:"method_validation"`
	Username         string `json:"username" csv:"username" xml:"username"`
	ResponseStatus   int    `json:"response_status" csv:"response_status" xml:"response_status"`
	SiteName         string `json:"site_name" csv:"site_name" xml:"site_name"`
	IDSite           int    `json:"id_site" csv:"id_site" xml:"id_site"`
	IDDownload       string `json:"id_download" csv:"id_download" xml:"id_download"`
}

type UsernameResponseDatabase struct {
	gorm.Model `json:"model"`

	UsernameResponse UsernameResponse `json:"username_response"`
}
