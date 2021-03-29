package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"strconv"

	"github.com/bonedaddy/go-compound/v2/models"
)

// Client is used to interact with the compound.finance api
type Client struct {
	url    string
	client *http.Client
}

// NewClient is used to instantiate a new go-compound client
func NewClient(url string) *Client {
	return &Client{url: url, client: &http.Client{}}
}

// GetTotalCollateralValueInEth is used to retrieve the total collateral value
// in eth that is owned by this account
func (c *Client) GetTotalCollateralValueInEth(address string) (float64, error) {
	resp, err := c.GetAccount(address)
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(resp.Accounts[0].TotalCollateralValueInEth.Value, 64)
}

// GetTotalBorrowValueInEth is used to retrieve the total collateral value
// in eth that is owned by this account
func (c *Client) GetTotalBorrowValueInEth(address string) (float64, error) {
	resp, err := c.GetAccount(address)
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(resp.Accounts[0].TotalBorrowValueInEth.Value, 64)
}

// GetAccount is used to retrieve information on a single account
func (c *Client) GetAccount(address string) (*models.AccountResponse, error) {
	call := fmt.Sprintf("/account?addresses[]=%s", address)
	apiURL := fmt.Sprintf("%s/%s", c.url, call)
	bodyBytes, err := c.sendRequest(apiURL)
	if err != nil {
		return nil, err
	}
	response := &models.AccountResponse{}
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetAccounts is used to retrieve information on many accounts
func (c *Client) GetAccounts(pageSize, pageNum string) (*models.AccountResponse, error) {
	// https://api.compound.finance/api/v2/account
	if pageSize == "" {
		pageSize = "10"
	}
	if pageNum == "" {
		pageNum = "0"
	}
	apiURL := fmt.Sprintf("%s/account?page_size=%s&page_number=%s", c.url, pageSize, pageNum)
	bodyBytes, err := c.sendRequest(apiURL)
	if err != nil {
		return nil, err
	}
	response := &models.AccountResponse{}
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetCToken is used to get information on a particular ctoken
func (c *Client) GetCToken(address string) (*models.CTokenResponse, error) {
	call := fmt.Sprintf("/ctoken?addresses[]=%s", address)
	apiURL := fmt.Sprintf("%s/%s", c.url, call)
	bodyBytes, err := c.sendRequest(apiURL)
	if err != nil {
		return nil, err
	}
	response := &models.CTokenResponse{}
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetCTokens is used to retrieve information on many ctokens
func (c *Client) GetCTokens() (*models.CTokenResponse, error) {
	apiURL := fmt.Sprintf("%s/ctoken", c.url)
	bodyBytes, err := c.sendRequest(apiURL)
	if err != nil {
		return nil, err
	}
	response := &models.CTokenResponse{}
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}
	return response, nil
}
