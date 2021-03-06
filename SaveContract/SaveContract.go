// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractapi

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

// ContractapiMetaData contains all meta data concerning the Contractapi contract.
var ContractapiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ex1\",\"type\":\"string\"}],\"name\":\"getData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ex1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ex2\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ex3\",\"type\":\"string\"}],\"name\":\"setData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506106cb806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063280af7d31461003b578063ae55c88814610050575b600080fd5b61004e6100493660046104da565b61007b565b005b61006361005e366004610562565b610154565b604051610072939291906105fb565b60405180910390f35b60008360405161008b919061063e565b9081526040519081900360200190205460ff16156100a857600080fd5b6040518060800160405280600115158152602001848152602001838152602001828152506000846040516100dc919061063e565b908152604051602091819003820190208251815460ff191690151517815582820151805191926101149260018501929091019061039e565b506040820151805161013091600284019160209091019061039e565b506060820151805161014c91600384019160209091019061039e565b505050505050565b6060806060600084604051610169919061063e565b9081526040519081900360200190205460ff16151560011461018a57600080fd5b60008460405161019a919061063e565b90815260200160405180910390206001016000856040516101bb919061063e565b90815260200160405180910390206002016000866040516101dc919061063e565b90815260200160405180910390206003018280546101f99061065a565b80601f01602080910402602001604051908101604052809291908181526020018280546102259061065a565b80156102725780601f1061024757610100808354040283529160200191610272565b820191906000526020600020905b81548152906001019060200180831161025557829003601f168201915b505050505092508180546102859061065a565b80601f01602080910402602001604051908101604052809291908181526020018280546102b19061065a565b80156102fe5780601f106102d3576101008083540402835291602001916102fe565b820191906000526020600020905b8154815290600101906020018083116102e157829003601f168201915b505050505091508080546103119061065a565b80601f016020809104026020016040519081016040528092919081815260200182805461033d9061065a565b801561038a5780601f1061035f5761010080835404028352916020019161038a565b820191906000526020600020905b81548152906001019060200180831161036d57829003601f168201915b505050505090509250925092509193909250565b8280546103aa9061065a565b90600052602060002090601f0160209004810192826103cc5760008555610412565b82601f106103e557805160ff1916838001178555610412565b82800160010185558215610412579182015b828111156104125782518255916020019190600101906103f7565b5061041e929150610422565b5090565b5b8082111561041e5760008155600101610423565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261045e57600080fd5b813567ffffffffffffffff8082111561047957610479610437565b604051601f8301601f19908116603f011681019082821181831017156104a1576104a1610437565b816040528381528660208588010111156104ba57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000606084860312156104ef57600080fd5b833567ffffffffffffffff8082111561050757600080fd5b6105138783880161044d565b9450602086013591508082111561052957600080fd5b6105358783880161044d565b9350604086013591508082111561054b57600080fd5b506105588682870161044d565b9150509250925092565b60006020828403121561057457600080fd5b813567ffffffffffffffff81111561058b57600080fd5b6105978482850161044d565b949350505050565b60005b838110156105ba5781810151838201526020016105a2565b838111156105c9576000848401525b50505050565b600081518084526105e781602086016020860161059f565b601f01601f19169290920160200192915050565b60608152600061060e60608301866105cf565b828103602084015261062081866105cf565b9050828103604084015261063481856105cf565b9695505050505050565b6000825161065081846020870161059f565b9190910192915050565b600181811c9082168061066e57607f821691505b6020821081141561068f57634e487b7160e01b600052602260045260246000fd5b5091905056fea2646970667358221220e30e8e1288eec3b8f7d95f37cb0be8b9544837e4e9a4e903c72b7fbf5e81f59764736f6c634300080c0033",
}

// ContractapiABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractapiMetaData.ABI instead.
var ContractapiABI = ContractapiMetaData.ABI

// ContractapiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractapiMetaData.Bin instead.
var ContractapiBin = ContractapiMetaData.Bin

// DeployContractapi deploys a new Ethereum contract, binding an instance of Contractapi to it.
func DeployContractapi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contractapi, error) {
	parsed, err := ContractapiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractapiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contractapi{ContractapiCaller: ContractapiCaller{contract: contract}, ContractapiTransactor: ContractapiTransactor{contract: contract}, ContractapiFilterer: ContractapiFilterer{contract: contract}}, nil
}

// Contractapi is an auto generated Go binding around an Ethereum contract.
type Contractapi struct {
	ContractapiCaller     // Read-only binding to the contract
	ContractapiTransactor // Write-only binding to the contract
	ContractapiFilterer   // Log filterer for contract events
}

// ContractapiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractapiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractapiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractapiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractapiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractapiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractapiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractapiSession struct {
	Contract     *Contractapi      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractapiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractapiCallerSession struct {
	Contract *ContractapiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ContractapiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractapiTransactorSession struct {
	Contract     *ContractapiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ContractapiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractapiRaw struct {
	Contract *Contractapi // Generic contract binding to access the raw methods on
}

// ContractapiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractapiCallerRaw struct {
	Contract *ContractapiCaller // Generic read-only contract binding to access the raw methods on
}

// ContractapiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractapiTransactorRaw struct {
	Contract *ContractapiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractapi creates a new instance of Contractapi, bound to a specific deployed contract.
func NewContractapi(address common.Address, backend bind.ContractBackend) (*Contractapi, error) {
	contract, err := bindContractapi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contractapi{ContractapiCaller: ContractapiCaller{contract: contract}, ContractapiTransactor: ContractapiTransactor{contract: contract}, ContractapiFilterer: ContractapiFilterer{contract: contract}}, nil
}

// NewContractapiCaller creates a new read-only instance of Contractapi, bound to a specific deployed contract.
func NewContractapiCaller(address common.Address, caller bind.ContractCaller) (*ContractapiCaller, error) {
	contract, err := bindContractapi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractapiCaller{contract: contract}, nil
}

// NewContractapiTransactor creates a new write-only instance of Contractapi, bound to a specific deployed contract.
func NewContractapiTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractapiTransactor, error) {
	contract, err := bindContractapi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractapiTransactor{contract: contract}, nil
}

// NewContractapiFilterer creates a new log filterer instance of Contractapi, bound to a specific deployed contract.
func NewContractapiFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractapiFilterer, error) {
	contract, err := bindContractapi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractapiFilterer{contract: contract}, nil
}

// bindContractapi binds a generic wrapper to an already deployed contract.
func bindContractapi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractapiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractapi *ContractapiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractapi.Contract.ContractapiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractapi *ContractapiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractapi.Contract.ContractapiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractapi *ContractapiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractapi.Contract.ContractapiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contractapi *ContractapiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contractapi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contractapi *ContractapiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contractapi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contractapi *ContractapiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contractapi.Contract.contract.Transact(opts, method, params...)
}

// GetData is a free data retrieval call binding the contract method 0xae55c888.
//
// Solidity: function getData(string _ex1) view returns(string, string, string)
func (_Contractapi *ContractapiCaller) GetData(opts *bind.CallOpts, _ex1 string) (string, string, string, error) {
	var out []interface{}
	err := _Contractapi.contract.Call(opts, &out, "getData", _ex1)

	if err != nil {
		return *new(string), *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)

	return out0, out1, out2, err

}

// GetData is a free data retrieval call binding the contract method 0xae55c888.
//
// Solidity: function getData(string _ex1) view returns(string, string, string)
func (_Contractapi *ContractapiSession) GetData(_ex1 string) (string, string, string, error) {
	return _Contractapi.Contract.GetData(&_Contractapi.CallOpts, _ex1)
}

// GetData is a free data retrieval call binding the contract method 0xae55c888.
//
// Solidity: function getData(string _ex1) view returns(string, string, string)
func (_Contractapi *ContractapiCallerSession) GetData(_ex1 string) (string, string, string, error) {
	return _Contractapi.Contract.GetData(&_Contractapi.CallOpts, _ex1)
}

// SetData is a paid mutator transaction binding the contract method 0x280af7d3.
//
// Solidity: function setData(string _ex1, string _ex2, string _ex3) returns()
func (_Contractapi *ContractapiTransactor) SetData(opts *bind.TransactOpts, _ex1 string, _ex2 string, _ex3 string) (*types.Transaction, error) {
	return _Contractapi.contract.Transact(opts, "setData", _ex1, _ex2, _ex3)
}

// SetData is a paid mutator transaction binding the contract method 0x280af7d3.
//
// Solidity: function setData(string _ex1, string _ex2, string _ex3) returns()
func (_Contractapi *ContractapiSession) SetData(_ex1 string, _ex2 string, _ex3 string) (*types.Transaction, error) {
	return _Contractapi.Contract.SetData(&_Contractapi.TransactOpts, _ex1, _ex2, _ex3)
}

// SetData is a paid mutator transaction binding the contract method 0x280af7d3.
//
// Solidity: function setData(string _ex1, string _ex2, string _ex3) returns()
func (_Contractapi *ContractapiTransactorSession) SetData(_ex1 string, _ex2 string, _ex3 string) (*types.Transaction, error) {
	return _Contractapi.Contract.SetData(&_Contractapi.TransactOpts, _ex1, _ex2, _ex3)
}
