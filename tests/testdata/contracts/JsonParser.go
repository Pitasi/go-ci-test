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

// JsonParserMetaData contains all meta data concerning the JsonParser contract.
var JsonParserMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"run\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234601c57600e6020565b61148161002c823961148190f35b6026565b60405190565b600080fdfe60806040526004361015610013575b61008c565b61001e60003561002d565b63c04062260361000e57610059565b60e01c90565b60405190565b600080fd5b600080fd5b600091031261004e57565b61003e565b60000190565b3461008757610069366004610043565b610071610f32565b610079610033565b8061008381610053565b0390f35b610039565b600080fd5b601f801991011690565b634e487b7160e01b600052604160045260246000fd5b906100bb90610091565b810190811067ffffffffffffffff8211176100d557604052565b61009b565b906100ed6100e6610033565b92836100b1565b565b67ffffffffffffffff811161010d57610109602091610091565b0190565b61009b565b9061012461011f836100ef565b6100da565b918252565b60607f3a317d2c202020207b2273223a2262222c2276223a327d20205d7d0000000000917f7b202022737472223a2022666f6f222c202022696e74223a20313233342c202060008201527f22626f6f6c223a20747275652c202022666c6f6174223a2031322e33342c202060208201527f226e65737465646172726179223a205b202020207b2273223a2261222c22762260408201520152565b6101cd607b610112565b906101da60208301610129565b565b6101e46101c3565b90565b60007f7374723a737472696e672c696e743a696e743235362c626f6f6c3a626f6f6c00910152565b610219601f610112565b90610226602083016101e7565b565b61023061020f565b90565b61090490565b60018060a01b031690565b90565b61025b61025661026092610239565b610244565b610239565b90565b61026c90610247565b90565b61027890610263565b90565b61028b610286610233565b61026f565b90565b61029790610247565b90565b6102a39061028e565b90565b90565b60e01b90565b600080fd5b600080fd5b600080fd5b67ffffffffffffffff81116102dc576102d8602091610091565b0190565b61009b565b60005b8381106102f5575050906000910152565b8060209183015181850152016102e4565b9092919261031b610316826102be565b6100da565b9381855260208501908284011161033757610335926102e1565b565b6102b9565b9080601f8301121561035a5781602061035793519101610306565b90565b6102b4565b9060208282031261039057600082015167ffffffffffffffff811161038b57610388920161033c565b90565b6102af565b61003e565b5190565b60209181520190565b6103c16103ca6020936103cf936103b881610395565b93848093610399565b958691016102e1565b610091565b0190565b90916103ee6103fc93604084019084820360008601526103a2565b9160208184039101526103a2565b90565b610407610033565b3d6000823e3d90fd5b90929192610425610420826100ef565b6100da565b938185526020850190828401116104415761043f926102e1565b565b6102b9565b9080601f830112156104645781602061046193519101610410565b90565b6102b4565b90565b61047581610469565b0361047c57565b600080fd5b9050519061048e8261046c565b565b151590565b61049e81610490565b036104a557565b600080fd5b905051906104b782610495565b565b909160608284031261050757600082015167ffffffffffffffff8111610502576104e8846104ff928501610446565b936104f68160208601610481565b936040016104aa565b90565b6102af565b61003e565b5190565b905090565b61053a610531926020926105288161050c565b94858093610510565b938491016102e1565b0190565b61054791610515565b90565b60200190565b60007f666f6f0000000000000000000000000000000000000000000000000000000000910152565b61058460038092610510565b61058d81610550565b0190565b61059a90610578565b90565b90565b60209181520190565b60007f77726f6e67207374720000000000000000000000000000000000000000000000910152565b6105de60096020926105a0565b6105e7816105a9565b0190565b61060190602081019060008183039101526105d1565b90565b1561060b57565b610613610033565b62461bcd60e51b815280610629600482016105eb565b0390fd5b90565b61064461063f6106499261062d565b610244565b610469565b90565b60007f77726f6e6720696e743235360000000000000000000000000000000000000000910152565b610681600c6020926105a0565b61068a8161064c565b0190565b6106a49060208101906000818303910152610674565b90565b156106ae57565b6106b6610033565b62461bcd60e51b8152806106cc6004820161068e565b0390fd5b60007f77726f6e6720626f6f6c00000000000000000000000000000000000000000000910152565b610705600a6020926105a0565b61070e816106d0565b0190565b61072890602081019060008183039101526106f8565b90565b1561073257565b61073a610033565b62461bcd60e51b81528061075060048201610712565b0390fd5b60007f666c6f61743a6670000000000000000000000000000000000000000000000000910152565b6107866008610112565b9061079360208301610754565b565b61079d61077c565b90565b600080fd5b60000b90565b6107b4816107a5565b036107bb57565b600080fd5b905051906107cd826107ab565b565b919060408382031261080b57610804906107e960406100da565b936107f78260008301610481565b60008601526020016107c0565b6020830152565b6107a0565b9060408282031261082a57610827916000016107cf565b90565b61003e565b6108399051610469565b90565b90565b61085361084e6108589261083c565b610244565b610469565b90565b60007f77726f6e67206d616e7469737361000000000000000000000000000000000000910152565b610890600e6020926105a0565b6108998161085b565b0190565b6108b39060208101906000818303910152610883565b90565b156108bd57565b6108c5610033565b62461bcd60e51b8152806108db6004820161089d565b0390fd5b6108e990516107a5565b90565b90565b6109036108fe610908926108ec565b610244565b6107a5565b90565b60007f77726f6e67206578706f6e656e74000000000000000000000000000000000000910152565b610940600e6020926105a0565b6109498161090b565b0190565b6109639060208101906000818303910152610933565b90565b1561096d57565b610975610033565b62461bcd60e51b81528061098b6004820161094d565b0390fd5b60207f5d00000000000000000000000000000000000000000000000000000000000000917f6e657374656461727261793a28733a737472696e672c763a696e74323536295b60008201520152565b6109e76021610112565b906109f46020830161098f565b565b6109fe6109dd565b90565b67ffffffffffffffff8111610a195760208091020190565b61009b565b600080fd5b929190610a37610a3282610a01565b6100da565b9381855260208086019202810191838311610a8e5781905b838210610a5d575050505050565b815167ffffffffffffffff8111610a8957602091610a7e878493870161033c565b815201910190610a4f565b6102b4565b610a1e565b9080601f83011215610ab157816020610aae93519101610a23565b90565b6102b4565b90602082820312610ae757600082015167ffffffffffffffff8111610ae257610adf9201610a93565b90565b6102af565b61003e565b5190565b90565b90565b610b0a610b05610b0f92610af3565b610244565b610af0565b90565b60007f77726f6e6720656c656d656e7473206c656e6774680000000000000000000000910152565b610b4760156020926105a0565b610b5081610b12565b0190565b610b6a9060208101906000818303910152610b3a565b90565b15610b7457565b610b7c610033565b62461bcd60e51b815280610b9260048201610b54565b0390fd5b634e487b7160e01b600052603260045260246000fd5b90610bb682610aec565b811015610bc7576020809102010190565b610b96565b90565b610be3610bde610be892610bcc565b610244565b610af0565b90565b9190604083820312610c2c5760008301519067ffffffffffffffff8211610c2757610c1b81610c24938601610446565b93602001610481565b90565b6102af565b61003e565b90565b610c48610c43610c4d92610c31565b610244565b610469565b90565b60007f77726f6e67207631000000000000000000000000000000000000000000000000910152565b610c8560086020926105a0565b610c8e81610c50565b0190565b610ca89060208101906000818303910152610c78565b90565b15610cb257565b610cba610033565b62461bcd60e51b815280610cd060048201610c92565b0390fd5b60007f6100000000000000000000000000000000000000000000000000000000000000910152565b610d0860018092610510565b610d1181610cd4565b0190565b610d1e90610cfc565b90565b60007f77726f6e67207331000000000000000000000000000000000000000000000000910152565b610d5660086020926105a0565b610d5f81610d21565b0190565b610d799060208101906000818303910152610d49565b90565b15610d8357565b610d8b610033565b62461bcd60e51b815280610da160048201610d63565b0390fd5b610db9610db4610dbe92610c31565b610244565b610af0565b90565b610dd5610dd0610dda92610af3565b610244565b610469565b90565b60007f77726f6e67207632000000000000000000000000000000000000000000000000910152565b610e1260086020926105a0565b610e1b81610ddd565b0190565b610e359060208101906000818303910152610e05565b90565b15610e3f57565b610e47610033565b62461bcd60e51b815280610e5d60048201610e1f565b0390fd5b60007f6200000000000000000000000000000000000000000000000000000000000000910152565b610e9560018092610510565b610e9e81610e61565b0190565b610eab90610e89565b90565b60007f77726f6e67207332000000000000000000000000000000000000000000000000910152565b610ee360086020926105a0565b610eec81610eae565b0190565b610f069060208101906000818303910152610ed6565b90565b15610f1057565b610f18610033565b62461bcd60e51b815280610f2e60048201610ef0565b0390fd5b610f3a6101dc565b610f42610228565b6000610f54610f4f61027b565b61029a565b916388956bb892610f88610f70610f6a876102a6565b936102a6565b94610f93610f7c610033565b968795869485946102a9565b8452600484016103d3565b03915afa90811561144657611078610feb61105e610fd161109195610ffa95600091611423575b506020610fc682610395565b8183010191016104b9565b939195909395610fdf610033565b9283916020830161053e565b602082018103825203826100b1565b61100c61100682610395565b9161054a565b2061105861105261101b610033565b61103a8161102b60208201610591565b602082018103825203826100b1565b61104c61104682610395565b9161054a565b2061059d565b9161059d565b14610604565b61107261106c6104d2610630565b91610469565b146106a7565b61108b6110856001610490565b91610490565b1461072b565b611099610795565b60006110ab6110a661027b565b61029a565b916388956bb8926110df6110c76110c1876102a6565b936102a6565b946110ea6110d3610033565b968795869485946102a9565b8452600484016103d3565b03915afa90811561141e57602061112061116e93611155936000916113fb575b508261111582610395565b818301019101610810565b61114f61112f6000830161082f565b61114961114367ab407c9eb052000061083f565b91610469565b146108b6565b016108df565b61116861116260126108ef565b916107a5565b14610966565b60006111786109f6565b61118861118361027b565b61029a565b6111ba6111a261119c6388956bb8966102a6565b936102a6565b946111c56111ae610033565b968795869485946102a9565b8452600484016103d3565b03915afa80156113f65761136d61134a61132e61131761120561135e956113d1976000916113d3575b5060206111fa82610395565b818301019101610ab6565b61122a61121182610aec565b61122461121e6002610af6565b91610af0565b14610b6d565b6113076112946112a361128061126461124d866112476000610bcf565b90610bac565b51602061125982610395565b818301019101610beb565b91909161127a6112746001610c34565b91610469565b14610cab565b611288610033565b9283916020830161053e565b602082018103825203826100b1565b6112b56112af82610395565b9161054a565b206113016112fb6112c4610033565b6112e3816112d460208201610d15565b602082018103825203826100b1565b6112f56112ef82610395565b9161054a565b2061059d565b9161059d565b14610d7c565b6113116001610da5565b90610bac565b51602061132382610395565b818301019101610beb565b91909161134461133e6002610dc1565b91610469565b14610e38565b611352610033565b9283916020830161053e565b602082018103825203826100b1565b61137f61137982610395565b9161054a565b206113cb6113c561138e610033565b6113ad8161139e60208201610ea2565b602082018103825203826100b1565b6113bf6113b982610395565b9161054a565b2061059d565b9161059d565b14610f09565b565b6113f091503d806000833e6113e881836100b1565b81019061035f565b386111ee565b6103ff565b61141891503d806000833e61141081836100b1565b81019061035f565b3861110a565b6103ff565b61144091503d806000833e61143881836100b1565b81019061035f565b38610fba565b6103ff56fea26469706673582212204a886eb4960cd5b3c89728e9b1b332d56f67e1b17c6b655c17337b7557bb8ac564736f6c634300081e0033",
}

// JsonParserABI is the input ABI used to generate the binding from.
// Deprecated: Use JsonParserMetaData.ABI instead.
var JsonParserABI = JsonParserMetaData.ABI

// JsonParserBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use JsonParserMetaData.Bin instead.
var JsonParserBin = JsonParserMetaData.Bin

// DeployJsonParser deploys a new Ethereum contract, binding an instance of JsonParser to it.
func DeployJsonParser(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *JsonParser, error) {
	parsed, err := JsonParserMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(JsonParserBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &JsonParser{JsonParserCaller: JsonParserCaller{contract: contract}, JsonParserTransactor: JsonParserTransactor{contract: contract}, JsonParserFilterer: JsonParserFilterer{contract: contract}}, nil
}

// JsonParser is an auto generated Go binding around an Ethereum contract.
type JsonParser struct {
	JsonParserCaller     // Read-only binding to the contract
	JsonParserTransactor // Write-only binding to the contract
	JsonParserFilterer   // Log filterer for contract events
}

// JsonParserCaller is an auto generated read-only Go binding around an Ethereum contract.
type JsonParserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JsonParserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JsonParserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JsonParserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JsonParserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JsonParserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JsonParserSession struct {
	Contract     *JsonParser       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JsonParserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JsonParserCallerSession struct {
	Contract *JsonParserCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// JsonParserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JsonParserTransactorSession struct {
	Contract     *JsonParserTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// JsonParserRaw is an auto generated low-level Go binding around an Ethereum contract.
type JsonParserRaw struct {
	Contract *JsonParser // Generic contract binding to access the raw methods on
}

// JsonParserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JsonParserCallerRaw struct {
	Contract *JsonParserCaller // Generic read-only contract binding to access the raw methods on
}

// JsonParserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JsonParserTransactorRaw struct {
	Contract *JsonParserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJsonParser creates a new instance of JsonParser, bound to a specific deployed contract.
func NewJsonParser(address common.Address, backend bind.ContractBackend) (*JsonParser, error) {
	contract, err := bindJsonParser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JsonParser{JsonParserCaller: JsonParserCaller{contract: contract}, JsonParserTransactor: JsonParserTransactor{contract: contract}, JsonParserFilterer: JsonParserFilterer{contract: contract}}, nil
}

// NewJsonParserCaller creates a new read-only instance of JsonParser, bound to a specific deployed contract.
func NewJsonParserCaller(address common.Address, caller bind.ContractCaller) (*JsonParserCaller, error) {
	contract, err := bindJsonParser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JsonParserCaller{contract: contract}, nil
}

// NewJsonParserTransactor creates a new write-only instance of JsonParser, bound to a specific deployed contract.
func NewJsonParserTransactor(address common.Address, transactor bind.ContractTransactor) (*JsonParserTransactor, error) {
	contract, err := bindJsonParser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JsonParserTransactor{contract: contract}, nil
}

// NewJsonParserFilterer creates a new log filterer instance of JsonParser, bound to a specific deployed contract.
func NewJsonParserFilterer(address common.Address, filterer bind.ContractFilterer) (*JsonParserFilterer, error) {
	contract, err := bindJsonParser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JsonParserFilterer{contract: contract}, nil
}

// bindJsonParser binds a generic wrapper to an already deployed contract.
func bindJsonParser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := JsonParserMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JsonParser *JsonParserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JsonParser.Contract.JsonParserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JsonParser *JsonParserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JsonParser.Contract.JsonParserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JsonParser *JsonParserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JsonParser.Contract.JsonParserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JsonParser *JsonParserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JsonParser.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JsonParser *JsonParserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JsonParser.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JsonParser *JsonParserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JsonParser.Contract.contract.Transact(opts, method, params...)
}

// Run is a free data retrieval call binding the contract method 0xc0406226.
//
// Solidity: function run() view returns()
func (_JsonParser *JsonParserCaller) Run(opts *bind.CallOpts) error {
	var out []interface{}
	err := _JsonParser.contract.Call(opts, &out, "run")

	if err != nil {
		return err
	}

	return err

}

// Run is a free data retrieval call binding the contract method 0xc0406226.
//
// Solidity: function run() view returns()
func (_JsonParser *JsonParserSession) Run() error {
	return _JsonParser.Contract.Run(&_JsonParser.CallOpts)
}

// Run is a free data retrieval call binding the contract method 0xc0406226.
//
// Solidity: function run() view returns()
func (_JsonParser *JsonParserCallerSession) Run() error {
	return _JsonParser.Contract.Run(&_JsonParser.CallOpts)
}
