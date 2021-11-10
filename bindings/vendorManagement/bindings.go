// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vendormanagement

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

// VendormanagementABI is the input ABI used to generate the binding from.
const VendormanagementABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"products\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawFunds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"soldAt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawLock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"}],\"name\":\"removeProductLocation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"}],\"name\":\"addProductLocation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_locations\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"_cost\",\"type\":\"uint256\"}],\"name\":\"registerProduct\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"_locations\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_cost\",\"type\":\"uint256\"}],\"name\":\"ProductRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"}],\"name\":\"ProductLocationAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_location\",\"type\":\"string\"}],\"name\":\"ProductLocationRemoved\",\"type\":\"event\"}]"

// VendormanagementBin is the compiled bytecode used for deploying new contracts.
var VendormanagementBin = "0x608060405234801561001057600080fd5b5032604051602001610022919061006b565b60408051601f198184030181529190528051602090910120600155600080546001600160a01b031916321790556100af565b61006561006082610080565b61009d565b82525050565b60006100778284610054565b50601401919050565b600061008b82610091565b92915050565b6001600160a01b031690565b600061008b82600061008b8260601b90565b610cbd806100be6000396000f3fe6080604052600436106100865760003560e01c80635d85ed13116100595780635d85ed13146101165780636c8e745c146101365780638da5cb5b14610156578063af640d0f14610178578063c43df6aa1461019a57610086565b80630186a4231461008857806324600fc3146100bf578063338a6d10146100e15780635c388ca614610101575b005b34801561009457600080fd5b506100a86100a3366004610791565b6101ba565b6040516100b6929190610b19565b60405180910390f35b3480156100cb57600080fd5b506100d461026d565b6040516100b69190610aa4565b3480156100ed57600080fd5b506100d46100fc36600461084a565b610329565b34801561010d57600080fd5b506100d4610366565b34801561012257600080fd5b506100d461013136600461084a565b61036f565b34801561014257600080fd5b506100d461015136600461084a565b610419565b34801561016257600080fd5b5061016b6104bd565b6040516100b69190610a96565b34801561018457600080fd5b5061018d6104cc565b6040516100b69190610ab2565b3480156101a657600080fd5b506100d46101b53660046107ce565b6104d2565b805160208183018101805160038252928201938201939093209190925280546040805160026001841615610100026000190190931692909204601f810185900485028301850190915280825291929091839183018282801561025d5780601f106102325761010080835404028352916020019161025d565b820191906000526020600020905b81548152906001019060200180831161024057829003601f168201915b5050505050908060010154905082565b6000610277610611565b61029c5760405162461bcd60e51b815260040161029390610b69565b60405180910390fd5b30316102ba5760405162461bcd60e51b815260040161029390610b39565b60025460ff16156102dd5760405162461bcd60e51b815260040161029390610b59565b6002805460ff191660011790556040513390303180156108fc02916000818181858888f19350505050158015610317573d6000