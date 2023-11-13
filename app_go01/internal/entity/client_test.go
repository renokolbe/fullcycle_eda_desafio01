package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("Cliente 1", "c@c.com")
	assert.Nil(t, err)
	assert.NotNil(t, client.ID)
	assert.Equal(t, "Cliente 1", client.Name)
	assert.Equal(t, "c@c.com", client.Email)
}

func TestCreateNewClientWithInvalidArguments(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("Cliente 1", "c@c.com")
	err := client.Update("Cliente 1 Atualizado", "novo_c@c.com")
	assert.Nil(t, err)
	assert.Equal(t, "Cliente 1 Atualizado", client.Name)
	assert.Equal(t, "novo_c@c.com", client.Email)
}

func TestUpdateClientWithInvalidArguments(t *testing.T) {
	client, _ := NewClient("Cliente 1", "c@c.com")
	err := client.Update("", "c@c.com")
	assert.Error(t, err, "name is required")
}

func TestAddAccounttoClient(t *testing.T) {
	client, _ := NewClient("Cliente 1", "c@c.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}

func TestAddAccounttoWrongClient(t *testing.T) {
	client1, _ := NewClient("Cliente 1", "c@c.com")
	client2, _ := NewClient("Cliente 2", "c@c.com")
	account := NewAccount(client1)
	err := client2.AddAccount(account)
	assert.Error(t, err, "account does not belong to Client")
}
