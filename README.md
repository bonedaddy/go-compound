# go-compound

`go-compound` is a Golang client, and library for interacting with the compound.finace API, and eventually providing means to execute on-chain transactions. It comes with a small CLI providing access to convience functions like retrieving account health, etc...

Eventually it will also ship with a [sampler](https://github.com/sqshq/sampler) config for dank console based metrics.

The `pb` folder contains the protobufs the compound API which was sent to me by a compound.finance developer through discord. At the moment there appears ot be some issue with using the API and unmarshaling into protobus so don't use them for now. Instead use the `models` folder which was done using https://mholt.github.io/json-to-go/.

# Links

* [API Docs](https://compound.finance/developers/api)
* [Compound GitHub](https://github.com/compound-finance/)
* [Compound App](https://app.compound.finance/)
* [Compound Website](https://compound.finance/)
* [comproi](https://www.comproi.com/#)