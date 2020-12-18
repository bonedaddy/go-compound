package client

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"math/big"

	cbat "github.com/bonedaddy/go-compound/v2/bindings/cbat"
	cdai "github.com/bonedaddy/go-compound/v2/bindings/cdai"
	ceth "github.com/bonedaddy/go-compound/v2/bindings/ceth"
	comptroller "github.com/bonedaddy/go-compound/v2/bindings/comptroller"
	crep "github.com/bonedaddy/go-compound/v2/bindings/crep"
	csai "github.com/bonedaddy/go-compound/v2/bindings/csai"
	cusdc "github.com/bonedaddy/go-compound/v2/bindings/cusdc"
	cwbtc "github.com/bonedaddy/go-compound/v2/bindings/cusdc"
	czrx "github.com/bonedaddy/go-compound/v2/bindings/cusdc"
	"github.com/bonedaddy/go-compound/v2/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// BClient is an ethereum blockchain client
type BClient struct {
	auth   *bind.TransactOpts
	client *ethclient.Client
}

// ConfigToOpts structs the optios needed for a bclient
func ConfigToOpts(cfg *config.Config) (*bind.TransactOpts, *ethclient.Client, error) {
	client, err := ethclient.Dial(cfg.Blockchain.Endpoint)
	if err != nil {
		return nil, nil, err
	}
	keyFileBytes, err := ioutil.ReadFile(cfg.Blockchain.KeyFile)
	if err != nil {
		return nil, nil, err
	}
	auth, err := bind.NewTransactor(bytes.NewReader(keyFileBytes), cfg.Blockchain.KeyPass)
	if err != nil {
		client.Close()
	}
	return auth, client, err
}

// NewBClient registers a new blockchain client
func NewBClient(auth *bind.TransactOpts, client *ethclient.Client) *BClient {
	return &BClient{auth: auth, client: client}
}

// GetPrice returns the price/exchange rate of the cToken (WIP)
func (bc *BClient) GetPrice(ctx context.Context, address Address) (*big.Int, error) {
	/*
		no. Firstly, a cToken is worth significantly less than the underlying asset. This conversion is given by ‘exchangeRateCurrent’. To get the dollar value, you will need to get exchange rates from the oracle which get you to eth value. Then get the exchange rate to usdc to get dollar value.
		Here is the oracle address: 0xddc46a3b076aec7ab3fc37420a8edd2959764ec4
		You need to fetch an exchange rate for your cToken, multiply by the underlying quantity and then divide the total by 1e18 to get eth value
		If you want to get to $ value, just do the opposite with usdc. Fetch cUSDC exchange rate from oracle, multiply previous product by 1e18, then divide this by the usdc exchange rate. You should get usdc value
	*/
	switch address {
	case CompoundBAT:
		contract, err := cbat.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return nil, err
		}
		return contract.ExchangeRateStored(nil)
	case CompoundDAI:
		contract, err := cdai.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return nil, err
		}
		return contract.ExchangeRateStored(nil)
	case CompoundSAI:
		contract, err := csai.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return nil, err
		}
		return contract.ExchangeRateStored(nil)
	case CompoundETH:
		contract, err := ceth.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return nil, err
		}
		return contract.ExchangeRateStored(nil)
	case CompoundREP:
		contract, err := crep.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return nil, err
		}
		return contract.ExchangeRateStored(nil)
	case CompoundUSDC:
		contract, err := cusdc.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return nil, err
		}
		return contract.ExchangeRateStored(nil)
	case CompoundWBTC:
		contract, err := cwbtc.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return nil, err
		}
		return contract.ExchangeRateStored(nil)
	case CompoundZRX:
		contract, err := czrx.NewBindings(address.EthAddress(), bc.client)
		if err != nil {
			return nil, err
		}
		return contract.ExchangeRateStored(nil)
	default:
		return nil, errors.New("invalid contract type")
	}
}

// CanLiquidate is used to check whether or not the given address can be liquidated
func (bc *BClient) CanLiquidate(ctx context.Context, account common.Address) (bool, error) {
	contract, err := comptroller.NewBindings(Comptroller.EthAddress(), bc.client)
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
	case CompoundSAI:
		contract, err := csai.NewBindings(address.EthAddress(), bc.client)
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
	case CompoundSAI:
		contract, err := csai.NewBindings(address.EthAddress(), bc.client)
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
	case CompoundSAI:
		contract, err := csai.NewBindings(borrowToken.EthAddress(), bc.client)
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

// GetAccountSnapshot returns a snapshot of the account state
func (bc *BClient) GetAccountSnapshot(ctx context.Context, address Address) {
	/*
		* https://github.com/compound-finance/compound-protocol/blob/f37f7bf8b6e1fb39239e43a4fb12fd15cf8b7c69/contracts/CErc20Delegator.sol#L226
		* https://github.com/compound-finance/compound-protocol/blob/f37f7bf8b6e1fb39239e43a4fb12fd15cf8b7c69/contracts/CToken.sol#L203
		(possible error, token balance, borrow balance, exchange rate mantissa)
	*/
}
