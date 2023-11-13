package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccountBalance(t *testing.T) {
	// Given
	id := "123"
	balance := 100.0

	// When
	accountBalance := NewAccountBalance(id, balance)

	// Then
	assert.Equal(t, id, accountBalance.id)
	assert.Equal(t, balance, accountBalance.balance)
}

func TestUpdateAccountBalance(t *testing.T) {
	// Given
	id := "123"
	balance := 100.0
	accountBalance := NewAccountBalance(id, balance)

	// When
	newBalance := 200.0
	accountBalance.UpdateBalance(newBalance)

	// Then
	assert.Equal(t, id, accountBalance.id)
	assert.Equal(t, newBalance, accountBalance.balance)
}
