package models

// CTokenResponse is a response to a
// https://api.compound.finance/api/v2/ctoken?addresses[]= call
type CTokenResponse struct {
	CToken []struct {
		BorrowRate struct {
			Value string `json:"value"`
		} `json:"borrow_rate"`
		Cash struct {
			Value string `json:"value"`
		} `json:"cash"`
		CollateralFactor struct {
			Value string `json:"value"`
		} `json:"collateral_factor"`
		ExchangeRate struct {
			Value string `json:"value"`
		} `json:"exchange_rate"`
		InterestRateModelAddress string `json:"interest_rate_model_address"`
		Name                     string `json:"name"`
		NumberOfBorrowers        int    `json:"number_of_borrowers"`
		NumberOfSuppliers        int    `json:"number_of_suppliers"`
		Reserves                 struct {
			Value string `json:"value"`
		} `json:"reserves"`
		SupplyRate struct {
			Value string `json:"value"`
		} `json:"supply_rate"`
		Symbol       string `json:"symbol"`
		TokenAddress string `json:"token_address"`
		TotalBorrows struct {
			Value string `json:"value"`
		} `json:"total_borrows"`
		TotalSupply struct {
			Value string `json:"value"`
		} `json:"total_supply"`
		UnderlyingAddress string `json:"underlying_address"`
		UnderlyingName    string `json:"underlying_name"`
		UnderlyingPrice   struct {
			Value string `json:"value"`
		} `json:"underlying_price"`
		UnderlyingSymbol string `json:"underlying_symbol"`
	} `json:"cToken"`
	Error   interface{} `json:"error"`
	Request struct {
		Addresses      []string `json:"addresses"`
		BlockNumber    int      `json:"block_number"`
		BlockTimestamp int      `json:"block_timestamp"`
	} `json:"request"`
}
