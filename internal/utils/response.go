package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response is the response for the API
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

// PaginatedResponse is the response for the paginated API
type PaginatedResponse struct {
	Response
	Meta PaginatedMeta `json:"meta"`
}

// PaginatedMeta is the meta for the paginated response
type PaginatedMeta struct {
	Total      int `json:"total"`
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"total_pages"`
}

// SuccessResponse is the response for the success API
func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse is the response for the error API
func ErrorResponse(c *gin.Context, statusCode int, message string, err error) {
	resp := Response{
		Success: false,
		Message: message,
	}
	if err != nil {
		resp.Error = err.Error()
	}

	c.JSON(statusCode, resp)
}

// CreatedResponse is the response for the created API
func CreatedResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// NotFoundResponse is the response for the not found API
func NotFoundResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, message, nil)
}

// BadRequestResponse is the response for the bad request API
func BadRequestResponse(c *gin.Context, message string, err error) {
	ErrorResponse(c, http.StatusBadRequest, message, err)
}

// UnauthorizedResponse is the response for the unauthorized API
func UnauthorizedResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusUnauthorized, message, nil)
}

// ForbiddenResponse is the response for the forbidden API
func ForbiddenResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusForbidden, message, nil)
}

// InternalServerErrorResponse is the response for the internal server error API
func InternalServerErrorResponse(c *gin.Context, message string, err error) {
	ErrorResponse(c, http.StatusInternalServerError, message, err)
}

// PaginatedSuccessResponse is the response for the paginated success API
func PaginatedSuccessResponse(c *gin.Context, message string, data interface{}, meta PaginatedMeta) {
	c.JSON(http.StatusOK, PaginatedResponse{
		Response: Response{
			Success: true,
			Message: message,
			Data:    data,
		},
		Meta: meta,
	})
}
