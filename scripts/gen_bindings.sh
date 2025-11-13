#!/bin/sh

source ./check_solc_version.sh
check_solc_version

contractroot=../../prover-network-contracts
out=$contractroot/out
contractdir=$contractroot/src

echo "run solc"
solc --overwrite --optimize --via-ir --pretty-json \
  --base-path $contractroot \
  --allow-paths $contractdir \
  --optimize-runs 20000 \
  --combined-json abi,bin \
  --evm-version london \
  -o $contractroot/out \
  '@openzeppelin/'=./lib/openzeppelin/ \
  '@security/'=./lib/security/src/ \
  $contractdir/market/BrevisMarket.sol \
  $contractdir/market/MarketViewer.sol

echo "run abigen"
abigen -combined-json $out/combined.json -pkg eth -out ../eth/bindings.go
echo "clean up"
rm -rf $contractroot/out
echo "done"
