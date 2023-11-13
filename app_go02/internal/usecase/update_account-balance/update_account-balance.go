package update_accountbalance

import "github.com/renokolbe/fc-ms-wallet-balance/internal/gateway"

type UpdateAccountBalanceInputDTO struct {
	ID      string  `json:"id"`
	Balance float64 `json:"balance"`
}

type UpdateAccountBalanceUseCase struct {
	AccountBalanceGateway gateway.AccountBalanceGateway
}

func NewUpdateAccountBalanceUseCase(accountBalanceGateway gateway.AccountBalanceGateway) *UpdateAccountBalanceUseCase {
	return &UpdateAccountBalanceUseCase{AccountBalanceGateway: accountBalanceGateway}
}

func (uc *UpdateAccountBalanceUseCase) Execute(inputDTO UpdateAccountBalanceInputDTO) error {
	err := uc.AccountBalanceGateway.UpdateAccountBalance(inputDTO.ID, inputDTO.Balance)
	if err != nil {
		return err
	}
	return nil
}
