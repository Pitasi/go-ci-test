package async

import (
	"embed"
	"fmt"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	evmcmn "github.com/cosmos/evm/precompiles/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcmn "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/vm"

	"github.com/warden-protocol/wardenprotocol/precompiles/common"
	asyncmodulekeeper "github.com/warden-protocol/wardenprotocol/warden/x/async/keeper"
	types "github.com/warden-protocol/wardenprotocol/warden/x/async/types/v1beta1"
)

var _ vm.PrecompiledContract = &Precompile{}

const PrecompileAddress = "0x0000000000000000000000000000000000000903"

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

// Precompile defines the precompiled contract for x/async.
type Precompile struct {
	evmcmn.Precompile

	asyncmodulekeeper asyncmodulekeeper.Keeper
	eventsRegistry    *common.EthEventsRegistry
	queryServer       types.QueryServer
}

// LoadABI loads the x/async ABI from the embedded abi.json file
// for the x/async precompile.
func LoadABI() (abi.ABI, error) {
	return evmcmn.LoadABI(f, "abi.json")
}

func NewPrecompile(asynckeeper asyncmodulekeeper.Keeper, er *common.EthEventsRegistry) (*Precompile, error) {
	abi, err := LoadABI()
	if err != nil {
		return nil, err
	}

	p := Precompile{
		Precompile: evmcmn.Precompile{
			ABI:                  abi,
			KvGasConfig:          storetypes.KVGasConfig(),
			TransientKVGasConfig: storetypes.TransientGasConfig(),
		},
		asyncmodulekeeper: asynckeeper,
		eventsRegistry:    er,
		queryServer:       asyncmodulekeeper.NewQueryServerImpl(asynckeeper),
	}

	p.SetAddress(p.Address())

	return &p, nil
}

// Address implements vm.PrecompiledContract.
func (p *Precompile) Address() ethcmn.Address {
	return ethcmn.HexToAddress(PrecompileAddress)
}

// RequiredGas implements vm.PrecompiledContract.
// Subtle: this method shadows the method (Precompile).RequiredGas of Precompile.Precompile.
func (p *Precompile) RequiredGas(input []byte) uint64 {
	// NOTE: This check avoid panicking when trying to decode the method ID
	if len(input) < 4 {
		return 0
	}

	methodID := input[:4]

	method, err := p.MethodById(methodID)
	if err != nil {
		// This should never happen since this method is going to fail during Run
		return 0
	}

	return p.Precompile.RequiredGas(input, p.IsTransaction(method))
}

// Run implements vm.PrecompiledContract.
func (p *Precompile) Run(evm *vm.EVM, contract *vm.Contract, readonly bool) (bz []byte, err error) {
	ctx, stateDB, method, initialGas, args, err := p.RunSetup(evm, contract, readonly, p.IsTransaction)
	if err != nil {
		return nil, err
	}

	p.GetBalanceHandler().BeforeBalanceChange(ctx)

	// This handles any out of gas errors that may occur during the execution of a precompile tx or query.
	// It avoids panics and returns the out of gas error so the EVM can continue gracefully.
	defer evmcmn.HandleGasError(ctx, contract, initialGas, &err)()

	switch method.Name {
	// transactions
	case AddTaskMethod:
		bz, err = p.AddTaskMethod(ctx, evm.Origin, stateDB, method, args)
	// queries
	case TaskByIdMethod:
		bz, err = p.TaskByIdMethod(ctx, method, args)
	case TasksMethod:
		bz, err = p.TasksMethod(ctx, method, args)
	case PendingTasksMethod:
		bz, err = p.PendingTasksMethod(ctx, method, args)
	case PluginsMethod:
		bz, err = p.PluginsMethod(ctx, method, args)
	}

	if err != nil {
		return nil, err
	}

	cost := ctx.GasMeter().GasConsumed() - initialGas

	if !contract.UseGas(cost, nil, tracing.GasChangeCallPrecompiledContract) {
		return nil, vm.ErrOutOfGas
	}

	if err = p.GetBalanceHandler().AfterBalanceChange(ctx, stateDB); err != nil {
		return nil, err
	}

	return bz, nil
}

func (p *Precompile) IsTransaction(method *abi.Method) bool {
	switch method.Name {
	// transactions
	case AddTaskMethod:
		return true
	// queries
	case TaskByIdMethod,
		TasksMethod,
		PendingTasksMethod,
		PluginsMethod:
		return false
	}

	panic(fmt.Errorf("async precompile: method not exists: %s", method))
}

// Logger returns a precompile-specific logger.
func (p *Precompile) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("evm extension", "async")
}
