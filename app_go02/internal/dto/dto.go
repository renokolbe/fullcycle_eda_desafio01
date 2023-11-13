package dto

type Payload struct {
	AccountIDFrom      string  `json:"account_id_from"`
	AccountIDTo        string  `json:"account_id_to"`
	BalanceAccountFrom float64 `json:"balance_account_from"`
	BalanceAccountTo   float64 `json:"balance_account_to"`
}

type TransactionInputDTO struct {
	Name    string  `json:"Name"`
	Payload Payload `json:"Payload"`
}
