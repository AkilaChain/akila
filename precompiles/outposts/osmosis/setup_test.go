// Copyright Tharsis Labs Ltd.(Akila)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/AkilaChain/akila/blob/main/LICENSE)

package osmosis_test

import (
	"testing"

	"github.com/AkilaChain/akila/precompiles/erc20"

	"github.com/AkilaChain/akila/precompiles/outposts/osmosis"
	"github.com/AkilaChain/akila/testutil/integration/akila/grpc"
	testkeyring "github.com/AkilaChain/akila/testutil/integration/akila/keyring"
	"github.com/AkilaChain/akila/testutil/integration/akila/network"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/suite"
)

const (
	PortID      = "transfer"
	ChannelID   = "channel-0"
	XCSContract = "osmo1a34wxsxjwvtz3ua4hnkh4lv3d4qrgry0fhkasppplphwu5k538tqcyms9x"
)

type PrecompileTestSuite struct {
	suite.Suite

	unitNetwork *network.UnitTestNetwork
	grpcHandler grpc.Handler
	keyring     testkeyring.Keyring

	precompile *osmosis.Precompile
}

func TestPrecompileTestSuite(t *testing.T) {
	suite.Run(t, new(PrecompileTestSuite))
}

func (s *PrecompileTestSuite) SetupTest() {
	keyring := testkeyring.New(2)
	unitNetwork := network.NewUnitTestNetwork(
		network.WithPreFundedAccounts(keyring.GetAllAccAddrs()...),
	)

	precompile, err := osmosis.NewPrecompile(
		common.HexToAddress(erc20.WAKILAContractTestnet),
		unitNetwork.App.AuthzKeeper,
		unitNetwork.App.BankKeeper,
		unitNetwork.App.TransferKeeper,
		unitNetwork.App.StakingKeeper,
		unitNetwork.App.Erc20Keeper,
		unitNetwork.App.IBCKeeper.ChannelKeeper,
	)
	s.Require().NoError(err, "expected no error during precompile creation")

	grpcHandler := grpc.NewIntegrationHandler(unitNetwork)

	s.unitNetwork = unitNetwork
	s.grpcHandler = grpcHandler
	s.keyring = keyring
	s.precompile = precompile
}
