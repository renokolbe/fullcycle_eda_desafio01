package entity

type AccountBalance struct {
	id      string
	balance float64
}

func NewAccountBalance(id string, balance float64) *AccountBalance {
	return &AccountBalance{id: id, balance: balance}
}

func (a *AccountBalance) UpdateBalance(amount float64) {
	a.balance = amount
}
