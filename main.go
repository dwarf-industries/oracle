package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"oracle/commands"
	"oracle/di"
)

func main() {
	di.SetupServices()
	var rootCmd = &cobra.Command{
		Use: "oracle",
	}

	addWalletcommand := commands.AddWalletCommand{
		WalletService:   di.WalletService(),
		RegisterService: di.RegisterService(),
		PasswordManager: di.GetPasswordManager(),
	}
	generateWalletCommand := commands.GenerateWalletCommand{
		WalletService: di.WalletService(),
	}
	rpcCommand := commands.SetRpcCommand{
		RpcService: di.RpcService(),
	}
	registerCommand := commands.RegisterCommand{
		WalletService:   di.WalletService(),
		RegisterService: di.RegisterService(),
		PasswordManager: di.GetPasswordManager(),
	}
	syncCommand := commands.SyncCommand{
		RegisterService:     di.RegisterService(),
		WalletService:       di.WalletService(),
		VerificationService: di.VerificationService(),
		IdentityService:     di.GetIdentityService(),
		PasswordManager:     di.GetPasswordManager(),
	}

	rootCmd.AddCommand(addWalletcommand.Executable())
	rootCmd.AddCommand(rpcCommand.Executable())
	rootCmd.AddCommand(registerCommand.Executable())
	rootCmd.AddCommand(syncCommand.Executable())
	rootCmd.AddCommand(generateWalletCommand.Executable())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
