package model

type ErrorResponse struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Id      string `json:"id,omitempty"`
}
