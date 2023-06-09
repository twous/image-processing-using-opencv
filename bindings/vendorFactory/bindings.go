
// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vendorfactory

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

// VendorfactoryABI is the input ABI used to generate the binding from.
const VendorfactoryABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vendorContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vendorContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"newVendor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registeredVendors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VendorfactoryBin is the compiled bytecode used for deploying new contracts.
var VendorfactoryBin = "0x608060405234801561001057600080fd5b5061100c806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80631e6765c71461005157806340b90f4b1461008a578063c27f2112146100b0578063e7b8c718146100cc575b600080fd5b61006e6004803603602081101561006757600080fd5b50356100f2565b604080516001600160a01b039092168252519081900360200190f35b61006e600480360360208110156100a057600080fd5b50356001600160a01b0316610119565b6100b8610134565b604080519115158252519081900360200190f35b6100b8600480360360208110156100e257600080fd5b50356001600160a01b031661023a565b600081815481106100ff57fe5b6000918252602090912001546001600160a01b0316905081565b6002602052600090815260409020546001600160a01b031681565b3360009081526001602052604081205460ff1615610192576040805162461bcd60e51b81526020600482015260166024820152751b5d5cdd081b9bdd081899481c9959da5cdd195c995960521b604482015290519081900360640190fd5b60006040516101a09061024f565b604051809103906000f0801580156101bc573d6000803e3d6000fd5b5060008054600181810183557f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56390910180546001600160a01b039094166001600160a01b0319948516811790915533835260208281526040808520805460ff1916851790556002909152909220805490931690911790915591505090565b60016020526000908152604090205460ff1681565b610d7b8061025d8339019056fe608060405234801561001057600080fd5b5032604051602001610022919061006b565b60408051601f198184030181529190528051602090910120600155600080546001600160a01b031916321790556100af565b61006561006082610080565b61009d565b82525050565b60006100778284610054565b50601401919050565b600061008b82610091565b92915050565b6001600160a01b031690565b600061008b82600061008b8260601b90565b610cbd806100be6000396000f3fe6080604052600436106100865760003560e01c80635d85ed13116100595780635d85ed13146101165780636c8e745c146101365780638da5cb5b14610156578063af640d0f14610178578063c43df6aa1461019a57610086565b80630186a4231461008857806324600fc3146100bf578063338a6d10146100e15780635c388ca614610101575b005b34801561009457600080fd5b506100a86100a3366004610791565b6101ba565b6040516100b6929190610b19565b60405180910390f35b3480156100cb57600080fd5b506100d461026d565b6040516100b69190610aa4565b3480156100ed57600080fd5b506100d46100fc36600461084a565b610329565b34801561010d57600080fd5b506100d4610366565b34801561012257600080fd5b506100d461013136600461084a565b61036f565b34801561014257600080fd5b506100d461015136600461084a565b610419565b34801561016257600080fd5b5061016b6104bd565b6040516100b69190610a96565b34801561018457600080fd5b5061018d6104cc565b6040516100b69190610ab2565b3480156101a657600080fd5b506100d46101b53660046107ce565b6104d2565b805160208183018101805160038252928201938201939093209190925280546040805160026001841615610100026000190190931692909204601f810185900485028301850190915280825291929091839183018282801561025d5780601f106102325761010080835404028352916020019161025d565b820191906000526020600020905b81548152906001019060200180831161024057829003601f168201915b5050505050908060010154905082565b6000610277610611565b61029c5760405162461bcd60e51b815260040161029390610b69565b60405180910390fd5b30316102ba5760405162461bcd60e51b815260040161029390610b39565b60025460ff16156102dd5760405162461bcd60e51b815260040161029390610b59565b6002805460ff191660011790556040513390303180156108fc02916000818181858888f19350505050158015610317573d6000803e3d6000fd5b50506002805460ff1916905560015b90565b8151602081840181018051600482529282019482019490942091909352815180830184018051928152908401929093019190912091525460ff1681565b60025460ff1681565b6000610379610611565b6103955760405162461bcd60e51b815260040161029390610b49565b6004836040516103a59190610a8a565b9081526020016040518091039020826040516103c19190610a8a565b908152604051908190036020018120805460ff191690557f6f87e5eed57f39feb4e7480e0eaa01353c1806e3a63675ebe5095b8f338cd62f906104079085908590610af4565b60405180910390a15060015b92915050565b6000610423610611565b61043f5760405162461bcd60e51b815260040161029390610b49565b60016004846040516104519190610a8a565b90815260200160405180910390208360405161046d9190610a8a565b908152604051908190036020018120805492151560ff19909316929092179091557f20eca5c8f895538bd583d98f6dc289ffd029bf3d0a3b5362dffeace9ab6634dc906104079085908590610af4565b6000546001600160a01b031681565b60015481565b60006104dc610611565b6104f85760405162461bcd60e51b815260040161029390610b49565b60405180604001604052808581526020018381525060038560405161051d9190610a8a565b90815260200160405180910390206000820151816000019080519060200190610547929190610633565b506020919091015160019091015560005b83518110156105ca5760016004866040516105739190610a8a565b908152602001604051809103902085838151811061058d57fe5b60200260200101516040516105a29190610a8a565b908152604051908190036020019020805491151560ff19909216919091179055600101610558565b507f3c16f324d383d6feb7b8939d45155634c4ba851efa4ea07b79cc280489899d1d8484846040516105fe93929190610ac0565b60405180910390a15060015b9392505050565b600080546001600160a01b031633141561062d57506001610326565b50600090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061067457805160ff19168380011785556106a1565b828001600101855582156106a1579182015b828111156106a1578251825591602001919060010190610686565b506106ad9291506106b1565b5090565b61032691905b808211156106ad57600081556001016106b7565b600082601f8301126106dc57600080fd5b81356106ef6106ea82610ba0565b610b79565b81815260209384019390925082018360005b8381101561072d57813586016107178882610737565b8452506020928301929190910190600101610701565b5050505092915050565b600082601f83011261074857600080fd5b81356107566106ea82610bc1565b9150808252602083016020830185838301111561077257600080fd5b61077d838284610c1d565b50505092915050565b803561041381610c63565b6000602082840312156107a357600080fd5b813567ffffffffffffffff8111156107ba57600080fd5b6107c684828501610737565b949350505050565b6000806000606084860312156107e357600080fd5b833567ffffffffffffffff8111156107fa57600080fd5b61080686828701610737565b935050602084013567ffffffffffffffff81111561082357600080fd5b61082f868287016106cb565b925050604061084086828701610786565b9150509250925092565b6000806040838503121561085d57600080fd5b823567ffffffffffffffff81111561087457600080fd5b61088085828601610737565b925050602083013567ffffffffffffffff81111561089d57600080fd5b6108a985828601610737565b9150509250929050565b600061060a838361094e565b6108c881610c01565b82525050565b60006108d982610bef565b6108e38185610bf3565b9350836020820285016108f585610be9565b8060005b8581101561092f578484038952815161091285826108b3565b945061091d83610be9565b60209a909a01999250506001016108f9565b5091979650505050505050565b6108c881610c0c565b6108c881610326565b600061095982610bef565b6109638185610bf3565b9350610973818560208601610c29565b61097c81610c59565b9093019392505050565b600061099182610bef565b61099b8185610bfc565b93506109ab818560208601610c29565b9290920192915050565b60006109c2601683610bf3565b751b9bc818985b185b98d9481a5b8818dbdb9d1c9858dd60521b815260200192915050565b60006109f4601783610bf3565b7f63616c6c6572206d7573742062652076656e646f726564000000000000000000815260200192915050565b6000610a2d602083610bf3565b7f636f6e7472616374206d757374206e6f74206265207769746864726177696e67815260200192915050565b6000610a66601583610bf3565b7431b0b63632b91036bab9ba103132903b32b73237b960591b815260200192915050565b600061060a8284610986565b6020810161041382846108bf565b60208101610413828461093c565b602081016104138284610945565b60608082528101610ad1818661094e565b90508181036020830152610ae581856108ce565b90506107c66040830184610945565b60408082528101610b05818561094e565b905081810360208301526107c6818461094e565b60408082528101610b2a818561094e565b905061060a6020830184610945565b60208082528101610413816109b5565b60208082528101610413816109e7565b6020808252810161041381610a20565b6020808252810161041381610a59565b60405181810167ffffffffffffffff81118282101715610b9857600080fd5b604052919050565b600067ffffffffffffffff821115610bb757600080fd5b5060209081020190565b600067ffffffffffffffff821115610bd857600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b919050565b600061041382610c11565b151590565b6001600160a01b031690565b82818337506000910152565b60005b83811015610c44578181015183820152602001610c2c565b83811115610c53576000848401525b50505050565b601f01601f191690565b610c6c81610326565b8114610c7757600080fd5b5056fea365627a7a72315820a9a2c9094e1a0cb1bcc7077e51dc09400585ef41c2da7265cb78aec80889d5e76c6578706572696d656e74616cf564736f6c634300050b0040a265627a7a72315820b51ac43d4203de7cf7c2829633cf284481134eaa22336a2e50bf9294f8de3a6264736f6c634300050b0032"

// DeployVendorfactory deploys a new Ethereum contract, binding an instance of Vendorfactory to it.
func DeployVendorfactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Vendorfactory, error) {
	parsed, err := abi.JSON(strings.NewReader(VendorfactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VendorfactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vendorfactory{VendorfactoryCaller: VendorfactoryCaller{contract: contract}, VendorfactoryTransactor: VendorfactoryTransactor{contract: contract}, VendorfactoryFilterer: VendorfactoryFilterer{contract: contract}}, nil
}

// Vendorfactory is an auto generated Go binding around an Ethereum contract.
type Vendorfactory struct {
	VendorfactoryCaller     // Read-only binding to the contract
	VendorfactoryTransactor // Write-only binding to the contract
	VendorfactoryFilterer   // Log filterer for contract events
}

// VendorfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type VendorfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VendorfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VendorfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VendorfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VendorfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VendorfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VendorfactorySession struct {
	Contract     *Vendorfactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VendorfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VendorfactoryCallerSession struct {
	Contract *VendorfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// VendorfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VendorfactoryTransactorSession struct {
	Contract     *VendorfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VendorfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type VendorfactoryRaw struct {
	Contract *Vendorfactory // Generic contract binding to access the raw methods on
}

// VendorfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VendorfactoryCallerRaw struct {
	Contract *VendorfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// VendorfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VendorfactoryTransactorRaw struct {
	Contract *VendorfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVendorfactory creates a new instance of Vendorfactory, bound to a specific deployed contract.
func NewVendorfactory(address common.Address, backend bind.ContractBackend) (*Vendorfactory, error) {
	contract, err := bindVendorfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vendorfactory{VendorfactoryCaller: VendorfactoryCaller{contract: contract}, VendorfactoryTransactor: VendorfactoryTransactor{contract: contract}, VendorfactoryFilterer: VendorfactoryFilterer{contract: contract}}, nil
}

// NewVendorfactoryCaller creates a new read-only instance of Vendorfactory, bound to a specific deployed contract.
func NewVendorfactoryCaller(address common.Address, caller bind.ContractCaller) (*VendorfactoryCaller, error) {
	contract, err := bindVendorfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VendorfactoryCaller{contract: contract}, nil
}

// NewVendorfactoryTransactor creates a new write-only instance of Vendorfactory, bound to a specific deployed contract.
func NewVendorfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*VendorfactoryTransactor, error) {
	contract, err := bindVendorfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VendorfactoryTransactor{contract: contract}, nil
}

// NewVendorfactoryFilterer creates a new log filterer instance of Vendorfactory, bound to a specific deployed contract.
func NewVendorfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*VendorfactoryFilterer, error) {
	contract, err := bindVendorfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VendorfactoryFilterer{contract: contract}, nil
}

// bindVendorfactory binds a generic wrapper to an already deployed contract.
func bindVendorfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VendorfactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vendorfactory *VendorfactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vendorfactory.Contract.VendorfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vendorfactory *VendorfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vendorfactory.Contract.VendorfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vendorfactory *VendorfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vendorfactory.Contract.VendorfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vendorfactory *VendorfactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vendorfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vendorfactory *VendorfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vendorfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vendorfactory *VendorfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vendorfactory.Contract.contract.Transact(opts, method, params...)
}

// RegisteredVendors is a free data retrieval call binding the contract method 0xe7b8c718.
//
// Solidity: function registeredVendors(address ) constant returns(bool)
func (_Vendorfactory *VendorfactoryCaller) RegisteredVendors(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vendorfactory.contract.Call(opts, out, "registeredVendors", arg0)
	return *ret0, err
}

// RegisteredVendors is a free data retrieval call binding the contract method 0xe7b8c718.
//
// Solidity: function registeredVendors(address ) constant returns(bool)
func (_Vendorfactory *VendorfactorySession) RegisteredVendors(arg0 common.Address) (bool, error) {
	return _Vendorfactory.Contract.RegisteredVendors(&_Vendorfactory.CallOpts, arg0)
}

// RegisteredVendors is a free data retrieval call binding the contract method 0xe7b8c718.
//
// Solidity: function registeredVendors(address ) constant returns(bool)
func (_Vendorfactory *VendorfactoryCallerSession) RegisteredVendors(arg0 common.Address) (bool, error) {
	return _Vendorfactory.Contract.RegisteredVendors(&_Vendorfactory.CallOpts, arg0)
}

// VendorContract is a free data retrieval call binding the contract method 0x40b90f4b.
//
// Solidity: function vendorContract(address ) constant returns(address)
func (_Vendorfactory *VendorfactoryCaller) VendorContract(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vendorfactory.contract.Call(opts, out, "vendorContract", arg0)
	return *ret0, err
}

// VendorContract is a free data retrieval call binding the contract method 0x40b90f4b.
//
// Solidity: function vendorContract(address ) constant returns(address)
func (_Vendorfactory *VendorfactorySession) VendorContract(arg0 common.Address) (common.Address, error) {
	return _Vendorfactory.Contract.VendorContract(&_Vendorfactory.CallOpts, arg0)
}

// VendorContract is a free data retrieval call binding the contract method 0x40b90f4b.
//
// Solidity: function vendorContract(address ) constant returns(address)
func (_Vendorfactory *VendorfactoryCallerSession) VendorContract(arg0 common.Address) (common.Address, error) {
	return _Vendorfactory.Contract.VendorContract(&_Vendorfactory.CallOpts, arg0)
}

// VendorContracts is a free data retrieval call binding the contract method 0x1e6765c7.
//
// Solidity: function vendorContracts(uint256 ) constant returns(address)
func (_Vendorfactory *VendorfactoryCaller) VendorContracts(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vendorfactory.contract.Call(opts, out, "vendorContracts", arg0)
	return *ret0, err
}

// VendorContracts is a free data retrieval call binding the contract method 0x1e6765c7.
//
// Solidity: function vendorContracts(uint256 ) constant returns(address)
func (_Vendorfactory *VendorfactorySession) VendorContracts(arg0 *big.Int) (common.Address, error) {
	return _Vendorfactory.Contract.VendorContracts(&_Vendorfactory.CallOpts, arg0)
}

// VendorContracts is a free data retrieval call binding the contract method 0x1e6765c7.
//
// Solidity: function vendorContracts(uint256 ) constant returns(address)
func (_Vendorfactory *VendorfactoryCallerSession) VendorContracts(arg0 *big.Int) (common.Address, error) {
	return _Vendorfactory.Contract.VendorContracts(&_Vendorfactory.CallOpts, arg0)
}

// NewVendor is a paid mutator transaction binding the contract method 0xc27f2112.
//
// Solidity: function newVendor() returns(bool)
func (_Vendorfactory *VendorfactoryTransactor) NewVendor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vendorfactory.contract.Transact(opts, "newVendor")
}

// NewVendor is a paid mutator transaction binding the contract method 0xc27f2112.
//
// Solidity: function newVendor() returns(bool)
func (_Vendorfactory *VendorfactorySession) NewVendor() (*types.Transaction, error) {
	return _Vendorfactory.Contract.NewVendor(&_Vendorfactory.TransactOpts)
}

// NewVendor is a paid mutator transaction binding the contract method 0xc27f2112.
//
// Solidity: function newVendor() returns(bool)
func (_Vendorfactory *VendorfactoryTransactorSession) NewVendor() (*types.Transaction, error) {
	return _Vendorfactory.Contract.NewVendor(&_Vendorfactory.TransactOpts)
}