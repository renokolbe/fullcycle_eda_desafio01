package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewAccount(t *testing.T) {
	client, _ := NewClient("Cliente 1", "c@c.com")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateAccountWithoutClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestAccountCredit(t *testing.T) {
	client, _ := NewClient("Cliente 1", "c@c.com")
	account := NewAccount(client)
	account.Credit(10)
	assert.Equal(t, account.Balance, 10.00)
}

func TestAccountDebit(t *testing.T) {
	client, _ := NewClient("Cliente 1", "c@c.com")
	account := NewAccount(client)
	account.Credit(10)
	account.Debit(5)
	assert.Equal(t, account.Balance, 5.00)
}
