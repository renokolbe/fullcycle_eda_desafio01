package transformer

import (
	"github.com/renokolbe/fc-ms-wallet-balance/internal/dto"
	"github.com/renokolbe/fc-ms-wallet-balance/internal/entity"
)

func TransformInputAccountFrom(input dto.TransactionInputDTO) *entity.AccountBalance {
	return entity.NewAccountBalance(input.Payload.AccountIDFrom, input.Payload.BalanceAccountFrom)
}

func TransformInputAccountTo(input dto.TransactionInputDTO) *entity.AccountBalance {
	return entity.NewAccountBalance(input.Payload.AccountIDTo, input.Payload.BalanceAccountTo)
}
