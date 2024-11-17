package main

import (
	"os"

	"github.com/spf13/cobra"

	"oracle/commands"
	"oracle/services"
)

func main() {
	var rootCmd = &cobra.Command{Use: "oracle"}
	addWalletcommand := commands.AddWalletCommand{
		WalletService: &services.WalletService{
			PasswordManager: &services.PasswordManager{},
		},
	}
	rpcCommand := commands.SetRpcCommand{
		RpcService: &services.RpcService{},
	}

	rootCmd.AddCommand(addWalletcommand.Executable())
	rootCmd.AddCommand(rpcCommand.Executable())

	if err := rootCmd.Execute(); err != nil {
		// fmt.Println(err)
		os.Exit(1)
	}
}
