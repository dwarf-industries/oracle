package models

import "math/big"

type Oracle struct {
	Name             string
	Ip               string
	Port             string
	Reputation       big.Int
	ConnectionStatus bool
}
