// Copyright Tharsis Labs Ltd.(Akila)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/AkilaChain/akila/blob/main/LICENSE)

package keeper

import (
	"github.com/AkilaChain/akila/x/inflation/v1/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetPeriod gets current period
func (k Keeper) GetPeriod(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixPeriod)
	if len(bz) == 0 {
		return 0
	}

	return sdk.BigEndianToUint64(bz)
}

// SetPeriod stores the current period
func (k Keeper) SetPeriod(ctx sdk.Context, period uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefixPeriod, sdk.Uint64ToBigEndian(period))
}
