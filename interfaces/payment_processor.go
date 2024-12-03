package interfaces

import (
	"oracle/models"
)

type PaymentProcessor interface {
	GeneratePaymentRequest(dataSize int) models.PaymentRequest
	VerifyPayment(expectedPaymentID string) (bool, error)
	Init()
}
