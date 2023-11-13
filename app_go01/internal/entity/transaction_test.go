package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := NewClient("Cliente 1", "c1@c.com")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Cliente 2", "c2@c.com")
	account2 := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 100)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 1100.0, account2.Balance)
	assert.Equal(t, 900.0, account1.Balance)

}

func TestCreateTransactionWithNegativeAmount(t *testing.T) {
	client1, _ := NewClient("Cliente 1", "c1@c.com")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Cliente 2", "c2@c.com")
	account2 := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, -100)
	assert.NotNil(t, err)
	assert.Error(t, err, "amount must be greater than zero")
	assert.Nil(t, transaction)
	assert.Equal(t, account1.Balance, account2.Balance)

}

func TestCreateTransactionWithInsufficientBalance(t *testing.T) {
	client1, _ := NewClient("Cliente 1", "c1@c.com")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Cliente 2", "c2@c.com")
	account2 := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 2000)
	assert.NotNil(t, err)
	assert.Error(t, err, "insufficient funds")
	assert.Nil(t, transaction)
	assert.Equal(t, account1.Balance, account2.Balance)

}
