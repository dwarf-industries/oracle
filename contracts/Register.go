// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// Oracle is an auto generated low-level Go binding around an user-defined struct.
type Oracle struct {
	Name       common.Address
	Ip         string
	Port       string
	Reputation *big.Int
	IsOnline   bool
}

// RegisterMetaData contains all meta data concerning the Register contract.
var RegisterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_registerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reportFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_feeSetter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daoOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"blacklistedOracle\",\"type\":\"address\"}],\"name\":\"OracleBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"}],\"name\":\"OracleOffline\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"}],\"name\":\"OracleOnline\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTax\",\"type\":\"uint256\"}],\"name\":\"OracleRegisterTaxUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"}],\"name\":\"OracleRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reportedOracle\",\"type\":\"address\"}],\"name\":\"OracleReported\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LogOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Login\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tax\",\"type\":\"uint256\"}],\"name\":\"changeRegisterTax\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"getOracle\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"name\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"reputation\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOnline\",\"type\":\"bool\"}],\"internalType\":\"structOracle\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracles\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"name\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"reputation\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOnline\",\"type\":\"bool\"}],\"internalType\":\"structOracle[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegistrationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReportFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOracleRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokeOwnership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"self\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"name\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"reputation\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOnline\",\"type\":\"bool\"}],\"internalType\":\"structOracle\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"updateRegistrationFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"updateReportFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260056002555f600560146101000a81548160ff02191690831515021790555034801561002e575f5ffd5b5060405161279338038061279383398181016040528101906100509190610177565b8360038190555082600481905550815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050506101db565b5f5ffd5b5f819050919050565b6100fc816100ea565b8114610106575f5ffd5b50565b5f81519050610117816100f3565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101468261011d565b9050919050565b6101568161013c565b8114610160575f5ffd5b50565b5f815190506101718161014d565b92915050565b5f5f5f5f6080858703121561018f5761018e6100e6565b5b5f61019c87828801610109565b94505060206101ad87828801610109565b93505060406101be87828801610163565b92505060606101cf87828801610163565b91505092959194509250565b6125ab806101e85f395ff3fe6080604052600436106100c1575f3560e01c806369366c481161007e5780639a6bea18116100585780639a6bea181461027b578063a208a898146102a5578063af582c6b146102cf578063f00ac1da1461030b576100c1565b806369366c48146101d95780637104ddb2146102155780637571bdd01461023f576100c1565b80630946e807146100c55780630d1db1f0146100ef57806310d3d22e146101195780632b968958146101555780633ffbd47f1461017f57806340884c52146101af575b5f5ffd5b3480156100d0575f5ffd5b506100d9610335565b6040516100e691906118c7565b60405180910390f35b3480156100fa575f5ffd5b5061010361033e565b60405161011091906118fa565b60405180910390f35b348015610124575f5ffd5b5061013f600480360381019061013a919061197e565b610406565b60405161014c9190611ac0565b60405180910390f35b348015610160575f5ffd5b50610169610799565b60405161017691906118fa565b60405180910390f35b61019960048036038101906101949190611c0c565b61089b565b6040516101a691906118fa565b60405180910390f35b3480156101ba575f5ffd5b506101c3610bd8565b6040516101d09190611db7565b60405180910390f35b3480156101e4575f5ffd5b506101ff60048036038101906101fa9190611e01565b610dcc565b60405161020c91906118fa565b60405180910390f35b348015610220575f5ffd5b50610229610ebc565b6040516102369190611ac0565b60405180910390f35b34801561024a575f5ffd5b5061026560048036038101906102609190611e01565b61124d565b60405161027291906118fa565b60405180910390f35b348015610286575f5ffd5b5061028f611374565b60405161029c91906118c7565b60405180910390f35b3480156102b0575f5ffd5b506102b961137d565b6040516102c691906118fa565b60405180910390f35b3480156102da575f5ffd5b506102f560048036038101906102f09190611e01565b61157c565b60405161030291906118fa565b60405180910390f35b348015610316575f5ffd5b5061031f61166c565b60405161032c91906118fa565b60405180910390f35b5f600354905090565b5f5f60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490506001805490508110801561040057503373ffffffffffffffffffffffffffffffffffffffff16600182815481106103ba576103b9611e2c565b5b905f5260205f2090600502015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b91505090565b61040e61186b565b5f60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490506001805490508110610496576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161048d90611eb3565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16600182815481106104c1576104c0611e2c565b5b905f5260205f2090600502015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610545576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053c90611f1b565b60405180910390fd5b5f60085f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905060018054905081106105cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105c490611f83565b60405180910390fd5b600181815481106105e1576105e0611e2c565b5b905f5260205f2090600502016040518060a00160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461065c90611fce565b80601f016020809104026020016040519081016040528092919081815260200182805461068890611fce565b80156106d35780601f106106aa576101008083540402835291602001916106d3565b820191905f5260205f20905b8154815290600101906020018083116106b657829003601f168201915b505050505081526020016002820180546106ec90611fce565b80601f016020809104026020016040519081016040528092919081815260200182805461071890611fce565b80156107635780601f1061073a57610100808354040283529160200191610763565b820191905f5260205f20905b81548152906001019060200180831161074657829003601f168201915b5050505050815260200160038201548152602001600482015f9054906101000a900460ff16151515158152505092505050919050565b5f600560149054906101000a900460ff16156107ea576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107e19061206e565b60405180910390fd5b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610879576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610870906120d6565b60405180910390fd5b6001600560146101000a81548160ff0219169083151502179055506001905090565b5f6003543410156108e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d89061213e565b60405180910390fd5b60016040518060a001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020016001815260200160011515815250908060018154018082558091505060019003905f5260205f2090600502015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010190816109a391906122fc565b5060408201518160020190816109b991906122fc565b50606082015181600301556080820151816004015f6101000a81548160ff0219169083151502179055505050600180805490506109f691906123f8565b60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505f6064600254600354610a4a919061242b565b610a549190612499565b90505f8134610a6391906123f8565b90505f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318a9055c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610acf573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610af391906124dd565b90508073ffffffffffffffffffffffffffffffffffffffff166108fc8490811502906040515f60405180830381858888f19350505050158015610b38573d5f5f3e3d5ffd5b5081600a5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055503373ffffffffffffffffffffffffffffffffffffffff167f8c1c8966369199dc710cd8b615bf16a86520dd165679cb74aa5f427d2bf8c0c98787604051610bc3929190612540565b60405180910390a26001935050505092915050565b60606001805480602002602001604051908101604052809291908181526020015f905b82821015610dc3578382905f5260205f2090600502016040518060a00160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182018054610c8090611fce565b80601f0160208091040260200160405190810160405280929190818152602001828054610cac90611fce565b8015610cf75780601f10610cce57610100808354040283529160200191610cf7565b820191905f5260205f20905b815481529060010190602001808311610cda57829003601f168201915b50505050508152602001600282018054610d1090611fce565b80601f0160208091040260200160405190810160405280929190818152602001828054610d3c90611fce565b8015610d875780601f10610d5e57610100808354040283529160200191610d87565b820191905f5260205f20905b815481529060010190602001808311610d6a57829003601f168201915b5050505050815260200160038201548152602001600482015f9054906101000a900460ff16151515158152505081526020019060010190610bfb565b50505050905090565b5f600560149054906101000a900460ff1615610e1d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e149061206e565b60405180910390fd5b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610eac576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ea3906120d6565b60405180910390fd5b8160048190555060019050919050565b610ec461186b565b5f60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490506001805490508110610f4c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f4390611eb3565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff1660018281548110610f7757610f76611e2c565b5b905f5260205f2090600502015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610ffb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ff290611f1b565b60405180910390fd5b5f60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490506001805490508110611083576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161107a90611f83565b60405180910390fd5b6001818154811061109757611096611e2c565b5b905f5260205f2090600502016040518060a00160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461111290611fce565b80601f016020809104026020016040519081016040528092919081815260200182805461113e90611fce565b80156111895780601f1061116057610100808354040283529160200191611189565b820191905f5260205f20905b81548152906001019060200180831161116c57829003601f168201915b505050505081526020016002820180546111a290611fce565b80601f01602080910402602001604051908101604052809291908181526020018280546111ce90611fce565b80156112195780601f106111f057610100808354040283529160200191611219565b820191905f5260205f20905b8154815290600101906020018083116111fc57829003601f168201915b5050505050815260200160038201548152602001600482015f9054906101000a900460ff1615151515815250509250505090565b5f600560149054906101000a900460ff161561129e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112959061206e565b60405180910390fd5b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461132d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611324906120d6565b60405180910390fd5b816002819055507fa288a9ac9f74fe4c85e8263739645d49297e82ffa010aac11899ba671a54d1ac8260405161136391906118c7565b60405180910390a160019050919050565b5f600454905090565b5f5f60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490506001805490508110611406576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113fd90611eb3565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166001828154811061143157611430611e2c565b5b905f5260205f2090600502015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146114b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114ac90611f1b565b60405180910390fd5b5f60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f6001828154811061150b5761150a611e2c565b5b905f5260205f2090600502016004015f6101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f6eccb14325d715b33601268fe5c7d8a1618261e4d242e890c9e2e372adaf1fbc60405160405180910390a260019250505090565b5f600560149054906101000a900460ff16156115cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115c49061206e565b60405180910390fd5b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461165c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611653906120d6565b60405180910390fd5b8160038190555060019050919050565b5f5f60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905060018054905081106116f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116ec90611eb3565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16600182815481106117205761171f611e2c565b5b905f5260205f2090600502015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146117a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161179b90611f1b565b60405180910390fd5b5f60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905060018082815481106117fa576117f9611e2c565b5b905f5260205f2090600502016004015f6101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f3a21ecfcdc410d1e3552f24e84bbf13af57e47436f3b9ac5045ae9f3adcdc41160405160405180910390a260019250505090565b6040518060a001604052805f73ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001606081526020015f81526020015f151581525090565b5f819050919050565b6118c1816118af565b82525050565b5f6020820190506118da5f8301846118b8565b92915050565b5f8115159050919050565b6118f4816118e0565b82525050565b5f60208201905061190d5f8301846118eb565b92915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61194d82611924565b9050919050565b61195d81611943565b8114611967575f5ffd5b50565b5f8135905061197881611954565b92915050565b5f602082840312156119935761199261191c565b5b5f6119a08482850161196a565b91505092915050565b6119b281611943565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6119fa826119b8565b611a0481856119c2565b9350611a148185602086016119d2565b611a1d816119e0565b840191505092915050565b611a31816118af565b82525050565b611a40816118e0565b82525050565b5f60a083015f830151611a5b5f8601826119a9565b5060208301518482036020860152611a7382826119f0565b91505060408301518482036040860152611a8d82826119f0565b9150506060830151611aa26060860182611a28565b506080830151611ab56080860182611a37565b508091505092915050565b5f6020820190508181035f830152611ad88184611a46565b905092915050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611b1e826119e0565b810181811067ffffffffffffffff82111715611b3d57611b3c611ae8565b5b80604052505050565b5f611b4f611913565b9050611b5b8282611b15565b919050565b5f67ffffffffffffffff821115611b7a57611b79611ae8565b5b611b83826119e0565b9050602081019050919050565b828183375f83830152505050565b5f611bb0611bab84611b60565b611b46565b905082815260208101848484011115611bcc57611bcb611ae4565b5b611bd7848285611b90565b509392505050565b5f82601f830112611bf357611bf2611ae0565b5b8135611c03848260208601611b9e565b91505092915050565b5f5f60408385031215611c2257611c2161191c565b5b5f83013567ffffffffffffffff811115611c3f57611c3e611920565b5b611c4b85828601611bdf565b925050602083013567ffffffffffffffff811115611c6c57611c6b611920565b5b611c7885828601611bdf565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f60a083015f830151611cc05f8601826119a9565b5060208301518482036020860152611cd882826119f0565b91505060408301518482036040860152611cf282826119f0565b9150506060830151611d076060860182611a28565b506080830151611d1a6080860182611a37565b508091505092915050565b5f611d308383611cab565b905092915050565b5f602082019050919050565b5f611d4e82611c82565b611d588185611c8c565b935083602082028501611d6a85611c9c565b805f5b85811015611da55784840389528151611d868582611d25565b9450611d9183611d38565b925060208a01995050600181019050611d6d565b50829750879550505050505092915050565b5f6020820190508181035f830152611dcf8184611d44565b905092915050565b611de0816118af565b8114611dea575f5ffd5b50565b5f81359050611dfb81611dd7565b92915050565b5f60208284031215611e1657611e1561191c565b5b5f611e2384828501611ded565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82825260208201905092915050565b7f4e6f7420616e20617574686f72697a6564206f7261636c6500000000000000005f82015250565b5f611e9d601883611e59565b9150611ea882611e69565b602082019050919050565b5f6020820190508181035f830152611eca81611e91565b9050919050565b7f4f7261636c65206d69736d6174636820617420696e64657800000000000000005f82015250565b5f611f05601883611e59565b9150611f1082611ed1565b602082019050919050565b5f6020820190508181035f830152611f3281611ef9565b9050919050565b7f496e76616c6964206f7261636c6520696e6465780000000000000000000000005f82015250565b5f611f6d601483611e59565b9150611f7882611f39565b602082019050919050565b5f6020820190508181035f830152611f9a81611f61565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611fe557607f821691505b602082108103611ff857611ff7611fa1565b5b50919050565b7f4f776e657273686970206973207265766f6b65642c206e6f2061646d696e69735f8201527f74726174697665206368616e67657320616c6c6f776564210000000000000000602082015250565b5f612058603883611e59565b915061206382611ffe565b604082019050919050565b5f6020820190508181035f8301526120858161204c565b9050919050565b7f4e6f7420617574686f72697a65640000000000000000000000000000000000005f82015250565b5f6120c0600e83611e59565b91506120cb8261208c565b602082019050919050565b5f6020820190508181035f8301526120ed816120b4565b9050919050565b7f496e73756666696369656e7420726567697374726174696f6e206665650000005f82015250565b5f612128601d83611e59565b9150612133826120f4565b602082019050919050565b5f6020820190508181035f8301526121558161211c565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026121b87fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261217d565b6121c2868361217d565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6121fd6121f86121f3846118af565b6121da565b6118af565b9050919050565b5f819050919050565b612216836121e3565b61222a61222282612204565b848454612189565b825550505050565b5f5f905090565b612241612232565b61224c81848461220d565b505050565b5b8181101561226f576122645f82612239565b600181019050612252565b5050565b601f8211156122b4576122858161215c565b61228e8461216e565b8101602085101561229d578190505b6122b16122a98561216e565b830182612251565b50505b505050565b5f82821c905092915050565b5f6122d45f19846008026122b9565b1980831691505092915050565b5f6122ec83836122c5565b9150826002028217905092915050565b612305826119b8565b67ffffffffffffffff81111561231e5761231d611ae8565b5b6123288254611fce565b612333828285612273565b5f60209050601f831160018114612364575f8415612352578287015190505b61235c85826122e1565b8655506123c3565b601f1984166123728661215c565b5f5b8281101561239957848901518255600182019150602085019450602081019050612374565b868310156123b657848901516123b2601f8916826122c5565b8355505b6001600288020188555050505b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f612402826118af565b915061240d836118af565b9250828203905081811115612425576124246123cb565b5b92915050565b5f612435826118af565b9150612440836118af565b925082820261244e816118af565b91508282048414831517612465576124646123cb565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6124a3826118af565b91506124ae836118af565b9250826124be576124bd61246c565b5b828204905092915050565b5f815190506124d781611954565b92915050565b5f602082840312156124f2576124f161191c565b5b5f6124ff848285016124c9565b91505092915050565b5f612512826119b8565b61251c8185611e59565b935061252c8185602086016119d2565b612535816119e0565b840191505092915050565b5f6040820190508181035f8301526125588185612508565b9050818103602083015261256c8184612508565b9050939250505056fea26469706673582212209c7a571b79cd8c405bf666e4075870e26c225996e5816520051fd96ed39370e764736f6c634300081c0033",
}

// RegisterABI is the input ABI used to generate the binding from.
// Deprecated: Use RegisterMetaData.ABI instead.
var RegisterABI = RegisterMetaData.ABI

// RegisterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RegisterMetaData.Bin instead.
var RegisterBin = RegisterMetaData.Bin

// DeployRegister deploys a new Ethereum contract, binding an instance of Register to it.
func DeployRegister(auth *bind.TransactOpts, backend bind.ContractBackend, _registerFee *big.Int, _reportFee *big.Int, _feeSetter common.Address, daoOwner common.Address) (common.Address, *types.Transaction, *Register, error) {
	parsed, err := RegisterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RegisterBin), backend, _registerFee, _reportFee, _feeSetter, daoOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Register{RegisterCaller: RegisterCaller{contract: contract}, RegisterTransactor: RegisterTransactor{contract: contract}, RegisterFilterer: RegisterFilterer{contract: contract}}, nil
}

// Register is an auto generated Go binding around an Ethereum contract.
type Register struct {
	RegisterCaller     // Read-only binding to the contract
	RegisterTransactor // Write-only binding to the contract
	RegisterFilterer   // Log filterer for contract events
}

// RegisterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegisterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegisterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegisterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegisterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegisterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegisterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegisterSession struct {
	Contract     *Register         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegisterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegisterCallerSession struct {
	Contract *RegisterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegisterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegisterTransactorSession struct {
	Contract     *RegisterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegisterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegisterRaw struct {
	Contract *Register // Generic contract binding to access the raw methods on
}

// RegisterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegisterCallerRaw struct {
	Contract *RegisterCaller // Generic read-only contract binding to access the raw methods on
}

// RegisterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegisterTransactorRaw struct {
	Contract *RegisterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegister creates a new instance of Register, bound to a specific deployed contract.
func NewRegister(address common.Address, backend bind.ContractBackend) (*Register, error) {
	contract, err := bindRegister(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Register{RegisterCaller: RegisterCaller{contract: contract}, RegisterTransactor: RegisterTransactor{contract: contract}, RegisterFilterer: RegisterFilterer{contract: contract}}, nil
}

// NewRegisterCaller creates a new read-only instance of Register, bound to a specific deployed contract.
func NewRegisterCaller(address common.Address, caller bind.ContractCaller) (*RegisterCaller, error) {
	contract, err := bindRegister(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegisterCaller{contract: contract}, nil
}

// NewRegisterTransactor creates a new write-only instance of Register, bound to a specific deployed contract.
func NewRegisterTransactor(address common.Address, transactor bind.ContractTransactor) (*RegisterTransactor, error) {
	contract, err := bindRegister(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegisterTransactor{contract: contract}, nil
}

// NewRegisterFilterer creates a new log filterer instance of Register, bound to a specific deployed contract.
func NewRegisterFilterer(address common.Address, filterer bind.ContractFilterer) (*RegisterFilterer, error) {
	contract, err := bindRegister(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegisterFilterer{contract: contract}, nil
}

// bindRegister binds a generic wrapper to an already deployed contract.
func bindRegister(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegisterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Register *RegisterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Register.Contract.RegisterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Register *RegisterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Register.Contract.RegisterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Register *RegisterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Register.Contract.RegisterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Register *RegisterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Register.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Register *RegisterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Register.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Register *RegisterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Register.Contract.contract.Transact(opts, method, params...)
}

// GetOracle is a free data retrieval call binding the contract method 0x10d3d22e.
//
// Solidity: function getOracle(address _oracle) view returns((address,string,string,uint256,bool))
func (_Register *RegisterCaller) GetOracle(opts *bind.CallOpts, _oracle common.Address) (Oracle, error) {
	var out []interface{}
	err := _Register.contract.Call(opts, &out, "getOracle", _oracle)

	if err != nil {
		return *new(Oracle), err
	}

	out0 := *abi.ConvertType(out[0], new(Oracle)).(*Oracle)

	return out0, err

}

// GetOracle is a free data retrieval call binding the contract method 0x10d3d22e.
//
// Solidity: function getOracle(address _oracle) view returns((address,string,string,uint256,bool))
func (_Register *RegisterSession) GetOracle(_oracle common.Address) (Oracle, error) {
	return _Register.Contract.GetOracle(&_Register.CallOpts, _oracle)
}

// GetOracle is a free data retrieval call binding the contract method 0x10d3d22e.
//
// Solidity: function getOracle(address _oracle) view returns((address,string,string,uint256,bool))
func (_Register *RegisterCallerSession) GetOracle(_oracle common.Address) (Oracle, error) {
	return _Register.Contract.GetOracle(&_Register.CallOpts, _oracle)
}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns((address,string,string,uint256,bool)[])
func (_Register *RegisterCaller) GetOracles(opts *bind.CallOpts) ([]Oracle, error) {
	var out []interface{}
	err := _Register.contract.Call(opts, &out, "getOracles")

	if err != nil {
		return *new([]Oracle), err
	}

	out0 := *abi.ConvertType(out[0], new([]Oracle)).(*[]Oracle)

	return out0, err

}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns((address,string,string,uint256,bool)[])
func (_Register *RegisterSession) GetOracles() ([]Oracle, error) {
	return _Register.Contract.GetOracles(&_Register.CallOpts)
}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns((address,string,string,uint256,bool)[])
func (_Register *RegisterCallerSession) GetOracles() ([]Oracle, error) {
	return _Register.Contract.GetOracles(&_Register.CallOpts)
}

// GetRegistrationFee is a free data retrieval call binding the contract method 0x0946e807.
//
// Solidity: function getRegistrationFee() view returns(uint256)
func (_Register *RegisterCaller) GetRegistrationFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Register.contract.Call(opts, &out, "getRegistrationFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRegistrationFee is a free data retrieval call binding the contract method 0x0946e807.
//
// Solidity: function getRegistrationFee() view returns(uint256)
func (_Register *RegisterSession) GetRegistrationFee() (*big.Int, error) {
	return _Register.Contract.GetRegistrationFee(&_Register.CallOpts)
}

// GetRegistrationFee is a free data retrieval call binding the contract method 0x0946e807.
//
// Solidity: function getRegistrationFee() view returns(uint256)
func (_Register *RegisterCallerSession) GetRegistrationFee() (*big.Int, error) {
	return _Register.Contract.GetRegistrationFee(&_Register.CallOpts)
}

// GetReportFee is a free data retrieval call binding the contract method 0x9a6bea18.
//
// Solidity: function getReportFee() view returns(uint256)
func (_Register *RegisterCaller) GetReportFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Register.contract.Call(opts, &out, "getReportFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReportFee is a free data retrieval call binding the contract method 0x9a6bea18.
//
// Solidity: function getReportFee() view returns(uint256)
func (_Register *RegisterSession) GetReportFee() (*big.Int, error) {
	return _Register.Contract.GetReportFee(&_Register.CallOpts)
}

// GetReportFee is a free data retrieval call binding the contract method 0x9a6bea18.
//
// Solidity: function getReportFee() view returns(uint256)
func (_Register *RegisterCallerSession) GetReportFee() (*big.Int, error) {
	return _Register.Contract.GetReportFee(&_Register.CallOpts)
}

// IsOracleRegistered is a free data retrieval call binding the contract method 0x0d1db1f0.
//
// Solidity: function isOracleRegistered() view returns(bool)
func (_Register *RegisterCaller) IsOracleRegistered(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Register.contract.Call(opts, &out, "isOracleRegistered")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOracleRegistered is a free data retrieval call binding the contract method 0x0d1db1f0.
//
// Solidity: function isOracleRegistered() view returns(bool)
func (_Register *RegisterSession) IsOracleRegistered() (bool, error) {
	return _Register.Contract.IsOracleRegistered(&_Register.CallOpts)
}

// IsOracleRegistered is a free data retrieval call binding the contract method 0x0d1db1f0.
//
// Solidity: function isOracleRegistered() view returns(bool)
func (_Register *RegisterCallerSession) IsOracleRegistered() (bool, error) {
	return _Register.Contract.IsOracleRegistered(&_Register.CallOpts)
}

// Self is a free data retrieval call binding the contract method 0x7104ddb2.
//
// Solidity: function self() view returns((address,string,string,uint256,bool))
func (_Register *RegisterCaller) Self(opts *bind.CallOpts) (Oracle, error) {
	var out []interface{}
	err := _Register.contract.Call(opts, &out, "self")

	if err != nil {
		return *new(Oracle), err
	}

	out0 := *abi.ConvertType(out[0], new(Oracle)).(*Oracle)

	return out0, err

}

// Self is a free data retrieval call binding the contract method 0x7104ddb2.
//
// Solidity: function self() view returns((address,string,string,uint256,bool))
func (_Register *RegisterSession) Self() (Oracle, error) {
	return _Register.Contract.Self(&_Register.CallOpts)
}

// Self is a free data retrieval call binding the contract method 0x7104ddb2.
//
// Solidity: function self() view returns((address,string,string,uint256,bool))
func (_Register *RegisterCallerSession) Self() (Oracle, error) {
	return _Register.Contract.Self(&_Register.CallOpts)
}

// LogOut is a paid mutator transaction binding the contract method 0xa208a898.
//
// Solidity: function LogOut() returns(bool)
func (_Register *RegisterTransactor) LogOut(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Register.contract.Transact(opts, "LogOut")
}

// LogOut is a paid mutator transaction binding the contract method 0xa208a898.
//
// Solidity: function LogOut() returns(bool)
func (_Register *RegisterSession) LogOut() (*types.Transaction, error) {
	return _Register.Contract.LogOut(&_Register.TransactOpts)
}

// LogOut is a paid mutator transaction binding the contract method 0xa208a898.
//
// Solidity: function LogOut() returns(bool)
func (_Register *RegisterTransactorSession) LogOut() (*types.Transaction, error) {
	return _Register.Contract.LogOut(&_Register.TransactOpts)
}

// Login is a paid mutator transaction binding the contract method 0xf00ac1da.
//
// Solidity: function Login() returns(bool)
func (_Register *RegisterTransactor) Login(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Register.contract.Transact(opts, "Login")
}

// Login is a paid mutator transaction binding the contract method 0xf00ac1da.
//
// Solidity: function Login() returns(bool)
func (_Register *RegisterSession) Login() (*types.Transaction, error) {
	return _Register.Contract.Login(&_Register.TransactOpts)
}

// Login is a paid mutator transaction binding the contract method 0xf00ac1da.
//
// Solidity: function Login() returns(bool)
func (_Register *RegisterTransactorSession) Login() (*types.Transaction, error) {
	return _Register.Contract.Login(&_Register.TransactOpts)
}

// ChangeRegisterTax is a paid mutator transaction binding the contract method 0x7571bdd0.
//
// Solidity: function changeRegisterTax(uint256 tax) returns(bool)
func (_Register *RegisterTransactor) ChangeRegisterTax(opts *bind.TransactOpts, tax *big.Int) (*types.Transaction, error) {
	return _Register.contract.Transact(opts, "changeRegisterTax", tax)
}

// ChangeRegisterTax is a paid mutator transaction binding the contract method 0x7571bdd0.
//
// Solidity: function changeRegisterTax(uint256 tax) returns(bool)
func (_Register *RegisterSession) ChangeRegisterTax(tax *big.Int) (*types.Transaction, error) {
	return _Register.Contract.ChangeRegisterTax(&_Register.TransactOpts, tax)
}

// ChangeRegisterTax is a paid mutator transaction binding the contract method 0x7571bdd0.
//
// Solidity: function changeRegisterTax(uint256 tax) returns(bool)
func (_Register *RegisterTransactorSession) ChangeRegisterTax(tax *big.Int) (*types.Transaction, error) {
	return _Register.Contract.ChangeRegisterTax(&_Register.TransactOpts, tax)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string ip, string port) payable returns(bool)
func (_Register *RegisterTransactor) Register(opts *bind.TransactOpts, ip string, port string) (*types.Transaction, error) {
	return _Register.contract.Transact(opts, "register", ip, port)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string ip, string port) payable returns(bool)
func (_Register *RegisterSession) Register(ip string, port string) (*types.Transaction, error) {
	return _Register.Contract.Register(&_Register.TransactOpts, ip, port)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string ip, string port) payable returns(bool)
func (_Register *RegisterTransactorSession) Register(ip string, port string) (*types.Transaction, error) {
	return _Register.Contract.Register(&_Register.TransactOpts, ip, port)
}

// RevokeOwnership is a paid mutator transaction binding the contract method 0x2b968958.
//
// Solidity: function revokeOwnership() returns(bool)
func (_Register *RegisterTransactor) RevokeOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Register.contract.Transact(opts, "revokeOwnership")
}

// RevokeOwnership is a paid mutator transaction binding the contract method 0x2b968958.
//
// Solidity: function revokeOwnership() returns(bool)
func (_Register *RegisterSession) RevokeOwnership() (*types.Transaction, error) {
	return _Register.Contract.RevokeOwnership(&_Register.TransactOpts)
}

// RevokeOwnership is a paid mutator transaction binding the contract method 0x2b968958.
//
// Solidity: function revokeOwnership() returns(bool)
func (_Register *RegisterTransactorSession) RevokeOwnership() (*types.Transaction, error) {
	return _Register.Contract.RevokeOwnership(&_Register.TransactOpts)
}

// UpdateRegistrationFee is a paid mutator transaction binding the contract method 0xaf582c6b.
//
// Solidity: function updateRegistrationFee(uint256 fee) returns(bool)
func (_Register *RegisterTransactor) UpdateRegistrationFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Register.contract.Transact(opts, "updateRegistrationFee", fee)
}

// UpdateRegistrationFee is a paid mutator transaction binding the contract method 0xaf582c6b.
//
// Solidity: function updateRegistrationFee(uint256 fee) returns(bool)
func (_Register *RegisterSession) UpdateRegistrationFee(fee *big.Int) (*types.Transaction, error) {
	return _Register.Contract.UpdateRegistrationFee(&_Register.TransactOpts, fee)
}

// UpdateRegistrationFee is a paid mutator transaction binding the contract method 0xaf582c6b.
//
// Solidity: function updateRegistrationFee(uint256 fee) returns(bool)
func (_Register *RegisterTransactorSession) UpdateRegistrationFee(fee *big.Int) (*types.Transaction, error) {
	return _Register.Contract.UpdateRegistrationFee(&_Register.TransactOpts, fee)
}

// UpdateReportFee is a paid mutator transaction binding the contract method 0x69366c48.
//
// Solidity: function updateReportFee(uint256 fee) returns(bool)
func (_Register *RegisterTransactor) UpdateReportFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Register.contract.Transact(opts, "updateReportFee", fee)
}

// UpdateReportFee is a paid mutator transaction binding the contract method 0x69366c48.
//
// Solidity: function updateReportFee(uint256 fee) returns(bool)
func (_Register *RegisterSession) UpdateReportFee(fee *big.Int) (*types.Transaction, error) {
	return _Register.Contract.UpdateReportFee(&_Register.TransactOpts, fee)
}

// UpdateReportFee is a paid mutator transaction binding the contract method 0x69366c48.
//
// Solidity: function updateReportFee(uint256 fee) returns(bool)
func (_Register *RegisterTransactorSession) UpdateReportFee(fee *big.Int) (*types.Transaction, error) {
	return _Register.Contract.UpdateReportFee(&_Register.TransactOpts, fee)
}

// RegisterOracleBlacklistedIterator is returned from FilterOracleBlacklisted and is used to iterate over the raw logs and unpacked data for OracleBlacklisted events raised by the Register contract.
type RegisterOracleBlacklistedIterator struct {
	Event *RegisterOracleBlacklisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegisterOracleBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegisterOracleBlacklisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegisterOracleBlacklisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegisterOracleBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegisterOracleBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegisterOracleBlacklisted represents a OracleBlacklisted event raised by the Register contract.
type RegisterOracleBlacklisted struct {
	BlacklistedOracle common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOracleBlacklisted is a free log retrieval operation binding the contract event 0x8f4f70e0f343350f6df22d754271f375cf94e4a4d04ccb29df713c6ffce14dbf.
//
// Solidity: event OracleBlacklisted(address indexed blacklistedOracle)
func (_Register *RegisterFilterer) FilterOracleBlacklisted(opts *bind.FilterOpts, blacklistedOracle []common.Address) (*RegisterOracleBlacklistedIterator, error) {

	var blacklistedOracleRule []interface{}
	for _, blacklistedOracleItem := range blacklistedOracle {
		blacklistedOracleRule = append(blacklistedOracleRule, blacklistedOracleItem)
	}

	logs, sub, err := _Register.contract.FilterLogs(opts, "OracleBlacklisted", blacklistedOracleRule)
	if err != nil {
		return nil, err
	}
	return &RegisterOracleBlacklistedIterator{contract: _Register.contract, event: "OracleBlacklisted", logs: logs, sub: sub}, nil
}

// WatchOracleBlacklisted is a free log subscription operation binding the contract event 0x8f4f70e0f343350f6df22d754271f375cf94e4a4d04ccb29df713c6ffce14dbf.
//
// Solidity: event OracleBlacklisted(address indexed blacklistedOracle)
func (_Register *RegisterFilterer) WatchOracleBlacklisted(opts *bind.WatchOpts, sink chan<- *RegisterOracleBlacklisted, blacklistedOracle []common.Address) (event.Subscription, error) {

	var blacklistedOracleRule []interface{}
	for _, blacklistedOracleItem := range blacklistedOracle {
		blacklistedOracleRule = append(blacklistedOracleRule, blacklistedOracleItem)
	}

	logs, sub, err := _Register.contract.WatchLogs(opts, "OracleBlacklisted", blacklistedOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegisterOracleBlacklisted)
				if err := _Register.contract.UnpackLog(event, "OracleBlacklisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOracleBlacklisted is a log parse operation binding the contract event 0x8f4f70e0f343350f6df22d754271f375cf94e4a4d04ccb29df713c6ffce14dbf.
//
// Solidity: event OracleBlacklisted(address indexed blacklistedOracle)
func (_Register *RegisterFilterer) ParseOracleBlacklisted(log types.Log) (*RegisterOracleBlacklisted, error) {
	event := new(RegisterOracleBlacklisted)
	if err := _Register.contract.UnpackLog(event, "OracleBlacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegisterOracleOfflineIterator is returned from FilterOracleOffline and is used to iterate over the raw logs and unpacked data for OracleOffline events raised by the Register contract.
type RegisterOracleOfflineIterator struct {
	Event *RegisterOracleOffline // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegisterOracleOfflineIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegisterOracleOffline)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegisterOracleOffline)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegisterOracleOfflineIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegisterOracleOfflineIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegisterOracleOffline represents a OracleOffline event raised by the Register contract.
type RegisterOracleOffline struct {
	OracleAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOracleOffline is a free log retrieval operation binding the contract event 0x6eccb14325d715b33601268fe5c7d8a1618261e4d242e890c9e2e372adaf1fbc.
//
// Solidity: event OracleOffline(address indexed oracleAddress)
func (_Register *RegisterFilterer) FilterOracleOffline(opts *bind.FilterOpts, oracleAddress []common.Address) (*RegisterOracleOfflineIterator, error) {

	var oracleAddressRule []interface{}
	for _, oracleAddressItem := range oracleAddress {
		oracleAddressRule = append(oracleAddressRule, oracleAddressItem)
	}

	logs, sub, err := _Register.contract.FilterLogs(opts, "OracleOffline", oracleAddressRule)
	if err != nil {
		return nil, err
	}
	return &RegisterOracleOfflineIterator{contract: _Register.contract, event: "OracleOffline", logs: logs, sub: sub}, nil
}

// WatchOracleOffline is a free log subscription operation binding the contract event 0x6eccb14325d715b33601268fe5c7d8a1618261e4d242e890c9e2e372adaf1fbc.
//
// Solidity: event OracleOffline(address indexed oracleAddress)
func (_Register *RegisterFilterer) WatchOracleOffline(opts *bind.WatchOpts, sink chan<- *RegisterOracleOffline, oracleAddress []common.Address) (event.Subscription, error) {

	var oracleAddressRule []interface{}
	for _, oracleAddressItem := range oracleAddress {
		oracleAddressRule = append(oracleAddressRule, oracleAddressItem)
	}

	logs, sub, err := _Register.contract.WatchLogs(opts, "OracleOffline", oracleAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegisterOracleOffline)
				if err := _Register.contract.UnpackLog(event, "OracleOffline", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOracleOffline is a log parse operation binding the contract event 0x6eccb14325d715b33601268fe5c7d8a1618261e4d242e890c9e2e372adaf1fbc.
//
// Solidity: event OracleOffline(address indexed oracleAddress)
func (_Register *RegisterFilterer) ParseOracleOffline(log types.Log) (*RegisterOracleOffline, error) {
	event := new(RegisterOracleOffline)
	if err := _Register.contract.UnpackLog(event, "OracleOffline", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegisterOracleOnlineIterator is returned from FilterOracleOnline and is used to iterate over the raw logs and unpacked data for OracleOnline events raised by the Register contract.
type RegisterOracleOnlineIterator struct {
	Event *RegisterOracleOnline // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegisterOracleOnlineIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegisterOracleOnline)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegisterOracleOnline)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegisterOracleOnlineIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegisterOracleOnlineIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegisterOracleOnline represents a OracleOnline event raised by the Register contract.
type RegisterOracleOnline struct {
	OracleAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOracleOnline is a free log retrieval operation binding the contract event 0x3a21ecfcdc410d1e3552f24e84bbf13af57e47436f3b9ac5045ae9f3adcdc411.
//
// Solidity: event OracleOnline(address indexed oracleAddress)
func (_Register *RegisterFilterer) FilterOracleOnline(opts *bind.FilterOpts, oracleAddress []common.Address) (*RegisterOracleOnlineIterator, error) {

	var oracleAddressRule []interface{}
	for _, oracleAddressItem := range oracleAddress {
		oracleAddressRule = append(oracleAddressRule, oracleAddressItem)
	}

	logs, sub, err := _Register.contract.FilterLogs(opts, "OracleOnline", oracleAddressRule)
	if err != nil {
		return nil, err
	}
	return &RegisterOracleOnlineIterator{contract: _Register.contract, event: "OracleOnline", logs: logs, sub: sub}, nil
}

// WatchOracleOnline is a free log subscription operation binding the contract event 0x3a21ecfcdc410d1e3552f24e84bbf13af57e47436f3b9ac5045ae9f3adcdc411.
//
// Solidity: event OracleOnline(address indexed oracleAddress)
func (_Register *RegisterFilterer) WatchOracleOnline(opts *bind.WatchOpts, sink chan<- *RegisterOracleOnline, oracleAddress []common.Address) (event.Subscription, error) {

	var oracleAddressRule []interface{}
	for _, oracleAddressItem := range oracleAddress {
		oracleAddressRule = append(oracleAddressRule, oracleAddressItem)
	}

	logs, sub, err := _Register.contract.WatchLogs(opts, "OracleOnline", oracleAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegisterOracleOnline)
				if err := _Register.contract.UnpackLog(event, "OracleOnline", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOracleOnline is a log parse operation binding the contract event 0x3a21ecfcdc410d1e3552f24e84bbf13af57e47436f3b9ac5045ae9f3adcdc411.
//
// Solidity: event OracleOnline(address indexed oracleAddress)
func (_Register *RegisterFilterer) ParseOracleOnline(log types.Log) (*RegisterOracleOnline, error) {
	event := new(RegisterOracleOnline)
	if err := _Register.contract.UnpackLog(event, "OracleOnline", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegisterOracleRegisterTaxUpdatedIterator is returned from FilterOracleRegisterTaxUpdated and is used to iterate over the raw logs and unpacked data for OracleRegisterTaxUpdated events raised by the Register contract.
type RegisterOracleRegisterTaxUpdatedIterator struct {
	Event *RegisterOracleRegisterTaxUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegisterOracleRegisterTaxUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegisterOracleRegisterTaxUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegisterOracleRegisterTaxUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegisterOracleRegisterTaxUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegisterOracleRegisterTaxUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegisterOracleRegisterTaxUpdated represents a OracleRegisterTaxUpdated event raised by the Register contract.
type RegisterOracleRegisterTaxUpdated struct {
	NewTax *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOracleRegisterTaxUpdated is a free log retrieval operation binding the contract event 0xa288a9ac9f74fe4c85e8263739645d49297e82ffa010aac11899ba671a54d1ac.
//
// Solidity: event OracleRegisterTaxUpdated(uint256 newTax)
func (_Register *RegisterFilterer) FilterOracleRegisterTaxUpdated(opts *bind.FilterOpts) (*RegisterOracleRegisterTaxUpdatedIterator, error) {

	logs, sub, err := _Register.contract.FilterLogs(opts, "OracleRegisterTaxUpdated")
	if err != nil {
		return nil, err
	}
	return &RegisterOracleRegisterTaxUpdatedIterator{contract: _Register.contract, event: "OracleRegisterTaxUpdated", logs: logs, sub: sub}, nil
}

// WatchOracleRegisterTaxUpdated is a free log subscription operation binding the contract event 0xa288a9ac9f74fe4c85e8263739645d49297e82ffa010aac11899ba671a54d1ac.
//
// Solidity: event OracleRegisterTaxUpdated(uint256 newTax)
func (_Register *RegisterFilterer) WatchOracleRegisterTaxUpdated(opts *bind.WatchOpts, sink chan<- *RegisterOracleRegisterTaxUpdated) (event.Subscription, error) {

	logs, sub, err := _Register.contract.WatchLogs(opts, "OracleRegisterTaxUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegisterOracleRegisterTaxUpdated)
				if err := _Register.contract.UnpackLog(event, "OracleRegisterTaxUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOracleRegisterTaxUpdated is a log parse operation binding the contract event 0xa288a9ac9f74fe4c85e8263739645d49297e82ffa010aac11899ba671a54d1ac.
//
// Solidity: event OracleRegisterTaxUpdated(uint256 newTax)
func (_Register *RegisterFilterer) ParseOracleRegisterTaxUpdated(log types.Log) (*RegisterOracleRegisterTaxUpdated, error) {
	event := new(RegisterOracleRegisterTaxUpdated)
	if err := _Register.contract.UnpackLog(event, "OracleRegisterTaxUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegisterOracleRegisteredIterator is returned from FilterOracleRegistered and is used to iterate over the raw logs and unpacked data for OracleRegistered events raised by the Register contract.
type RegisterOracleRegisteredIterator struct {
	Event *RegisterOracleRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegisterOracleRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegisterOracleRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegisterOracleRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegisterOracleRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegisterOracleRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegisterOracleRegistered represents a OracleRegistered event raised by the Register contract.
type RegisterOracleRegistered struct {
	OracleAddress common.Address
	Ip            string
	Port          string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOracleRegistered is a free log retrieval operation binding the contract event 0x8c1c8966369199dc710cd8b615bf16a86520dd165679cb74aa5f427d2bf8c0c9.
//
// Solidity: event OracleRegistered(address indexed oracleAddress, string ip, string port)
func (_Register *RegisterFilterer) FilterOracleRegistered(opts *bind.FilterOpts, oracleAddress []common.Address) (*RegisterOracleRegisteredIterator, error) {

	var oracleAddressRule []interface{}
	for _, oracleAddressItem := range oracleAddress {
		oracleAddressRule = append(oracleAddressRule, oracleAddressItem)
	}

	logs, sub, err := _Register.contract.FilterLogs(opts, "OracleRegistered", oracleAddressRule)
	if err != nil {
		return nil, err
	}
	return &RegisterOracleRegisteredIterator{contract: _Register.contract, event: "OracleRegistered", logs: logs, sub: sub}, nil
}

// WatchOracleRegistered is a free log subscription operation binding the contract event 0x8c1c8966369199dc710cd8b615bf16a86520dd165679cb74aa5f427d2bf8c0c9.
//
// Solidity: event OracleRegistered(address indexed oracleAddress, string ip, string port)
func (_Register *RegisterFilterer) WatchOracleRegistered(opts *bind.WatchOpts, sink chan<- *RegisterOracleRegistered, oracleAddress []common.Address) (event.Subscription, error) {

	var oracleAddressRule []interface{}
	for _, oracleAddressItem := range oracleAddress {
		oracleAddressRule = append(oracleAddressRule, oracleAddressItem)
	}

	logs, sub, err := _Register.contract.WatchLogs(opts, "OracleRegistered", oracleAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegisterOracleRegistered)
				if err := _Register.contract.UnpackLog(event, "OracleRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOracleRegistered is a log parse operation binding the contract event 0x8c1c8966369199dc710cd8b615bf16a86520dd165679cb74aa5f427d2bf8c0c9.
//
// Solidity: event OracleRegistered(address indexed oracleAddress, string ip, string port)
func (_Register *RegisterFilterer) ParseOracleRegistered(log types.Log) (*RegisterOracleRegistered, error) {
	event := new(RegisterOracleRegistered)
	if err := _Register.contract.UnpackLog(event, "OracleRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegisterOracleReportedIterator is returned from FilterOracleReported and is used to iterate over the raw logs and unpacked data for OracleReported events raised by the Register contract.
type RegisterOracleReportedIterator struct {
	Event *RegisterOracleReported // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegisterOracleReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegisterOracleReported)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegisterOracleReported)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegisterOracleReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegisterOracleReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegisterOracleReported represents a OracleReported event raised by the Register contract.
type RegisterOracleReported struct {
	Reporter       common.Address
	ReportedOracle common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOracleReported is a free log retrieval operation binding the contract event 0xc288b29f22724d188803361dc865debd5a98e3f2c073a19c9fe8639a6274d23d.
//
// Solidity: event OracleReported(address indexed reporter, address indexed reportedOracle)
func (_Register *RegisterFilterer) FilterOracleReported(opts *bind.FilterOpts, reporter []common.Address, reportedOracle []common.Address) (*RegisterOracleReportedIterator, error) {

	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var reportedOracleRule []interface{}
	for _, reportedOracleItem := range reportedOracle {
		reportedOracleRule = append(reportedOracleRule, reportedOracleItem)
	}

	logs, sub, err := _Register.contract.FilterLogs(opts, "OracleReported", reporterRule, reportedOracleRule)
	if err != nil {
		return nil, err
	}
	return &RegisterOracleReportedIterator{contract: _Register.contract, event: "OracleReported", logs: logs, sub: sub}, nil
}

// WatchOracleReported is a free log subscription operation binding the contract event 0xc288b29f22724d188803361dc865debd5a98e3f2c073a19c9fe8639a6274d23d.
//
// Solidity: event OracleReported(address indexed reporter, address indexed reportedOracle)
func (_Register *RegisterFilterer) WatchOracleReported(opts *bind.WatchOpts, sink chan<- *RegisterOracleReported, reporter []common.Address, reportedOracle []common.Address) (event.Subscription, error) {

	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}
	var reportedOracleRule []interface{}
	for _, reportedOracleItem := range reportedOracle {
		reportedOracleRule = append(reportedOracleRule, reportedOracleItem)
	}

	logs, sub, err := _Register.contract.WatchLogs(opts, "OracleReported", reporterRule, reportedOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegisterOracleReported)
				if err := _Register.contract.UnpackLog(event, "OracleReported", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOracleReported is a log parse operation binding the contract event 0xc288b29f22724d188803361dc865debd5a98e3f2c073a19c9fe8639a6274d23d.
//
// Solidity: event OracleReported(address indexed reporter, address indexed reportedOracle)
func (_Register *RegisterFilterer) ParseOracleReported(log types.Log) (*RegisterOracleReported, error) {
	event := new(RegisterOracleReported)
	if err := _Register.contract.UnpackLog(event, "OracleReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
