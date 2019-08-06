package client

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// provides account monitoring utilities

// WatchHealth is a helper function used to watch account health and send signals on different states.
// the riskChan is a channel used to signal when an account health is at 1.0 or lower, which means it is at risk of liquidation
// warnChan is a channel used to signal when an account is nearing liquidation risk, and has a health of 1.2 or lower
func (c *Client) WatchHealth(ctx context.Context, address string, riskChan, warnChan chan float64) error {
	var (
		ticker   = time.NewTicker(time.Second * 15)
		errCount int
	)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			if errCount > 3 {
				return errors.New("too many errors")
			}
			resp, err := c.GetAccount(address)
			if err != nil {
				fmt.Println("got error", err.Error())
				errCount++
				continue
			}
			if len(resp.Accounts) == 0 {
				fmt.Println("unexpected error")
				errCount++
				continue
			}
			health, err := strconv.ParseFloat(resp.Accounts[0].Health.Value, 64)
			if err != nil {
				return err
			}
			if health <= 1.0 {
				riskChan <- health
				fmt.Println("account at risk of liquidation")
			} else if health <= 1.2 {
				warnChan <- health
				fmt.Println("account nearing liquidation risk")
			}
		}
	}
}
