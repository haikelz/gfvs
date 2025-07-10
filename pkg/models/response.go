package models

type APIHomeResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"string"`
}

type APIErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Stack      string `json:"stack"`
}
