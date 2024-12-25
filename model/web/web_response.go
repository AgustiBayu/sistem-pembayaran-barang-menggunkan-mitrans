package web

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type DeleteResponse struct {
	Code int `json:"code"`
}

type DataResponse struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Links []Link      `json:"links"`
}

type ErrorResponse struct {
	Code  int         `json:"code"`
	Error interface{} `json:"error"`
}
