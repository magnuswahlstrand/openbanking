package seb

type Account struct {
	Accounts []struct {
		ResourceID             string    `json:"resourceId"`
		Iban                   string    `json:"iban"`
		Bban                   string    `json:"bban"`
		Currency               string    `json:"currency"`
		OwnerName              string    `json:"ownerName"`
		Balances               []Balance `json:"balances"`
		CreditLine             string    `json:"creditLine"`
		Product                string    `json:"product"`
		Name                   string    `json:"name"`
		Status                 string    `json:"status"`
		StatusDate             string    `json:"statusDate"`
		Bic                    string    `json:"bic"`
		BicAddress             string    `json:"bicAddress"`
		AccountInterest        string    `json:"accountInterest"`
		CardLinkedToTheAccount bool      `json:"cardLinkedToTheAccount"`
		PaymentService         bool      `json:"paymentService"`
		BankgiroNumber         string    `json:"bankgiroNumber"`
		// AccountOwners          struct {
		// 	Name string `json:"name"`
		// } `json:"accountOwners"`
		// Interests struct {
		// 	Posted []struct {
		// 		InterestCapitalizationPostingDate string `json:"interestCapitalizationPostingDate"`
		// 		TransactionAccountID              string `json:"transactionAccountId"`
		// 		InterestCapitalizationAccountID   string `json:"interestCapitalizationAccountId"`
		// 		InterestType                      string `json:"interestType"`
		// 		InterestCapitalizationAmount      string `json:"interestCapitalizationAmount"`
		// 		PdAccountability                  string `json:"pdAccountability"`
		// 	} `json:"posted"`
		// 	Accrued []struct {
		// 		InterestDate                           string `json:"interestDate"`
		// 		AccruedDebitInterestAmount             string `json:"accruedDebitInterestAmount"`
		// 		AccruedCreditInterestAmount            string `json:"accruedCreditInterestAmount"`
		// 		AccruedPenaltyInterestAmount           string `json:"accruedPenaltyInterestAmount"`
		// 		AccruedDebitInterestAmountAdjusted     string `json:"accruedDebitInterestAmountAdjusted"`
		// 		AccruedCreditInterestAmountAdjusted    string `json:"accruedCreditInterestAmountAdjusted"`
		// 		AccruedPenaltyInterestAmountAdjusted   string `json:"accruedPenaltyInterestAmountAdjusted"`
		// 		AccruedPenaltyInterestAmountLP         string `json:"accruedPenaltyInterestAmountLP"`
		// 		AccruedPenaltyInterestAmountAdjustedLP string `json:"accruedPenaltyInterestAmountAdjustedLP"`
		// 		PdAccountability                       string `json:"pdAccountability"`
		// 		AccruedCreditInterestIndicator         string `json:"accruedCreditInterestIndicator"`
		// 		AccruedDebitInterestIndicator          string `json:"accruedDebitInterestIndicator"`
		// 	} `json:"accrued"`
		// } `json:"interests"`
		// InterestConditions struct {
		// 	Interests []struct {
		// 		InterestType     string `json:"interestType"`
		// 		InterestRateType string `json:"interestRateType"`
		// 		CalculationBase  string `json:"calculationBase"`
		// 		PdAccountability string `json:"pdAccountability"`
		// 		Tiers            []struct {
		// 			UpperBalance                    int    `json:"upperBalance"`
		// 			EffectiveDate                   string `json:"effectiveDate"`
		// 			AbsoluteRate                    string `json:"absoluteRate"`
		// 			ReferenceRateBaseRate           string `json:"referenceRateBaseRate"`
		// 			ReferenceRateMarginRate         string `json:"referenceRateMarginRate"`
		// 			EffectedInterestRate            string `json:"effectedInterestRate"`
		// 			ReferenceRateType               string `json:"referenceRateType"`
		// 			InterestCapitalizationFrequence int    `json:"interestCapitalizationFrequence"`
		// 			DayCountConvention              string `json:"dayCountConvention"`
		// 			PegAmount                       string `json:"pegAmount"`
		// 		} `json:"tiers"`
		// 	} `json:"interests"`
		// } `json:"interestConditions"`
		// Links struct {
		// 	Transactions struct {
		// 		Href string `json:"href"`
		// 	} `json:"transactions"`
		// } `json:"_links"`
	} `json:"accounts"`
}

type Balance struct {
	BalanceType         string `json:"balanceType"`
	CreditLimitIncluded bool   `json:"creditLimitIncluded"`
	BalanceAmount       struct {
		Currency string `json:"currency"`
		Amount   string `json:"amount"`
	} `json:"balanceAmount"`
}
