package models

import "errors"

// AccountResponse is a response to a
// https://api.compound.finance/api/v2/account?addresses[]= call
type AccountResponse struct {
	Accounts             []Account         `json:"accounts"`
	CloseFactor          float64           `json:"close_factor"`
	Error                interface{}       `json:"error"`
	LiquidationIncentive float64           `json:"liquidation_incentive"`
	PaginationSummary    PaginationSummary `json:"pagination_summary"`
	Request              Request           `json:"request"`
}

type Account struct {
	Address                   string      `json:"address"`
	BlockUpdated              interface{} `json:"block_updated"`
	Health                    Value       `json:"health"`
	Tokens                    []Token     `json:"tokens"`
	TotalBorrowValueInEth     Value       `json:"total_borrow_value_in_eth"`
	TotalCollateralValueInEth Value       `json:"total_collateral_value_in_eth"`
}

type PaginationSummary struct {
	PageNumber   int `json:"page_number"`
	PageSize     int `json:"page_size"`
	TotalEntries int `json:"total_entries"`
	TotalPages   int `json:"total_pages"`
}

type Request struct {
	Addresses           []string    `json:"addresses"`
	BlockNumber         int         `json:"block_number"`
	BlockTimestamp      int         `json:"block_timestamp"`
	MaxHealth           interface{} `json:"max_health"`
	MinBorrowValueInEth interface{} `json:"min_borrow_value_in_eth"`
	PageNumber          int         `json:"page_number"`
	PageSize            int         `json:"page_size"`
}

type Value struct {
	Value string `json:"value"`
}

// Token is an individual token
type Token struct {
	Address                       string      `json:"address"`
	BorrowBalanceUnderlying       Value       `json:"borrow_balance_underlying"`
	LifetimeBorrowInterestAccrued Value       `json:"lifetime_borrow_interest_accrued"`
	LifetimeSupplyInterestAccrued Value       `json:"lifetime_supply_interest_accrued"`
	SupplyBalanceUnderlying       Value       `json:"supply_balance_underlying"`
	Symbol                        interface{} `json:"symbol"`
}

// GetTokenByAddress is used to retrieve a token by its address from an AccountResponse type
func GetTokenByAddress(address string, resp *AccountResponse) (Token, error) {
	for _, token := range resp.Accounts[0].Tokens {
		if token.Address == address {
			return token, nil
		}
	}
	return Token{}, errors.New("token not found")
}
