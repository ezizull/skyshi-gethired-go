package controllers

// JSONSwagger is a struct that contains the swagger documentation
type JSONSwagger struct {
}

// MessageResponse is a struct that contains the response body for the message
type MessageResponse struct {
	Message string `json:"message"`
}

// ErrorResponse is a error response struct for error handler
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// DefaultResponse is a default struct for response body
type DefaultResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
