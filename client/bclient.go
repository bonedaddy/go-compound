package client

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	cbat "github.com/postables/go-compound/bindings/cbat"
	cdai "github.com/postables/go-compound/bindings/cdai"
	ceth "github.com/postables/go-compound/bindings/ceth"
	comptroller "github.com/postables/go-compound/bindings/comptroller"
	crep "github.com/postables/go-compound/bindings/crep"
	cusdc "github.com/postables/go-compound/bindings/cusdc"
	cwbtc "github.com/postables/go-compound/bindings/cusdc"
	czrx "github.com/postables/go-compound/bindings/cusdc"
)

// BClient is an ethereum blockchain client
type BClient struct {
	auth     *bind.TransactOpts
	client   *ethclient.Client
	readOnly bool
}

// NewBClient registers a new blockchain client
func NewBClient(auth *bind.TransactOpts, address string, readOnly bool) (*BClient, error) {
	client, err := ethclient.Dial(address)
	if err != nil {
		return nil, err
	}
	bc := &BClient{client: client, readOnly: readOnly}
	if !readOnly {
		bc.auth = auth
	}
	return bc, nil
}

// CanLiquidate is used to check whether or not the given address can be liquidated
func (bc *BClient) CanLiquidate(ctx context.Context, account common.Address) (bool, error) {
	contract, err := comptroller.NewBindings(Unitroller.EthAddress(), bc.client)
	if err != nil {
		return false, err
	}

	errCode, liquidity, shortfall, err := contract.GetAccountLiquidity(nil, account)
	if err != nil {
		return false, err
	}
	if errCode.Int64() != 0 {
		return false, errors.New("smart contract return a non 0 error code")
	}
	// liquidity at 0 means the account has no liquidity
	// shortfall should be 1, indicating they are at risk of being liquidated
	if liquidity.Int64() == 0 {
		// shortfall needs to be 1 or higher to be liquidatable
		if shortfall.Int64() == 0 {
			return false, errors.New("both shortfall and liquidity are 0")
		} else if shortfall.Int64() >= 1 {
			return true, nil
		}
	}
	if liquidity.Int64() > 0 {
		if shortfall.Int64() > 0 {
			return false, errors.New("both liquidity and shortfall are greater than 0")
		}
	}
	return false, nil
}

// Borrow is used to borrow a particular address
func (bc *BClient) Borrow(ctx context.Context, address Address, borrowAmount *big.Int) error {
	var (
		tx  *types.Transaction
		err error
	)
	switch address {
	case CompoundBAT:
		contract, err := cbat.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return err
		}
		tx, err = contract.Borrow(bc.auth, borrowAmount)
	case CompoundDAI:
		contract, err := cdai.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return err
		}
		tx, err = contract.Borrow(bc.auth, borrowAmount)
	case CompoundETH:
		contract, err := ceth.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return err
		}
		tx, err = contract.Borrow(bc.auth, borrowAmount)
	case CompoundREP:
		contract, err := crep.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return err
		}
		tx, err = contract.Borrow(bc.auth, borrowAmount)
	case CompoundUSDC:
		contract, err := cusdc.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return err
		}
		tx, err = contract.Borrow(bc.auth, borrowAmount)
	case CompoundWBTC:
		contract, err := cwbtc.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return err
		}
		tx, err = contract.Borrow(bc.auth, borrowAmount)
	case CompoundZRX:
		contract, err := czrx.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return err
		}
		tx, err = contract.Borrow(bc.auth, borrowAmount)
	}
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

// GetBorrowRate calls BorrowRatePerBlock to retrieve the current borrow interest rate
func (bc *BClient) GetBorrowRate(ctx context.Context, address Address) (*big.Int, error) {
	var (
		rate *big.Int
		err  error
	)
	switch address {
	case CompoundBAT:
		contract, err := cbat.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return nil, err
		}
		rate, err = contract.BorrowRatePerBlock(nil)
	case CompoundDAI:
		contract, err := cdai.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return nil, err
		}
		rate, err = contract.BorrowRatePerBlock(nil)
	case CompoundETH:
		contract, err := ceth.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return nil, err
		}
		rate, err = contract.BorrowRatePerBlock(nil)
	case CompoundREP:
		contract, err := crep.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return nil, err
		}
		rate, err = contract.BorrowRatePerBlock(nil)
	case CompoundUSDC:
		contract, err := cusdc.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return nil, err
		}
		rate, err = contract.BorrowRatePerBlock(nil)
	case CompoundWBTC:
		contract, err := cwbtc.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return nil, err
		}
		rate, err = contract.BorrowRatePerBlock(nil)
	case CompoundZRX:
		contract, err := czrx.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return nil, err
		}
		rate, err = contract.BorrowRatePerBlock(nil)
	}
	if err != nil {
		return nil, err
	}
	return rate, nil
}

// GetSupplyRate calls SupplyRatePerBlock to retrieve the current borrow interest rate
func (bc *BClient) GetSupplyRate(ctx context.Context, address Address) (*big.Int, error) {
	var (
		rate *big.Int
		err  error
	)
	switch address {
	case CompoundBAT:
		contract, err := cbat.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return nil, err
		}
		rate, err = contract.SupplyRatePerBlock(nil)
	case CompoundDAI:
		contract, err := cdai.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return nil, err
		}
		rate, err = contract.SupplyRatePerBlock(nil)
	case CompoundETH:
		contract, err := ceth.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return nil, err
		}
		rate, err = contract.SupplyRatePerBlock(nil)
	case CompoundREP:
		contract, err := crep.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return nil, err
		}
		rate, err = contract.SupplyRatePerBlock(nil)
	case CompoundUSDC:
		contract, err := cusdc.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return nil, err
		}
		rate, err = contract.SupplyRatePerBlock(nil)
	case CompoundWBTC:
		contract, err := cwbtc.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return nil, err
		}
		rate, err = contract.SupplyRatePerBlock(nil)
	case CompoundZRX:
		contract, err := czrx.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return nil, err
		}
		rate, err = contract.SupplyRatePerBlock(nil)
	}
	if err != nil {
		return nil, err
	}
	return rate, nil
}

// LiquidateOpts is used to provide input parameters to
// LiquidateBorrow functinos
type LiquidateOpts struct {
	Borrower         common.Address
	RepayAmount      *big.Int
	CTokenCollateral Address
}

// GetLiqd is used to liquidate a borrower
func (bc *BClient) GetLiqd(ctx context.Context, borrowToken Address, opts LiquidateOpts) error {
	var (
		tx  *types.Transaction
		err error
	)
	switch borrowToken {
	case CompoundBAT:
		contract, err := cbat.NewBindings(borrowToken.EthAddress(), bc.client)
		if err != nil {
			return err
		}
		tx, err = contract.LiquidateBorrow(bc.auth, opts.Borrower, opts.RepayAmount, opts.CTokenCollateral.EthAddress())
	case CompoundDAI:
		contract, err := cdai.NewBindings(borrowToken.EthAddress(), bc.client)
		if err != nil {
			return err
		}
		tx, err = contract.LiquidateBorrow(bc.auth, opts.Borrower, opts.RepayAmount, opts.CTokenCollateral.EthAddress())
	case CompoundETH:
		contract, err := ceth.NewBindings(borrowToken.EthAddress(), bc.client)
		if err != nil {
			return err
		}
		tx, err = contract.LiquidateBorrow(bc.auth, opts.Borrower, opts.CTokenCollateral.EthAddress())
	case CompoundREP:
		contract, err := crep.NewBindings(borrowToken.EthAddress(), bc.client)
		if err != nil {
			return err
		}
		tx, err = contract.LiquidateBorrow(bc.auth, opts.Borrower, opts.RepayAmount, opts.CTokenCollateral.EthAddress())
	case CompoundUSDC:
		contract, err := cusdc.NewBindings(borrowToken.EthAddress(), bc.client)
		if err != nil {
			return err
		}
		tx, err = contract.LiquidateBorrow(bc.auth, opts.Borrower, opts.RepayAmount, opts.CTokenCollateral.EthAddress())
	case CompoundWBTC:
		contract, err := cwbtc.NewBindings(borrowToken.EthAddress(), bc.client)
		if err != nil {
			return err
		}
		tx, err = contract.LiquidateBorrow(bc.auth, opts.Borrower, opts.RepayAmount, opts.CTokenCollateral.EthAddress())
	case CompoundZRX:
		contract, err := czrx.NewBindings(borrowToken.EthAddress(), bc.client)
		if err != nil {
			return err
		}
		tx, err = contract.LiquidateBorrow(bc.auth, opts.Borrower, opts.RepayAmount, opts.CTokenCollateral.EthAddress())
	}
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
