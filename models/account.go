package models

// AccountResponse is a response to a 
// https://api.compound.finance/api/v2/account?addresses[]= call
type AccountResponse struct {
	Accounts []struct {
		Address      string      `json:"address"`
		BlockUpdated interface{} `json:"block_updated"`
		Health       struct {
			Value string `json:"value"`
		} `json:"health"`
		Tokens []struct {
			Address                 string `json:"address"`
			BorrowBalanceUnderlying struct {
				Value string `json:"value"`
			} `json:"borrow_balance_underlying"`
			LifetimeBorrowInterestAccrued struct {
				Value string `json:"value"`
			} `json:"lifetime_borrow_interest_accrued"`
			LifetimeSupplyInterestAccrued struct {
				Value string `json:"value"`
			} `json:"lifetime_supply_interest_accrued"`
			SupplyBalanceUnderlying struct {
				Value string `json:"value"`
			} `json:"supply_balance_underlying"`
			Symbol interface{} `json:"symbol"`
		} `json:"tokens"`
		TotalBorrowValueInEth struct {
			Value string `json:"value"`
		} `json:"total_borrow_value_in_eth"`
		TotalCollateralValueInEth struct {
			Value string `json:"value"`
		} `json:"total_collateral_value_in_eth"`
	} `json:"accounts"`
	CloseFactor          float64     `json:"close_factor"`
	Error                interface{} `json:"error"`
	LiquidationIncentive float64     `json:"liquidation_incentive"`
	PaginationSummary    struct {
		PageNumber   int `json:"page_number"`
		PageSize     int `json:"page_size"`
		TotalEntries int `json:"total_entries"`
		TotalPages   int `json:"total_pages"`
	} `json:"pagination_summary"`
	Request struct {
		Addresses           []string    `json:"addresses"`
		BlockNumber         int         `json:"block_number"`
		BlockTimestamp      int         `json:"block_timestamp"`
		MaxHealth           interface{} `json:"max_health"`
		MinBorrowValueInEth interface{} `json:"min_borrow_value_in_eth"`
		PageNumber          int         `json:"page_number"`
		PageSize            int         `json:"page_size"`
	} `json:"request"`
}
