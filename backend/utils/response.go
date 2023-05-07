package utils

import "github.com/gin-gonic/gin"

// ErrorResponse is a generic error response
func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
