package bchain

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ceth "github.com/postables/go-compound/bindings/ceth"
)

// BClient is an ethereum blockchain client
type BClient struct {
	auth   *bind.TransactOpts
	client *ethclient.Client
}

// NewBClient registers a new blockchain client
func NewBClient(auth *bind.TransactOpts, client *ethclient.Client) *BClient {
	return &BClient{auth: auth, client: client}
}

// BorrowETH is used to borrow eth from the cETH contract
func (bc *BClient) BorrowETH(ctx context.Context, address string, borrowAmount *big.Int) error {
	contract, err := ceth.NewBindings(common.HexToAddress(address), bc.client)
	if err != nil {
		return err
	}
	tx, err := contract.Borrow(bc.auth, borrowAmount)
	if err != nil {
		return err
	}
	rcpt, err := bind.WaitMined(ctx, bc.client, tx)
	if err != nil {
		return err
	}
	if rcpt.Status != 1 {
		return errors.New("tx receipt status is not 1, indicating a failure occurred")
	}
	return nil
}
