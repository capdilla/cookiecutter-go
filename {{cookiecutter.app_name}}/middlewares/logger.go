package middlewares

import (
	"math"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var timeFormat = "02/Jan/2006:15:04:05 -0700"

// Logger will return a middleware who will log every request
func Logger(lg *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Setting up variables previous calling next handler
		start := time.Now()

		// Call next handler, it might be another middleware of the endpoint handler
		c.Next()

		// Log request with its response
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()

		fields := []zap.Field{
			zap.String("method", c.Request.Method),
			zap.String("request_id", requestid.Get(c)),
			zap.String("time", time.Now().Format(timeFormat)),
			zap.String("route", c.Request.URL.Path),
			zap.Int("status_code", statusCode),
			zap.String("referer", c.Request.Referer()),
			zap.String("clientUser_agent", c.Request.UserAgent()),
			zap.Int("latency", latency),
		}

		if statusCode > 499 {
			lg.Error("Request finished with server error: 5xx", fields...)
		} else if statusCode > 399 {
			lg.Warn("Request finished with client error: 4xx", fields...)
		} else {
			lg.Info("Request finished succesfully", fields...)
		}
	}
}
