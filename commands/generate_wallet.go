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
		Use:   "new-wallet",
		Short: "Generates a new empty wallet",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}

func (g *GenerateWalletCommand) Execute() {
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
	var password string
	fmt.Scanln(&password)
	g.WalletService.SetWallet(&privateKeyHex, &password)

}
