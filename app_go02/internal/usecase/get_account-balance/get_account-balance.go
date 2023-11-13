package get_accountbalance

import "github.com/renokolbe/fc-ms-wallet-balance/internal/gateway"

type GetAccountBalanceInputDTO struct {
	ID string `json:"id"`
}

type GetAccountBalanceOutputDTO struct {
	Balance float64 `json:"balance"`
}

type GetAccountBalanceUseCase struct {
	AccountBalanceGateway gateway.AccountBalanceGateway
}

func NewGetAccountBalanceUseCase(accountBalanceGateway gateway.AccountBalanceGateway) *GetAccountBalanceUseCase {
	return &GetAccountBalanceUseCase{AccountBalanceGateway: accountBalanceGateway}
}

func (uc *GetAccountBalanceUseCase) Execute(inputDTO GetAccountBalanceInputDTO) (*GetAccountBalanceOutputDTO, error) {
	balance, err := uc.AccountBalanceGateway.GetAccountBalance(inputDTO.ID)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			return nil, err
		}
		return &GetAccountBalanceOutputDTO{
			Balance: -1,
		}, nil
	}
	return &GetAccountBalanceOutputDTO{
		Balance: balance,
	}, nil
}
