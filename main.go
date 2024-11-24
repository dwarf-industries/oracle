package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	"oracle/commands"
	"oracle/services"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rpc := getRpc()

	var rootCmd = &cobra.Command{Use: "oracle"}
	walletService := &services.WalletService{
		PasswordManager: &services.PasswordManager{},
	}
	rpcService := &services.RpcService{}
	rpcService.SetClient(rpc)

	registerService := &services.RegisterService{
		ContractAddr:  os.Getenv("CONTRACT_ADDRESS"),
		WalletService: walletService,
		RpcService:    rpcService,
		VerificationService: &services.IdentityService{
			WalletService: walletService,
		},
	}
	addWalletcommand := commands.AddWalletCommand{
		WalletService: walletService,
	}
	generateWalletCommand := commands.GenerateWalletCommand{
		WalletService: walletService,
	}
	rpcCommand := commands.SetRpcCommand{
		RpcService: rpcService,
	}
	registerCommand := commands.RegisterCommand{
		WalletService: walletService,

		RegisterService: registerService,
	}
	syncCommand := commands.SyncCommand{
		RegisterService: registerService,
	}

	generateWalletCommand.Execute()

	rootCmd.AddCommand(addWalletcommand.Executable())
	rootCmd.AddCommand(rpcCommand.Executable())
	rootCmd.AddCommand(registerCommand.Executable())
	rootCmd.AddCommand(syncCommand.Executable())
	rootCmd.AddCommand(generateWalletCommand.Executable())

	if err := rootCmd.Execute(); err != nil {
		// fmt.Println(err)
		os.Exit(1)
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
