// Copyright Tharsis Labs Ltd.(Akila)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/AkilaChain/akila/blob/main/LICENSE)
package types

import (
	"math/big"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// AttoAkila defines the default coin denomination used in Akila in:
	//
	// - Staking parameters: denomination used as stake in the dPoS chain
	// - Mint parameters: denomination minted due to fee distribution rewards
	// - Governance parameters: denomination used for spam prevention in proposal deposits
	// - Crisis parameters: constant fee denomination used for spam prevention to check broken invariant
	// - EVM parameters: denomination used for running EVM state transitions in Akila.
	AttoAkila string = "aakila"

	// BaseDenomUnit defines the base denomination unit for Akila.
	// 1 akila = 1x10^{BaseDenomUnit} aakila
	BaseDenomUnit = 18

	// DefaultGasPrice is default gas price for evm transactions
	DefaultGasPrice = 20
)

// PowerReduction defines the default power reduction value for staking
var PowerReduction = sdkmath.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(BaseDenomUnit), nil))

// NewAkilaCoin is a utility function that returns an "aakila" coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewAkilaCoin(amount sdkmath.Int) sdk.Coin {
	return sdk.NewCoin(AttoAkila, amount)
}

// NewAkilaDecCoin is a utility function that returns an "aakila" decimal coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewAkilaDecCoin(amount sdkmath.Int) sdk.DecCoin {
	return sdk.NewDecCoin(AttoAkila, amount)
}

// NewAkilaCoinInt64 is a utility function that returns an "aakila" coin with the given int64 amount.
// The function will panic if the provided amount is negative.
func NewAkilaCoinInt64(amount int64) sdk.Coin {
	return sdk.NewInt64Coin(AttoAkila, amount)
}
