package utils

// CreateErrorResponse is a generic error response
func CreateErrorResponse(err error) ErrorResponse {
	return ErrorResponse{Error: err.Error()}
}

type ErrorResponse struct {
	Error string `json:"error"`
}
