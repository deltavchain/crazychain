// Code generated
// This file is a generated precompile contract with stubbed abstract functions.
// The file is generated by a template. Please inspect every code and comment in this file before use.

// There are some must-be-done changes waiting in the file. Each area requiring you to add your code is marked with CUSTOM CODE to make them easy to find and modify.
// Additionally there are other files you need to edit to activate your precompile.
// These areas are highlighted with comments "ADD YOUR PRECOMPILE HERE".
// For testing take a look at other precompile tests in core/stateful_precompile_test.go

/* General guidelines for precompile development:
1- Read the comment and set a suitable contract address in precompile/params.go. E.g:
	CrazyWithPowerAddress = common.HexToAddress("ASUITABLEHEXADDRESS")
2- Set gas costs here
3- It is recommended to only modify code in the highlighted areas marked with "CUSTOM CODE STARTS HERE". Modifying code outside of these areas should be done with caution and with a deep understanding of how these changes may impact the EVM.
Typically, custom codes are required in only those areas.
4- Add your upgradable config in params/precompile_config.go
5- Add your precompile upgrade in params/config.go
6- Add your solidity interface and test contract to contract-examples/contracts
7- Write solidity tests for your precompile in contract-examples/test
8- Create your genesis with your precompile enabled in tests/e2e/genesis/
9- Create e2e test for your solidity test in tests/e2e/solidity/suites.go
10- Run your e2e precompile Solidity tests with 'E2E=true ./scripts/run.sh'

*/

package precompile

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/vmerrs"

	"github.com/ethereum/go-ethereum/common"
)

const (
	SetProtectionGasCost uint64 = 0 // SET A GAS COST HERE
	StealGasCost         uint64 = 0 // SET A GAS COST HERE
	UncertainFateGasCost uint64 = 0 // SET A GAS COST HERE

	// CrazyWithPowerRawABI contains the raw ABI of CrazyWithPower contract.
	CrazyWithPowerRawABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"protection\",\"type\":\"uint256\"}],\"name\":\"setProtection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"enemy\",\"type\":\"address\"}],\"name\":\"steal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uncertainFate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
)

// CUSTOM CODE STARTS HERE
// Reference imports to suppress errors from unused imports. This code and any unnecessary imports can be removed.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = fmt.Printf
)

// Singleton StatefulPrecompiledContract and signatures.
var (
	_ StatefulPrecompileConfig = &CrazyWithPowerConfig{}

	CrazyWithPowerABI abi.ABI // will be initialized by init function

	CrazyWithPowerPrecompile StatefulPrecompiledContract // will be initialized by init function
)

// CrazyWithPowerConfig implements the StatefulPrecompileConfig
// interface while adding in the CrazyWithPower specific precompile address.
type CrazyWithPowerConfig struct {
	UpgradeableConfig
}

func init() {
	parsed, err := abi.JSON(strings.NewReader(CrazyWithPowerRawABI))
	if err != nil {
		panic(err)
	}
	CrazyWithPowerABI = parsed

	CrazyWithPowerPrecompile = createCrazyWithPowerPrecompile(CrazyWithPowerAddress)
}

// NewCrazyWithPowerConfig returns a config for a network upgrade at [blockTimestamp] that enables
// CrazyWithPower .
func NewCrazyWithPowerConfig(blockTimestamp *big.Int) *CrazyWithPowerConfig {
	return &CrazyWithPowerConfig{

		UpgradeableConfig: UpgradeableConfig{BlockTimestamp: blockTimestamp},
	}
}

// NewDisableCrazyWithPowerConfig returns config for a network upgrade at [blockTimestamp]
// that disables CrazyWithPower.
func NewDisableCrazyWithPowerConfig(blockTimestamp *big.Int) *CrazyWithPowerConfig {
	return &CrazyWithPowerConfig{
		UpgradeableConfig: UpgradeableConfig{
			BlockTimestamp: blockTimestamp,
			Disable:        true,
		},
	}
}

// Equal returns true if [s] is a [*CrazyWithPowerConfig] and it has been configured identical to [c].
func (c *CrazyWithPowerConfig) Equal(s StatefulPrecompileConfig) bool {
	// typecast before comparison
	other, ok := (s).(*CrazyWithPowerConfig)
	if !ok {
		return false
	}
	// CUSTOM CODE STARTS HERE
	// modify this boolean accordingly with your custom CrazyWithPowerConfig, to check if [other] and the current [c] are equal
	// if CrazyWithPowerConfig contains only UpgradeableConfig  you can skip modifying it.
	equals := c.UpgradeableConfig.Equal(&other.UpgradeableConfig)
	return equals
}

// Address returns the address of the CrazyWithPower. Addresses reside under the precompile/params.go
// Select a non-conflicting address and set it in the params.go.
func (c *CrazyWithPowerConfig) Address() common.Address {
	return CrazyWithPowerAddress
}

// Configure configures [state] with the initial configuration.
func (c *CrazyWithPowerConfig) Configure(_ ChainConfig, state StateDB, _ BlockContext) {

	// CUSTOM CODE STARTS HERE
}

// Contract returns the singleton stateful precompiled contract to be used for CrazyWithPower.
func (c *CrazyWithPowerConfig) Contract() StatefulPrecompiledContract {
	return CrazyWithPowerPrecompile
}

// Verify tries to verify CrazyWithPowerConfig and returns an error accordingly.
func (c *CrazyWithPowerConfig) Verify() error {

	// CUSTOM CODE STARTS HERE
	// Add your own custom verify code for CrazyWithPowerConfig here
	// and return an error accordingly
	return nil
}

// UnpackSetProtectionInput attempts to unpack [input] into the *big.Int type argument
// assumes that [input] does not include selector (omits first 4 func signature bytes)
func UnpackSetProtectionInput(input []byte) (*big.Int, error) {
	res, err := CrazyWithPowerABI.UnpackInput("setProtection", input)
	if err != nil {
		return 0, err
	}
	unpacked := *abi.ConvertType(res[0], new(*big.Int)).(**big.Int)
	return unpacked, nil
}

// PackSetProtection packs [protection] of type *big.Int into the appropriate arguments for setProtection.
// the packed bytes include selector (first 4 func signature bytes).
// This function is mostly used for tests.
func PackSetProtection(protection *big.Int) ([]byte, error) {
	return CrazyWithPowerABI.Pack("setProtection", protection)
}

func setProtection(accessibleState PrecompileAccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	if remainingGas, err = deductGas(suppliedGas, SetProtectionGasCost); err != nil {
		return nil, 0, err
	}
	if readOnly {
		return nil, remainingGas, vmerrs.ErrWriteProtection
	}
	// attempts to unpack [input] into the arguments to the SetProtectionInput.
	// Assumes that [input] does not include selector
	// You can use unpacked [inputStruct] variable in your code
	inputStruct, err := UnpackSetProtectionInput(input)
	if err != nil {
		return nil, remainingGas, err
	}

	// CUSTOM CODE STARTS HERE
	_ = inputStruct // CUSTOM CODE OPERATES ON INPUT
	// this function does not return an output, leave this one as is
	packedOutput := []byte{}

	// Return the packed output and the remaining gas
	return packedOutput, remainingGas, nil
}

// UnpackStealInput attempts to unpack [input] into the common.Address type argument
// assumes that [input] does not include selector (omits first 4 func signature bytes)
func UnpackStealInput(input []byte) (common.Address, error) {
	res, err := CrazyWithPowerABI.UnpackInput("steal", input)
	if err != nil {
		return common.Address{}, err
	}
	unpacked := *abi.ConvertType(res[0], new(common.Address)).(*common.Address)
	return unpacked, nil
}

// PackSteal packs [enemy] of type common.Address into the appropriate arguments for steal.
// the packed bytes include selector (first 4 func signature bytes).
// This function is mostly used for tests.
func PackSteal(enemy common.Address) ([]byte, error) {
	return CrazyWithPowerABI.Pack("steal", enemy)
}

func steal(accessibleState PrecompileAccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	if remainingGas, err = deductGas(suppliedGas, StealGasCost); err != nil {
		return nil, 0, err
	}
	if readOnly {
		return nil, remainingGas, vmerrs.ErrWriteProtection
	}
	// attempts to unpack [input] into the arguments to the StealInput.
	// Assumes that [input] does not include selector
	// You can use unpacked [inputStruct] variable in your code
	inputStruct, err := UnpackStealInput(input)
	if err != nil {
		return nil, remainingGas, err
	}

	// CUSTOM CODE STARTS HERE
	_ = inputStruct // CUSTOM CODE OPERATES ON INPUT
	// this function does not return an output, leave this one as is
	packedOutput := []byte{}

	// Return the packed output and the remaining gas
	return packedOutput, remainingGas, nil
}

// PackUncertainFate packs the include selector (first 4 func signature bytes).
// This function is mostly used for tests.
func PackUncertainFate() ([]byte, error) {
	return CrazyWithPowerABI.Pack("uncertainFate")
}

func uncertainFate(accessibleState PrecompileAccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	if remainingGas, err = deductGas(suppliedGas, UncertainFateGasCost); err != nil {
		return nil, 0, err
	}
	if readOnly {
		return nil, remainingGas, vmerrs.ErrWriteProtection
	}
	// no input provided for this function

	// CUSTOM CODE STARTS HERE
	// this function does not return an output, leave this one as is
	packedOutput := []byte{}

	// Return the packed output and the remaining gas
	return packedOutput, remainingGas, nil
}

// createCrazyWithPowerPrecompile returns a StatefulPrecompiledContract with getters and setters for the precompile.

func createCrazyWithPowerPrecompile(precompileAddr common.Address) StatefulPrecompiledContract {
	var functions []*statefulPrecompileFunction

	methodSetProtection, ok := CrazyWithPowerABI.Methods["setProtection"]
	if !ok {
		panic("given method does not exist in the ABI")
	}
	functions = append(functions, newStatefulPrecompileFunction(methodSetProtection.ID, setProtection))

	methodSteal, ok := CrazyWithPowerABI.Methods["steal"]
	if !ok {
		panic("given method does not exist in the ABI")
	}
	functions = append(functions, newStatefulPrecompileFunction(methodSteal.ID, steal))

	methodUncertainFate, ok := CrazyWithPowerABI.Methods["uncertainFate"]
	if !ok {
		panic("given method does not exist in the ABI")
	}
	functions = append(functions, newStatefulPrecompileFunction(methodUncertainFate.ID, uncertainFate))

	// Construct the contract with no fallback function.
	contract := newStatefulPrecompileWithFunctionSelectors(nil, functions)
	return contract
}
