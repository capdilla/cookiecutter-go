package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthController type
type HealthController struct{
	healthService: *services.HealthService
}

// GenericError is the default error message that is generated.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    string `json:"code"`
		Message error  `json:"message"`
	} `json:"body"`
}

// StatusOK a healthcheck response model
//
// swagger:response statusOK
type StatusOK struct {
	Body string
}

// Status swagger:route GET /health health checkStatus
//
// Reports app's healthcheck.
//
// Deprecated: false
// Responses:
//		default: genericError
// 		    200: statusOK "OK"
func (h HealthController) Status(c *gin.Context) {
	c.String(http.StatusOK, h.healthService.GetStatus())
}
