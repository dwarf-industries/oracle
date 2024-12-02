package commands

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"

	"oracle/interfaces"
)

type AddWalletCommand struct {
	WalletService   interfaces.WalletService
	RegisterService interfaces.RegisterService
	PasswordManager interfaces.PasswordManager
}

func (aw *AddWalletCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "setup [private key]  ",
		Short: "configure your oracle account",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			privateKey := args[0]
			aw.Execute(&privateKey)
		},
	}
}

func (aw *AddWalletCommand) Execute(wallet *string) {
	password := aw.PasswordManager.Input()
	privateKey, err := aw.WalletService.SetWallet(wallet, password)
	if err != nil {
		fmt.Println("Failed to import wallet!")
		fmt.Println(err)
		os.Exit(1)
		return
	}
	fmt.Println("")
	publicKeyHex := aw.WalletService.GetAddressForPrivateKey(&privateKey)
	privateKeyHex := privateKey.D.Text(16)

	header := color.New(color.FgCyan, color.Bold).SprintFunc()
	fmt.Println(header("🚀 Wallet Imported Successful!"))

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Key Type", "Key Value"})
	t.AppendRow(table.Row{"Public Key", publicKeyHex})
	t.AppendRow(table.Row{"Private Key", privateKeyHex})
	t.Render()

	warning := color.New(color.FgRed, color.Bold).SprintFunc()
	fmt.Println(warning("\n⚠️  IMPORTANT: Keep a secure copy of your private key. It is required for wallet recovery and cannot be retrieved if lost."))

	walletBalance := aw.WalletService.GetBalance((*common.Address)(common.FromHex(publicKeyHex)))

	if walletBalance.Int64() == 0 {
		fmt.Println("\n⚠️  ATTENTION: it appears that you wallet balance is empty, you will need to fund it in order to use it as operations wallet!")
	}

	registerFee, err := aw.RegisterService.GetRegistrationFee()

	if err != nil {
		fmt.Println("Failed to get registration fee, can't estimate if balance is less then the requirment, interactions will likely fail!")
		fmt.Println("This error has nothing to do with your wallet, it's either caused by an issue with the RPC client or the blockchain smart contract")
		return
	}

	if walletBalance.Cmp(registerFee) <= 0 {
		fmt.Println("It appears that your balance is less then the required register fee, if you attempt to register your wallet you will likely have to add more funds.")

		b := table.NewWriter()
		b.SetOutputMirror(os.Stdout)
		b.AppendHeader(table.Row{"Type", "Amount"})
		b.AppendRow(table.Row{"Wallet balance", walletBalance.String()})
		b.AppendRow(table.Row{"Registration Fee", registerFee.String()})
		b.Render()
	}

	fmt.Println("you can learn more about funding and why it's required under 'oracle network-operation-info'")
	infoHeader := color.New(color.BgCyan, color.Bold).SprintFunc()
	fmt.Println(infoHeader("for more information regarding wallets 'oracle wallets-info'"))

	os.Exit(0)
}
