# go-compound

> **WARNING: this code deals with money both by making blockchain calls and returning information that can be used to lose/gain money. Please use cautiously and dont come complaining if you end up losing money.**

`go-compound` is a Golang client, and library for interacting with the compound.finace API, and interacting with the compound contracts, allowing for the creation of bots trading with the Compound protocol.

# Contents

* `abi` contains json abi definitions for various compound smart contracts
* `bindigns` contains `abigen` generated golang bindings for the various abi's
* `client` contains a client library to build applications that use the compound.finance API and interact with the smart contracts
* `cmd` contains a small command-line client
* `models` contains Golang types for the various responses that the API gives. Currently it has types for `CTokenService` and `AccountService` responses.
* `pb` contains protobuf definitions for the compound APIs. Do not use
* `sampler` contains [sampler](https://github.com/sqshq/sampler) configurations to enable console based monitoring of your compound accounts

# Current Capabilities

## APIs

* Complete [AccountService](https://compound.finance/developers/api#AccountService) calls
* Complete [CTokenService](https://compound.finance/developers/api#CTokenService) calls

## Client Library

* Access to APIs via Golang programs, and not having to deal with raw http calls
* Watch account health, signalling and printing on different account health states
* Retrieve supply interest earned for a particular token
* Retrieve total supply interest earned
* Retrieve borrow interest owed for a particular token
* Borrow from any compound contract
* Get borrow rate for any compound contract
* Retrieve list of liquidatable addresses

## CLI

* Pretty print full  `AccountService::AccountResponse` information, suitable for piping to `jq`
* Retrieve account health
* Retrieve supply interest earned for a particular token
* Retrieve total supply interest earned
* Retrieve borrow interest owed for a particular token
* Retrieve a list of addresses that can be liquidated

## Monitoring

![](https://gateway.temporal.cloud/ipfs/QmZeLP8oN93H9DgCTp8G3C7cpPfw2xZ1rDeSFsLwa1ftTe)

* Enables monitoring your account with [sampler](https://github.com/sqshq/sampler).
* Default sampler config located in `sampler/sampler.yml` however you'll need to replace the addresses as needed.

# Long Term Goals

* Enable persisting retrieve data locally in a DB for fast lookups
* Enable report generation of your holdings
* Enable the `MarketHistoryService` API
* Enable graphing of `MarketHistoryService` metrics
* Enable arbitrage between uniswap and compound
  * Enable things like spotting interest rate arbitrages
  * Repayment arbitrages
  * ???

# Links

* [API Docs](https://compound.finance/developers/api)
* [Compound GitHub](https://github.com/compound-finance/)
* [Compound App](https://app.compound.finance/)
* [Compound Website](https://compound.finance/)
* [comproi](https://www.comproi.com/#)

# Support

If you like this application feel free to send me some ETH, or whatever (shit)coins you want to get rid of `0xc7459562777DDf3A1A7afefBE515E8479Bd3FDBD`.

# Misc Q&A's

* Q: on the compound token contracts, what's the different between `exchangeRateCurrent` and `exchangeRateStored`
* A: Current calls accrueInterest - stored returns the last computed value
