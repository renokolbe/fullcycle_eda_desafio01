package mocks

import (
	"context"

	"github.com/renokolbe/fc-ms-wallet/pkg/uow"
	"github.com/stretchr/testify/mock"
)

type UowMock struct {
	mock.Mock
}

/*
Register(name string, fc RepositoryFactory)
	GetRepository(ctx context.Context, name string) (interface{}, error)
	Do(ctx context.Context, fn func(uow *Uow) error) error
	CommitOrRollback() error
	Rollback() error
	UnRegister(name string)
*/

func (m *UowMock) Register(name string, rfc uow.RepositoryFactory) {
	m.Called(name, rfc)
}

func (m *UowMock) GetRepository(ctx context.Context, name string) (interface{}, error) {
	args := m.Called(ctx, name)
	return args.Get(0), args.Error(1)
}

func (m *UowMock) Do(ctx context.Context, fn func(uow *uow.Uow) error) error {
	args := m.Called(fn)
	return args.Error(0)

}

func (m *UowMock) CommitOrRollback() error {
	args := m.Called()
	return args.Error(0)

}

func (m *UowMock) Rollback() error {
	args := m.Called()
	return args.Error(0)

}

func (m *UowMock) UnRegister(name string) {
	m.Called(name)
}
