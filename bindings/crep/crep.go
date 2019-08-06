// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BindingsABI is the input ABI used to generate the binding from.
const BindingsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x06fdde03\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x095ea7b3\"},{\"constant\":false,\"inputs\":[{\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrow\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x0e752702\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserveFactorMantissa\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x173b9904\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceCurrent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x17bfdfbc\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x18160ddd\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchangeRateStored\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x182df0f5\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x23b872dd\"},{\"constant\":false,\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\"},{\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowBehalf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x2608f818\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x26782247\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x313ce567\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x3af9e669\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x3b1d21a2\"},{\"constant\":false,\"inputs\":[{\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"_setComptroller\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x4576b5db\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x47bd3718\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptroller\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x5fe3b567\"},{\"constant\":false,\"inputs\":[{\"name\":\"reduceAmount\",\"type\":\"uint256\"}],\"name\":\"_reduceReserves\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x601a0bf1\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialExchangeRateMantissa\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x675d972c\"},{\"constant\":true,\"inputs\":[],\"name\":\"accrualBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x6c540baf\"},{\"constant\":true,\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x6f307dc3\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x70a08231\"},{\"constant\":false,\"inputs\":[],\"name\":\"totalBorrowsCurrent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x73acee98\"},{\"constant\":false,\"inputs\":[{\"name\":\"redeemAmount\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x852a12e3\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalReserves\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x8f840ddd\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x95d89b41\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceStored\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x95dd9193\"},{\"constant\":false,\"inputs\":[{\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xa0712d68\"},{\"constant\":false,\"inputs\":[],\"name\":\"accrueInterest\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xa6afed95\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xa9059cbb\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xaa5af0fd\"},{\"constant\":true,\"inputs\":[],\"name\":\"supplyRatePerBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xae9d70b0\"},{\"constant\":false,\"inputs\":[{\"name\":\"liquidator\",\"type\":\"address\"},{\"name\":\"borrower\",\"type\":\"address\"},{\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xb2a02ff1\"},{\"constant\":false,\"inputs\":[{\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xb71d1a0c\"},{\"constant\":false,\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xbd6d894d\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountSnapshot\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xc37f68e2\"},{\"constant\":false,\"inputs\":[{\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xc5ebeaec\"},{\"constant\":false,\"inputs\":[{\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xdb006a75\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xdd62ed3e\"},{\"constant\":false,\"inputs\":[],\"name\":\"_acceptAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xe9c714f2\"},{\"constant\":false,\"inputs\":[{\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"_setInterestRateModel\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xf2b3abbd\"},{\"constant\":true,\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xf3fdb15a\"},{\"constant\":false,\"inputs\":[{\"name\":\"borrower\",\"type\":\"address\"},{\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"name\":\"cTokenCollateral\",\"type\":\"address\"}],\"name\":\"liquidateBorrow\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xf5e3c462\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xf851a440\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowRatePerBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xf8f9da28\"},{\"constant\":false,\"inputs\":[{\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setReserveFactor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xfca7820b\"},{\"constant\":true,\"inputs\":[],\"name\":\"isCToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xfe9c44ae\"},{\"inputs\":[{\"name\":\"underlying_\",\"type\":\"address\"},{\"name\":\"comptroller_\",\"type\":\"address\"},{\"name\":\"interestRateModel_\",\"type\":\"address\"},{\"name\":\"initialExchangeRateMantissa_\",\"type\":\"uint256\"},{\"name\":\"name_\",\"type\":\"string\"},{\"name\":\"symbol_\",\"type\":\"string\"},{\"name\":\"decimals_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"signature\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"interestAccumulated\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"borrowIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"AccrueInterest\",\"type\":\"event\",\"signature\":\"0x875352fb3fadeb8c0be7cbbe8ff761b308fa7033470cd0287f02f3436fd76cb9\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\",\"signature\":\"0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\",\"signature\":\"0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\",\"signature\":\"0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"RepayBorrow\",\"type\":\"event\",\"signature\":\"0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"LiquidateBorrow\",\"type\":\"event\",\"signature\":\"0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldPendingAdmin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\",\"signature\":\"0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\",\"signature\":\"0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldComptroller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"NewComptroller\",\"type\":\"event\",\"signature\":\"0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldInterestRateModel\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"NewMarketInterestRateModel\",\"type\":\"event\",\"signature\":\"0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldReserveFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewReserveFactor\",\"type\":\"event\",\"signature\":\"0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"reduceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesReduced\",\"type\":\"event\",\"signature\":\"0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\",\"signature\":\"0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\",\"signature\":\"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\",\"signature\":\"0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925\"}]"

// Bindings is an auto generated Go binding around an Ethereum contract.
type Bindings struct {
	BindingsCaller     // Read-only binding to the contract
	BindingsTransactor // Write-only binding to the contract
	BindingsFilterer   // Log filterer for contract events
}

// BindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingsSession struct {
	Contract     *Bindings         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingsCallerSession struct {
	Contract *BindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingsTransactorSession struct {
	Contract     *BindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingsRaw struct {
	Contract *Bindings // Generic contract binding to access the raw methods on
}

// BindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingsCallerRaw struct {
	Contract *BindingsCaller // Generic read-only contract binding to access the raw methods on
}

// BindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingsTransactorRaw struct {
	Contract *BindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindings creates a new instance of Bindings, bound to a specific deployed contract.
func NewBindings(address common.Address, backend bind.ContractBackend) (*Bindings, error) {
	contract, err := bindBindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// NewBindingsCaller creates a new read-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsCaller(address common.Address, caller bind.ContractCaller) (*BindingsCaller, error) {
	contract, err := bindBindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsCaller{contract: contract}, nil
}

// NewBindingsTransactor creates a new write-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingsTransactor, error) {
	contract, err := bindBindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsTransactor{contract: contract}, nil
}

// NewBindingsFilterer creates a new log filterer instance of Bindings, bound to a specific deployed contract.
func NewBindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingsFilterer, error) {
	contract, err := bindBindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingsFilterer{contract: contract}, nil
}

// bindBindings binds a generic wrapper to an already deployed contract.
func bindBindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BindingsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.BindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transact(opts, method, params...)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() constant returns(uint256)
func (_Bindings *BindingsCaller) AccrualBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "accrualBlockNumber")
	return *ret0, err
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() constant returns(uint256)
func (_Bindings *BindingsSession) AccrualBlockNumber() (*big.Int, error) {
	return _Bindings.Contract.AccrualBlockNumber(&_Bindings.CallOpts)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() constant returns(uint256)
func (_Bindings *BindingsCallerSession) AccrualBlockNumber() (*big.Int, error) {
	return _Bindings.Contract.AccrualBlockNumber(&_Bindings.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Bindings *BindingsCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Bindings *BindingsSession) Admin() (common.Address, error) {
	return _Bindings.Contract.Admin(&_Bindings.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Bindings *BindingsCallerSession) Admin() (common.Address, error) {
	return _Bindings.Contract.Admin(&_Bindings.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_Bindings *BindingsCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_Bindings *BindingsSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Bindings.Contract.Allowance(&_Bindings.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_Bindings *BindingsCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Bindings.Contract.Allowance(&_Bindings.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_Bindings *BindingsCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_Bindings *BindingsSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Bindings.Contract.BalanceOf(&_Bindings.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_Bindings *BindingsCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Bindings.Contract.BalanceOf(&_Bindings.CallOpts, owner)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) constant returns(uint256)
func (_Bindings *BindingsCaller) BorrowBalanceStored(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "borrowBalanceStored", account)
	return *ret0, err
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) constant returns(uint256)
func (_Bindings *BindingsSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _Bindings.Contract.BorrowBalanceStored(&_Bindings.CallOpts, account)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) constant returns(uint256)
func (_Bindings *BindingsCallerSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _Bindings.Contract.BorrowBalanceStored(&_Bindings.CallOpts, account)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() constant returns(uint256)
func (_Bindings *BindingsCaller) BorrowIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "borrowIndex")
	return *ret0, err
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() constant returns(uint256)
func (_Bindings *BindingsSession) BorrowIndex() (*big.Int, error) {
	return _Bindings.Contract.BorrowIndex(&_Bindings.CallOpts)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() constant returns(uint256)
func (_Bindings *BindingsCallerSession) BorrowIndex() (*big.Int, error) {
	return _Bindings.Contract.BorrowIndex(&_Bindings.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() constant returns(uint256)
func (_Bindings *BindingsCaller) BorrowRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "borrowRatePerBlock")
	return *ret0, err
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() constant returns(uint256)
func (_Bindings *BindingsSession) BorrowRatePerBlock() (*big.Int, error) {
	return _Bindings.Contract.BorrowRatePerBlock(&_Bindings.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() constant returns(uint256)
func (_Bindings *BindingsCallerSession) BorrowRatePerBlock() (*big.Int, error) {
	return _Bindings.Contract.BorrowRatePerBlock(&_Bindings.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() constant returns(address)
func (_Bindings *BindingsCaller) Comptroller(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "comptroller")
	return *ret0, err
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() constant returns(address)
func (_Bindings *BindingsSession) Comptroller() (common.Address, error) {
	return _Bindings.Contract.Comptroller(&_Bindings.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() constant returns(address)
func (_Bindings *BindingsCallerSession) Comptroller() (common.Address, error) {
	return _Bindings.Contract.Comptroller(&_Bindings.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Bindings *BindingsCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Bindings *BindingsSession) Decimals() (*big.Int, error) {
	return _Bindings.Contract.Decimals(&_Bindings.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Bindings *BindingsCallerSession) Decimals() (*big.Int, error) {
	return _Bindings.Contract.Decimals(&_Bindings.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() constant returns(uint256)
func (_Bindings *BindingsCaller) ExchangeRateStored(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "exchangeRateStored")
	return *ret0, err
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() constant returns(uint256)
func (_Bindings *BindingsSession) ExchangeRateStored() (*big.Int, error) {
	return _Bindings.Contract.ExchangeRateStored(&_Bindings.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() constant returns(uint256)
func (_Bindings *BindingsCallerSession) ExchangeRateStored() (*big.Int, error) {
	return _Bindings.Contract.ExchangeRateStored(&_Bindings.CallOpts)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) constant returns(uint256, uint256, uint256, uint256)
func (_Bindings *BindingsCaller) GetAccountSnapshot(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Bindings.contract.Call(opts, out, "getAccountSnapshot", account)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) constant returns(uint256, uint256, uint256, uint256)
func (_Bindings *BindingsSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Bindings.Contract.GetAccountSnapshot(&_Bindings.CallOpts, account)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) constant returns(uint256, uint256, uint256, uint256)
func (_Bindings *BindingsCallerSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Bindings.Contract.GetAccountSnapshot(&_Bindings.CallOpts, account)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() constant returns(uint256)
func (_Bindings *BindingsCaller) GetCash(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "getCash")
	return *ret0, err
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() constant returns(uint256)
func (_Bindings *BindingsSession) GetCash() (*big.Int, error) {
	return _Bindings.Contract.GetCash(&_Bindings.CallOpts)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() constant returns(uint256)
func (_Bindings *BindingsCallerSession) GetCash() (*big.Int, error) {
	return _Bindings.Contract.GetCash(&_Bindings.CallOpts)
}

// InitialExchangeRateMantissa is a free data retrieval call binding the contract method 0x675d972c.
//
// Solidity: function initialExchangeRateMantissa() constant returns(uint256)
func (_Bindings *BindingsCaller) InitialExchangeRateMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "initialExchangeRateMantissa")
	return *ret0, err
}

// InitialExchangeRateMantissa is a free data retrieval call binding the contract method 0x675d972c.
//
// Solidity: function initialExchangeRateMantissa() constant returns(uint256)
func (_Bindings *BindingsSession) InitialExchangeRateMantissa() (*big.Int, error) {
	return _Bindings.Contract.InitialExchangeRateMantissa(&_Bindings.CallOpts)
}

// InitialExchangeRateMantissa is a free data retrieval call binding the contract method 0x675d972c.
//
// Solidity: function initialExchangeRateMantissa() constant returns(uint256)
func (_Bindings *BindingsCallerSession) InitialExchangeRateMantissa() (*big.Int, error) {
	return _Bindings.Contract.InitialExchangeRateMantissa(&_Bindings.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() constant returns(address)
func (_Bindings *BindingsCaller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "interestRateModel")
	return *ret0, err
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() constant returns(address)
func (_Bindings *BindingsSession) InterestRateModel() (common.Address, error) {
	return _Bindings.Contract.InterestRateModel(&_Bindings.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() constant returns(address)
func (_Bindings *BindingsCallerSession) InterestRateModel() (common.Address, error) {
	return _Bindings.Contract.InterestRateModel(&_Bindings.CallOpts)
}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() constant returns(bool)
func (_Bindings *BindingsCaller) IsCToken(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "isCToken")
	return *ret0, err
}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() constant returns(bool)
func (_Bindings *BindingsSession) IsCToken() (bool, error) {
	return _Bindings.Contract.IsCToken(&_Bindings.CallOpts)
}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() constant returns(bool)
func (_Bindings *BindingsCallerSession) IsCToken() (bool, error) {
	return _Bindings.Contract.IsCToken(&_Bindings.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Bindings *BindingsCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Bindings *BindingsSession) Name() (string, error) {
	return _Bindings.Contract.Name(&_Bindings.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Bindings *BindingsCallerSession) Name() (string, error) {
	return _Bindings.Contract.Name(&_Bindings.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_Bindings *BindingsCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "pendingAdmin")
	return *ret0, err
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_Bindings *BindingsSession) PendingAdmin() (common.Address, error) {
	return _Bindings.Contract.PendingAdmin(&_Bindings.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_Bindings *BindingsCallerSession) PendingAdmin() (common.Address, error) {
	return _Bindings.Contract.PendingAdmin(&_Bindings.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() constant returns(uint256)
func (_Bindings *BindingsCaller) ReserveFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "reserveFactorMantissa")
	return *ret0, err
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() constant returns(uint256)
func (_Bindings *BindingsSession) ReserveFactorMantissa() (*big.Int, error) {
	return _Bindings.Contract.ReserveFactorMantissa(&_Bindings.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() constant returns(uint256)
func (_Bindings *BindingsCallerSession) ReserveFactorMantissa() (*big.Int, error) {
	return _Bindings.Contract.ReserveFactorMantissa(&_Bindings.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() constant returns(uint256)
func (_Bindings *BindingsCaller) SupplyRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "supplyRatePerBlock")
	return *ret0, err
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() constant returns(uint256)
func (_Bindings *BindingsSession) SupplyRatePerBlock() (*big.Int, error) {
	return _Bindings.Contract.SupplyRatePerBlock(&_Bindings.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() constant returns(uint256)
func (_Bindings *BindingsCallerSession) SupplyRatePerBlock() (*big.Int, error) {
	return _Bindings.Contract.SupplyRatePerBlock(&_Bindings.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Bindings *BindingsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Bindings *BindingsSession) Symbol() (string, error) {
	return _Bindings.Contract.Symbol(&_Bindings.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Bindings *BindingsCallerSession) Symbol() (string, error) {
	return _Bindings.Contract.Symbol(&_Bindings.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() constant returns(uint256)
func (_Bindings *BindingsCaller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "totalBorrows")
	return *ret0, err
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() constant returns(uint256)
func (_Bindings *BindingsSession) TotalBorrows() (*big.Int, error) {
	return _Bindings.Contract.TotalBorrows(&_Bindings.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() constant returns(uint256)
func (_Bindings *BindingsCallerSession) TotalBorrows() (*big.Int, error) {
	return _Bindings.Contract.TotalBorrows(&_Bindings.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() constant returns(uint256)
func (_Bindings *BindingsCaller) TotalReserves(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "totalReserves")
	return *ret0, err
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() constant returns(uint256)
func (_Bindings *BindingsSession) TotalReserves() (*big.Int, error) {
	return _Bindings.Contract.TotalReserves(&_Bindings.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() constant returns(uint256)
func (_Bindings *BindingsCallerSession) TotalReserves() (*big.Int, error) {
	return _Bindings.Contract.TotalReserves(&_Bindings.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Bindings *BindingsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Bindings *BindingsSession) TotalSupply() (*big.Int, error) {
	return _Bindings.Contract.TotalSupply(&_Bindings.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Bindings *BindingsCallerSession) TotalSupply() (*big.Int, error) {
	return _Bindings.Contract.TotalSupply(&_Bindings.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() constant returns(address)
func (_Bindings *BindingsCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bindings.contract.Call(opts, out, "underlying")
	return *ret0, err
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() constant returns(address)
func (_Bindings *BindingsSession) Underlying() (common.Address, error) {
	return _Bindings.Contract.Underlying(&_Bindings.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() constant returns(address)
func (_Bindings *BindingsCallerSession) Underlying() (common.Address, error) {
	return _Bindings.Contract.Underlying(&_Bindings.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Bindings *BindingsTransactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "_acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Bindings *BindingsSession) AcceptAdmin() (*types.Transaction, error) {
	return _Bindings.Contract.AcceptAdmin(&_Bindings.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Bindings *BindingsTransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _Bindings.Contract.AcceptAdmin(&_Bindings.TransactOpts)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_Bindings *BindingsTransactor) ReduceReserves(opts *bind.TransactOpts, reduceAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "_reduceReserves", reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_Bindings *BindingsSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.ReduceReserves(&_Bindings.TransactOpts, reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_Bindings *BindingsTransactorSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.ReduceReserves(&_Bindings.TransactOpts, reduceAmount)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Bindings *BindingsTransactor) SetComptroller(opts *bind.TransactOpts, newComptroller common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "_setComptroller", newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Bindings *BindingsSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetComptroller(&_Bindings.TransactOpts, newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Bindings *BindingsTransactorSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetComptroller(&_Bindings.TransactOpts, newComptroller)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Bindings *BindingsTransactor) SetInterestRateModel(opts *bind.TransactOpts, newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "_setInterestRateModel", newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Bindings *BindingsSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetInterestRateModel(&_Bindings.TransactOpts, newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Bindings *BindingsTransactorSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetInterestRateModel(&_Bindings.TransactOpts, newInterestRateModel)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Bindings *BindingsTransactor) SetPendingAdmin(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "_setPendingAdmin", newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Bindings *BindingsSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetPendingAdmin(&_Bindings.TransactOpts, newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Bindings *BindingsTransactorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetPendingAdmin(&_Bindings.TransactOpts, newPendingAdmin)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_Bindings *BindingsTransactor) SetReserveFactor(opts *bind.TransactOpts, newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "_setReserveFactor", newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_Bindings *BindingsSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetReserveFactor(&_Bindings.TransactOpts, newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_Bindings *BindingsTransactorSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.SetReserveFactor(&_Bindings.TransactOpts, newReserveFactorMantissa)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Bindings *BindingsTransactor) AccrueInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "accrueInterest")
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Bindings *BindingsSession) AccrueInterest() (*types.Transaction, error) {
	return _Bindings.Contract.AccrueInterest(&_Bindings.TransactOpts)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Bindings *BindingsTransactorSession) AccrueInterest() (*types.Transaction, error) {
	return _Bindings.Contract.AccrueInterest(&_Bindings.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bindings *BindingsTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bindings *BindingsSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Approve(&_Bindings.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bindings *BindingsTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Approve(&_Bindings.TransactOpts, spender, amount)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Bindings *BindingsTransactor) BalanceOfUnderlying(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "balanceOfUnderlying", owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Bindings *BindingsSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.BalanceOfUnderlying(&_Bindings.TransactOpts, owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Bindings *BindingsTransactorSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.BalanceOfUnderlying(&_Bindings.TransactOpts, owner)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Bindings *BindingsTransactor) Borrow(opts *bind.TransactOpts, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "borrow", borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Bindings *BindingsSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Borrow(&_Bindings.TransactOpts, borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Bindings *BindingsTransactorSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Borrow(&_Bindings.TransactOpts, borrowAmount)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Bindings *BindingsTransactor) BorrowBalanceCurrent(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "borrowBalanceCurrent", account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Bindings *BindingsSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.BorrowBalanceCurrent(&_Bindings.TransactOpts, account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Bindings *BindingsTransactorSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.BorrowBalanceCurrent(&_Bindings.TransactOpts, account)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Bindings *BindingsTransactor) ExchangeRateCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "exchangeRateCurrent")
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Bindings *BindingsSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _Bindings.Contract.ExchangeRateCurrent(&_Bindings.TransactOpts)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Bindings *BindingsTransactorSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _Bindings.Contract.ExchangeRateCurrent(&_Bindings.TransactOpts)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address cTokenCollateral) returns(uint256)
func (_Bindings *BindingsTransactor) LiquidateBorrow(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "liquidateBorrow", borrower, repayAmount, cTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address cTokenCollateral) returns(uint256)
func (_Bindings *BindingsSession) LiquidateBorrow(borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.LiquidateBorrow(&_Bindings.TransactOpts, borrower, repayAmount, cTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address cTokenCollateral) returns(uint256)
func (_Bindings *BindingsTransactorSession) LiquidateBorrow(borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.LiquidateBorrow(&_Bindings.TransactOpts, borrower, repayAmount, cTokenCollateral)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_Bindings *BindingsTransactor) Mint(opts *bind.TransactOpts, mintAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "mint", mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_Bindings *BindingsSession) Mint(mintAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Mint(&_Bindings.TransactOpts, mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_Bindings *BindingsTransactorSession) Mint(mintAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Mint(&_Bindings.TransactOpts, mintAmount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Bindings *BindingsTransactor) Redeem(opts *bind.TransactOpts, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "redeem", redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Bindings *BindingsSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Redeem(&_Bindings.TransactOpts, redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Bindings *BindingsTransactorSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Redeem(&_Bindings.TransactOpts, redeemTokens)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Bindings *BindingsTransactor) RedeemUnderlying(opts *bind.TransactOpts, redeemAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "redeemUnderlying", redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Bindings *BindingsSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.RedeemUnderlying(&_Bindings.TransactOpts, redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Bindings *BindingsTransactorSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.RedeemUnderlying(&_Bindings.TransactOpts, redeemAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_Bindings *BindingsTransactor) RepayBorrow(opts *bind.TransactOpts, repayAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "repayBorrow", repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_Bindings *BindingsSession) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.RepayBorrow(&_Bindings.TransactOpts, repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_Bindings *BindingsTransactorSession) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.RepayBorrow(&_Bindings.TransactOpts, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_Bindings *BindingsTransactor) RepayBorrowBehalf(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "repayBorrowBehalf", borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_Bindings *BindingsSession) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.RepayBorrowBehalf(&_Bindings.TransactOpts, borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_Bindings *BindingsTransactorSession) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.RepayBorrowBehalf(&_Bindings.TransactOpts, borrower, repayAmount)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Bindings *BindingsTransactor) Seize(opts *bind.TransactOpts, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "seize", liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Bindings *BindingsSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Seize(&_Bindings.TransactOpts, liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Bindings *BindingsTransactorSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Seize(&_Bindings.TransactOpts, liquidator, borrower, seizeTokens)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Bindings *BindingsTransactor) TotalBorrowsCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "totalBorrowsCurrent")
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Bindings *BindingsSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _Bindings.Contract.TotalBorrowsCurrent(&_Bindings.TransactOpts)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Bindings *BindingsTransactorSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _Bindings.Contract.TotalBorrowsCurrent(&_Bindings.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Bindings *BindingsTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Bindings *BindingsSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Transfer(&_Bindings.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Bindings *BindingsTransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Transfer(&_Bindings.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Bindings *BindingsTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Bindings *BindingsSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.TransferFrom(&_Bindings.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Bindings *BindingsTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.TransferFrom(&_Bindings.TransactOpts, src, dst, amount)
}

// BindingsAccrueInterestIterator is returned from FilterAccrueInterest and is used to iterate over the raw logs and unpacked data for AccrueInterest events raised by the Bindings contract.
type BindingsAccrueInterestIterator struct {
	Event *BindingsAccrueInterest // Event containing the contract specifics and raw log

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
func (it *BindingsAccrueInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsAccrueInterest)
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
		it.Event = new(BindingsAccrueInterest)
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
func (it *BindingsAccrueInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsAccrueInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsAccrueInterest represents a AccrueInterest event raised by the Bindings contract.
type BindingsAccrueInterest struct {
	InterestAccumulated *big.Int
	BorrowIndex         *big.Int
	TotalBorrows        *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterAccrueInterest is a free log retrieval operation binding the contract event 0x875352fb3fadeb8c0be7cbbe8ff761b308fa7033470cd0287f02f3436fd76cb9.
//
// Solidity: event AccrueInterest(uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Bindings *BindingsFilterer) FilterAccrueInterest(opts *bind.FilterOpts) (*BindingsAccrueInterestIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return &BindingsAccrueInterestIterator{contract: _Bindings.contract, event: "AccrueInterest", logs: logs, sub: sub}, nil
}

// WatchAccrueInterest is a free log subscription operation binding the contract event 0x875352fb3fadeb8c0be7cbbe8ff761b308fa7033470cd0287f02f3436fd76cb9.
//
// Solidity: event AccrueInterest(uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Bindings *BindingsFilterer) WatchAccrueInterest(opts *bind.WatchOpts, sink chan<- *BindingsAccrueInterest) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsAccrueInterest)
				if err := _Bindings.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
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

// ParseAccrueInterest is a log parse operation binding the contract event 0x875352fb3fadeb8c0be7cbbe8ff761b308fa7033470cd0287f02f3436fd76cb9.
//
// Solidity: event AccrueInterest(uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Bindings *BindingsFilterer) ParseAccrueInterest(log types.Log) (*BindingsAccrueInterest, error) {
	event := new(BindingsAccrueInterest)
	if err := _Bindings.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BindingsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Bindings contract.
type BindingsApprovalIterator struct {
	Event *BindingsApproval // Event containing the contract specifics and raw log

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
func (it *BindingsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsApproval)
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
		it.Event = new(BindingsApproval)
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
func (it *BindingsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsApproval represents a Approval event raised by the Bindings contract.
type BindingsApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Bindings *BindingsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BindingsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BindingsApprovalIterator{contract: _Bindings.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Bindings *BindingsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BindingsApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsApproval)
				if err := _Bindings.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Bindings *BindingsFilterer) ParseApproval(log types.Log) (*BindingsApproval, error) {
	event := new(BindingsApproval)
	if err := _Bindings.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BindingsBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the Bindings contract.
type BindingsBorrowIterator struct {
	Event *BindingsBorrow // Event containing the contract specifics and raw log

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
func (it *BindingsBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsBorrow)
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
		it.Event = new(BindingsBorrow)
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
func (it *BindingsBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsBorrow represents a Borrow event raised by the Bindings contract.
type BindingsBorrow struct {
	Borrower       common.Address
	BorrowAmount   *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Bindings *BindingsFilterer) FilterBorrow(opts *bind.FilterOpts) (*BindingsBorrowIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return &BindingsBorrowIterator{contract: _Bindings.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Bindings *BindingsFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *BindingsBorrow) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsBorrow)
				if err := _Bindings.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Bindings *BindingsFilterer) ParseBorrow(log types.Log) (*BindingsBorrow, error) {
	event := new(BindingsBorrow)
	if err := _Bindings.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BindingsFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the Bindings contract.
type BindingsFailureIterator struct {
	Event *BindingsFailure // Event containing the contract specifics and raw log

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
func (it *BindingsFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsFailure)
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
		it.Event = new(BindingsFailure)
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
func (it *BindingsFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsFailure represents a Failure event raised by the Bindings contract.
type BindingsFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Bindings *BindingsFilterer) FilterFailure(opts *bind.FilterOpts) (*BindingsFailureIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &BindingsFailureIterator{contract: _Bindings.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Bindings *BindingsFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *BindingsFailure) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsFailure)
				if err := _Bindings.contract.UnpackLog(event, "Failure", log); err != nil {
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

// ParseFailure is a log parse operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Bindings *BindingsFilterer) ParseFailure(log types.Log) (*BindingsFailure, error) {
	event := new(BindingsFailure)
	if err := _Bindings.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BindingsLiquidateBorrowIterator is returned from FilterLiquidateBorrow and is used to iterate over the raw logs and unpacked data for LiquidateBorrow events raised by the Bindings contract.
type BindingsLiquidateBorrowIterator struct {
	Event *BindingsLiquidateBorrow // Event containing the contract specifics and raw log

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
func (it *BindingsLiquidateBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsLiquidateBorrow)
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
		it.Event = new(BindingsLiquidateBorrow)
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
func (it *BindingsLiquidateBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsLiquidateBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsLiquidateBorrow represents a LiquidateBorrow event raised by the Bindings contract.
type BindingsLiquidateBorrow struct {
	Liquidator       common.Address
	Borrower         common.Address
	RepayAmount      *big.Int
	CTokenCollateral common.Address
	SeizeTokens      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLiquidateBorrow is a free log retrieval operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_Bindings *BindingsFilterer) FilterLiquidateBorrow(opts *bind.FilterOpts) (*BindingsLiquidateBorrowIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return &BindingsLiquidateBorrowIterator{contract: _Bindings.contract, event: "LiquidateBorrow", logs: logs, sub: sub}, nil
}

// WatchLiquidateBorrow is a free log subscription operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_Bindings *BindingsFilterer) WatchLiquidateBorrow(opts *bind.WatchOpts, sink chan<- *BindingsLiquidateBorrow) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsLiquidateBorrow)
				if err := _Bindings.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
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

// ParseLiquidateBorrow is a log parse operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_Bindings *BindingsFilterer) ParseLiquidateBorrow(log types.Log) (*BindingsLiquidateBorrow, error) {
	event := new(BindingsLiquidateBorrow)
	if err := _Bindings.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BindingsMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Bindings contract.
type BindingsMintIterator struct {
	Event *BindingsMint // Event containing the contract specifics and raw log

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
func (it *BindingsMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsMint)
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
		it.Event = new(BindingsMint)
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
func (it *BindingsMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsMint represents a Mint event raised by the Bindings contract.
type BindingsMint struct {
	Minter     common.Address
	MintAmount *big.Int
	MintTokens *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Bindings *BindingsFilterer) FilterMint(opts *bind.FilterOpts) (*BindingsMintIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &BindingsMintIterator{contract: _Bindings.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Bindings *BindingsFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *BindingsMint) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsMint)
				if err := _Bindings.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Bindings *BindingsFilterer) ParseMint(log types.Log) (*BindingsMint, error) {
	event := new(BindingsMint)
	if err := _Bindings.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BindingsNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Bindings contract.
type BindingsNewAdminIterator struct {
	Event *BindingsNewAdmin // Event containing the contract specifics and raw log

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
func (it *BindingsNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsNewAdmin)
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
		it.Event = new(BindingsNewAdmin)
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
func (it *BindingsNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsNewAdmin represents a NewAdmin event raised by the Bindings contract.
type BindingsNewAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Bindings *BindingsFilterer) FilterNewAdmin(opts *bind.FilterOpts) (*BindingsNewAdminIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return &BindingsNewAdminIterator{contract: _Bindings.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Bindings *BindingsFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *BindingsNewAdmin) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsNewAdmin)
				if err := _Bindings.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Bindings *BindingsFilterer) ParseNewAdmin(log types.Log) (*BindingsNewAdmin, error) {
	event := new(BindingsNewAdmin)
	if err := _Bindings.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BindingsNewComptrollerIterator is returned from FilterNewComptroller and is used to iterate over the raw logs and unpacked data for NewComptroller events raised by the Bindings contract.
type BindingsNewComptrollerIterator struct {
	Event *BindingsNewComptroller // Event containing the contract specifics and raw log

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
func (it *BindingsNewComptrollerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsNewComptroller)
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
		it.Event = new(BindingsNewComptroller)
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
func (it *BindingsNewComptrollerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsNewComptrollerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsNewComptroller represents a NewComptroller event raised by the Bindings contract.
type BindingsNewComptroller struct {
	OldComptroller common.Address
	NewComptroller common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewComptroller is a free log retrieval operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Bindings *BindingsFilterer) FilterNewComptroller(opts *bind.FilterOpts) (*BindingsNewComptrollerIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return &BindingsNewComptrollerIterator{contract: _Bindings.contract, event: "NewComptroller", logs: logs, sub: sub}, nil
}

// WatchNewComptroller is a free log subscription operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Bindings *BindingsFilterer) WatchNewComptroller(opts *bind.WatchOpts, sink chan<- *BindingsNewComptroller) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsNewComptroller)
				if err := _Bindings.contract.UnpackLog(event, "NewComptroller", log); err != nil {
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

// ParseNewComptroller is a log parse operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Bindings *BindingsFilterer) ParseNewComptroller(log types.Log) (*BindingsNewComptroller, error) {
	event := new(BindingsNewComptroller)
	if err := _Bindings.contract.UnpackLog(event, "NewComptroller", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BindingsNewMarketInterestRateModelIterator is returned from FilterNewMarketInterestRateModel and is used to iterate over the raw logs and unpacked data for NewMarketInterestRateModel events raised by the Bindings contract.
type BindingsNewMarketInterestRateModelIterator struct {
	Event *BindingsNewMarketInterestRateModel // Event containing the contract specifics and raw log

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
func (it *BindingsNewMarketInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsNewMarketInterestRateModel)
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
		it.Event = new(BindingsNewMarketInterestRateModel)
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
func (it *BindingsNewMarketInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsNewMarketInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsNewMarketInterestRateModel represents a NewMarketInterestRateModel event raised by the Bindings contract.
type BindingsNewMarketInterestRateModel struct {
	OldInterestRateModel common.Address
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewMarketInterestRateModel is a free log retrieval operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Bindings *BindingsFilterer) FilterNewMarketInterestRateModel(opts *bind.FilterOpts) (*BindingsNewMarketInterestRateModelIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return &BindingsNewMarketInterestRateModelIterator{contract: _Bindings.contract, event: "NewMarketInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchNewMarketInterestRateModel is a free log subscription operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Bindings *BindingsFilterer) WatchNewMarketInterestRateModel(opts *bind.WatchOpts, sink chan<- *BindingsNewMarketInterestRateModel) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsNewMarketInterestRateModel)
				if err := _Bindings.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
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

// ParseNewMarketInterestRateModel is a log parse operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Bindings *BindingsFilterer) ParseNewMarketInterestRateModel(log types.Log) (*BindingsNewMarketInterestRateModel, error) {
	event := new(BindingsNewMarketInterestRateModel)
	if err := _Bindings.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BindingsNewPendingAdminIterator is returned from FilterNewPendingAdmin and is used to iterate over the raw logs and unpacked data for NewPendingAdmin events raised by the Bindings contract.
type BindingsNewPendingAdminIterator struct {
	Event *BindingsNewPendingAdmin // Event containing the contract specifics and raw log

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
func (it *BindingsNewPendingAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsNewPendingAdmin)
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
		it.Event = new(BindingsNewPendingAdmin)
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
func (it *BindingsNewPendingAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsNewPendingAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsNewPendingAdmin represents a NewPendingAdmin event raised by the Bindings contract.
type BindingsNewPendingAdmin struct {
	OldPendingAdmin common.Address
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAdmin is a free log retrieval operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Bindings *BindingsFilterer) FilterNewPendingAdmin(opts *bind.FilterOpts) (*BindingsNewPendingAdminIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return &BindingsNewPendingAdminIterator{contract: _Bindings.contract, event: "NewPendingAdmin", logs: logs, sub: sub}, nil
}

// WatchNewPendingAdmin is a free log subscription operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Bindings *BindingsFilterer) WatchNewPendingAdmin(opts *bind.WatchOpts, sink chan<- *BindingsNewPendingAdmin) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsNewPendingAdmin)
				if err := _Bindings.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
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

// ParseNewPendingAdmin is a log parse operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Bindings *BindingsFilterer) ParseNewPendingAdmin(log types.Log) (*BindingsNewPendingAdmin, error) {
	event := new(BindingsNewPendingAdmin)
	if err := _Bindings.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BindingsNewReserveFactorIterator is returned from FilterNewReserveFactor and is used to iterate over the raw logs and unpacked data for NewReserveFactor events raised by the Bindings contract.
type BindingsNewReserveFactorIterator struct {
	Event *BindingsNewReserveFactor // Event containing the contract specifics and raw log

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
func (it *BindingsNewReserveFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsNewReserveFactor)
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
		it.Event = new(BindingsNewReserveFactor)
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
func (it *BindingsNewReserveFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsNewReserveFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsNewReserveFactor represents a NewReserveFactor event raised by the Bindings contract.
type BindingsNewReserveFactor struct {
	OldReserveFactorMantissa *big.Int
	NewReserveFactorMantissa *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNewReserveFactor is a free log retrieval operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Bindings *BindingsFilterer) FilterNewReserveFactor(opts *bind.FilterOpts) (*BindingsNewReserveFactorIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return &BindingsNewReserveFactorIterator{contract: _Bindings.contract, event: "NewReserveFactor", logs: logs, sub: sub}, nil
}

// WatchNewReserveFactor is a free log subscription operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Bindings *BindingsFilterer) WatchNewReserveFactor(opts *bind.WatchOpts, sink chan<- *BindingsNewReserveFactor) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsNewReserveFactor)
				if err := _Bindings.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
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

// ParseNewReserveFactor is a log parse operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Bindings *BindingsFilterer) ParseNewReserveFactor(log types.Log) (*BindingsNewReserveFactor, error) {
	event := new(BindingsNewReserveFactor)
	if err := _Bindings.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BindingsRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the Bindings contract.
type BindingsRedeemIterator struct {
	Event *BindingsRedeem // Event containing the contract specifics and raw log

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
func (it *BindingsRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRedeem)
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
		it.Event = new(BindingsRedeem)
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
func (it *BindingsRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRedeem represents a Redeem event raised by the Bindings contract.
type BindingsRedeem struct {
	Redeemer     common.Address
	RedeemAmount *big.Int
	RedeemTokens *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Bindings *BindingsFilterer) FilterRedeem(opts *bind.FilterOpts) (*BindingsRedeemIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return &BindingsRedeemIterator{contract: _Bindings.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Bindings *BindingsFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *BindingsRedeem) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRedeem)
				if err := _Bindings.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Bindings *BindingsFilterer) ParseRedeem(log types.Log) (*BindingsRedeem, error) {
	event := new(BindingsRedeem)
	if err := _Bindings.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BindingsRepayBorrowIterator is returned from FilterRepayBorrow and is used to iterate over the raw logs and unpacked data for RepayBorrow events raised by the Bindings contract.
type BindingsRepayBorrowIterator struct {
	Event *BindingsRepayBorrow // Event containing the contract specifics and raw log

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
func (it *BindingsRepayBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRepayBorrow)
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
		it.Event = new(BindingsRepayBorrow)
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
func (it *BindingsRepayBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRepayBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRepayBorrow represents a RepayBorrow event raised by the Bindings contract.
type BindingsRepayBorrow struct {
	Payer          common.Address
	Borrower       common.Address
	RepayAmount    *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRepayBorrow is a free log retrieval operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Bindings *BindingsFilterer) FilterRepayBorrow(opts *bind.FilterOpts) (*BindingsRepayBorrowIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return &BindingsRepayBorrowIterator{contract: _Bindings.contract, event: "RepayBorrow", logs: logs, sub: sub}, nil
}

// WatchRepayBorrow is a free log subscription operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Bindings *BindingsFilterer) WatchRepayBorrow(opts *bind.WatchOpts, sink chan<- *BindingsRepayBorrow) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRepayBorrow)
				if err := _Bindings.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
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

// ParseRepayBorrow is a log parse operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Bindings *BindingsFilterer) ParseRepayBorrow(log types.Log) (*BindingsRepayBorrow, error) {
	event := new(BindingsRepayBorrow)
	if err := _Bindings.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BindingsReservesReducedIterator is returned from FilterReservesReduced and is used to iterate over the raw logs and unpacked data for ReservesReduced events raised by the Bindings contract.
type BindingsReservesReducedIterator struct {
	Event *BindingsReservesReduced // Event containing the contract specifics and raw log

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
func (it *BindingsReservesReducedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsReservesReduced)
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
		it.Event = new(BindingsReservesReduced)
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
func (it *BindingsReservesReducedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsReservesReducedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsReservesReduced represents a ReservesReduced event raised by the Bindings contract.
type BindingsReservesReduced struct {
	Admin            common.Address
	ReduceAmount     *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesReduced is a free log retrieval operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Bindings *BindingsFilterer) FilterReservesReduced(opts *bind.FilterOpts) (*BindingsReservesReducedIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return &BindingsReservesReducedIterator{contract: _Bindings.contract, event: "ReservesReduced", logs: logs, sub: sub}, nil
}

// WatchReservesReduced is a free log subscription operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Bindings *BindingsFilterer) WatchReservesReduced(opts *bind.WatchOpts, sink chan<- *BindingsReservesReduced) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsReservesReduced)
				if err := _Bindings.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
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

// ParseReservesReduced is a log parse operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Bindings *BindingsFilterer) ParseReservesReduced(log types.Log) (*BindingsReservesReduced, error) {
	event := new(BindingsReservesReduced)
	if err := _Bindings.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BindingsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Bindings contract.
type BindingsTransferIterator struct {
	Event *BindingsTransfer // Event containing the contract specifics and raw log

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
func (it *BindingsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsTransfer)
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
		it.Event = new(BindingsTransfer)
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
func (it *BindingsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsTransfer represents a Transfer event raised by the Bindings contract.
type BindingsTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Bindings *BindingsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BindingsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BindingsTransferIterator{contract: _Bindings.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Bindings *BindingsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BindingsTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsTransfer)
				if err := _Bindings.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Bindings *BindingsFilterer) ParseTransfer(log types.Log) (*BindingsTransfer, error) {
	event := new(BindingsTransfer)
	if err := _Bindings.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
