# go-compound

> **WARNING: this code deals with money both by making blockchain calls and returning information that can be used to lose/gain money. Please use cautiously and dont come complaining if you end up losing money.**

`go-compound` is a Golang client, and library for interacting with the compound.finace API, and eventually providing means to execute on-chain transactions. It comes with a small CLI providing access to convience functions like retrieving account health, etc...

# Contents

* `client` contains a client library to build applications that use the compound.finance API
* `cmd` contains a small command-line client
* `models` contains Golang types for the various responses that the API gives. Currently it has types for `CTokenService` and `AccountService` responses.
* `pb` contains protobuf definitions for the compound APIs. Do not use

# Current Capabilities

## APIs

* Complete [AccountService](https://compound.finance/developers/api#AccountService) calls
* Complete [CTokenService](https://compound.finance/developers/api#CTokenService) calls

## Client Library

* Access to APIs via Golang programs
* Watch account health, printing if account is at liquidation risk, or nearing liquidation risk.
* Helper functions to rerieve supply interested, and borrow interest

## CLI

* Pretty print full  `AccountService::AccountResponse` information, suitable for piping to `jq`
* Retrieve account health
* Retrieve supply interest earned for a particular token
* Retrieve borrow interest owed for a particular token

# Long Term Goals

* Include a [sampler](https://github.com/sqshq/sampler) config for dank console based metrics.
* Enable persisting retrieve data locally in a DB for fast lookups
* Enable report generation of your holdings
* Enable the `MarketHistoryService` API
* Enable graphing of `MarketHistoryService` metrics

# Links

* [API Docs](https://compound.finance/developers/api)
* [Compound GitHub](https://github.com/compound-finance/)
* [Compound App](https://app.compound.finance/)
* [Compound Website](https://compound.finance/)
* [comproi](https://www.comproi.com/#)

# Support

If you like this application feel free to send me some ETH, or whatever (shit)coins you want to get rid of `0xc7459562777DDf3A1A7afefBE515E8479Bd3FDBD`.