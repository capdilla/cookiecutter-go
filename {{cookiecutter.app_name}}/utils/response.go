package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseCode represents a list of available response codes
type ResponseCode string

// NoOp handles case in which const can't be addressable
const NoOp ResponseCode = ""

// Bad request errors
const (
	// ErrorInvalidQueryType will be used when query string value is wrong
	ErrorInvalidQueryType ResponseCode = "error.invalid_query_type"
	// ErrorInvalidPayload will be used when an error parsing payload occurs
	ErrorInvalidPayload ResponseCode = "error.invalid_payload"
	// ErrorInvalidQueryParam used when a query param has an invalid type
	ErrorInvalidParam ResponseCode = "error.invalid_param"
	// ErrorValidation will be used when a validation doesn't pass
	ErrorValidation ResponseCode = "error.validation"
	// ErrorPrivateKey can be used with any internal error
	ErrorPrivateKey ResponseCode = "error.private_key"
)

// Unauthorized errors
const (
	// ErrorNotAllowed used when user has nos permission to access resource
	ErrorNotAllowed ResponseCode = "error.not_allowed"
	// ErrorInvalidToken used when token is invalid or has invalid claim
	ErrorInvalidToken ResponseCode = "error.invalid_token"
)

// Internal server errors
const (
	// ErrorInternal can be used with any internal error
	ErrorInternal ResponseCode = "error.internal"
	// ErrorDB will be used when a error on db accurs
	ErrorDB ResponseCode = "error.db"
)

// ErrorFunc ...
type ErrorFunc func(c *gin.Context, customCode ResponseCode, err error)

// ResponseCodeToErrorRespondFunc get function to appropiately handle error
var ResponseCodeToErrorRespondFunc = map[ResponseCode]ErrorFunc{
	ErrorInvalidQueryType: RespondBadRequestJSON,
	ErrorInvalidPayload:   RespondBadRequestJSON,
	ErrorInvalidParam:     RespondBadRequestJSON,
	ErrorValidation:       RespondBadRequestJSON,
	ErrorNotAllowed:       RespondUnauthorized,
	ErrorInvalidToken:     RespondUnauthorized,
	ErrorInternal:         RespondInternalSrvErrJSON,
	ErrorDB:               RespondInternalSrvErrJSON,
}

// RespondJSON will respond a JSON response with the provided values
func RespondJSON(c *gin.Context, httpCode int, customCode ResponseCode, err error) {
	// TODO: Change c.Error for log.Error when zap is added
	c.Error(err)
	c.AbortWithStatusJSON(httpCode, gin.H{
		"code":    customCode,
		"message": err.Error(),
	})
}

// RespondUnauthorized returns "Unauthorized"
func RespondUnauthorized(c *gin.Context, customCode ResponseCode, err error) {
	RespondJSON(c, http.StatusUnauthorized, customCode, err)
}

// RespondNotFoundJSON will return a "NotFound" error in JSON format
func RespondNotFoundJSON(c *gin.Context, customCode ResponseCode, err error) {
	RespondJSON(c, http.StatusNotFound, customCode, err)
}

// RespondBadRequestJSON will return a "BadRequest" error in JSON format
func RespondBadRequestJSON(c *gin.Context, customCode ResponseCode, err error) {
	RespondJSON(c, http.StatusBadRequest, customCode, err)
}

// RespondUnprocessableEntityJSON will respond an "UnprocessableEntity" error in JSON format
func RespondUnprocessableEntityJSON(c *gin.Context, customCode ResponseCode, err error) {
	RespondJSON(c, http.StatusUnprocessableEntity, customCode, err)
}

// RespondInternalSrvErrJSON will respond an "StatusInternalServerError" in JSON format
func RespondInternalSrvErrJSON(c *gin.Context, customCode ResponseCode, err error) {
	RespondJSON(c, http.StatusInternalServerError, customCode, err)
}
