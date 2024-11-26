package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"oracle/interfaces"
	"oracle/services"
)

var walletService interfaces.WalletService
var registerService interfaces.RegisterService
var rpcService interfaces.RpcService
var verificationService interfaces.VerificationService

func setupServices() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rpc := getRpc()

	rpcService = &services.RpcService{}
	rpcService.SetClient(rpc)

	walletService = &services.WalletService{
		PasswordManager: &services.PasswordManager{},
		RpcService:      rpcService,
	}
	registerService = &services.RegisterService{
		ContractAddr:  os.Getenv("CONTRACT_ADDRESS"),
		WalletService: walletService,
		RpcService:    rpcService,
		VerificationService: &services.IdentityService{
			WalletService: walletService,
		},
	}
	verificationService = &services.VerificationService{}
	verificationService.Init()
}
func getRpc() *string {
	rpc := os.Getenv("RPC")

	overriden, err := os.ReadFile("oracle-rpc")
	if err != nil {
		return &rpc
	}

	converted := string(overriden)
	return &converted
}
