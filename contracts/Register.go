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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_registerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reportFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_feeSetter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daoOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"blacklistedOracle\",\"type\":\"address\"}],\"name\":\"OracleBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"}],\"name\":\"OracleOffline\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"}],\"name\":\"OracleOnline\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTax\",\"type\":\"uint256\"}],\"name\":\"OracleRegisterTaxUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"}],\"name\":\"OracleRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reportedOracle\",\"type\":\"address\"}],\"name\":\"OracleReported\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LogOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Login\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tax\",\"type\":\"uint256\"}],\"name\":\"changeRegisterTax\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracles\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"name\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"reputation\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOnline\",\"type\":\"bool\"}],\"internalType\":\"structOracle[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegistrationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReportFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"}],\"name\":\"isBlacklisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOracleRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"}],\"name\":\"reportOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokeOwnership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"self\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"name\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"reputation\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOnline\",\"type\":\"bool\"}],\"internalType\":\"structOracle\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"updateRegistrationFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"updateReportFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405260056002555f600560146101000a81548160ff02191690831515021790555034801561002e575f5ffd5b50604051612c9e380380612c9e83398181016040528101906100509190610177565b8360038190555082600481905550815f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050506101db565b5f5ffd5b5f819050919050565b6100fc816100ea565b8114610106575f5ffd5b50565b5f81519050610117816100f3565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101468261011d565b9050919050565b6101568161013c565b8114610160575f5ffd5b50565b5f815190506101718161014d565b92915050565b5f5f5f5f6080858703121561018f5761018e6100e6565b5b5f61019c87828801610109565b94505060206101ad87828801610109565b93505060406101be87828801610163565b92505060606101cf87828801610163565b91505092959194509250565b612ab6806101e85f395ff3fe6080604052600436106100dc575f3560e01c80637571bdd01161007e578063af582c6b11610058578063af582c6b14610296578063e1704d83146102c6578063f00ac1da146102f6578063fe575a8714610314576100dc565b80637571bdd0146102125780639a6bea181461024e578063a208a89814610278576100dc565b80633ffbd47f116100ba5780633ffbd47f1461015e57806340884c521461018e57806369366c48146101b85780637104ddb2146101e8576100dc565b80630946e807146100e05780630d1db1f01461010a5780632b96895814610134575b5f5ffd5b3480156100eb575f5ffd5b506100f4610350565b6040516101019190611bfa565b60405180910390f35b348015610115575f5ffd5b5061011e610359565b60405161012b9190611c2d565b60405180910390f35b34801561013f575f5ffd5b506101486103a5565b6040516101559190611c2d565b60405180910390f35b61017860048036038101906101739190611d93565b6104a7565b6040516101859190611c2d565b60405180910390f35b348015610199575f5ffd5b506101a2610863565b6040516101af9190611ffb565b60405180910390f35b6101d260048036038101906101cd9190612045565b610a57565b6040516101df9190611c2d565b60405180910390f35b3480156101f3575f5ffd5b506101fc610b47565b60405161020991906120ea565b60405180910390f35b34801561021d575f5ffd5b5061023860048036038101906102339190612045565b610de2565b6040516102459190611c2d565b60405180910390f35b348015610259575f5ffd5b50610262610f09565b60405161026f9190611bfa565b60405180910390f35b610280610f12565b60405161028d9190611c2d565b60405180910390f35b6102b060048036038101906102ab9190612045565b61105c565b6040516102bd9190611c2d565b60405180910390f35b6102e060048036038101906102db9190612134565b61114c565b6040516102ed9190611c2d565b60405180910390f35b6102fe611670565b60405161030b9190611c2d565b60405180910390f35b34801561031f575f5ffd5b5061033a60048036038101906103359190612134565b6117bb565b6040516103479190611c2d565b60405180910390f35b5f600354905090565b5f5f60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f81141591505090565b5f600560149054906101000a900460ff16156103f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ed906121df565b60405180910390fd5b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610485576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047c90612247565b60405180910390fd5b6001600560146101000a81548160ff0219169083151502179055506001905090565b5f6003543410156104ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104e4906122af565b60405180910390fd5b60095f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615610577576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056e9061233d565b60405180910390fd5b60016040518060a001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020016001815260200160011515815250908060018154018082558091505060019003905f5260205f2090600502015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010190816106399190612558565b50604082015181600201908161064f9190612558565b50606082015181600301556080820151816004015f6101000a81548160ff021916908315150217905550505060018054905060085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505f60646002546003546106d59190612654565b6106df91906126c2565b90505f81346106ee91906126f2565b90505f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318a9055c6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561075a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061077e9190612739565b90508073ffffffffffffffffffffffffffffffffffffffff166108fc8490811502906040515f60405180830381858888f193505050501580156107c3573d5f5f3e3d5ffd5b5081600b5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055503373ffffffffffffffffffffffffffffffffffffffff167f8c1c8966369199dc710cd8b615bf16a86520dd165679cb74aa5f427d2bf8c0c9878760405161084e92919061279c565b60405180910390a26001935050505092915050565b60606001805480602002602001604051908101604052809291908181526020015f905b82821015610a4e578382905f5260205f2090600502016040518060a00160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461090b90612388565b80601f016020809104026020016040519081016040528092919081815260200182805461093790612388565b80156109825780601f1061095957610100808354040283529160200191610982565b820191905f5260205f20905b81548152906001019060200180831161096557829003601f168201915b5050505050815260200160028201805461099b90612388565b80601f01602080910402602001604051908101604052809291908181526020018280546109c790612388565b8015610a125780601f106109e957610100808354040283529160200191610a12565b820191905f5260205f20905b8154815290600101906020018083116109f557829003601f168201915b5050505050815260200160038201548152602001600482015f9054906101000a900460ff16151515158152505081526020019060010190610886565b50505050905090565b5f600560149054906101000a900460ff1615610aa8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9f906121df565b60405180910390fd5b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b37576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b2e90612247565b60405180910390fd5b8160048190555060019050919050565b610b4f611b9e565b5f60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f8103610bd2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bc99061281b565b60405180910390fd5b5f60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f60018281548110610c2857610c27612839565b5b905f5260205f2090600502019050806040518060a00160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182018054610ca690612388565b80601f0160208091040260200160405190810160405280929190818152602001828054610cd290612388565b8015610d1d5780601f10610cf457610100808354040283529160200191610d1d565b820191905f5260205f20905b815481529060010190602001808311610d0057829003601f168201915b50505050508152602001600282018054610d3690612388565b80601f0160208091040260200160405190810160405280929190818152602001828054610d6290612388565b8015610dad5780601f10610d8457610100808354040283529160200191610dad565b820191905f5260205f20905b815481529060010190602001808311610d9057829003601f168201915b5050505050815260200160038201548152602001600482015f9054906101000a900460ff161515151581525050935050505090565b5f600560149054906101000a900460ff1615610e33576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e2a906121df565b60405180910390fd5b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ec2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610eb990612247565b60405180910390fd5b816002819055507fa288a9ac9f74fe4c85e8263739645d49297e82ffa010aac11899ba671a54d1ac82604051610ef89190611bfa565b60405180910390a160019050919050565b5f600454905090565b5f5f60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f8103610f96576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f8d9061281b565b60405180910390fd5b5f60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f60018281548110610fec57610feb612839565b5b905f5260205f2090600502016004015f6101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f6eccb14325d715b33601268fe5c7d8a1618261e4d242e890c9e2e372adaf1fbc60405160405180910390a25f9250505090565b5f600560149054906101000a900460ff16156110ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110a4906121df565b60405180910390fd5b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461113c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161113390612247565b60405180910390fd5b8160038190555060019050919050565b5f5f60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f81036111d0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111c79061281b565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361123e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611235906128b0565b60405180910390fd5b60095f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16156112c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112bf90612918565b60405180910390fd5b60065f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff161561138d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161138490612980565b60405180910390fd5b6004543410156113d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113c9906129e8565b60405180910390fd5b600160065f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555060075f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8154809291906114af90612a06565b91905055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fc288b29f22724d188803361dc865debd5a98e3f2c073a19c9fe8639a6274d23d60405160405180910390a36001600260018054905061152291906126c2565b61152c9190612a4d565b60075f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541061166657600160095f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508273ffffffffffffffffffffffffffffffffffffffff167f8f4f70e0f343350f6df22d754271f375cf94e4a4d04ccb29df713c6ffce14dbf60405160405180910390a2600454600a5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546116559190612a4d565b925050819055506116658361180d565b5b6001915050919050565b5f5f60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f81036116f4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116eb9061281b565b60405180910390fd5b5f60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f6001828154811061174a57611749612839565b5b905f5260205f2090600502016004015f6101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f3a21ecfcdc410d1e3552f24e84bbf13af57e47436f3b9ac5045ae9f3adcdc41160405160405180910390a260019250505090565b5f60095f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff169050919050565b5f5f90505f5f90505b60018054905081101561190a5760065f6001838154811061183a57611839612839565b5b905f5260205f2090600502015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16156118fd5781806118f990612a06565b9250505b8080600101915050611816565b505f8160035461191a91906126c2565b90505f5f90505b600180549050811015611b985760065f6001838154811061194557611944612839565b5b905f5260205f2090600502015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff1615611b8b5760018181548110611a0d57611a0c612839565b5b905f5260205f2090600502015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc83600a5f60018681548110611a6c57611a6b612839565b5b905f5260205f2090600502015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054611ade9190612a4d565b90811502906040515f60405180830381858888f19350505050158015611b06573d5f5f3e3d5ffd5b505f600a5f60018481548110611b1f57611b1e612839565b5b905f5260205f2090600502015f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505b8080600101915050611921565b50505050565b6040518060a001604052805f73ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001606081526020015f81526020015f151581525090565b5f819050919050565b611bf481611be2565b82525050565b5f602082019050611c0d5f830184611beb565b92915050565b5f8115159050919050565b611c2781611c13565b82525050565b5f602082019050611c405f830184611c1e565b92915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611ca582611c5f565b810181811067ffffffffffffffff82111715611cc457611cc3611c6f565b5b80604052505050565b5f611cd6611c46565b9050611ce28282611c9c565b919050565b5f67ffffffffffffffff821115611d0157611d00611c6f565b5b611d0a82611c5f565b9050602081019050919050565b828183375f83830152505050565b5f611d37611d3284611ce7565b611ccd565b905082815260208101848484011115611d5357611d52611c5b565b5b611d5e848285611d17565b509392505050565b5f82601f830112611d7a57611d79611c57565b5b8135611d8a848260208601611d25565b91505092915050565b5f5f60408385031215611da957611da8611c4f565b5b5f83013567ffffffffffffffff811115611dc657611dc5611c53565b5b611dd285828601611d66565b925050602083013567ffffffffffffffff811115611df357611df2611c53565b5b611dff85828601611d66565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611e5b82611e32565b9050919050565b611e6b81611e51565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f611ea382611e71565b611ead8185611e7b565b9350611ebd818560208601611e8b565b611ec681611c5f565b840191505092915050565b611eda81611be2565b82525050565b611ee981611c13565b82525050565b5f60a083015f830151611f045f860182611e62565b5060208301518482036020860152611f1c8282611e99565b91505060408301518482036040860152611f368282611e99565b9150506060830151611f4b6060860182611ed1565b506080830151611f5e6080860182611ee0565b508091505092915050565b5f611f748383611eef565b905092915050565b5f602082019050919050565b5f611f9282611e09565b611f9c8185611e13565b935083602082028501611fae85611e23565b805f5b85811015611fe95784840389528151611fca8582611f69565b9450611fd583611f7c565b925060208a01995050600181019050611fb1565b50829750879550505050505092915050565b5f6020820190508181035f8301526120138184611f88565b905092915050565b61202481611be2565b811461202e575f5ffd5b50565b5f8135905061203f8161201b565b92915050565b5f6020828403121561205a57612059611c4f565b5b5f61206784828501612031565b91505092915050565b5f60a083015f8301516120855f860182611e62565b506020830151848203602086015261209d8282611e99565b915050604083015184820360408601526120b78282611e99565b91505060608301516120cc6060860182611ed1565b5060808301516120df6080860182611ee0565b508091505092915050565b5f6020820190508181035f8301526121028184612070565b905092915050565b61211381611e51565b811461211d575f5ffd5b50565b5f8135905061212e8161210a565b92915050565b5f6020828403121561214957612148611c4f565b5b5f61215684828501612120565b91505092915050565b5f82825260208201905092915050565b7f4f776e657273686970206973207265766f6b65642c206e6f2061646d696e69735f8201527f74726174697665206368616e67657320616c6c6f776564210000000000000000602082015250565b5f6121c960388361215f565b91506121d48261216f565b604082019050919050565b5f6020820190508181035f8301526121f6816121bd565b9050919050565b7f4e6f7420617574686f72697a65640000000000000000000000000000000000005f82015250565b5f612231600e8361215f565b915061223c826121fd565b602082019050919050565b5f6020820190508181035f83015261225e81612225565b9050919050565b7f496e73756666696369656e7420726567697374726174696f6e206665650000005f82015250565b5f612299601d8361215f565b91506122a482612265565b602082019050919050565b5f6020820190508181035f8301526122c68161228d565b9050919050565b7f43616e6e6f74207265676973746572206120626c61636b6c6973746564206f725f8201527f61636c6500000000000000000000000000000000000000000000000000000000602082015250565b5f61232760248361215f565b9150612332826122cd565b604082019050919050565b5f6020820190508181035f8301526123548161231b565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061239f57607f821691505b6020821081036123b2576123b161235b565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026124147fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826123d9565b61241e86836123d9565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61245961245461244f84611be2565b612436565b611be2565b9050919050565b5f819050919050565b6124728361243f565b61248661247e82612460565b8484546123e5565b825550505050565b5f5f905090565b61249d61248e565b6124a8818484612469565b505050565b5b818110156124cb576124c05f82612495565b6001810190506124ae565b5050565b601f821115612510576124e1816123b8565b6124ea846123ca565b810160208510156124f9578190505b61250d612505856123ca565b8301826124ad565b50505b505050565b5f82821c905092915050565b5f6125305f1984600802612515565b1980831691505092915050565b5f6125488383612521565b9150826002028217905092915050565b61256182611e71565b67ffffffffffffffff81111561257a57612579611c6f565b5b6125848254612388565b61258f8282856124cf565b5f60209050601f8311600181146125c0575f84156125ae578287015190505b6125b8858261253d565b86555061261f565b601f1984166125ce866123b8565b5f5b828110156125f5578489015182556001820191506020850194506020810190506125d0565b86831015612612578489015161260e601f891682612521565b8355505b6001600288020188555050505b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61265e82611be2565b915061266983611be2565b925082820261267781611be2565b9150828204841483151761268e5761268d612627565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6126cc82611be2565b91506126d783611be2565b9250826126e7576126e6612695565b5b828204905092915050565b5f6126fc82611be2565b915061270783611be2565b925082820390508181111561271f5761271e612627565b5b92915050565b5f815190506127338161210a565b92915050565b5f6020828403121561274e5761274d611c4f565b5b5f61275b84828501612725565b91505092915050565b5f61276e82611e71565b612778818561215f565b9350612788818560208601611e8b565b61279181611c5f565b840191505092915050565b5f6040820190508181035f8301526127b48185612764565b905081810360208301526127c88184612764565b90509392505050565b7f4e6f7420616e20617574686f72697a6564206f7261636c6500000000000000005f82015250565b5f61280560188361215f565b9150612810826127d1565b602082019050919050565b5f6020820190508181035f830152612832816127f9565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f43616e6e6f74207265706f727420796f757273656c66000000000000000000005f82015250565b5f61289a60168361215f565b91506128a582612866565b602082019050919050565b5f6020820190508181035f8301526128c78161288e565b9050919050565b7f4f7261636c6520697320616c726561647920626c61636b6c69737465640000005f82015250565b5f612902601d8361215f565b915061290d826128ce565b602082019050919050565b5f6020820190508181035f83015261292f816128f6565b9050919050565b7f4f7261636c6520616c7265616479207265706f7274656420627920796f7500005f82015250565b5f61296a601e8361215f565b915061297582612936565b602082019050919050565b5f6020820190508181035f8301526129978161295e565b9050919050565b7f496e73756666696369656e74207265706f7274206665650000000000000000005f82015250565b5f6129d260178361215f565b91506129dd8261299e565b602082019050919050565b5f6020820190508181035f8301526129ff816129c6565b9050919050565b5f612a1082611be2565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203612a4257612a41612627565b5b600182019050919050565b5f612a5782611be2565b9150612a6283611be2565b9250828201905080821115612a7a57612a79612627565b5b9291505056fea26469706673582212206ad36085d44235be45e88c7908a552cdb5efc00e30dabdc48de316229bdc059964736f6c634300081c0033",
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
// Solidity: function LogOut() payable returns(bool)
func (_Register *RegisterTransactor) LogOut(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Register.contract.Transact(opts, "LogOut")
}

// LogOut is a paid mutator transaction binding the contract method 0xa208a898.
//
// Solidity: function LogOut() payable returns(bool)
func (_Register *RegisterSession) LogOut() (*types.Transaction, error) {
	return _Register.Contract.LogOut(&_Register.TransactOpts)
}

// LogOut is a paid mutator transaction binding the contract method 0xa208a898.
//
// Solidity: function LogOut() payable returns(bool)
func (_Register *RegisterTransactorSession) LogOut() (*types.Transaction, error) {
	return _Register.Contract.LogOut(&_Register.TransactOpts)
}

// Login is a paid mutator transaction binding the contract method 0xf00ac1da.
//
// Solidity: function Login() payable returns(bool)
func (_Register *RegisterTransactor) Login(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Register.contract.Transact(opts, "Login")
}

// Login is a paid mutator transaction binding the contract method 0xf00ac1da.
//
// Solidity: function Login() payable returns(bool)
func (_Register *RegisterSession) Login() (*types.Transaction, error) {
	return _Register.Contract.Login(&_Register.TransactOpts)
}

// Login is a paid mutator transaction binding the contract method 0xf00ac1da.
//
// Solidity: function Login() payable returns(bool)
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
