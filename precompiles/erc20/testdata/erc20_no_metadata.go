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
	//go:embed ERC20NoMetadata.json
	ERC20NoMetadataJSON []byte //nolint: golint

	// ERC20NoMetadataContract is the compiled erc20 contract
	ERC20NoMetadataContract evmtypes.CompiledContract

	// ERC20NoMetadataAddress is the erc20 module address
	ERC20NoMetadataAddress common.Address
)

func init() {
	ERC20NoMetadataAddress = types.ModuleAddress

	err := json.Unmarshal(ERC20NoMetadataJSON, &ERC20NoMetadataContract)
	if err != nil {
		panic(err)
	}

	if len(ERC20NoMetadataContract.Bin) == 0 {
		panic("load contract failed")
	}
}
