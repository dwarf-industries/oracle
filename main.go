package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"

	"oracle/commands"
	"oracle/di"
)

func logo() {
	in := `
     ████████▄     ▄████████    ▄████████    ▄████████ ███▄▄▄▄   ████████▄
     ███   ▀███   ███    ███   ███    ███   ███    ███ ███▀▀▀██▄ ███   ▀███
     ███    ███   ███    █▀    ███    █▀    ███    █▀  ███   ███ ███    ███
     ███    ███  ▄███▄▄▄      ▄███▄▄▄      ▄███▄▄▄     ███   ███ ███    ███
     ███    ███ ▀▀███▀▀▀     ▀▀███▀▀▀     ▀▀███▀▀▀     ███   ███ ███    ███
     ███    ███   ███    █▄    ███          ███    █▄  ███   ███ ███    ███
     ███   ▄███   ███    ███   ███          ███    ███ ███   ███ ███   ▄███
     ████████▀    ██████████   ███          ██████████  ▀█   █▀  ████████▀

     ████████▄     ▄████████ ███▄▄▄▄   ▄██   ▄
     ███   ▀███   ███    ███ ███▀▀▀██▄ ███   ██▄
     ███    ███   ███    █▀  ███   ███ ███▄▄▄███
     ███    ███  ▄███▄▄▄     ███   ███ ▀▀▀▀▀▀███
     ███    ███ ▀▀███▀▀▀     ███   ███ ▄██   ███
     ███    ███   ███    █▄  ███   ███ ███   ███
     ███   ▄███   ███    ███ ███   ███ ███   ███
     ████████▀    ██████████  ▀█   █▀   ▀█████▀

     ████████▄     ▄████████    ▄███████▄  ▄██████▄     ▄████████    ▄████████
     ███   ▀███   ███    ███   ███    ███ ███    ███   ███    ███   ███    ███
     ███    ███   ███    █▀    ███    ███ ███    ███   ███    █▀    ███    █▀
     ███    ███  ▄███▄▄▄       ███    ███ ███    ███   ███         ▄███▄▄▄
     ███    ███ ▀▀███▀▀▀     ▀█████████▀  ███    ███ ▀███████████ ▀▀███▀▀▀
     ███    ███   ███    █▄    ███        ███    ███          ███   ███    █▄
     ███   ▄███   ███    ███   ███        ███    ███    ▄█    ███   ███    ███
     ████████▀    ██████████  ▄████▀       ▀██████▀   ▄████████▀    ██████████

	`

	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(120))

	out, _ := r.Render(in)
	fmt.Print(out)
}

func main() {
	logo()
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
		oracle, err := di.RegisterService().GetOracle()
		if err != nil {
			fmt.Println("failed to shutdown gracefully.")
			fmt.Println("Can't retrive oracle state on chain.")
			fmt.Println(err)
			os.Exit(1)
		}

		if !oracle.ConnectionStatus {
			fmt.Println("Shutting down, node state is already offline on chain")
			os.Exit(0)
		}

		err = di.RegisterService().LogOut()
		if err != nil {
			fmt.Println("failed to shutdown gracefully.")
			fmt.Println("Can't update oracle state on chain.")
			fmt.Println("This is a dangerious situtation, oracles that remain online after being shutdown will get peanalized reach out to the development team!")
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(err)
		os.Exit(1)
	}
}
