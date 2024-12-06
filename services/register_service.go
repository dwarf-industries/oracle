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

	registerFee, err := r.GetRegistrationFee()
	if err != nil {
		return fmt.Errorf("failed to get registration fee: %v", err)
	}

	auth, err := r.newTransactor(wallet)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %v", err)
	}

	auth.Value = registerFee
	fmt.Println("Value:", auth.Value)
	balance := r.WalletService.GetBalance(&auth.From)
	fmt.Println(balance)
	fmt.Println(balance.Cmp(registerFee))
	tx, err := contract.Register(auth, ip, "0")
	if err != nil {
		fmt.Println(auth.From.String())

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

	wallet, err := r.WalletService.ActiveWallet()
	if err != nil {
		fmt.Println("Failed to retrive credentails")
		return []models.Oracle{}, err
	}
	auth, err := r.newTransactor(wallet)
	if err != nil {
		fmt.Println("Failed to Assign credentails")
		return []models.Oracle{}, err
	}

	var oracles []models.Oracle
	for _, o := range oracleResult {

		if o.Name == auth.From {
			oracles = append(oracles, models.Oracle{
				Name:       o.Name.Hex(),
				Ip:         o.Ip,
				Port:       o.Port,
				Reputation: *o.Reputation,
			})
			continue
		}

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

func (r *RegisterService) Registered() bool {
	contractAddress := common.HexToAddress(r.ContractAddr)
	contract, err := register.NewRegister(contractAddress, r.RpcService.GetClient())
	if err != nil {
		fmt.Println("Failed to load contract:", err)
		return false
	}
	activeWallet, err := r.WalletService.ActiveWallet()
	if err != nil {
		fmt.Println("Wallet is locked or not setup")
		return false
	}
	auth, err := r.newTransactor(activeWallet)
	if err != nil {
		fmt.Println("Failed to setup credentails")
	}

	registered, err := contract.IsOracleRegistered(&bind.CallOpts{
		From: auth.From,
	})

	if err != nil {
		fmt.Println("Smart contract check reverted, contract an administrator.:", err)
		return false
	}
	return registered
}

func (r *RegisterService) GetRegistrationFee() (*big.Int, error) {
	contractAddress := common.HexToAddress(r.ContractAddr)
	contract, err := register.NewRegister(contractAddress, r.RpcService.GetClient())
	if err != nil {
		fmt.Println("Failed to load contract:", err)
		return nil, err
	}

	return contract.GetRegistrationFee(&bind.CallOpts{})
}

func (r *RegisterService) GetReportFee() (*big.Int, error) {
	contractAddress := common.HexToAddress(r.ContractAddr)
	contract, err := register.NewRegister(contractAddress, r.RpcService.GetClient())
	if err != nil {
		fmt.Println("Failed to load contract:", err)
		return nil, err
	}

	return contract.GetReportFee(&bind.CallOpts{})
}
func (r *RegisterService) newTransactor(privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	chainID, err := r.RpcService.GetClient().ChainID(context.Background())
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
		auth.GasPrice = big.NewInt(20000000000)
		fmt.Println("Warning: Failed to suggest gas price, using fallback value")
	}

	block, err := r.RpcService.GetClient().BlockNumber(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to retrive current block: %v", err)
	}
	nonce, err := r.RpcService.GetClient().NonceAt(context.Background(), auth.From, big.NewInt(int64(block)))
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	fmt.Println("Transaction details:")
	fmt.Println("ChainID:", chainID)
	fmt.Println("GasLimit:", auth.GasLimit)
	fmt.Println("GasPrice:", auth.GasPrice)
	fmt.Println("Nonce:", auth.Nonce)
	fmt.Println("Block:", block)

	return auth, nil
}
