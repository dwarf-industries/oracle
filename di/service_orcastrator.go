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
var identityService interfaces.IdentityVerificationService
var passwordManager interfaces.PasswordManager
var shutdownService interfaces.ShutdownService

func SetupServices() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rpc := getRpc()
	rpcService = &services.RpcService{}
	rpcService.SetClient(rpc)
	passwordManager = &services.PasswordManager{}
	walletService = &services.WalletService{
		PasswordManager: passwordManager,
		RpcService:      rpcService,
	}
	identityService = &services.IdentityService{
		WalletService: walletService,
	}
	registerService = &services.RegisterService{
		ContractAddr:        os.Getenv("CONTRACT_ADDRESS"),
		WalletService:       walletService,
		RpcService:          rpcService,
		VerificationService: identityService,
	}
	verificationService = &services.VerificationService{}
	verificationService.Init()
	paymentProcessorService = &services.PaymentProcessor{
		RpcService:    rpcService,
		WalletService: walletService,
	}
	paymentProcessorService.Init()
	shutdownService = &services.ShutdownService{
		RegisterService: registerService,
	}
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
func GetIdentityService() interfaces.IdentityVerificationService {
	return identityService
}
func GetPasswordManager() interfaces.PasswordManager {
	return passwordManager
}
func GetShutdownService() interfaces.ShutdownService {
	return shutdownService
}
