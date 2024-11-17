package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"oracle/interfaces"
)

type SetRpcCommand struct {
	RpcService interfaces.RpcService
}

func (s *SetRpcCommand) Executable() *cobra.Command {
	return &cobra.Command{
		Use:   "use-rpc [rpc url]",
		Short: "configure custom RPC to connect to the blockchain",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			rpc := args[0]
			s.Execute(&rpc)
		},
	}
}

func (s *SetRpcCommand) Execute(rpc *string) {
	err := s.RpcService.OverrideClient(rpc)
	if err != nil {
		fmt.Printf("Failed to override existing client, error: %v", err)
	}
}
