package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"os"

	"strconv"

	"context"

	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/postables/go-compound/client"
	"github.com/urfave/cli"
)

var (
	url    string
	ethURL string
)

func main() {
	app := newApp()
	if err := app.Run(os.Args); err != nil {
		log.Println("error running application: ", err.Error())
		os.Exit(1)
	}
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "go-compound cli"
	app.Authors = []cli.Author{
		{
			Name:  "Alex Trottier (postables)",
			Email: "postables@rtradetechnologies.com",
		},
	}
	app.Flags = loadFlags()
	app.Commands = loadCommands()
	return app
}

func loadFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "api.url, au",
			Usage:       "the compound api url",
			Value:       "https://api.compound.finance/api/v2",
			Destination: &url,
		},
		cli.StringFlag{
			Name:        "eth.url, eu",
			Usage:       "ethereum rpc endpoint",
			Value:       "http://127.0.0.1:8545",
			Destination: &ethURL,
		},
	}
}
func loadCommands() cli.Commands {
	commands := loadAccountCommands()
	commands = append(commands, loadBlockchainCommands()...)
	return commands
}

func loadBlockchainCommands() cli.Commands {
	return cli.Commands{
		cli.Command{
			Name:        "blockchain",
			Usage:       "interact with compound contracts",
			Description: "used to interact with the actual compound contracts",
			Subcommands: cli.Commands{
				cli.Command{
					Name:  "can-liquidate",
					Usage: "check if an account is liquidatable",
					Action: func(c *cli.Context) error {
						if c.String("eth.address") == "" {
							return errors.New("eth.address flag is empty")
						}
						bc, err := client.NewBClient(nil, ethURL, true)
						if err != nil {
							return err
						}
						canLiq, err := bc.CanLiquidate(
							context.Background(),
							common.HexToAddress(c.String("eth.address")),
						)
						if err != nil {
							return err
						}
						fmt.Println(canLiq)
						return nil
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "eth.address, ea",
							Usage: "eth address to lookup",
						},
					},
				},
				cli.Command{
					Name:  "exchange-rate",
					Usage: "get stored exchange rate",
					Action: func(c *cli.Context) error {
						if c.String("asset.name") == "" {
							return errors.New("asset.name flag is empty")
						}
						bc, err := client.NewBClient(nil, ethURL, true)
						if err != nil {
							return err
						}
						resp, err := bc.GetExchangeRate(
							context.Background(),
							client.CompoundTokens[c.String("asset.name")],
						)
						if err != nil {
							return err
						}
						fmt.Println(resp.Int64())
						exchangeRate := new(big.Float)
						exchangeRate.SetString(resp.String())
						fmt.Println(exchangeRate.String())
						ethValue := new(big.Float).Quo(exchangeRate, big.NewFloat(math.Pow10(18)))
						fmt.Println(ethValue)
						return nil
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "asset.name, an",
							Usage: "asset name, ie cDAI",
						},
					},
				},
				cli.Command{
					Name:  "borrow-rate",
					Usage: "get asset borrow rate",
					Action: func(c *cli.Context) error {
						if c.String("asset.name") == "" {
							return errors.New("asset.name flag is empty")
						}
						bc, err := client.NewBClient(nil, ethURL, true)
						if err != nil {
							return err
						}
						resp, err := bc.GetBorrowRate(
							context.Background(),
							client.CompoundTokens[c.String("asset.name")],
						)
						if err != nil {
							return err
						}
						yearRate := new(big.Int).Mul(resp, big.NewInt(2102400))
						apr := new(big.Float)
						apr.SetString(yearRate.String())
						ethValue := new(big.Float).Quo(apr, big.NewFloat(math.Pow10(18)))
						fmt.Println(ethValue)
						return nil
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "asset.name, an",
							Usage: "asset name, ie cDAI",
						},
					},
				},
				cli.Command{
					Name:  "supply-rate",
					Usage: "get asset supply rate",
					Action: func(c *cli.Context) error {
						if c.String("asset.name") == "" {
							return errors.New("asset.name flag is empty")
						}
						bc, err := client.NewBClient(nil, ethURL, true)
						if err != nil {
							return err
						}
						resp, err := bc.GetSupplyRate(
							context.Background(),
							client.CompoundTokens[c.String("asset.name")],
						)
						if err != nil {
							return err
						}
						yearRate := new(big.Int).Mul(resp, big.NewInt(2102400))
						apr := new(big.Float)
						apr.SetString(yearRate.String())
						ethValue := new(big.Float).Quo(apr, big.NewFloat(math.Pow10(18)))
						fmt.Println(ethValue)
						return nil
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "asset.name, an",
							Usage: "asset name, ie cDAI",
						},
					},
				},
			},
		},
	}
}

func loadAccountCommands() cli.Commands {
	return cli.Commands{
		cli.Command{
			Name:        "account",
			Usage:       "account management commands",
			Description: "calling account by itself, intended for piping to jq command",
			Action: func(c *cli.Context) error {
				if c.String("eth.address") == "" {
					return errors.New("eth.address flag is empty")
				}
				cl := client.NewClient(url)
				resp, err := cl.GetAccount(c.String("eth.address"))
				if err != nil {
					return err
				}
				if len(resp.Accounts) == 0 {
					return errors.New("an unexpected error occurred")
				}
				respBytes, err := json.Marshal(resp)
				if err != nil {
					return err
				}
				var pretty bytes.Buffer
				if err = json.Indent(&pretty, respBytes, "", "\t"); err != nil {
					return err
				}
				if err := json.Unmarshal(pretty.Bytes(), resp); err != nil {
					return err
				}
				fmt.Println(string(pretty.Bytes()))
				return nil
			},
			Aliases: []string{"acct"},
			Subcommands: cli.Commands{
				cli.Command{
					Name:  "collateral-value",
					Usage: "get account collateral value in eth",
					Action: func(c *cli.Context) error {
						if c.String("eth.address") == "" {
							return errors.New("eth.address flag is empty")
						}
						cl := client.NewClient(url)
						value, err := cl.GetTotalCollateralValueInEth(c.String("eth.address"))
						if err != nil {
							return err
						}
						fmt.Println("collateral value: ", value)
						return nil
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "eth.address",
							Usage: "the address to lookup",
						},
					},
				},
				cli.Command{
					Name:  "borrow-value",
					Usage: "get account borrow value in eth",
					Action: func(c *cli.Context) error {
						if c.String("eth.address") == "" {
							return errors.New("eth.address flag is empty")
						}
						cl := client.NewClient(url)
						value, err := cl.GetTotalBorrowValueInEth(c.String("eth.address"))
						if err != nil {
							return err
						}
						fmt.Println("borrow value: ", value)
						return nil
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "eth.address",
							Usage: "the address to lookup",
						},
					},
				},
				cli.Command{
					Name:  "health",
					Usage: "get account health",
					Action: func(c *cli.Context) error {
						if c.String("eth.address") == "" {
							return errors.New("eth.address flag is empty")
						}
						cl := client.NewClient(url)
						resp, err := cl.GetAccount(c.String("eth.address"))
						if err != nil {
							return err
						}
						if len(resp.Accounts) == 0 {
							return errors.New("an unexpected error occurred")
						}
						health, err := strconv.ParseFloat(resp.Accounts[0].Health.Value, 64)
						if err != nil {
							return err
						}
						fmt.Println("account health: ", health)
						return nil
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "eth.address",
							Usage: "the address to lookup",
						},
					},
				},
				cli.Command{
					Name:  "supply-interest",
					Usage: "get supply-interest earned",
					Action: func(c *cli.Context) error {
						if !c.Bool("total") {
							if c.String("eth.address") == "" {
								return errors.New("eth.address flag is empty")
							}
							if c.String("token.name") == "" {
								return errors.New("token.name flag is empty")
							}
						}
						cl := client.NewClient(url)
						resp, err := cl.GetAccount(c.String("eth.address"))
						if err != nil {
							return err
						}
						if len(resp.Accounts) == 0 {
							return errors.New("an unexpected error occurred")
						}
						var interest float64
						if !c.Bool("total") {
							interest, err = cl.GetSupplyInterestEarned(
								client.CompoundTokens[c.String("token.name")],
								resp,
							)
						} else {
							interest, err = cl.GetTotalSupplyInterestedEarned(resp)
						}
						if err != nil {
							return err
						}
						fmt.Println("interest earned: ", interest)
						return nil
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "eth.address, ea",
							Usage: "the address to lookup",
						},
						cli.StringFlag{
							Name:  "token.name, tn",
							Usage: "the compound token being supplied",
						},
						cli.BoolFlag{
							Name:  "total",
							Usage: "whether or not to collect total interest",
						},
					},
				},
				cli.Command{
					Name:  "borrow-interest",
					Usage: "get borrow-interest owed",
					Action: func(c *cli.Context) error {
						if c.String("eth.address") == "" {
							return errors.New("eth.address flag is empty")
						}
						if c.String("token.name") == "" {
							return errors.New("token.name flag is empty")
						}
						cl := client.NewClient(url)
						resp, err := cl.GetAccount(c.String("eth.address"))
						if err != nil {
							return err
						}
						if len(resp.Accounts) == 0 {
							return errors.New("an unexpected error occurred")
						}
						interest, err := cl.GetBorrowInterestedAccrued(
							client.CompoundTokens[c.String("token.name")],
							resp,
						)
						if err != nil {
							return err
						}
						fmt.Println("interest owed: ", interest)
						return nil
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "eth.address, ea",
							Usage: "the address to lookup",
						},
						cli.StringFlag{
							Name:  "token.name, tn",
							Usage: "the compound token being supplied",
						},
					},
				},
				cli.Command{
					Name:    "liquidatable",
					Aliases: []string{"liqqable"},
					Usage:   "get all liquidatable accounts",
					Action: func(c *cli.Context) error {
						cl := client.NewClient(url)
						accts, err := cl.GetLiquidatableAccounts()
						if err != nil {
							return err
						}
						for k, v := range accts {
							fmt.Printf("account %s health %v\n", k, v)
						}
						return nil
					},
				},
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "eth.address",
					Usage: "the address to lookup",
				},
			},
		},
	}
}
