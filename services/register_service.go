package services

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	register "oracle/contracts"
	"oracle/interfaces"
	"oracle/models"
)

type RegisterService struct {
	WalletService       interfaces.WalletService
	RpcService          interfaces.RpcService
	ContractAddr        string
	VerificationService interfaces.IdentityVerificationService
}

func (r *RegisterService) Register(ip string) error {
	wallet, err := r.WalletService.ActiveWallet()
	if err != nil {
		return err
	}

	contractAddress := common.HexToAddress(r.ContractAddr)
	contract, err := register.NewRegister(contractAddress, r.RpcService.GetClient())
	if err != nil {
		return fmt.Errorf("failed to load contract: %v", err)
	}

	auth, err := r.newTransactor(wallet)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %v", err)
	}

	auth.Value = big.NewInt(1000000000000000)

	tx, err := contract.Register(auth, ip, "0")
	if err != nil {
		return fmt.Errorf("failed to execute register function: %v", err)
	}

	fmt.Printf("Register transaction sent: %s\n", tx.Hash().Hex())

	return nil
}

func (r *RegisterService) Oracles() ([]models.Oracle, error) {
	contractAddress := common.HexToAddress(r.ContractAddr)
	contract, err := register.NewRegister(contractAddress, r.RpcService.GetClient())
	if err != nil {
		fmt.Println("Failed to load contract:", err)
		return []models.Oracle{}, err
	}

	oracleResult, err := contract.GetOracles(&bind.CallOpts{})
	if err != nil {
		return []models.Oracle{}, nil
	}

	var oracles []models.Oracle
	for _, o := range oracleResult {
		//Verify that the node is running with the expected public address
		verified := r.VerificationService.Verify(o.Ip, o.Name.Hex())
		if !verified {
			continue
		}

		oracles = append(oracles, models.Oracle{
			Name:       o.Name.Hex(),
			Ip:         o.Ip,
			Port:       o.Port,
			Reputation: *o.Reputation,
		})
	}
	return oracles, nil
}

func (r *RegisterService) newTransactor(privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	chainID, err := r.RpcService.GetClient().NetworkID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get network ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %v", err)
	}

	auth.GasLimit = uint64(300000)
	auth.GasPrice, err = r.RpcService.GetClient().SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to suggest gas price: %v", err)
	}

	return auth, nil
}
