package models

type UsernameRespnse struct {
	Found            bool   `json:"found"`
	URI              string `json:"uri"`
	MethodValidation string `json:"method_validation"`
	Username         string `json:"username"`
	ResponseStatus   int    `json:"response_status"`
	SiteName         string `json:"site_name"`
}
