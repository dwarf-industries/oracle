package services

import (
	"errors"
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
	oracles             []models.Oracle
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

	auth, err := r.WalletService.NewTransactor(wallet)
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

	if len(r.oracles) > 0 {
		return r.oracles, nil
	}

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
	auth, err := r.WalletService.NewTransactor(wallet)
	if err != nil {
		fmt.Println("Failed to Assign credentails")
		return []models.Oracle{}, err
	}

	var oracles []models.Oracle
	for _, o := range oracleResult {
		if !o.IsOnline {
			continue
		}

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

	r.oracles = oracles
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
	auth, err := r.WalletService.NewTransactor(activeWallet)
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

func (r *RegisterService) GetOracle() (*models.Oracle, error) {
	contractAddress := common.HexToAddress(r.ContractAddr)
	contract, err := register.NewRegister(contractAddress, r.RpcService.GetClient())
	if err != nil {
		fmt.Println("Failed to load contract:", err)
		return nil, err
	}

	wallet, err := r.WalletService.ActiveWallet()
	if err != nil {
		return nil, err
	}

	transOps, err := r.WalletService.NewTransactor(wallet)
	if err != nil {
		return nil, err
	}

	exists, err := contract.IsOracleRegistered(&bind.CallOpts{
		From: transOps.From,
	})

	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("oracle not doesn't exist")
	}

	oracle, err := contract.Self(&bind.CallOpts{
		From: transOps.From,
	})

	if err != nil {
		return nil, err
	}

	return &models.Oracle{
		Name:             oracle.Name.Hex(),
		Ip:               oracle.Ip,
		Port:             oracle.Ip,
		Reputation:       *oracle.Reputation,
		ConnectionStatus: oracle.IsOnline,
	}, nil
}

func (r *RegisterService) LogOut() error {
	wallet, err := r.WalletService.ActiveWallet()
	if err != nil {
		return err
	}

	contractAddress := common.HexToAddress(r.ContractAddr)
	contract, err := register.NewRegister(contractAddress, r.RpcService.GetClient())
	if err != nil {
		return fmt.Errorf("failed to load contract: %v", err)
	}

	auth, err := r.WalletService.NewTransactor(wallet)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %v", err)
	}

	_, err = contract.LogOut(auth)
	return err
}

func (r *RegisterService) Login() error {
	wallet, err := r.WalletService.ActiveWallet()
	if err != nil {
		return err
	}

	contractAddress := common.HexToAddress(r.ContractAddr)
	contract, err := register.NewRegister(contractAddress, r.RpcService.GetClient())
	if err != nil {
		return fmt.Errorf("failed to load contract: %v", err)
	}

	auth, err := r.WalletService.NewTransactor(wallet)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %v", err)
	}

	_, err = contract.Login(auth)
	return err
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

func (r *RegisterService) AddOracle(oracle common.Address) error {
	wallet, err := r.WalletService.ActiveWallet()
	if err != nil {
		return err
	}

	contractAddress := common.HexToAddress(r.ContractAddr)
	contract, err := register.NewRegister(contractAddress, r.RpcService.GetClient())
	if err != nil {
		return fmt.Errorf("failed to load contract: %v", err)
	}

	auth, err := r.WalletService.NewTransactor(wallet)
	if err != nil {
		return fmt.Errorf("failed to create transactor: %v", err)
	}

	oracleData, err := contract.GetOracle(&bind.CallOpts{
		From: auth.From,
	}, oracle)

	if err != nil {
		return err
	}

	if !oracleData.IsOnline {
		return errors.New("oracle is offline, can't add it")
	}

	r.oracles = append(r.oracles, models.Oracle{
		Name:             oracleData.Name.Hex(),
		Ip:               oracleData.Ip,
		Port:             oracleData.Port,
		Reputation:       *oracleData.Reputation,
		ConnectionStatus: oracleData.IsOnline,
	})

	return nil
}

func (r *RegisterService) RemoveOracle(current common.Address) error {
	var filtered []models.Oracle
	for _, oracle := range r.oracles {
		if oracle.Name != current.Hex() {
			filtered = append(filtered, oracle)
		}
	}
	r.oracles = filtered
	return nil
}
