package client

import (
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/postables/go-compound/models"
)

// Address is a compound contract address type
type Address string

func (a Address) String() string {
	return string(a)
}

// EthAddress returns a typed ethereum address
func (a Address) EthAddress() common.Address {
	return common.HexToAddress(a.String())
}

const (
	// CompoundBAT is the address of the cBAT contract
	CompoundBAT = Address("0x6c8c6b02e7b2be14d4fa6022dfd6d75921d90e4e")
	// CompoundDAI is the address of the cDAI contract
	CompoundDAI = Address("0xf5dce57282a584d2746faf1593d3121fcac444dc")
	// CompoundETH is the address of the cETH contract
	CompoundETH = Address("0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5")
	// CompoundREP is the address of the cREP contract
	CompoundREP = Address("0x158079ee67fce2f58472a96584a73c7ab9ac95c1")
	// CompoundUSDC is the address of the cUSDC contract
	CompoundUSDC = Address("0x39aa39c021dfbae8fac545936693ac917d5e7563")
	// CompoundWBTC is the address of the cWBTC contract
	CompoundWBTC = Address("0xc11b1268c1a384e55c48c2391d8d480264a3a7f4")
	// CompoundZRX is th eaddress of the cZRX contract
	CompoundZRX = Address("0xb3319f5d18bc0d84dd1b4825dcde5d5f7266d407")
)

var (
	// CompoundTokens is map containing the name, and address of all compound tokens
	CompoundTokens = map[string]Address{
		"cBAT":  CompoundBAT,
		"cDAI":  CompoundDAI,
		"cETH":  CompoundETH,
		"cREP":  CompoundREP,
		"cUSDC": CompoundUSDC,
		"cWBTC": CompoundWBTC,
		"cZRX":  CompoundZRX,
	}
)

// GetSupplyInterestEarned is used to retrieve the interest earned by supply a particular token
func (c *Client) GetSupplyInterestEarned(token Address, resp *models.AccountResponse) (string, error) {
	tkn, err := models.GetTokenByAddress(token.String(), resp)
	if err != nil {
		return "", err
	}
	return tkn.LifetimeSupplyInterestAccrued.Value, nil
}

// GetBorrowInterestedAccrued is used to retrieve the interest you owe for borrowing
func (c *Client) GetBorrowInterestedAccrued(token Address, resp *models.AccountResponse) (string, error) {
	tkn, err := models.GetTokenByAddress(token.String(), resp)
	if err != nil {
		return "", err
	}
	return tkn.LifetimeBorrowInterestAccrued.Value, nil
}

// sendRequest is used to send a request, and return the given body bytes
func (c *Client) sendRequest(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("content-type", "application/json")
	resp, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}
