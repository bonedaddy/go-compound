package client

import (
	"io/ioutil"
	"net/http"
)


// Address is a compound contract address type
type Address string

func (a Address) String() string {
	return string(a)
}

const (
	// CompoundDAI is the address of the cDAI contract
	CompoundDAI Address = Address("0xf5dce57282a584d2746faf1593d3121fcac444dc")
)

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
