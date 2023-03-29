/*
Package response provides helpers and utils for working with HTTP response
*/
package response

import (
	http_errors "web3-with-go/pkg/http/errors"

	"github.com/gin-gonic/gin"
)

// SendJson returns data as json response
func SendJson(c *gin.Context, statusCode int, payload interface{}) {
	if payload == nil {
		c.Status(statusCode)
	} else {
		c.JSON(statusCode, payload)
	}
}

func SendJsonError(c *gin.Context, errorName string, code int, message string) {
	err := &http_errors.HttpError{
		StatusCode: code,
		Message:    message,
		ErrorName:  errorName,
	}
	c.JSON(err.StatusCode, err)
}
func SendJsonHttpError(c *gin.Context, err *http_errors.HttpError) {
	if err == nil {
		panic("JSONHttpError called for nil error")
	}
	c.JSON(err.StatusCode, err)
}

func SendSuccess(c *gin.Context, payload interface{}) {
	SendJson(c, 200, payload)
}
func SendCreated(c *gin.Context, payload interface{}) {
	SendJson(c, 201, payload)
}
func SendUnauthorized(c *gin.Context, message string) {
	SendJsonError(c, "Unauthorized", 401, message)
}
func SendBadRequest(c *gin.Context, message string) {
	SendJsonError(c, "Bad request", 400, message)
}
func SendNotFound(c *gin.Context, message string) {
	SendJsonError(c, "Not found", 404, message)
}
func SendInternalError(c *gin.Context, message string) {
	SendJsonError(c, "Internal error", 500, message)
}
