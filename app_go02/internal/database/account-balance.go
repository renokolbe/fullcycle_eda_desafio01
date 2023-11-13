package database

import "database/sql"

type AccountBalanceDB struct {
	DB *sql.DB
}

func NewAccountBalanceDB(db *sql.DB) *AccountBalanceDB {
	return &AccountBalanceDB{DB: db}
}

func (a *AccountBalanceDB) InsertAccountBalance(id string, amount float64) error {
	_, err := a.DB.Exec("INSERT INTO account_balance (id, balance) VALUES (?, ?)", id, amount)
	if err != nil {
		return err
	}
	return nil
}

func (a *AccountBalanceDB) GetAccountBalance(id string) (float64, error) {
	var balance float64
	err := a.DB.QueryRow("SELECT balance FROM account_balance WHERE id = ?", id).Scan(&balance)
	if err != nil {
		return -1, err
	}
	return balance, nil
}

func (a *AccountBalanceDB) UpdateAccountBalance(id string, amount float64) error {
	_, err := a.DB.Exec("UPDATE account_balance SET balance = ? WHERE id = ?", amount, id)
	if err != nil {
		return err
	}
	return nil
}
