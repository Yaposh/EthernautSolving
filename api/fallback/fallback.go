// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fallback

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FallbackABI is the input ABI used to generate the binding from.
const FallbackABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"contribute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"contributions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContribution\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// FallbackBin is the compiled bytecode used for deploying new contracts.
var FallbackBin = "0x608060405234801561001057600080fd5b5033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550683635c9adc5dea000006000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506106af806100ad6000396000f3fe60806040526004361061004e5760003560e01c80633ccfd60b146100f257806342e94c90146101095780638da5cb5b14610146578063d7bb99ba14610171578063f10fdf5c1461017b576100ed565b366100ed576000341180156100a1575060008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054115b6100aa57600080fd5b33600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550005b600080fd5b3480156100fe57600080fd5b506101076101a6565b005b34801561011557600080fd5b50610130600480360381019061012b919061048c565b6102a1565b60405161013d9190610535565b60405180910390f35b34801561015257600080fd5b5061015b6102b9565b60405161016891906104fa565b60405180910390f35b6101796102df565b005b34801561018757600080fd5b50610190610431565b60405161019d9190610535565b60405180910390f35b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610236576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161022d90610515565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f1935050505015801561029e573d6000803e3d6000fd5b50565b60006020528060005260406000206000915090505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b66038d7ea4c6800034106102f257600080fd5b346000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546103409190610561565b92505081905550600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054111561042f5733600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b60008135905061048681610662565b92915050565b6000602082840312156104a2576104a1610634565b5b60006104b084828501610477565b91505092915050565b6104c2816105c9565b82525050565b60006104d5601783610550565b91506104e082610639565b602082019050919050565b6104f4816105fb565b82525050565b600060208201905061050f60008301846104b9565b92915050565b6000602082019050818103600083015261052e816104c8565b9050919050565b600060208201905061054a60008301846104eb565b92915050565b600082825260208201905092915050565b600061056c826105fb565b9150610577836105fb565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156105ac576105ab610605565b5b828201905092915050565b60006105c2826105db565b9050919050565b60006105d4826105db565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600080fd5b7f63616c6c6572206973206e6f7420746865206f776e6572000000000000000000600082015250565b61066b816105b7565b811461067657600080fd5b5056fea26469706673582212201e115fee1b1d58c9dd5e3639d5c1d3b207f9c0963a3e6807f5af4c96e8b0788264736f6c63430008060033"

// DeployFallback deploys a new Ethereum contract, binding an instance of Fallback to it.
func DeployFallback(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Fallback, error) {
	parsed, err := abi.JSON(strings.NewReader(FallbackABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FallbackBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Fallback{FallbackCaller: FallbackCaller{contract: contract}, FallbackTransactor: FallbackTransactor{contract: contract}, FallbackFilterer: FallbackFilterer{contract: contract}}, nil
}

// Fallback is an auto generated Go binding around an Ethereum contract.
type Fallback struct {
	FallbackCaller     // Read-only binding to the contract
	FallbackTransactor // Write-only binding to the contract
	FallbackFilterer   // Log filterer for contract events
}

// FallbackCaller is an auto generated read-only Go binding around an Ethereum contract.
type FallbackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FallbackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FallbackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FallbackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FallbackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FallbackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FallbackSession struct {
	Contract     *Fallback         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FallbackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FallbackCallerSession struct {
	Contract *FallbackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FallbackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FallbackTransactorSession struct {
	Contract     *FallbackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FallbackRaw is an auto generated low-level Go binding around an Ethereum contract.
type FallbackRaw struct {
	Contract *Fallback // Generic contract binding to access the raw methods on
}

// FallbackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FallbackCallerRaw struct {
	Contract *FallbackCaller // Generic read-only contract binding to access the raw methods on
}

// FallbackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FallbackTransactorRaw struct {
	Contract *FallbackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFallback creates a new instance of Fallback, bound to a specific deployed contract.
func NewFallback(address common.Address, backend bind.ContractBackend) (*Fallback, error) {
	contract, err := bindFallback(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fallback{FallbackCaller: FallbackCaller{contract: contract}, FallbackTransactor: FallbackTransactor{contract: contract}, FallbackFilterer: FallbackFilterer{contract: contract}}, nil
}

// NewFallbackCaller creates a new read-only instance of Fallback, bound to a specific deployed contract.
func NewFallbackCaller(address common.Address, caller bind.ContractCaller) (*FallbackCaller, error) {
	contract, err := bindFallback(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FallbackCaller{contract: contract}, nil
}

// NewFallbackTransactor creates a new write-only instance of Fallback, bound to a specific deployed contract.
func NewFallbackTransactor(address common.Address, transactor bind.ContractTransactor) (*FallbackTransactor, error) {
	contract, err := bindFallback(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FallbackTransactor{contract: contract}, nil
}

// NewFallbackFilterer creates a new log filterer instance of Fallback, bound to a specific deployed contract.
func NewFallbackFilterer(address common.Address, filterer bind.ContractFilterer) (*FallbackFilterer, error) {
	contract, err := bindFallback(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FallbackFilterer{contract: contract}, nil
}

// bindFallback binds a generic wrapper to an already deployed contract.
func bindFallback(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FallbackABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fallback *FallbackRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fallback.Contract.FallbackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fallback *FallbackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fallback.Contract.FallbackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fallback *FallbackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fallback.Contract.FallbackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fallback *FallbackCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fallback.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fallback *FallbackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fallback.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fallback *FallbackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fallback.Contract.contract.Transact(opts, method, params...)
}

// Contributions is a free data retrieval call binding the contract method 0x42e94c90.
//
// Solidity: function contributions(address ) view returns(uint256)
func (_Fallback *FallbackCaller) Contributions(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Fallback.contract.Call(opts, &out, "contributions", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Contributions is a free data retrieval call binding the contract method 0x42e94c90.
//
// Solidity: function contributions(address ) view returns(uint256)
func (_Fallback *FallbackSession) Contributions(arg0 common.Address) (*big.Int, error) {
	return _Fallback.Contract.Contributions(&_Fallback.CallOpts, arg0)
}

// Contributions is a free data retrieval call binding the contract method 0x42e94c90.
//
// Solidity: function contributions(address ) view returns(uint256)
func (_Fallback *FallbackCallerSession) Contributions(arg0 common.Address) (*big.Int, error) {
	return _Fallback.Contract.Contributions(&_Fallback.CallOpts, arg0)
}

// GetContribution is a free data retrieval call binding the contract method 0xf10fdf5c.
//
// Solidity: function getContribution() view returns(uint256)
func (_Fallback *FallbackCaller) GetContribution(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Fallback.contract.Call(opts, &out, "getContribution")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContribution is a free data retrieval call binding the contract method 0xf10fdf5c.
//
// Solidity: function getContribution() view returns(uint256)
func (_Fallback *FallbackSession) GetContribution() (*big.Int, error) {
	return _Fallback.Contract.GetContribution(&_Fallback.CallOpts)
}

// GetContribution is a free data retrieval call binding the contract method 0xf10fdf5c.
//
// Solidity: function getContribution() view returns(uint256)
func (_Fallback *FallbackCallerSession) GetContribution() (*big.Int, error) {
	return _Fallback.Contract.GetContribution(&_Fallback.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Fallback *FallbackCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fallback.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Fallback *FallbackSession) Owner() (common.Address, error) {
	return _Fallback.Contract.Owner(&_Fallback.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Fallback *FallbackCallerSession) Owner() (common.Address, error) {
	return _Fallback.Contract.Owner(&_Fallback.CallOpts)
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_Fallback *FallbackTransactor) Contribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fallback.contract.Transact(opts, "contribute")
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_Fallback *FallbackSession) Contribute() (*types.Transaction, error) {
	return _Fallback.Contract.Contribute(&_Fallback.TransactOpts)
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_Fallback *FallbackTransactorSession) Contribute() (*types.Transaction, error) {
	return _Fallback.Contract.Contribute(&_Fallback.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Fallback *FallbackTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fallback.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Fallback *FallbackSession) Withdraw() (*types.Transaction, error) {
	return _Fallback.Contract.Withdraw(&_Fallback.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Fallback *FallbackTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Fallback.Contract.Withdraw(&_Fallback.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Fallback *FallbackTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fallback.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Fallback *FallbackSession) Receive() (*types.Transaction, error) {
	return _Fallback.Contract.Receive(&_Fallback.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Fallback *FallbackTransactorSession) Receive() (*types.Transaction, error) {
	return _Fallback.Contract.Receive(&_Fallback.TransactOpts)
}
