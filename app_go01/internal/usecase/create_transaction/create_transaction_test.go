package create_transaction

import (
	"context"
	"testing"

	"github.com/renokolbe/fc-ms-wallet/internal/entity"
	"github.com/renokolbe/fc-ms-wallet/internal/event"
	"github.com/renokolbe/fc-ms-wallet/internal/usecase/mocks"
	"github.com/renokolbe/fc-ms-wallet/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*
type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}
*/

func Test_CreateTransactionUseCase_Execute(t *testing.T) {
	client_from, _ := entity.NewClient("John Doe", "email@dominio.com")
	account_from := entity.NewAccount(client_from)
	account_from.Credit(1000)

	client_to, _ := entity.NewClient("Jane Doe", "email_jane@dominio.com")
	account_to := entity.NewAccount(client_to)
	account_to.Credit(1000)

	uowMock := &mocks.UowMock{}
	uowMock.On("Do", mock.Anything, mock.Anything).Return(nil)
	/*
		accountMock := &mocks.AccountGatewayMock{}
		accountMock.On("FindByID", account_from.ID).Return(account_from, nil)
		accountMock.On("FindByID", account_to.ID).Return(account_to, nil)

		transactionMock := &mocks.TransactionGatewayMock{}
		transactionMock.On("Create", mock.Anything).Return(nil)
	*/

	dispatcher := events.NewEventDispatcher()
	eventTransaction := event.NewTransactionCreated()
	eventBalance := event.NewBalanceUpdated()
	ctx := context.Background()

	//	uc := NewCreateTransactionUseCase(accountMock, transactionMock, dispatcher, event)
	uc := NewCreateTransactionUseCase(uowMock, dispatcher, eventTransaction, eventBalance)

	inputDTO := CreateTransactionInputDTO{
		AccountIDFrom: account_from.ID,
		AccountIDTo:   account_to.ID,
		Amount:        100,
	}

	output, err := uc.Execute(ctx, inputDTO)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	//assert.NotEmpty(t, output.TransactionID)
	//accountMock.AssertExpectations(t)
	//transactionMock.AssertExpectations(t)
	//accountMock.AssertNumberOfCalls(t, "FindByID", 2)
	//transactionMock.AssertNumberOfCalls(t, "Create", 1)
	uowMock.AssertExpectations(t)
	uowMock.AssertNumberOfCalls(t, "Do", 1)

}
