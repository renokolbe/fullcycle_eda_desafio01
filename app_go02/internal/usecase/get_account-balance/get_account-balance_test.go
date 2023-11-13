package get_accountbalance

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

func TestGetAcccountBalanceUseCase_Execute(t *testing.T) {
	accountBalanceMock := AccountBalanceMock{}
	accountBalanceMock.On("GetAccountBalance", mock.Anything).Return(100.00, nil)
	uc := NewGetAccountBalanceUseCase(&accountBalanceMock)
	output, err := uc.Execute(GetAccountBalanceInputDTO{
		ID: "123",
	})

	assert.Nil(t, err)
	assert.NotNil(t, output.Balance)
	accountBalanceMock.AssertExpectations(t)
	accountBalanceMock.AssertNumberOfCalls(t, "GetAccountBalance", 1)

}
