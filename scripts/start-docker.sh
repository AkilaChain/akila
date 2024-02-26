#!/bin/bash

KEY="dev0"
CHAINID="akila_9000-1"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t akila-datadir.XXXXX)

echo "create and add new keys"
./akilad keys add $KEY --home $DATA_DIR --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init Akila with moniker=$MONIKER and chain-id=$CHAINID"
./akilad init $MONIKER --chain-id $CHAINID --home $DATA_DIR
echo "prepare genesis: Allocate genesis accounts"
./akilad add-genesis-account \
"$(./akilad keys show $KEY -a --home $DATA_DIR --keyring-backend test)" 1000000000000000000aakila,1000000000000000000stake \
--home $DATA_DIR --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./akilad gentx $KEY 1000000000000000000stake --keyring-backend test --home $DATA_DIR --keyring-backend test --chain-id $CHAINID
echo "prepare genesis: Collect genesis tx"
./akilad collect-gentxs --home $DATA_DIR
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./akilad validate-genesis --home $DATA_DIR

echo "starting akila node $i in background ..."
./akilad start --pruning=nothing --rpc.unsafe \
--keyring-backend test --home $DATA_DIR \
>$DATA_DIR/node.log 2>&1 & disown

echo "started akila node"
tail -f /dev/null