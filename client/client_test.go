package client

import (
	"testing"
)

var (
	api     = "https://api.compound.finance/api/v2"
	account = "0xe18999d3f7e1a84e35bf9c15699b79e46eb44704"
	cdai    = "0xf5dce57282a584d2746faf1593d3121fcac444dc"
)

func Test_Client(t *testing.T) {
	client := NewClient(api)
	if _, err := client.GetAccount(account); err != nil {
		t.Fatal(err)
	}
	if _, err := client.GetAccounts(); err != nil {
		t.Fatal(err)
	}
	if _, err := client.GetCToken(cdai); err != nil {
		t.Fatal(err)
	}
	if _, err := client.GetCTokens(); err != nil {
		t.Fatal(err)
	}
}
