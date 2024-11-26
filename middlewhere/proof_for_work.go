package middlewhere

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"oracle/di"
)

func PFW() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payment := ctx.GetHeader("pop")

		allow, err := di.PaymentProcessor().VerifyPayment(payment)
		if !allow || err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Rejected, can't verify payment"})
			return
		}

		ctx.Next()
	}
}
