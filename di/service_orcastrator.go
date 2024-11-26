package di

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
var paymentProcessorService interfaces.PaymentProcessor

func SetupServices() {
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
	paymentProcessorService = &services.PaymentProcessor{}

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

func WalletService() interfaces.WalletService {
	return walletService
}
func RegisterService() interfaces.RegisterService {
	return registerService
}
func RpcService() interfaces.RpcService {
	return rpcService
}
func VerificationService() interfaces.VerificationService {
	return verificationService
}
func PaymentProcessor() interfaces.PaymentProcessor {
	return paymentProcessorService
}
