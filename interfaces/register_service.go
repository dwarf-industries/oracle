package interfaces

import (
	"math/big"

	"oracle/models"
)

type RegisterService interface {
	Register(ip string) error
	Oracles() ([]models.Oracle, error)
	Registered() bool
	GetRegistrationFee() (*big.Int, error)
	GetReportFee() (*big.Int, error)
	GetOracle() (*models.Oracle, error)
	LogOut() error
	Login() error
}
