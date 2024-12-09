package services

import (
	"fmt"
	"os"

	"oracle/interfaces"
)

type ShutdownService struct {
	RegisterService interfaces.RegisterService
}

func (s *ShutdownService) ShutdownServer() {
	oracle, err := s.RegisterService.GetOracle()
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

	err = s.RegisterService.LogOut()
	if err != nil {
		fmt.Println("failed to shutdown gracefully.")
		fmt.Println("Can't update oracle state on chain.")
		fmt.Println("This is a dangerious situtation, oracles that remain online after being shutdown will get peanalized reach out to the development team!")
		fmt.Println(err)
		os.Exit(1)
	}

}
