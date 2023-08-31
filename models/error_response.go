package models

type ErrorResponse struct {
	Success bool   `json:"success"`
	Field   string `json:"field"`
	Errors  string `json:"errors"`
}
