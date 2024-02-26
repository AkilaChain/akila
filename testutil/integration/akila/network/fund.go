// Copyright Tharsis Labs Ltd.(Akila)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/AkilaChain/akila/blob/main/LICENSE)

package network

import (
	sdkmath "cosmossdk.io/math"
	inflationtypes "github.com/AkilaChain/akila/x/inflation/v1/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// FundAccount funds the given account with the given amount of coins.
func (n *IntegrationNetwork) FundAccount(addr sdk.AccAddress, coins sdk.Coins) error {
	ctx := n.GetContext()

	if err := n.app.BankKeeper.MintCoins(ctx, inflationtypes.ModuleName, coins); err != nil {
		return err
	}

	return n.app.BankKeeper.SendCoinsFromModuleToAccount(ctx, inflationtypes.ModuleName, addr, coins)
}

// FundAccountWithBaseDenom funds the given account with the given amount of the network's
// base denomination.
func (n *IntegrationNetwork) FundAccountWithBaseDenom(addr sdk.AccAddress, amount sdkmath.Int) error {
	return n.FundAccount(addr, sdk.NewCoins(sdk.NewCoin(n.GetDenom(), amount)))
}
