package client

import (
	"testing"
)

var (
	api     = "https://api.compound.finance/api/v2"
	account = "0xe18999d3f7e1a84e35bf9c15699b79e46eb44704"
)

func Test_Client(t *testing.T) {
	client := NewClient(api)
	if _, err := client.GetAccount(account); err != nil {
		t.Fatal(err)
	}
	client.GetAccounts()
}
