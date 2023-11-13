package entity

import (
	"sync"

	"github.com/renokolbe/fc-ms-wallet-balance/internal/gateway"
	create_accountbalance "github.com/renokolbe/fc-ms-wallet-balance/internal/usecase/create_account-balance"
	get_accountbalance "github.com/renokolbe/fc-ms-wallet-balance/internal/usecase/get_account-balance"
	update_accountbalance "github.com/renokolbe/fc-ms-wallet-balance/internal/usecase/update_account-balance"
)

type AccountBalances struct {
	AccountBalance     []*AccountBalance
	AccountBalanceChan chan *AccountBalance
	Wg                 *sync.WaitGroup
}

func NewAccountBalances(abChan chan *AccountBalance, wg *sync.WaitGroup) *AccountBalances {
	return &AccountBalances{
		AccountBalance:     []*AccountBalance{},
		AccountBalanceChan: abChan,
		Wg:                 wg,
	}
}

func (t *AccountBalances) Registro(gw gateway.AccountBalanceGateway) {
	for ab := range t.AccountBalanceChan {
		ug := get_accountbalance.NewGetAccountBalanceUseCase(gw)
		balance, err := ug.Execute(get_accountbalance.GetAccountBalanceInputDTO{ID: ab.id})
		if err != nil {
			panic(err)
		}
		if balance.Balance == -1 {
			uc := create_accountbalance.NewAccountBalanceDBUseCase(gw)
			inputDTO := create_accountbalance.CreateAccountBalanceInputDTO{
				ID:     ab.id,
				Balace: ab.balance,
			}
			uc.Execute(inputDTO)
		}
		uu := update_accountbalance.NewUpdateAccountBalanceUseCase(gw)
		inputUpdateDTO := update_accountbalance.UpdateAccountBalanceInputDTO{
			ID:      ab.id,
			Balance: ab.balance,
		}
		err = uu.Execute(inputUpdateDTO)
		if err != nil {
			panic(err)
		}
	}
}
