package models

type AccountsResponse struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	Iban             string  `json:"iban"`
	Bban             string  `json:"bban"`
	Type             string  `json:"type"`
	AvailableBalance float64 `json:"available_balance"`
	Metadata         string  `json:"metadata"`
}
