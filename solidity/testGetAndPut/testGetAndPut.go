// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testGetAndPut

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
)

// TestGetAndPutMetaData contains all meta data concerning the TestGetAndPut contract.
var TestGetAndPutMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061012f806100206000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80632e64cec11460375780636057361d146051575b600080fd5b603d6069565b6040516048919060c2565b60405180910390f35b6067600480360381019060639190608f565b6072565b005b60008054905090565b8060008190555050565b60008135905060898160e5565b92915050565b60006020828403121560a057600080fd5b600060ac84828501607c565b91505092915050565b60bc8160db565b82525050565b600060208201905060d5600083018460b5565b92915050565b6000819050919050565b60ec8160db565b811460f657600080fd5b5056fea26469706673582212204f6065327406dd800b47cd88f01f94f99369c4f605861f49fb21846078528a5b64736f6c63430008040033",
}

// TestGetAndPutABI is the input ABI used to generate the binding from.
// Deprecated: Use TestGetAndPutMetaData.ABI instead.
var TestGetAndPutABI = TestGetAndPutMetaData.ABI

// TestGetAndPutBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestGetAndPutMetaData.Bin instead.
var TestGetAndPutBin = TestGetAndPutMetaData.Bin

// DeployTestGetAndPut deploys a new Ethereum contract, binding an instance of TestGetAndPut to it.
func DeployTestGetAndPut(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestGetAndPut, error) {
	parsed, err := TestGetAndPutMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestGetAndPutBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestGetAndPut{TestGetAndPutCaller: TestGetAndPutCaller{contract: contract}, TestGetAndPutTransactor: TestGetAndPutTransactor{contract: contract}, TestGetAndPutFilterer: TestGetAndPutFilterer{contract: contract}}, nil
}

// TestGetAndPut is an auto generated Go binding around an Ethereum contract.
type TestGetAndPut struct {
	TestGetAndPutCaller     // Read-only binding to the contract
	TestGetAndPutTransactor // Write-only binding to the contract
	TestGetAndPutFilterer   // Log filterer for contract events
}

// TestGetAndPutCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestGetAndPutCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestGetAndPutTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestGetAndPutTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestGetAndPutFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestGetAndPutFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestGetAndPutSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestGetAndPutSession struct {
	Contract     *TestGetAndPut    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestGetAndPutCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestGetAndPutCallerSession struct {
	Contract *TestGetAndPutCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TestGetAndPutTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestGetAndPutTransactorSession struct {
	Contract     *TestGetAndPutTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TestGetAndPutRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestGetAndPutRaw struct {
	Contract *TestGetAndPut // Generic contract binding to access the raw methods on
}

// TestGetAndPutCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestGetAndPutCallerRaw struct {
	Contract *TestGetAndPutCaller // Generic read-only contract binding to access the raw methods on
}

// TestGetAndPutTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestGetAndPutTransactorRaw struct {
	Contract *TestGetAndPutTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestGetAndPut creates a new instance of TestGetAndPut, bound to a specific deployed contract.
func NewTestGetAndPut(address common.Address, backend bind.ContractBackend) (*TestGetAndPut, error) {
	contract, err := bindTestGetAndPut(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestGetAndPut{TestGetAndPutCaller: TestGetAndPutCaller{contract: contract}, TestGetAndPutTransactor: TestGetAndPutTransactor{contract: contract}, TestGetAndPutFilterer: TestGetAndPutFilterer{contract: contract}}, nil
}

// NewTestGetAndPutCaller creates a new read-only instance of TestGetAndPut, bound to a specific deployed contract.
func NewTestGetAndPutCaller(address common.Address, caller bind.ContractCaller) (*TestGetAndPutCaller, error) {
	contract, err := bindTestGetAndPut(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestGetAndPutCaller{contract: contract}, nil
}

// NewTestGetAndPutTransactor creates a new write-only instance of TestGetAndPut, bound to a specific deployed contract.
func NewTestGetAndPutTransactor(address common.Address, transactor bind.ContractTransactor) (*TestGetAndPutTransactor, error) {
	contract, err := bindTestGetAndPut(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestGetAndPutTransactor{contract: contract}, nil
}

// NewTestGetAndPutFilterer creates a new log filterer instance of TestGetAndPut, bound to a specific deployed contract.
func NewTestGetAndPutFilterer(address common.Address, filterer bind.ContractFilterer) (*TestGetAndPutFilterer, error) {
	contract, err := bindTestGetAndPut(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestGetAndPutFilterer{contract: contract}, nil
}

// bindTestGetAndPut binds a generic wrapper to an already deployed contract.
func bindTestGetAndPut(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestGetAndPutABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestGetAndPut *TestGetAndPutRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestGetAndPut.Contract.TestGetAndPutCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestGetAndPut *TestGetAndPutRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestGetAndPut.Contract.TestGetAndPutTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestGetAndPut *TestGetAndPutRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestGetAndPut.Contract.TestGetAndPutTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestGetAndPut *TestGetAndPutCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestGetAndPut.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestGetAndPut *TestGetAndPutTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestGetAndPut.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestGetAndPut *TestGetAndPutTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestGetAndPut.Contract.contract.Transact(opts, method, params...)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_TestGetAndPut *TestGetAndPutCaller) Retrieve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestGetAndPut.contract.Call(opts, &out, "retrieve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_TestGetAndPut *TestGetAndPutSession) Retrieve() (*big.Int, error) {
	return _TestGetAndPut.Contract.Retrieve(&_TestGetAndPut.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_TestGetAndPut *TestGetAndPutCallerSession) Retrieve() (*big.Int, error) {
	return _TestGetAndPut.Contract.Retrieve(&_TestGetAndPut.CallOpts)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_TestGetAndPut *TestGetAndPutTransactor) Store(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _TestGetAndPut.contract.Transact(opts, "store", num)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_TestGetAndPut *TestGetAndPutSession) Store(num *big.Int) (*types.Transaction, error) {
	return _TestGetAndPut.Contract.Store(&_TestGetAndPut.TransactOpts, num)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_TestGetAndPut *TestGetAndPutTransactorSession) Store(num *big.Int) (*types.Transaction, error) {
	return _TestGetAndPut.Contract.Store(&_TestGetAndPut.TransactOpts, num)
}
