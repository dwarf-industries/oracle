package main

import (
	"os"

	"github.com/spf13/cobra"

	"oracle/commands"
	"oracle/services"
)

func main() {
	var rootCmd = &cobra.Command{Use: "oracle"}
	walletService := &services.WalletService{
		PasswordManager: &services.PasswordManager{},
	}
	rpcService := &services.RpcService{}
	addWalletcommand := commands.AddWalletCommand{
		WalletService: walletService,
	}
	rpcCommand := commands.SetRpcCommand{
		RpcService: rpcService,
	}
	registerCommand := commands.RegisterCommand{
		WalletService: walletService,
		RegisterService: &services.RegisterService{
			WalletService: walletService,
			RpcService:    rpcService,
			VerificationService: &services.IdentityService{
				WalletService: walletService,
			},
		},
	}
	syncCommand := commands.SyncCommand{}
	rootCmd.AddCommand(addWalletcommand.Executable())
	rootCmd.AddCommand(rpcCommand.Executable())
	rootCmd.AddCommand(registerCommand.Executable())
	rootCmd.AddCommand(syncCommand.Executable())

	if err := rootCmd.Execute(); err != nil {
		// fmt.Println(err)
		os.Exit(1)
	}
}
