#!/bin/sh

expected=0.8.29+commit.ab55807c

check_solc_version() {
  version=$(solc --version | tail -n 1)
  if [[ $version != *$expected* ]]; then
    echo "Solidity version mismatch. Expected $expected, but got $version"
    exit 1
  fi
}
