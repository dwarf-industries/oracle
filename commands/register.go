package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"oracle/interfaces"
)

type RegisterCommand struct {
	RegisterService interfaces.RegisterService
	WalletService   interfaces.WalletService
	PasswordManager interfaces.PasswordManager
}

func (r *RegisterCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "register [domain]",
		Short: "register the oracle with the blockchain",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			domain := args[0]
			r.Execute(&domain)
		},
	}
}

func (r *RegisterCommand) Execute(domain *string) {
	password := r.PasswordManager.Input()
	_, err := r.WalletService.GetWallet(password)
	if err != nil {
		fmt.Println("Failed to unlock wallet, either wrong password or not set.", err)
		return
	}

	_, err = r.WalletService.ActiveWallet()
	if err != nil {
		fmt.Println("Operations wallet is not set, can't register oracle.")
		fmt.Println("Use 'oracle add-wallet -h' for more details")
		return
	}

	err = r.RegisterService.Register(*domain)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to register oracle")
		return
	}
}
