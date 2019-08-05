# go-compound

Golang API for compound.finace

To get the API to respond using protobufs, we need to use `Content-Type` of `application/x-protobuf`.

The `pb` folder contains the protobufs the compound API which was sent to me by a compound.finance developer through discord. I have made a couple slight modifications, mostly to to leverage gogoprotobuf for faster unmarshaling/marshaling.

# Links

* [API Docs](https://compound.finance/developers/api)