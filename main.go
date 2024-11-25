package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"oracle/commands"
)

func main() {
	setupServices()
	var rootCmd = &cobra.Command{
		Use: "oracle",
	}
	addWalletcommand := commands.AddWalletCommand{
		WalletService:   walletService,
		RegisterService: registerService,
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
