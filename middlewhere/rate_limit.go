package middlewhere

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var requestCounts = make(map[string]int)
var requestTimestamps = make(map[string]time.Time)

const rateLimit = 10

func RateLimiterMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()

		lastRequestTime, exists := requestTimestamps[ip]
		if !exists || time.Since(lastRequestTime) > time.Minute {
			requestCounts[ip] = 0
		}

		if requestCounts[ip] >= rateLimit {
			ctx.JSON(http.StatusTooManyRequests, gin.H{"message": "Rate limit exceeded"})
			ctx.Abort()
			return
		}

		requestCounts[ip]++
		requestTimestamps[ip] = time.Now()

		ctx.Next()
	}
}
