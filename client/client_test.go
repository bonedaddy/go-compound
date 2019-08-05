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
	if _, err := client.GetAccounts(); err != nil {
		t.Fatal(err)
	}
	if _, err := client.GetCToken(CompoundDAI.String()); err != nil {
		t.Fatal(err)
	}
	if _, err := client.GetCTokens(); err != nil {
		t.Fatal(err)
	}
}


func Test_CompoundAddresses(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct{
		name string
		args args
		wantToken Address
	}{
		{"cBAT", args{"cBAT"}, CompoundBAT},
		{"cDAI", args{"cDAI"}, CompoundDAI},
		{"cETH", args{"cETH"}, CompoundETH},
		{"cREP", args{"cREP"}, CompoundREP},
		{"cUSDC", args{"cUSDC"}, CompoundUSDC},
		{"cWBTC", args{"cWBTC"}, CompoundWBTC},
		{"cZRX", args{"cZRX"}, CompoundZRX},
	}

	for _, tt := range tests {
		if tok := CompoundTokens[tt.args.token]; tok != tt.wantToken {
			t.Fatal("bad token type")
		}
	}
}