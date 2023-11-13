package create_accountbalance

import "github.com/renokolbe/fc-ms-wallet-balance/internal/gateway"

type CreateAccountBalanceInputDTO struct {
	ID     string  `json:"id"`
	Balace float64 `json:"balance"`
}

type CreateAccountBalanceUseCase struct {
	AccountBalanceGateway gateway.AccountBalanceGateway
}

func NewAccountBalanceDBUseCase(accountBalanceGateway gateway.AccountBalanceGateway) *CreateAccountBalanceUseCase {
	return &CreateAccountBalanceUseCase{AccountBalanceGateway: accountBalanceGateway}
}

func (uc *CreateAccountBalanceUseCase) Execute(inputDTO CreateAccountBalanceInputDTO) error {
	err := uc.AccountBalanceGateway.InsertAccountBalance(inputDTO.ID, inputDTO.Balace)
	if err != nil {
		return err
	}
	return nil
}
