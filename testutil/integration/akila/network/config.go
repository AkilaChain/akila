// Copyright Tharsis Labs Ltd.(Akila)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/AkilaChain/akila/blob/main/LICENSE)
package network

import (
	"math/big"

	testtx "github.com/AkilaChain/akila/testutil/tx"
	"github.com/AkilaChain/akila/utils"

	akilatypes "github.com/AkilaChain/akila/types"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
)

// Config defines the configuration for a chain.
// It allows for customization of the network to adjust to
// testing needs.
type Config struct {
	chainID            string
	eip155ChainID      *big.Int
	amountOfValidators int
	preFundedAccounts  []sdktypes.AccAddress
	denom              string
}

// DefaultConfig returns the default configuration for a chain.
func DefaultConfig() Config {
	account, _ := testtx.NewAccAddressAndKey()
	return Config{
		chainID:            utils.MainnetChainID + "-1",
		eip155ChainID:      big.NewInt(2927),
		amountOfValidators: 3,
		// No funded accounts besides the validators by default
		preFundedAccounts: []sdktypes.AccAddress{account},
		denom:             utils.BaseDenom,
	}
}

// ConfigOption defines a function that can modify the NetworkConfig.
// The purpose of this is to force to be declarative when the default configuration
// requires to be changed.
type ConfigOption func(*Config)

// WithChainID sets a custom chainID for the network. It panics if the chainID is invalid.
func WithChainID(chainID string) ConfigOption {
	chainIDNum, err := akilatypes.ParseChainID(chainID)
	if err != nil {
		panic(err)
	}
	return func(cfg *Config) {
		cfg.chainID = chainID
		cfg.eip155ChainID = chainIDNum
	}
}

// WithAmountOfValidators sets the amount of validators for the network.
func WithAmountOfValidators(amount int) ConfigOption {
	return func(cfg *Config) {
		cfg.amountOfValidators = amount
	}
}

// WithPreFundedAccounts sets the pre-funded accounts for the network.
func WithPreFundedAccounts(accounts ...sdktypes.AccAddress) ConfigOption {
	return func(cfg *Config) {
		cfg.preFundedAccounts = accounts
	}
}

// WithDenom sets the denom for the network.
func WithDenom(denom string) ConfigOption {
	return func(cfg *Config) {
		cfg.denom = denom
	}
}
