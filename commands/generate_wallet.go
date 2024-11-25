package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"oracle/interfaces"
)

type GenerateWalletCommand struct {
	WalletService interfaces.WalletService
}

func (g *GenerateWalletCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "new-wallet [password]",
		Short: "Generates a new empty wallet",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			password := args[0]
			g.Execute(&password)
		},
	}
}

func (g *GenerateWalletCommand) Execute(password *string) {
	wallet, err := g.WalletService.NewWallet()
	if err != nil {
		fmt.Println("Failed to create new wallet, please submit an issue or try again!")
		fmt.Println(err)
	}

	publicKeyHex := g.WalletService.GetAddressForPrivateKey(wallet)
	fmt.Println("Your public address: ", publicKeyHex)
	fmt.Println("Make sure to backup of your private key")

	privateKeyHex := wallet.D.Text(16)

	fmt.Println("Private key (raw hex):", privateKeyHex)

	fmt.Println("Please enter a password to protect the wallet for future use")

	created := g.WalletService.SetWallet(&privateKeyHex, password)

	if !created {
		fmt.Println("Failed to save wallet!")
	}
}
