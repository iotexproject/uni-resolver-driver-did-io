// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IoTeXDID

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

// IoTeXDIDABI is the input ABI used to generate the binding from.
const IoTeXDIDABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"did\",\"type\":\"string\"},{\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"updateURI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"createDID\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"did\",\"type\":\"string\"}],\"name\":\"deleteDID\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"did\",\"type\":\"string\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"updateHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"didString\",\"type\":\"string\"}],\"name\":\"CreateDID\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"didString\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"UpdateHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"didString\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"UpdateURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"didString\",\"type\":\"string\"}],\"name\":\"DeleteDID\",\"type\":\"event\"}]"

// IoTeXDIDFuncSigs maps the 4-byte function signature to its string representation.
var IoTeXDIDFuncSigs = map[string]string{
	"3e05e066": "createDID(string,bytes32,string)",
	"42e643bc": "deleteDID(string)",
	"5b6beeb9": "getHash(string)",
	"93ff5d3e": "getURI(string)",
	"78eab45a": "updateHash(string,bytes32)",
	"1789e2d8": "updateURI(string,string)",
}

// IoTeXDIDBin is the compiled bytecode used for deploying new contracts.
var IoTeXDIDBin = "0x608060405234801561001057600080fd5b506117ea806100206000396000f3006080604052600436106100775763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631789e2d8811461007c5780633e05e0661461011557806342e643bc146101b25780635b6beeb91461020b57806378eab45a1461027657806393ff5d3e146102d1575b600080fd5b34801561008857600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261011394369492936024939284019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375094975061039f9650505050505050565b005b34801561012157600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261011394369492936024939284019190819084018382808284375050604080516020601f818a01358b0180359182018390048302840183018552818452989b8a359b909a9099940197509195509182019350915081908401838280828437509497506106719650505050505050565b3480156101be57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526101139436949293602493928401919081908401838280828437509497506109219650505050505050565b34801561021757600080fd5b506040805160206004803580820135601f8101849004840285018401909552848452610264943694929360249392840191908190840183828082843750949750610b809650505050505050565b60408051918252519081900360200190f35b34801561028257600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526101139436949293602493928401919081908401838280828437509497505093359450610cb49350505050565b3480156102dd57600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261032a943694929360249392840191908190840183828082843750949750610f139650505050505050565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561036457818101518382015260200161034c565b50505050905090810190601f1680156103915780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b8160606103aa6110d2565b905060008251111561043c576103c082826111c8565b151561043c576040805160e560020a62461bcd02815260206004820152602160248201527f63616c6c657220646f6573206e6f74206f776e2074686520676976656e20646960448201527f6400000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000816040518082805190602001908083835b6020831061046e5780518252601f19909201916020918201910161044f565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff16151591506104f79050576040805160e560020a62461bcd02815260206004820152601960248201527f63616c6c6572206973206e6f74206120646964206f776e657200000000000000604482015290519081900360640190fd5b8260006105026110d2565b6040518082805190602001908083835b602083106105315780518252601f199092019160209182019101610512565b51815160209384036101000a60001901801990921691161790529201948552506040519384900381019093208451610576956002909201949190910192509050611726565b5061057f6110d2565b6040518082805190602001908083835b602083106105ae5780518252601f19909201916020918201910161058f565b51815160209384036101000a60001901801990921691161790526040805192909401829003822081835289518383015289519096507fa55ea16d9a68fded3c5b007d1fbba6ed1524d9354d44c66bcc54c721108586c495508994929350839283019185019080838360005b83811015610631578181015183820152602001610619565b50505050905090810190601f16801561065e5780820380516001836020036101000a031916815260200191505b509250505060405180910390a250505050565b60606000845111156106e55761068f8461068a3361136e565b6111c8565b15156106e5576040805160e560020a62461bcd02815260206004820152601960248201527f696420646f6573206e6f74206d617463682063726561746f7200000000000000604482015290519081900360640190fd5b6106ed6110d2565b90506000816040518082805190602001908083835b602083106107215780518252601f199092019160209182019101610702565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff161591506107a99050576040805160e560020a62461bcd02815260206004820152601260248201527f64696420616c7265616479206578697374730000000000000000000000000000604482015290519081900360640190fd5b60606040519081016040528060011515815260200184600019168152602001838152506000826040518082805190602001908083835b602083106107fe5780518252601f1990920191602091820191016107df565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382019094208551815460ff191690151517815585820151600182015593850151805161085b9450600286019350910190611726565b5090505061087061086b3361136e565b61157d565b6040518082805190602001908083835b6020831061089f5780518252601f199092019160209182019101610880565b51815160209384036101000a60001901801990921691161790526040805192909401829003822081835287518383015287519096507fedcad282d2fc969322aa68caeb2b9bf9de762094ec288a9965535b8968dca17f955087949293508392830191850190808383600083811015610631578181015183820152602001610619565b80606061092c6110d2565b90506000825111156109be5761094282826111c8565b15156109be576040805160e560020a62461bcd02815260206004820152602160248201527f63616c6c657220646f6573206e6f74206f776e2074686520676976656e20646960448201527f6400000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000816040518082805190602001908083835b602083106109f05780518252601f1990920191602091820191016109d1565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff1615159150610a799050576040805160e560020a62461bcd02815260206004820152601960248201527f63616c6c6572206973206e6f74206120646964206f776e657200000000000000604482015290519081900360640190fd5b600080610a846110d2565b6040518082805190602001908083835b60208310610ab35780518252601f199092019160209182019101610a94565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220805460ff19169315159390931790925550610af990506110d2565b6040518082805190602001908083835b60208310610b285780518252601f199092019160209182019101610b09565b5181516020939093036101000a60001901801990911692169190911790526040519201829003822093507ff8910781c5d425c8a634b0a60a6567ef96cae8b2aa97b0db44389fff2bbcc64f92506000919050a2505050565b60006060610b8d8361157d565b90506000816040518082805190602001908083835b60208310610bc15780518252601f199092019160209182019101610ba2565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff1615159150610c4a9050576040805160e560020a62461bcd02815260206004820152601260248201527f64696420646f6573206e6f742065786973740000000000000000000000000000604482015290519081900360640190fd5b6000816040518082805190602001908083835b60208310610c7c5780518252601f199092019160209182019101610c5d565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092206001015495945050505050565b816060610cbf6110d2565b9050600082511115610d5157610cd582826111c8565b1515610d51576040805160e560020a62461bcd02815260206004820152602160248201527f63616c6c657220646f6573206e6f74206f776e2074686520676976656e20646960448201527f6400000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000816040518082805190602001908083835b60208310610d835780518252601f199092019160209182019101610d64565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff1615159150610e0c9050576040805160e560020a62461bcd02815260206004820152601960248201527f63616c6c6572206973206e6f74206120646964206f776e657200000000000000604482015290519081900360640190fd5b826000610e176110d2565b6040518082805190602001908083835b60208310610e465780518252601f199092019160209182019101610e27565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092206001019290925550610e8390506110d2565b6040518082805190602001908083835b60208310610eb25780518252601f199092019160209182019101610e93565b51815160209384036101000a60001901801990921691161790526040805192909401829003822089835293519395507fad2d38b24c03f1f0db04b9345ca745dcf8904f27f8977599ec4769743d5e898d94509083900301919050a250505050565b606080610f1f8361157d565b90506000816040518082805190602001908083835b60208310610f535780518252601f199092019160209182019101610f34565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff1615159150610fdc9050576040805160e560020a62461bcd02815260206004820152601260248201527f64696420646f6573206e6f742065786973740000000000000000000000000000604482015290519081900360640190fd5b6000816040518082805190602001908083835b6020831061100e5780518252601f199092019160209182019101610fef565b518151600019602094850361010090810a8201928316921993909316919091179092529490920196875260408051978890038201882060029081018054601f6001821615909802909501909416049485018290048202880182019052838752909450919250508301828280156110c55780601f1061109a576101008083540402835291602001916110c5565b820191906000526020600020905b8154815290600101906020018083116110a857829003601f168201915b5050505050915050919050565b60606040805190810160405280600781526020017f6469643a696f3a000000000000000000000000000000000000000000000000008152506111133361136e565b6040516020018083805190602001908083835b602083106111455780518252601f199092019160209182019101611126565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b6020831061118d5780518252601f19909201916020918201910161116e565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405160208183030381529060405290505b90565b60006111d38261157d565b6040516020018082805190602001908083835b602083106112055780518252601f1990920191602091820191016111e6565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b602083106112685780518252601f199092019160209182019101611249565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020600019166112a28461157d565b6040516020018082805190602001908083835b602083106112d45780518252601f1990920191602091820191016112b5565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b602083106113375780518252601f199092019160209182019101611318565b5181516020939093036101000a60001901801990911692169190911790526040519201829003909120939093149695505050505050565b604080518082018252601081527f303132333435363738396162636465660000000000000000000000000000000060208201528151602a808252606082810190945273ffffffffffffffffffffffffffffffffffffffff851692918491600091908160200160208202803883390190505091507f300000000000000000000000000000000000000000000000000000000000000082600081518110151561141157fe5b906020010190600160f860020a031916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000082600181518110151561145957fe5b906020010190600160f860020a031916908160001a905350600090505b60148110156115745782600485600c84016020811061149157fe5b1a60f860020a02600160f860020a0319169060020a900460f860020a90048151811015156114bb57fe5b90602001015160f860020a900460f860020a0282826002026002018151811015156114e257fe5b906020010190600160f860020a031916908160001a9053508284600c83016020811061150a57fe5b1a60f860020a02600f60f860020a021660f860020a900481518110151561152d57fe5b90602001015160f860020a900460f860020a02828260020260030181518110151561155457fe5b906020010190600160f860020a031916908160001a905350600101611476565b50949350505050565b6060806060600084925082516040519080825280601f01601f1916602001820160405280156115b6578160200160208202803883390190505b509150600090505b825181101561171e5782517f4100000000000000000000000000000000000000000000000000000000000000908490839081106115f757fe5b90602001015160f860020a900460f860020a02600160f860020a0319161015801561166d575082517f5a000000000000000000000000000000000000000000000000000000000000009084908390811061164d57fe5b90602001015160f860020a900460f860020a02600160f860020a03191611155b156116ce57828181518110151561168057fe5b90602001015160f860020a900460f860020a0260f860020a900460200160f860020a0282828151811015156116b157fe5b906020010190600160f860020a031916908160001a905350611716565b82818151811015156116dc57fe5b90602001015160f860020a900460f860020a0282828151811015156116fd57fe5b906020010190600160f860020a031916908160001a9053505b6001016115be565b509392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061176757805160ff1916838001178555611794565b82800160010185558215611794579182015b82811115611794578251825591602001919060010190611779565b506117a09291506117a4565b5090565b6111c591905b808211156117a057600081556001016117aa5600a165627a7a7230582005003cdfcc875793c06262c4887198d0aa4cf5d1acd6c3a3beba7a6ab95695370029"

// DeployIoTeXDID deploys a new Ethereum contract, binding an instance of IoTeXDID to it.
func DeployIoTeXDID(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IoTeXDID, error) {
	parsed, err := abi.JSON(strings.NewReader(IoTeXDIDABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IoTeXDIDBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IoTeXDID{IoTeXDIDCaller: IoTeXDIDCaller{contract: contract}, IoTeXDIDTransactor: IoTeXDIDTransactor{contract: contract}, IoTeXDIDFilterer: IoTeXDIDFilterer{contract: contract}}, nil
}

// IoTeXDID is an auto generated Go binding around an Ethereum contract.
type IoTeXDID struct {
	IoTeXDIDCaller     // Read-only binding to the contract
	IoTeXDIDTransactor // Write-only binding to the contract
	IoTeXDIDFilterer   // Log filterer for contract events
}

// IoTeXDIDCaller is an auto generated read-only Go binding around an Ethereum contract.
type IoTeXDIDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoTeXDIDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IoTeXDIDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoTeXDIDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IoTeXDIDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoTeXDIDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IoTeXDIDSession struct {
	Contract     *IoTeXDID         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IoTeXDIDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IoTeXDIDCallerSession struct {
	Contract *IoTeXDIDCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IoTeXDIDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IoTeXDIDTransactorSession struct {
	Contract     *IoTeXDIDTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IoTeXDIDRaw is an auto generated low-level Go binding around an Ethereum contract.
type IoTeXDIDRaw struct {
	Contract *IoTeXDID // Generic contract binding to access the raw methods on
}

// IoTeXDIDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IoTeXDIDCallerRaw struct {
	Contract *IoTeXDIDCaller // Generic read-only contract binding to access the raw methods on
}

// IoTeXDIDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IoTeXDIDTransactorRaw struct {
	Contract *IoTeXDIDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIoTeXDID creates a new instance of IoTeXDID, bound to a specific deployed contract.
func NewIoTeXDID(address common.Address, backend bind.ContractBackend) (*IoTeXDID, error) {
	contract, err := bindIoTeXDID(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IoTeXDID{IoTeXDIDCaller: IoTeXDIDCaller{contract: contract}, IoTeXDIDTransactor: IoTeXDIDTransactor{contract: contract}, IoTeXDIDFilterer: IoTeXDIDFilterer{contract: contract}}, nil
}

// NewIoTeXDIDCaller creates a new read-only instance of IoTeXDID, bound to a specific deployed contract.
func NewIoTeXDIDCaller(address common.Address, caller bind.ContractCaller) (*IoTeXDIDCaller, error) {
	contract, err := bindIoTeXDID(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDCaller{contract: contract}, nil
}

// NewIoTeXDIDTransactor creates a new write-only instance of IoTeXDID, bound to a specific deployed contract.
func NewIoTeXDIDTransactor(address common.Address, transactor bind.ContractTransactor) (*IoTeXDIDTransactor, error) {
	contract, err := bindIoTeXDID(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDTransactor{contract: contract}, nil
}

// NewIoTeXDIDFilterer creates a new log filterer instance of IoTeXDID, bound to a specific deployed contract.
func NewIoTeXDIDFilterer(address common.Address, filterer bind.ContractFilterer) (*IoTeXDIDFilterer, error) {
	contract, err := bindIoTeXDID(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDFilterer{contract: contract}, nil
}

// bindIoTeXDID binds a generic wrapper to an already deployed contract.
func bindIoTeXDID(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IoTeXDIDABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IoTeXDID *IoTeXDIDRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IoTeXDID.Contract.IoTeXDIDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IoTeXDID *IoTeXDIDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IoTeXDID.Contract.IoTeXDIDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IoTeXDID *IoTeXDIDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IoTeXDID.Contract.IoTeXDIDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IoTeXDID *IoTeXDIDCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IoTeXDID.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IoTeXDID *IoTeXDIDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IoTeXDID.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IoTeXDID *IoTeXDIDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IoTeXDID.Contract.contract.Transact(opts, method, params...)
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(string did) view returns(bytes32)
func (_IoTeXDID *IoTeXDIDCaller) GetHash(opts *bind.CallOpts, did string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IoTeXDID.contract.Call(opts, out, "getHash", did)
	return *ret0, err
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(string did) view returns(bytes32)
func (_IoTeXDID *IoTeXDIDSession) GetHash(did string) ([32]byte, error) {
	return _IoTeXDID.Contract.GetHash(&_IoTeXDID.CallOpts, did)
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(string did) view returns(bytes32)
func (_IoTeXDID *IoTeXDIDCallerSession) GetHash(did string) ([32]byte, error) {
	return _IoTeXDID.Contract.GetHash(&_IoTeXDID.CallOpts, did)
}

// GetURI is a free data retrieval call binding the contract method 0x93ff5d3e.
//
// Solidity: function getURI(string did) view returns(string)
func (_IoTeXDID *IoTeXDIDCaller) GetURI(opts *bind.CallOpts, did string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IoTeXDID.contract.Call(opts, out, "getURI", did)
	return *ret0, err
}

// GetURI is a free data retrieval call binding the contract method 0x93ff5d3e.
//
// Solidity: function getURI(string did) view returns(string)
func (_IoTeXDID *IoTeXDIDSession) GetURI(did string) (string, error) {
	return _IoTeXDID.Contract.GetURI(&_IoTeXDID.CallOpts, did)
}

// GetURI is a free data retrieval call binding the contract method 0x93ff5d3e.
//
// Solidity: function getURI(string did) view returns(string)
func (_IoTeXDID *IoTeXDIDCallerSession) GetURI(did string) (string, error) {
	return _IoTeXDID.Contract.GetURI(&_IoTeXDID.CallOpts, did)
}

// CreateDID is a paid mutator transaction binding the contract method 0x3e05e066.
//
// Solidity: function createDID(string id, bytes32 hash, string uri) returns()
func (_IoTeXDID *IoTeXDIDTransactor) CreateDID(opts *bind.TransactOpts, id string, hash [32]byte, uri string) (*types.Transaction, error) {
	return _IoTeXDID.contract.Transact(opts, "createDID", id, hash, uri)
}

// CreateDID is a paid mutator transaction binding the contract method 0x3e05e066.
//
// Solidity: function createDID(string id, bytes32 hash, string uri) returns()
func (_IoTeXDID *IoTeXDIDSession) CreateDID(id string, hash [32]byte, uri string) (*types.Transaction, error) {
	return _IoTeXDID.Contract.CreateDID(&_IoTeXDID.TransactOpts, id, hash, uri)
}

// CreateDID is a paid mutator transaction binding the contract method 0x3e05e066.
//
// Solidity: function createDID(string id, bytes32 hash, string uri) returns()
func (_IoTeXDID *IoTeXDIDTransactorSession) CreateDID(id string, hash [32]byte, uri string) (*types.Transaction, error) {
	return _IoTeXDID.Contract.CreateDID(&_IoTeXDID.TransactOpts, id, hash, uri)
}

// DeleteDID is a paid mutator transaction binding the contract method 0x42e643bc.
//
// Solidity: function deleteDID(string did) returns()
func (_IoTeXDID *IoTeXDIDTransactor) DeleteDID(opts *bind.TransactOpts, did string) (*types.Transaction, error) {
	return _IoTeXDID.contract.Transact(opts, "deleteDID", did)
}

// DeleteDID is a paid mutator transaction binding the contract method 0x42e643bc.
//
// Solidity: function deleteDID(string did) returns()
func (_IoTeXDID *IoTeXDIDSession) DeleteDID(did string) (*types.Transaction, error) {
	return _IoTeXDID.Contract.DeleteDID(&_IoTeXDID.TransactOpts, did)
}

// DeleteDID is a paid mutator transaction binding the contract method 0x42e643bc.
//
// Solidity: function deleteDID(string did) returns()
func (_IoTeXDID *IoTeXDIDTransactorSession) DeleteDID(did string) (*types.Transaction, error) {
	return _IoTeXDID.Contract.DeleteDID(&_IoTeXDID.TransactOpts, did)
}

// UpdateHash is a paid mutator transaction binding the contract method 0x78eab45a.
//
// Solidity: function updateHash(string did, bytes32 hash) returns()
func (_IoTeXDID *IoTeXDIDTransactor) UpdateHash(opts *bind.TransactOpts, did string, hash [32]byte) (*types.Transaction, error) {
	return _IoTeXDID.contract.Transact(opts, "updateHash", did, hash)
}

// UpdateHash is a paid mutator transaction binding the contract method 0x78eab45a.
//
// Solidity: function updateHash(string did, bytes32 hash) returns()
func (_IoTeXDID *IoTeXDIDSession) UpdateHash(did string, hash [32]byte) (*types.Transaction, error) {
	return _IoTeXDID.Contract.UpdateHash(&_IoTeXDID.TransactOpts, did, hash)
}

// UpdateHash is a paid mutator transaction binding the contract method 0x78eab45a.
//
// Solidity: function updateHash(string did, bytes32 hash) returns()
func (_IoTeXDID *IoTeXDIDTransactorSession) UpdateHash(did string, hash [32]byte) (*types.Transaction, error) {
	return _IoTeXDID.Contract.UpdateHash(&_IoTeXDID.TransactOpts, did, hash)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x1789e2d8.
//
// Solidity: function updateURI(string did, string uri) returns()
func (_IoTeXDID *IoTeXDIDTransactor) UpdateURI(opts *bind.TransactOpts, did string, uri string) (*types.Transaction, error) {
	return _IoTeXDID.contract.Transact(opts, "updateURI", did, uri)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x1789e2d8.
//
// Solidity: function updateURI(string did, string uri) returns()
func (_IoTeXDID *IoTeXDIDSession) UpdateURI(did string, uri string) (*types.Transaction, error) {
	return _IoTeXDID.Contract.UpdateURI(&_IoTeXDID.TransactOpts, did, uri)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x1789e2d8.
//
// Solidity: function updateURI(string did, string uri) returns()
func (_IoTeXDID *IoTeXDIDTransactorSession) UpdateURI(did string, uri string) (*types.Transaction, error) {
	return _IoTeXDID.Contract.UpdateURI(&_IoTeXDID.TransactOpts, did, uri)
}

// IoTeXDIDCreateDIDIterator is returned from FilterCreateDID and is used to iterate over the raw logs and unpacked data for CreateDID events raised by the IoTeXDID contract.
type IoTeXDIDCreateDIDIterator struct {
	Event *IoTeXDIDCreateDID // Event containing the contract specifics and raw log

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
func (it *IoTeXDIDCreateDIDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoTeXDIDCreateDID)
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
		it.Event = new(IoTeXDIDCreateDID)
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
func (it *IoTeXDIDCreateDIDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoTeXDIDCreateDIDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoTeXDIDCreateDID represents a CreateDID event raised by the IoTeXDID contract.
type IoTeXDIDCreateDID struct {
	Id        common.Hash
	DidString string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCreateDID is a free log retrieval operation binding the contract event 0xedcad282d2fc969322aa68caeb2b9bf9de762094ec288a9965535b8968dca17f.
//
// Solidity: event CreateDID(string indexed id, string didString)
func (_IoTeXDID *IoTeXDIDFilterer) FilterCreateDID(opts *bind.FilterOpts, id []string) (*IoTeXDIDCreateDIDIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IoTeXDID.contract.FilterLogs(opts, "CreateDID", idRule)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDCreateDIDIterator{contract: _IoTeXDID.contract, event: "CreateDID", logs: logs, sub: sub}, nil
}

// WatchCreateDID is a free log subscription operation binding the contract event 0xedcad282d2fc969322aa68caeb2b9bf9de762094ec288a9965535b8968dca17f.
//
// Solidity: event CreateDID(string indexed id, string didString)
func (_IoTeXDID *IoTeXDIDFilterer) WatchCreateDID(opts *bind.WatchOpts, sink chan<- *IoTeXDIDCreateDID, id []string) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IoTeXDID.contract.WatchLogs(opts, "CreateDID", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoTeXDIDCreateDID)
				if err := _IoTeXDID.contract.UnpackLog(event, "CreateDID", log); err != nil {
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

// ParseCreateDID is a log parse operation binding the contract event 0xedcad282d2fc969322aa68caeb2b9bf9de762094ec288a9965535b8968dca17f.
//
// Solidity: event CreateDID(string indexed id, string didString)
func (_IoTeXDID *IoTeXDIDFilterer) ParseCreateDID(log types.Log) (*IoTeXDIDCreateDID, error) {
	event := new(IoTeXDIDCreateDID)
	if err := _IoTeXDID.contract.UnpackLog(event, "CreateDID", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IoTeXDIDDeleteDIDIterator is returned from FilterDeleteDID and is used to iterate over the raw logs and unpacked data for DeleteDID events raised by the IoTeXDID contract.
type IoTeXDIDDeleteDIDIterator struct {
	Event *IoTeXDIDDeleteDID // Event containing the contract specifics and raw log

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
func (it *IoTeXDIDDeleteDIDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoTeXDIDDeleteDID)
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
		it.Event = new(IoTeXDIDDeleteDID)
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
func (it *IoTeXDIDDeleteDIDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoTeXDIDDeleteDIDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoTeXDIDDeleteDID represents a DeleteDID event raised by the IoTeXDID contract.
type IoTeXDIDDeleteDID struct {
	DidString common.Hash
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeleteDID is a free log retrieval operation binding the contract event 0xf8910781c5d425c8a634b0a60a6567ef96cae8b2aa97b0db44389fff2bbcc64f.
//
// Solidity: event DeleteDID(string indexed didString)
func (_IoTeXDID *IoTeXDIDFilterer) FilterDeleteDID(opts *bind.FilterOpts, didString []string) (*IoTeXDIDDeleteDIDIterator, error) {

	var didStringRule []interface{}
	for _, didStringItem := range didString {
		didStringRule = append(didStringRule, didStringItem)
	}

	logs, sub, err := _IoTeXDID.contract.FilterLogs(opts, "DeleteDID", didStringRule)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDDeleteDIDIterator{contract: _IoTeXDID.contract, event: "DeleteDID", logs: logs, sub: sub}, nil
}

// WatchDeleteDID is a free log subscription operation binding the contract event 0xf8910781c5d425c8a634b0a60a6567ef96cae8b2aa97b0db44389fff2bbcc64f.
//
// Solidity: event DeleteDID(string indexed didString)
func (_IoTeXDID *IoTeXDIDFilterer) WatchDeleteDID(opts *bind.WatchOpts, sink chan<- *IoTeXDIDDeleteDID, didString []string) (event.Subscription, error) {

	var didStringRule []interface{}
	for _, didStringItem := range didString {
		didStringRule = append(didStringRule, didStringItem)
	}

	logs, sub, err := _IoTeXDID.contract.WatchLogs(opts, "DeleteDID", didStringRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoTeXDIDDeleteDID)
				if err := _IoTeXDID.contract.UnpackLog(event, "DeleteDID", log); err != nil {
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

// ParseDeleteDID is a log parse operation binding the contract event 0xf8910781c5d425c8a634b0a60a6567ef96cae8b2aa97b0db44389fff2bbcc64f.
//
// Solidity: event DeleteDID(string indexed didString)
func (_IoTeXDID *IoTeXDIDFilterer) ParseDeleteDID(log types.Log) (*IoTeXDIDDeleteDID, error) {
	event := new(IoTeXDIDDeleteDID)
	if err := _IoTeXDID.contract.UnpackLog(event, "DeleteDID", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IoTeXDIDUpdateHashIterator is returned from FilterUpdateHash and is used to iterate over the raw logs and unpacked data for UpdateHash events raised by the IoTeXDID contract.
type IoTeXDIDUpdateHashIterator struct {
	Event *IoTeXDIDUpdateHash // Event containing the contract specifics and raw log

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
func (it *IoTeXDIDUpdateHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoTeXDIDUpdateHash)
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
		it.Event = new(IoTeXDIDUpdateHash)
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
func (it *IoTeXDIDUpdateHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoTeXDIDUpdateHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoTeXDIDUpdateHash represents a UpdateHash event raised by the IoTeXDID contract.
type IoTeXDIDUpdateHash struct {
	DidString common.Hash
	Hash      [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateHash is a free log retrieval operation binding the contract event 0xad2d38b24c03f1f0db04b9345ca745dcf8904f27f8977599ec4769743d5e898d.
//
// Solidity: event UpdateHash(string indexed didString, bytes32 hash)
func (_IoTeXDID *IoTeXDIDFilterer) FilterUpdateHash(opts *bind.FilterOpts, didString []string) (*IoTeXDIDUpdateHashIterator, error) {

	var didStringRule []interface{}
	for _, didStringItem := range didString {
		didStringRule = append(didStringRule, didStringItem)
	}

	logs, sub, err := _IoTeXDID.contract.FilterLogs(opts, "UpdateHash", didStringRule)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDUpdateHashIterator{contract: _IoTeXDID.contract, event: "UpdateHash", logs: logs, sub: sub}, nil
}

// WatchUpdateHash is a free log subscription operation binding the contract event 0xad2d38b24c03f1f0db04b9345ca745dcf8904f27f8977599ec4769743d5e898d.
//
// Solidity: event UpdateHash(string indexed didString, bytes32 hash)
func (_IoTeXDID *IoTeXDIDFilterer) WatchUpdateHash(opts *bind.WatchOpts, sink chan<- *IoTeXDIDUpdateHash, didString []string) (event.Subscription, error) {

	var didStringRule []interface{}
	for _, didStringItem := range didString {
		didStringRule = append(didStringRule, didStringItem)
	}

	logs, sub, err := _IoTeXDID.contract.WatchLogs(opts, "UpdateHash", didStringRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoTeXDIDUpdateHash)
				if err := _IoTeXDID.contract.UnpackLog(event, "UpdateHash", log); err != nil {
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

// ParseUpdateHash is a log parse operation binding the contract event 0xad2d38b24c03f1f0db04b9345ca745dcf8904f27f8977599ec4769743d5e898d.
//
// Solidity: event UpdateHash(string indexed didString, bytes32 hash)
func (_IoTeXDID *IoTeXDIDFilterer) ParseUpdateHash(log types.Log) (*IoTeXDIDUpdateHash, error) {
	event := new(IoTeXDIDUpdateHash)
	if err := _IoTeXDID.contract.UnpackLog(event, "UpdateHash", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IoTeXDIDUpdateURIIterator is returned from FilterUpdateURI and is used to iterate over the raw logs and unpacked data for UpdateURI events raised by the IoTeXDID contract.
type IoTeXDIDUpdateURIIterator struct {
	Event *IoTeXDIDUpdateURI // Event containing the contract specifics and raw log

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
func (it *IoTeXDIDUpdateURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoTeXDIDUpdateURI)
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
		it.Event = new(IoTeXDIDUpdateURI)
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
func (it *IoTeXDIDUpdateURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoTeXDIDUpdateURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoTeXDIDUpdateURI represents a UpdateURI event raised by the IoTeXDID contract.
type IoTeXDIDUpdateURI struct {
	DidString common.Hash
	Uri       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateURI is a free log retrieval operation binding the contract event 0xa55ea16d9a68fded3c5b007d1fbba6ed1524d9354d44c66bcc54c721108586c4.
//
// Solidity: event UpdateURI(string indexed didString, string uri)
func (_IoTeXDID *IoTeXDIDFilterer) FilterUpdateURI(opts *bind.FilterOpts, didString []string) (*IoTeXDIDUpdateURIIterator, error) {

	var didStringRule []interface{}
	for _, didStringItem := range didString {
		didStringRule = append(didStringRule, didStringItem)
	}

	logs, sub, err := _IoTeXDID.contract.FilterLogs(opts, "UpdateURI", didStringRule)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDUpdateURIIterator{contract: _IoTeXDID.contract, event: "UpdateURI", logs: logs, sub: sub}, nil
}

// WatchUpdateURI is a free log subscription operation binding the contract event 0xa55ea16d9a68fded3c5b007d1fbba6ed1524d9354d44c66bcc54c721108586c4.
//
// Solidity: event UpdateURI(string indexed didString, string uri)
func (_IoTeXDID *IoTeXDIDFilterer) WatchUpdateURI(opts *bind.WatchOpts, sink chan<- *IoTeXDIDUpdateURI, didString []string) (event.Subscription, error) {

	var didStringRule []interface{}
	for _, didStringItem := range didString {
		didStringRule = append(didStringRule, didStringItem)
	}

	logs, sub, err := _IoTeXDID.contract.WatchLogs(opts, "UpdateURI", didStringRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoTeXDIDUpdateURI)
				if err := _IoTeXDID.contract.UnpackLog(event, "UpdateURI", log); err != nil {
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

// ParseUpdateURI is a log parse operation binding the contract event 0xa55ea16d9a68fded3c5b007d1fbba6ed1524d9354d44c66bcc54c721108586c4.
//
// Solidity: event UpdateURI(string indexed didString, string uri)
func (_IoTeXDID *IoTeXDIDFilterer) ParseUpdateURI(log types.Log) (*IoTeXDIDUpdateURI, error) {
	event := new(IoTeXDIDUpdateURI)
	if err := _IoTeXDID.contract.UnpackLog(event, "UpdateURI", log); err != nil {
		return nil, err
	}
	return event, nil
}
