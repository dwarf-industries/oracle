package commands

import (
	"log"

	"github.com/spf13/cobra"

	"oracle/interfaces"
)

type AddWalletCommand struct {
	WalletService interfaces.WalletService
}

func (aw *AddWalletCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "setup [private key] [password]",
		Short: "configure your oracle account",
		Args:  cobra.MaximumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				log.Fatalf("Parameters are required, expected 3 got %v", len(args))
			}
			privateKey := args[0]
			password := args[1]

			aw.Execute(&privateKey, &password)
		},
	}
}

func (aw *AddWalletCommand) Execute(wallet *string, password *string) {
	aw.WalletService.SetWallet(wallet, password)
}
