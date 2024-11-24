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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_registerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reportFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"daoOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"blacklistedOracle\",\"type\":\"address\"}],\"name\":\"OracleBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"}],\"name\":\"OracleRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reportedOracle\",\"type\":\"address\"}],\"name\":\"OracleReported\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getOracles\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"name\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"reputation\",\"type\":\"uint256\"}],\"internalType\":\"structOracle[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"}],\"name\":\"isBlacklisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOracleRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"}],\"name\":\"reportOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"updateRegistrationFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"updateReportFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040516120c83803806120c883398181016040528101906100319190610118565b82600181905550816002819055508060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050610168565b5f5ffd5b5f819050919050565b61009d8161008b565b81146100a7575f5ffd5b50565b5f815190506100b881610094565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100e7826100be565b9050919050565b6100f7816100dd565b8114610101575f5ffd5b50565b5f81519050610112816100ee565b92915050565b5f5f5f6060848603121561012f5761012e610087565b5b5f61013c868287016100aa565b935050602061014d868287016100aa565b925050604061015e86828701610104565b9150509250925092565b611f53806101755f395ff3fe60806040526004361061006f575f3560e01c806369366c481161004d57806369366c48146100f7578063af582c6b14610127578063e1704d8314610157578063fe575a87146101875761006f565b80630d1db1f0146100735780633ffbd47f1461009d57806340884c52146100cd575b5f5ffd5b34801561007e575f5ffd5b506100876101c3565b60405161009491906112be565b60405180910390f35b6100b760048036038101906100b29190611424565b6103d7565b6040516100c491906112be565b60405180910390f35b3480156100d8575f5ffd5b506100e16105dc565b6040516100ee9190611673565b60405180910390f35b610111600480360381019061010c91906116bd565b6107b5565b60405161011e91906112be565b60405180910390f35b610141600480360381019061013c91906116bd565b610855565b60405161014e91906112be565b60405180910390f35b610171600480360381019061016c9190611712565b6108f5565b60405161017e91906112be565b60405180910390f35b348015610192575f5ffd5b506101ad60048036038101906101a89190611712565b610ec8565b6040516101ba91906112be565b60405180910390f35b5f5f5f90505b5f805490508110156103cf575f5f82815481106101e9576101e861173d565b5b905f5260205f2090600402016040518060800160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461026490611797565b80601f016020809104026020016040519081016040528092919081815260200182805461029090611797565b80156102db5780601f106102b2576101008083540402835291602001916102db565b820191905f5260205f20905b8154815290600101906020018083116102be57829003601f168201915b505050505081526020016002820180546102f490611797565b80601f016020809104026020016040519081016040528092919081815260200182805461032090611797565b801561036b5780601f106103425761010080835404028352916020019161036b565b820191905f5260205f20905b81548152906001019060200180831161034e57829003601f168201915b5050505050815260200160038201548152505090503373ffffffffffffffffffffffffffffffffffffffff16815f015173ffffffffffffffffffffffffffffffffffffffff16036103c1576001925050506103d4565b5080806001019150506101c9565b505f90505b90565b5f60015434101561041d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041490611821565b60405180910390fd5b60065f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16156104a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161049e906118af565b60405180910390fd5b5f60405180608001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020016001815250908060018154018082558091505060019003905f5260205f2090600402015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101908161055f9190611a6d565b5060408201518160020190816105759190611a6d565b506060820151816003015550503373ffffffffffffffffffffffffffffffffffffffff167f8c1c8966369199dc710cd8b615bf16a86520dd165679cb74aa5f427d2bf8c0c984846040516105ca929190611b74565b60405180910390a26001905092915050565b60605f805480602002602001604051908101604052809291908181526020015f905b828210156107ac578382905f5260205f2090600402016040518060800160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461068390611797565b80601f01602080910402602001604051908101604052809291908181526020018280546106af90611797565b80156106fa5780601f106106d1576101008083540402835291602001916106fa565b820191905f5260205f20905b8154815290600101906020018083116106dd57829003601f168201915b5050505050815260200160028201805461071390611797565b80601f016020809104026020016040519081016040528092919081815260200182805461073f90611797565b801561078a5780601f106107615761010080835404028352916020019161078a565b820191905f5260205f20905b81548152906001019060200180831161076d57829003601f168201915b50505050508152602001600382015481525050815260200190600101906105fe565b50505050905090565b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610845576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161083c90611bf3565b60405180910390fd5b8160028190555060019050919050565b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108dc90611bf3565b60405180910390fd5b8160018190555060019050919050565b5f5f5f90505f5f90505b5f805490508110156109e8573373ffffffffffffffffffffffffffffffffffffffff165f82815481106109355761093461173d565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161480156109cd575060065f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16155b156109db57600191506109e8565b80806001019150506108ff565b5080610a29576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2090611c5b565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610a97576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8e90611cc3565b60405180910390fd5b60065f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615610b21576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b1890611d2b565b60405180910390fd5b60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615610be6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bdd90611d93565b60405180910390fd5b600254341015610c2b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c2290611dfb565b60405180910390fd5b600160045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555060055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f815480929190610d0890611e46565b91905055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fc288b29f22724d188803361dc865debd5a98e3f2c073a19c9fe8639a6274d23d60405160405180910390a3600160025f80549050610d7a9190611eba565b610d849190611eea565b60055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410610ebe57600160065f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508273ffffffffffffffffffffffffffffffffffffffff167f8f4f70e0f343350f6df22d754271f375cf94e4a4d04ccb29df713c6ffce14dbf60405160405180910390a260025460075f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254610ead9190611eea565b92505081905550610ebd83610f1a565b5b6001915050919050565b5f60065f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b5f5f90505f5f90505b5f805490508110156110155760045f5f8381548110610f4557610f4461173d565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff161561100857818061100490611e46565b9250505b8080600101915050610f23565b505f816001546110259190611eba565b90505f5f90505b5f8054905081101561129e5760045f5f838154811061104e5761104d61173d565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615611291575f81815481106111155761111461173d565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc8360075f5f86815481106111735761117261173d565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20546111e59190611eea565b90811502906040515f60405180830381858888f1935050505015801561120d573d5f5f3e3d5ffd5b505f60075f5f84815481106112255761122461173d565b5b905f5260205f2090600402015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505b808060010191505061102c565b50505050565b5f8115159050919050565b6112b8816112a4565b82525050565b5f6020820190506112d15f8301846112af565b92915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611336826112f0565b810181811067ffffffffffffffff8211171561135557611354611300565b5b80604052505050565b5f6113676112d7565b9050611373828261132d565b919050565b5f67ffffffffffffffff82111561139257611391611300565b5b61139b826112f0565b9050602081019050919050565b828183375f83830152505050565b5f6113c86113c384611378565b61135e565b9050828152602081018484840111156113e4576113e36112ec565b5b6113ef8482856113a8565b509392505050565b5f82601f83011261140b5761140a6112e8565b5b813561141b8482602086016113b6565b91505092915050565b5f5f6040838503121561143a576114396112e0565b5b5f83013567ffffffffffffffff811115611457576114566112e4565b5b611463858286016113f7565b925050602083013567ffffffffffffffff811115611484576114836112e4565b5b611490858286016113f7565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6114ec826114c3565b9050919050565b6114fc816114e2565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61153482611502565b61153e818561150c565b935061154e81856020860161151c565b611557816112f0565b840191505092915050565b5f819050919050565b61157481611562565b82525050565b5f608083015f83015161158f5f8601826114f3565b50602083015184820360208601526115a7828261152a565b915050604083015184820360408601526115c1828261152a565b91505060608301516115d6606086018261156b565b508091505092915050565b5f6115ec838361157a565b905092915050565b5f602082019050919050565b5f61160a8261149a565b61161481856114a4565b935083602082028501611626856114b4565b805f5b85811015611661578484038952815161164285826115e1565b945061164d836115f4565b925060208a01995050600181019050611629565b50829750879550505050505092915050565b5f6020820190508181035f83015261168b8184611600565b905092915050565b61169c81611562565b81146116a6575f5ffd5b50565b5f813590506116b781611693565b92915050565b5f602082840312156116d2576116d16112e0565b5b5f6116df848285016116a9565b91505092915050565b6116f1816114e2565b81146116fb575f5ffd5b50565b5f8135905061170c816116e8565b92915050565b5f60208284031215611727576117266112e0565b5b5f611734848285016116fe565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806117ae57607f821691505b6020821081036117c1576117c061176a565b5b50919050565b5f82825260208201905092915050565b7f496e73756666696369656e7420726567697374726174696f6e206665650000005f82015250565b5f61180b601d836117c7565b9150611816826117d7565b602082019050919050565b5f6020820190508181035f830152611838816117ff565b9050919050565b7f43616e6e6f74207265676973746572206120626c61636b6c6973746564206f725f8201527f61636c6500000000000000000000000000000000000000000000000000000000602082015250565b5f6118996024836117c7565b91506118a48261183f565b604082019050919050565b5f6020820190508181035f8301526118c68161188d565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026119297fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826118ee565b61193386836118ee565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61196e61196961196484611562565b61194b565b611562565b9050919050565b5f819050919050565b61198783611954565b61199b61199382611975565b8484546118fa565b825550505050565b5f5f905090565b6119b26119a3565b6119bd81848461197e565b505050565b5b818110156119e0576119d55f826119aa565b6001810190506119c3565b5050565b601f821115611a25576119f6816118cd565b6119ff846118df565b81016020851015611a0e578190505b611a22611a1a856118df565b8301826119c2565b50505b505050565b5f82821c905092915050565b5f611a455f1984600802611a2a565b1980831691505092915050565b5f611a5d8383611a36565b9150826002028217905092915050565b611a7682611502565b67ffffffffffffffff811115611a8f57611a8e611300565b5b611a998254611797565b611aa48282856119e4565b5f60209050601f831160018114611ad5575f8415611ac3578287015190505b611acd8582611a52565b865550611b34565b601f198416611ae3866118cd565b5f5b82811015611b0a57848901518255600182019150602085019450602081019050611ae5565b86831015611b275784890151611b23601f891682611a36565b8355505b6001600288020188555050505b505050505050565b5f611b4682611502565b611b5081856117c7565b9350611b6081856020860161151c565b611b69816112f0565b840191505092915050565b5f6040820190508181035f830152611b8c8185611b3c565b90508181036020830152611ba08184611b3c565b90509392505050565b7f4e6f7420617574686f72697a65640000000000000000000000000000000000005f82015250565b5f611bdd600e836117c7565b9150611be882611ba9565b602082019050919050565b5f6020820190508181035f830152611c0a81611bd1565b9050919050565b7f4e6f7420616e20617574686f72697a6564206f7261636c6500000000000000005f82015250565b5f611c456018836117c7565b9150611c5082611c11565b602082019050919050565b5f6020820190508181035f830152611c7281611c39565b9050919050565b7f43616e6e6f74207265706f727420796f757273656c66000000000000000000005f82015250565b5f611cad6016836117c7565b9150611cb882611c79565b602082019050919050565b5f6020820190508181035f830152611cda81611ca1565b9050919050565b7f4f7261636c6520697320616c726561647920626c61636b6c69737465640000005f82015250565b5f611d15601d836117c7565b9150611d2082611ce1565b602082019050919050565b5f6020820190508181035f830152611d4281611d09565b9050919050565b7f4f7261636c6520616c7265616479207265706f7274656420627920796f7500005f82015250565b5f611d7d601e836117c7565b9150611d8882611d49565b602082019050919050565b5f6020820190508181035f830152611daa81611d71565b9050919050565b7f496e73756666696369656e74207265706f7274206665650000000000000000005f82015250565b5f611de56017836117c7565b9150611df082611db1565b602082019050919050565b5f6020820190508181035f830152611e1281611dd9565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611e5082611562565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611e8257611e81611e19565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f611ec482611562565b9150611ecf83611562565b925082611edf57611ede611e8d565b5b828204905092915050565b5f611ef482611562565b9150611eff83611562565b9250828201905080821115611f1757611f16611e19565b5b9291505056fea2646970667358221220ab417f680be2b3112249d18002781ba8a97c71da1850251ad1af0dceea22972864736f6c634300081c0033",
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
