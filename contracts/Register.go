// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package register

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
}

// RegisterMetaData contains all meta data concerning the Register contract.
var RegisterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_registerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reportFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"daoOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"blacklistedOracle\",\"type\":\"address\"}],\"name\":\"OracleBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"}],\"name\":\"OracleRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reportedOracle\",\"type\":\"address\"}],\"name\":\"OracleReported\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getOracles\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"name\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"reputation\",\"type\":\"uint256\"}],\"internalType\":\"structOracle[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"}],\"name\":\"isBlacklisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"}],\"name\":\"reportOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"updateRegistrationFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"updateReportFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051611e6f380380611e6f83398181016040528101906100319190610118565b82600181905550816002819055508060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050610168565b5f5ffd5b5f819050919050565b61009d8161008b565b81146100a7575f5ffd5b50565b5f815190506100b881610094565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100e7826100be565b9050919050565b6100f7816100dd565b8114610101575f5ffd5b50565b5f81519050610112816100ee565b92915050565b5f5f5f6060848603121561012f5761012e610087565b5b5f61013c868287016100aa565b935050602061014d868287016100aa565b925050604061015e86828701610104565b9150509250925092565b611cfa806101755f395ff3fe608060405260043610610054575f3560e01c80633ffbd47f1461005857806340884c521461008857806369366c48146100b2578063af582c6b146100e2578063e1704d8314610112578063fe575a8714610142575b5f5ffd5b610072600480360381019061006d9190611198565b61017e565b60405161007f9190611228565b60405180910390f35b348015610093575f5ffd5b5061009c610383565b6040516100a9919061141a565b60405180910390f35b6100cc60048036038101906100c79190611464565b61055c565b6040516100d99190611228565b60405180910390f35b6100fc60048036038101906100f79190611464565b6105fc565b6040516101099190611228565b60405180910390f35b61012c600480360381019061012791906114b9565b61069c565b6040516101399190611228565b60405180910390f35b34801561014d575f5ffd5b50610168600480360381019061016391906114b9565b610c6f565b6040516101759190611228565b60405180910390f35b5f6001543410156101c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101bb9061153e565b60405180910390fd5b60065f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff161561024e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610245906115cc565b60405180910390fd5b5f60405180608001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020016001815250908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101908161030691906117e7565b50604082015181600201908161031c91906117e7565b506060820151816003015550503373ffffffffffffffffffffffffffffffffffffffff167f8c1c8966369199dc710cd8b615bf16a86520dd165679cb74aa5f427d2bf8c0c984846040516103719291906118ee565b60405180910390a26001905092915050565b60605f805480602002602001604051908101604052809291908181526020015f905b82821015610553578382905f5260205f2090600402016040518060800160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461042a90611617565b80601f016020809104026020016040519081016040528092919081815260200182805461045690611617565b80156104a15780601f10610478576101008083540402835291602001916104a1565b820191905f5260205f20905b81548152906001019060200180831161048457829003601f168201915b505050505081526020016002820180546104ba90611617565b80601f01602080910402602001604051908101604052809291908181526020018280546104e690611617565b80156105315780601f1061050857610100808354040283529160200191610531565b820191905f5260205f20905b81548152906001019060200180831161051457829003601f168201915b50505050508152602001600382015481525050815260200190600101906103a5565b50505050905090565b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105e39061196d565b60405180910390fd5b8160028190555060019050919050565b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461068c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106839061196d565b60405180910390fd5b8160018190555060019050919050565b5f5f5f90505f5f90505b5f8054905081101561078f573373ffffffffffffffffffffffffffffffffffffffff165f82815481106106dc576106db61198b565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16148015610774575060065f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16155b15610782576001915061078f565b80806001019150506106a6565b50806107d0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107c790611a02565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361083e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161083590611a6a565b60405180910390fd5b60065f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16156108c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108bf90611ad2565b60405180910390fd5b60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff161561098d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161098490611b3a565b60405180910390fd5b6002543410156109d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109c990611ba2565b60405180910390fd5b600160045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555060055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f815480929190610aaf90611bed565b91905055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fc288b29f22724d188803361dc865debd5a98e3f2c073a19c9fe8639a6274d23d60405160405180910390a3600160025f80549050610b219190611c61565b610b2b9190611c91565b60055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410610c6557600160065f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508273ffffffffffffffffffffffffffffffffffffffff167f8f4f70e0f343350f6df22d754271f375cf94e4a4d04ccb29df713c6ffce14dbf60405160405180910390a260025460075f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610c549190611c91565b92505081905550610c6483610cc1565b5b6001915050919050565b5f60065f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b5f5f90505f5f90505b5f80549050811015610dbc5760045f5f8381548110610cec57610ceb61198b565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615610daf578180610dab90611bed565b9250505b8080600101915050610cca565b505f81600154610dcc9190611c61565b90505f5f90505b5f805490508110156110455760045f5f8381548110610df557610df461198b565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615611038575f8181548110610ebc57610ebb61198b565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc8360075f5f8681548110610f1a57610f1961198b565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054610f8c9190611c91565b90811502906040515f60405180830381858888f19350505050158015610fb4573d5f5f3e3d5ffd5b505f60075f5f8481548110610fcc57610fcb61198b565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505b8080600101915050610dd3565b50505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6110aa82611064565b810181811067ffffffffffffffff821117156110c9576110c8611074565b5b80604052505050565b5f6110db61104b565b90506110e782826110a1565b919050565b5f67ffffffffffffffff82111561110657611105611074565b5b61110f82611064565b9050602081019050919050565b828183375f83830152505050565b5f61113c611137846110ec565b6110d2565b90508281526020810184848401111561115857611157611060565b5b61116384828561111c565b509392505050565b5f82601f83011261117f5761117e61105c565b5b813561118f84826020860161112a565b91505092915050565b5f5f604083850312156111ae576111ad611054565b5b5f83013567ffffffffffffffff8111156111cb576111ca611058565b5b6111d78582860161116b565b925050602083013567ffffffffffffffff8111156111f8576111f7611058565b5b6112048582860161116b565b9150509250929050565b5f8115159050919050565b6112228161120e565b82525050565b5f60208201905061123b5f830184611219565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6112938261126a565b9050919050565b6112a381611289565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6112db826112a9565b6112e581856112b3565b93506112f58185602086016112c3565b6112fe81611064565b840191505092915050565b5f819050919050565b61131b81611309565b82525050565b5f608083015f8301516113365f86018261129a565b506020830151848203602086015261134e82826112d1565b9150506040830151848203604086015261136882826112d1565b915050606083015161137d6060860182611312565b508091505092915050565b5f6113938383611321565b905092915050565b5f602082019050919050565b5f6113b182611241565b6113bb818561124b565b9350836020820285016113cd8561125b565b805f5b8581101561140857848403895281516113e98582611388565b94506113f48361139b565b925060208a019950506001810190506113d0565b50829750879550505050505092915050565b5f6020820190508181035f83015261143281846113a7565b905092915050565b61144381611309565b811461144d575f5ffd5b50565b5f8135905061145e8161143a565b92915050565b5f6020828403121561147957611478611054565b5b5f61148684828501611450565b91505092915050565b61149881611289565b81146114a2575f5ffd5b50565b5f813590506114b38161148f565b92915050565b5f602082840312156114ce576114cd611054565b5b5f6114db848285016114a5565b91505092915050565b5f82825260208201905092915050565b7f496e73756666696369656e7420726567697374726174696f6e206665650000005f82015250565b5f611528601d836114e4565b9150611533826114f4565b602082019050919050565b5f6020820190508181035f8301526115558161151c565b9050919050565b7f43616e6e6f74207265676973746572206120626c61636b6c6973746564206f725f8201527f61636c6500000000000000000000000000000000000000000000000000000000602082015250565b5f6115b66024836114e4565b91506115c18261155c565b604082019050919050565b5f6020820190508181035f8301526115e3816115aa565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061162e57607f821691505b602082108103611641576116406115ea565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026116a37fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611668565b6116ad8683611668565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6116e86116e36116de84611309565b6116c5565b611309565b9050919050565b5f819050919050565b611701836116ce565b61171561170d826116ef565b848454611674565b825550505050565b5f5f905090565b61172c61171d565b6117378184846116f8565b505050565b5b8181101561175a5761174f5f82611724565b60018101905061173d565b5050565b601f82111561179f5761177081611647565b61177984611659565b81016020851015611788578190505b61179c61179485611659565b83018261173c565b50505b505050565b5f82821c905092915050565b5f6117bf5f19846008026117a4565b1980831691505092915050565b5f6117d783836117b0565b9150826002028217905092915050565b6117f0826112a9565b67ffffffffffffffff81111561180957611808611074565b5b6118138254611617565b61181e82828561175e565b5f60209050601f83116001811461184f575f841561183d578287015190505b61184785826117cc565b8655506118ae565b601f19841661185d86611647565b5f5b828110156118845784890151825560018201915060208501945060208101905061185f565b868310156118a1578489015161189d601f8916826117b0565b8355505b6001600288020188555050505b505050505050565b5f6118c0826112a9565b6118ca81856114e4565b93506118da8185602086016112c3565b6118e381611064565b840191505092915050565b5f6040820190508181035f83015261190681856118b6565b9050818103602083015261191a81846118b6565b90509392505050565b7f4e6f7420617574686f72697a65640000000000000000000000000000000000005f82015250565b5f611957600e836114e4565b915061196282611923565b602082019050919050565b5f6020820190508181035f8301526119848161194b565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e6f7420616e20617574686f72697a6564206f7261636c6500000000000000005f82015250565b5f6119ec6018836114e4565b91506119f7826119b8565b602082019050919050565b5f6020820190508181035f830152611a19816119e0565b9050919050565b7f43616e6e6f74207265706f727420796f757273656c66000000000000000000005f82015250565b5f611a546016836114e4565b9150611a5f82611a20565b602082019050919050565b5f6020820190508181035f830152611a8181611a48565b9050919050565b7f4f7261636c6520697320616c726561647920626c61636b6c69737465640000005f82015250565b5f611abc601d836114e4565b9150611ac782611a88565b602082019050919050565b5f6020820190508181035f830152611ae981611ab0565b9050919050565b7f4f7261636c6520616c7265616479207265706f7274656420627920796f7500005f82015250565b5f611b24601e836114e4565b9150611b2f82611af0565b602082019050919050565b5f6020820190508181035f830152611b5181611b18565b9050919050565b7f496e73756666696369656e74207265706f7274206665650000000000000000005f82015250565b5f611b8c6017836114e4565b9150611b9782611b58565b602082019050919050565b5f6020820190508181035f830152611bb981611b80565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611bf782611309565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611c2957611c28611bc0565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f611c6b82611309565b9150611c7683611309565b925082611c8657611c85611c34565b5b828204905092915050565b5f611c9b82611309565b9150611ca683611309565b9250828201905080821115611cbe57611cbd611bc0565b5b9291505056fea26469706673582212204514d71751edb4bda1a819545e24f4c99ae2f16ec685acd77a7cd748fac0533f64736f6c634300081c0033",
}

// RegisterABI is the input ABI used to generate the binding from.
// Deprecated: Use RegisterMetaData.ABI instead.
var RegisterABI = RegisterMetaData.ABI

// RegisterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RegisterMetaData.Bin instead.
var RegisterBin = RegisterMetaData.Bin

// DeployRegister deploys a new Ethereum contract, binding an instance of Register to it.
func DeployRegister(auth *bind.TransactOpts, backend bind.ContractBackend, _registerFee *big.Int, _reportFee *big.Int, daoOwner common.Address) (common.Address, *types.Transaction, *Register, error) {
	parsed, err := RegisterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RegisterBin), backend, _registerFee, _reportFee, daoOwner)
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

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns((address,string,string,uint256)[])
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
// Solidity: function getOracles() view returns((address,string,string,uint256)[])
func (_Register *RegisterSession) GetOracles() ([]Oracle, error) {
	return _Register.Contract.GetOracles(&_Register.CallOpts)
}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns((address,string,string,uint256)[])
func (_Register *RegisterCallerSession) GetOracles() ([]Oracle, error) {
	return _Register.Contract.GetOracles(&_Register.CallOpts)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address oracleAddress) view returns(bool)
func (_Register *RegisterCaller) IsBlacklisted(opts *bind.CallOpts, oracleAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Register.contract.Call(opts, &out, "isBlacklisted", oracleAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address oracleAddress) view returns(bool)
func (_Register *RegisterSession) IsBlacklisted(oracleAddress common.Address) (bool, error) {
	return _Register.Contract.IsBlacklisted(&_Register.CallOpts, oracleAddress)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address oracleAddress) view returns(bool)
func (_Register *RegisterCallerSession) IsBlacklisted(oracleAddress common.Address) (bool, error) {
	return _Register.Contract.IsBlacklisted(&_Register.CallOpts, oracleAddress)
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

// ReportOracle is a paid mutator transaction binding the contract method 0xe1704d83.
//
// Solidity: function reportOracle(address oracleAddress) payable returns(bool)
func (_Register *RegisterTransactor) ReportOracle(opts *bind.TransactOpts, oracleAddress common.Address) (*types.Transaction, error) {
	return _Register.contract.Transact(opts, "reportOracle", oracleAddress)
}

// ReportOracle is a paid mutator transaction binding the contract method 0xe1704d83.
//
// Solidity: function reportOracle(address oracleAddress) payable returns(bool)
func (_Register *RegisterSession) ReportOracle(oracleAddress common.Address) (*types.Transaction, error) {
	return _Register.Contract.ReportOracle(&_Register.TransactOpts, oracleAddress)
}

// ReportOracle is a paid mutator transaction binding the contract method 0xe1704d83.
//
// Solidity: function reportOracle(address oracleAddress) payable returns(bool)
func (_Register *RegisterTransactorSession) ReportOracle(oracleAddress common.Address) (*types.Transaction, error) {
	return _Register.Contract.ReportOracle(&_Register.TransactOpts, oracleAddress)
}

// UpdateRegistrationFee is a paid mutator transaction binding the contract method 0xaf582c6b.
//
// Solidity: function updateRegistrationFee(uint256 fee) payable returns(bool)
func (_Register *RegisterTransactor) UpdateRegistrationFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Register.contract.Transact(opts, "updateRegistrationFee", fee)
}

// UpdateRegistrationFee is a paid mutator transaction binding the contract method 0xaf582c6b.
//
// Solidity: function updateRegistrationFee(uint256 fee) payable returns(bool)
func (_Register *RegisterSession) UpdateRegistrationFee(fee *big.Int) (*types.Transaction, error) {
	return _Register.Contract.UpdateRegistrationFee(&_Register.TransactOpts, fee)
}

// UpdateRegistrationFee is a paid mutator transaction binding the contract method 0xaf582c6b.
//
// Solidity: function updateRegistrationFee(uint256 fee) payable returns(bool)
func (_Register *RegisterTransactorSession) UpdateRegistrationFee(fee *big.Int) (*types.Transaction, error) {
	return _Register.Contract.UpdateRegistrationFee(&_Register.TransactOpts, fee)
}

// UpdateReportFee is a paid mutator transaction binding the contract method 0x69366c48.
//
// Solidity: function updateReportFee(uint256 fee) payable returns(bool)
func (_Register *RegisterTransactor) UpdateReportFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Register.contract.Transact(opts, "updateReportFee", fee)
}

// UpdateReportFee is a paid mutator transaction binding the contract method 0x69366c48.
//
// Solidity: function updateReportFee(uint256 fee) payable returns(bool)
func (_Register *RegisterSession) UpdateReportFee(fee *big.Int) (*types.Transaction, error) {
	return _Register.Contract.UpdateReportFee(&_Register.TransactOpts, fee)
}

// UpdateReportFee is a paid mutator transaction binding the contract method 0x69366c48.
//
// Solidity: function updateReportFee(uint256 fee) payable returns(bool)
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
