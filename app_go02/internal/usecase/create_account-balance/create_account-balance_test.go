package create_accountbalance

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AccountBalanceMock struct {
	mock.Mock
}

func (m *AccountBalanceMock) GetAccountBalance(id string) (float64, error) {
	args := m.Called(id)
	return args.Get(0).(float64), args.Error(1)
}

func (m *AccountBalanceMock) InsertAccountBalance(id string, amount float64) error {
	args := m.Called(id, amount)
	return args.Error(0)
}

func (m *AccountBalanceMock) UpdateAccountBalance(id string, amount float64) error {
	args := m.Called(id, amount)
	return args.Error(0)
}

func TestCreateAccountBalanceUsecase_Execute(t *testing.T) {
	accountBalanceMock := AccountBalanceMock{}
	accountBalanceMock.On("InsertAccountBalance", mock.Anything, mock.Anything).Return(nil)
	uc := NewAccountBalanceDBUseCase(&accountBalanceMock)

	err := uc.Execute(CreateAccountBalanceInputDTO{
		ID:     "123",
		Balace: 100,
	})

	assert.Nil(t, err)
	accountBalanceMock.AssertExpectations(t)
	accountBalanceMock.AssertNumberOfCalls(t, "InsertAccountBalance", 1)
}
