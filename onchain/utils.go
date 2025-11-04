package onchain

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/celer-network/goutils/eth"
	"github.com/cenkalti/backoff/v5"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TxMethod func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error)

func TransactWaitSuccess(
	ec *ethclient.Client, signer *bind.TransactOpts, method TxMethod) (*ethtypes.Transaction, *ethtypes.Receipt, error) {
	op := func() (*ethtypes.Transaction, error) {
		tx, err := method(ec, signer)
		if err != nil {
			// retry on common transient errors
			if isTransientError(err) {
				return nil, err
			}
			return nil, backoff.Permanent(err)
		}
		return tx, nil
	}
	tx, err := backoff.Retry(context.Background(), op, backoff.WithBackOff(backoff.NewExponentialBackOff()))
	if err != nil {
		return nil, nil, err
	}
	receipt, err := waitMinedTxSuccess(ec, tx)
	if err != nil {
		return nil, nil, fmt.Errorf("waitMinedTxSuccess err: %w", err)
	}
	return tx, receipt, nil
}

func waitMinedTxSuccess(ec *ethclient.Client, tx *ethtypes.Transaction) (*ethtypes.Receipt, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	receipt, err := bind.WaitMined(ctx, ec, tx)
	if err != nil {
		return nil, fmt.Errorf("WaitMined err: %w", err)
	}
	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return nil, fmt.Errorf("tx failed: %s", tx.Hash().Hex())
	}
	return receipt, nil
}

func isTransientError(err error) bool {
	errStr := err.Error()
	return strings.Contains(errStr, "temporarily") ||
		strings.Contains(errStr, "nonce too low") ||
		strings.Contains(errStr, "InvalidState") ||
		strings.Contains(errStr, "invalid state")
}

// example: awskms:us-west-2:alias/xxxxx
func CreateTransactOpts(ksfilePath, passphrase string, chainid *big.Int) (*bind.TransactOpts, common.Address, error) {
	if strings.HasPrefix(ksfilePath, "awskms") {
		kmskeyinfo := strings.SplitN(ksfilePath, ":", 3)
		if len(kmskeyinfo) != 3 {
			return nil, ZeroAddr, fmt.Errorf("%s has wrong format", ksfilePath)
		}
		awskeysec := []string{"", ""}
		if passphrase != "" {
			awskeysec = strings.SplitN(passphrase, ":", 2)
			if len(awskeysec) != 2 {
				return nil, ZeroAddr, fmt.Errorf("%s has wrong format", passphrase)
			}
		}
		kmsSigner, err := eth.NewKmsSigner(kmskeyinfo[1], kmskeyinfo[2], awskeysec[0], awskeysec[1], chainid)
		if err != nil {
			return nil, ZeroAddr, err
		}
		return kmsSigner.NewTransactOpts(), kmsSigner.Addr, nil
	}
	ksBytes, err := os.ReadFile(ksfilePath)
	if err != nil {
		return nil, ZeroAddr, err
	}

	key, err := keystore.DecryptKey(ksBytes, passphrase)
	if err != nil {
		return nil, ZeroAddr, err
	}

	submitChainAuth, err :=
		bind.NewTransactorWithChainID(strings.NewReader(string(ksBytes)), passphrase, chainid)
	if err != nil {
		return nil, ZeroAddr, err
	}

	return submitChainAuth, key.Address, err
}
