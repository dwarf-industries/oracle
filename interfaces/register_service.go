package interfaces

import "oracle/models"

type RegisterService interface {
	Register(ip string) error
	Oracles() ([]models.Oracle, error)
	Registered() bool
}
