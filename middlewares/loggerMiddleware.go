package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	appLogger "gogintemplate/logger"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	logger := appLogger.GetLogger()

	return func(context *gin.Context) {
		msg := fmt.Sprintf("[%s] %s %s\n", time.Now().Format("2006-01-02 15:04:05"), context.Request.Method, context.Request.URL.Path)

		// Record the start time to measure the duration of the request
		logger.Info(msg)
		start := time.Now()

		// Call the next handler in the chain
		context.Next()

		// Log information about the response
		duration := time.Since(start)
		msg = fmt.Sprintf("[%s] Completed in %v\n", time.Now().Format("2006-01-02 15:04:05"), duration)

		logger.Info(msg)

	}
}
