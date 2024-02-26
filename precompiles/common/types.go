// Copyright Tharsis Labs Ltd.(Akila)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/AkilaChain/akila/blob/main/LICENSE)

package common

import (
	"math/big"
	"strings"
	"time"

	"cosmossdk.io/math"
	akilautils "github.com/AkilaChain/akila/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

var (
	// TrueValue is the byte array representing a true value in solidity.
	TrueValue = []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1}
	// DefaultExpirationDuration is the default duration for an authorization to expire.
	DefaultExpirationDuration = time.Hour * 24 * 365
	// DefaultChainID is the standard chain id used for testing purposes
	DefaultChainID = akilautils.MainnetChainID + "-1"
	// DefaultPrecompilesBech32 is the standard bech32 address for the precompiles
	DefaultPrecompilesBech32 = []string{
		"akila1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqn8xnepl", // secp256r1 curve precompile
		"akila1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqpqq7k8qrt", // bech32 precompile
		"akila1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqzqqcsusp2", // Staking precompile
		"akila1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqzqp9xg9uc", // Distribution precompile
		"akila1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqzqzt4anj8", // ICS20 transfer precompile
		"akila1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqzqrkrfx04", // Vesting precompile
		"akila1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqzqyh67kws", // Bank precompile
	}
)

// ICS20Allocation defines the spend limit for a particular port and channel.
// We need this to be able to unpack to big.Int instead of math.Int.
type ICS20Allocation struct {
	SourcePort    string
	SourceChannel string
	SpendLimit    []Coin
	AllowList     []string
}

// Coin defines a struct that stores all needed information about a coin
// in types native to the EVM.
type Coin struct {
	Denom  string
	Amount *big.Int
}

// DecCoin defines a struct that stores all needed information about a decimal coin
// in types native to the EVM.
type DecCoin struct {
	Denom     string
	Amount    *big.Int
	Precision uint8
}

// Dec defines a struct that represents a decimal number of a given precision
// in types native to the EVM.
type Dec struct {
	Value     *big.Int
	Precision uint8
}

// ToSDKType converts the Coin to the Cosmos SDK representation.
func (c Coin) ToSDKType() sdk.Coin {
	return sdk.NewCoin(c.Denom, math.NewIntFromBigInt(c.Amount))
}

// NewCoinsResponse converts a response to an array of Coin.
func NewCoinsResponse(amount sdk.Coins) []Coin {
	// Create a new output for each coin and add it to the output array.
	outputs := make([]Coin, len(amount))
	for i, coin := range amount {
		outputs[i] = Coin{
			Denom:  coin.Denom,
			Amount: coin.Amount.BigInt(),
		}
	}
	return outputs
}

// NewDecCoinsResponse converts a response to an array of DecCoin.
func NewDecCoinsResponse(amount sdk.DecCoins) []DecCoin {
	// Create a new output for each coin and add it to the output array.
	outputs := make([]DecCoin, len(amount))
	for i, coin := range amount {
		outputs[i] = DecCoin{
			Denom:     coin.Denom,
			Amount:    coin.Amount.TruncateInt().BigInt(),
			Precision: math.LegacyPrecision,
		}
	}
	return outputs
}

// HexAddressFromBech32String converts a hex address to a bech32 encoded address.
func HexAddressFromBech32String(addr string) (res common.Address, err error) {
	if strings.Contains(addr, sdk.PrefixValidator) {
		valAddr, err := sdk.ValAddressFromBech32(addr)
		if err != nil {
			return res, err
		}
		return common.BytesToAddress(valAddr.Bytes()), nil
	}
	return common.BytesToAddress(sdk.MustAccAddressFromBech32(addr)), nil
}

// SafeAdd adds two integers and returns a boolean if an overflow occurs to avoid panic.
// TODO: Upstream this to the SDK math package.
func SafeAdd(a, b math.Int) (res *big.Int, overflow bool) {
	res = a.BigInt().Add(a.BigInt(), b.BigInt())
	return res, res.BitLen() > math.MaxBitLen
}