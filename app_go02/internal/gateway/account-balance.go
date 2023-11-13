package gateway

type AccountBalanceGateway interface {
	InsertAccountBalance(id string, amount float64) error
	GetAccountBalance(id string) (float64, error)
	UpdateAccountBalance(id string, amount float64) error
}
