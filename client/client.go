package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"io/ioutil"

	"github.com/postables/go-compound/models"
)

// Client is a helper struct that gives access to both
// the account service and ctoken service
type Client struct {
	url    string
	client *http.Client
}

// NewClient is used to instantiate a new go-compound client
func NewClient(url string) *Client {
	return &Client{url: url, client: &http.Client{}}
}

// GetAccount is used to retrieve information on a single account
func (c *Client) GetAccount(address string) (*models.AccountResponse, error) {
	call := fmt.Sprintf("/account?addresses[]=%s", address)
	apiURL := fmt.Sprintf("%s/%s", c.url, call)
	request, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("content-type", "application/json")
	resp, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
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
func (c *Client) GetAccounts() (*models.AccountResponse, error) {
	// https://api.compound.finance/api/v2/account
	apiURL := fmt.Sprintf("%s/account", c.url)
	request, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("content-type", "application/json")
	resp, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response := &models.AccountResponse{}
	if err := json.Unmarshal(bodyBytes, response); err != nil {
		return nil, err
	}
	return response, nil
}
