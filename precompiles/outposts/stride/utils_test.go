// Copyright Tharsis Labs Ltd.(Akila)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/AkilaChain/akila/blob/main/LICENSE)
package stride_test

import (
	"fmt"

	"cosmossdk.io/math"
	commonnetwork "github.com/AkilaChain/akila/testutil/integration/common/network"
	"github.com/AkilaChain/akila/testutil/integration/ibc/coordinator"

	erc20types "github.com/AkilaChain/akila/x/erc20/types"

	inflationtypes "github.com/AkilaChain/akila/x/inflation/v1/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
)

const (
	portID    = "transfer"
	channelID = "channel-0"
)

// registerStrideCoinERC20 registers stAkila and Akila coin as an ERC20 token
func (s *PrecompileTestSuite) registerStrideCoinERC20() {
	// Register AKILA ERC20 equivalent
	ctx := s.network.GetContext()
	bondDenom := s.network.App.StakingKeeper.BondDenom(ctx)
	akilaMetadata, found := s.network.App.BankKeeper.GetDenomMetaData(ctx, bondDenom)
	s.Require().True(found, "expected akila denom metadata")

	coin := sdk.NewCoin(akilaMetadata.Base, math.NewInt(2e18))
	err := s.network.App.BankKeeper.MintCoins(ctx, inflationtypes.ModuleName, sdk.NewCoins(coin))
	s.Require().NoError(err)

	// Register some Token Pairs
	_, err = s.network.App.Erc20Keeper.RegisterCoin(ctx, akilaMetadata)
	s.Require().NoError(err)

	// Register stAkila Token Pair
	denomTrace := transfertypes.DenomTrace{
		Path:      fmt.Sprintf("%s/%s", portID, channelID),
		BaseDenom: "st" + bondDenom,
	}
	s.network.App.TransferKeeper.SetDenomTrace(ctx, denomTrace)
	stAkilaMetadata := banktypes.Metadata{
		Description: "The native token of Akila",
		Base:        denomTrace.IBCDenom(),
		// NOTE: Denom units MUST be increasing
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    denomTrace.IBCDenom(),
				Exponent: 0,
				Aliases:  []string{denomTrace.BaseDenom},
			},
			{
				Denom:    denomTrace.BaseDenom,
				Exponent: 18,
			},
		},
		Name:    "stAkila",
		Symbol:  "STAKILA",
		Display: denomTrace.BaseDenom,
	}

	stAkila := sdk.NewCoin(stAkilaMetadata.Base, math.NewInt(9e18))
	err = s.network.App.BankKeeper.MintCoins(ctx, inflationtypes.ModuleName, sdk.NewCoins(stAkila))
	s.Require().NoError(err)
	err = s.network.App.BankKeeper.SendCoinsFromModuleToAccount(ctx, inflationtypes.ModuleName, s.keyring.GetAccAddr(0), sdk.NewCoins(stAkila))
	s.Require().NoError(err)

	// Register some Token Pairs
	_, err = s.network.App.Erc20Keeper.RegisterCoin(ctx, stAkilaMetadata)
	s.Require().NoError(err)

	convertCoin := erc20types.NewMsgConvertCoin(
		stAkila,
		s.keyring.GetAddr(0),
		s.keyring.GetAccAddr(0),
	)

	_, err = s.network.App.Erc20Keeper.ConvertCoin(ctx, convertCoin)
	s.Require().NoError(err)
}

// setupIBCCoordinator sets up the IBC coordinator
func (s *PrecompileTestSuite) setupIBCCoordinator() {
	ibcSender, ibcSenderPrivKey := s.keyring.GetAccAddr(0), s.keyring.GetPrivKey(0)
	ibcAcc, err := s.grpcHandler.GetAccount(ibcSender.String())
	s.Require().NoError(err)

	IBCCoordinator := coordinator.NewIntegrationCoordinator(
		s.T(),
		[]commonnetwork.Network{s.network},
	)

	IBCCoordinator.SetDefaultSignerForChain(s.network.GetChainID(), ibcSenderPrivKey, ibcAcc)
	IBCCoordinator.Setup(s.network.GetChainID(), IBCCoordinator.GetDummyChainsIds()[0])

	err = IBCCoordinator.CommitAll()
	s.Require().NoError(err)
}
