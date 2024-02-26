// Copyright Tharsis Labs Ltd.(Akila)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/AkilaChain/akila/blob/main/LICENSE)

package testdata

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	evmtypes "github.com/AkilaChain/akila/x/evm/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/AkilaChain/akila/x/erc20/types"
)

var (
	//go:embed ERC20Minter_OpenZeppelinV5.json
	ERC20MinterV5JSON []byte //nolint: golint

	// ERC20MinterV5Contract is the compiled erc20 contract
	ERC20MinterV5Contract evmtypes.CompiledContract

	// ERC20MinterV5Address is the erc20 module address
	ERC20MinterV5Address common.Address
)

func init() {
	ERC20MinterV5Address = types.ModuleAddress

	err := json.Unmarshal(ERC20MinterV5JSON, &ERC20MinterV5Contract)
	if err != nil {
		panic(err)
	}

	if len(ERC20MinterV5Contract.Bin) == 0 {
		panic("load contract failed")
	}
}
