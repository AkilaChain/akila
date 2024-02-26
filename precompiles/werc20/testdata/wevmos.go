// Copyright Tharsis Labs Ltd.(Akila)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/AkilaChain/akila/blob/main/LICENSE)

package testdata

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	evmtypes "github.com/AkilaChain/akila/x/evm/types"
)

var (
	//go:embed WAKILA.json
	WakilaJSON []byte

	// WAKILAContract is the compiled contract of WAKILA
	WAKILAContract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(WakilaJSON, &WAKILAContract)
	if err != nil {
		panic(err)
	}

	if len(WAKILAContract.Bin) == 0 {
		panic("failed to load WAKILA smart contract")
	}
}
