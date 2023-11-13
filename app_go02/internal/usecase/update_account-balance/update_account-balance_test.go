package update_accountbalance

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

func TestUpdateAccountBalanceUseCase_Execute(t *testing.T) {
	accountBalanceMock := AccountBalanceMock{}
	accountBalanceMock.On("UpdateAccountBalance", mock.Anything, mock.Anything).Return(nil)
	uc := NewUpdateAccountBalanceUseCase(&accountBalanceMock)
	err := uc.Execute(UpdateAccountBalanceInputDTO{
		ID:      "123",
		Balance: 100.00,
	})

	assert.Nil(t, err)
	accountBalanceMock.AssertExpectations(t)
	accountBalanceMock.AssertNumberOfCalls(t, "UpdateAccountBalance", 1)

}
