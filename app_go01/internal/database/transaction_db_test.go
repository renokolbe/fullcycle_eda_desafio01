package database

import (
	"database/sql"
	"testing"

	"github.com/renokolbe/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	clientFrom    *entity.Client
	clientTo      *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (t *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	t.Nil(err)
	t.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance float, created_at date)")
	db.Exec("CREATE TABLE transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount float, created_at date)")
	t.transactionDB = NewTransactionDB(db)
	t.clientFrom, err = entity.NewClient("John Doe", "john@domnio.com")
	t.Nil(err)
	t.accountFrom = entity.NewAccount(t.clientFrom)
	t.accountFrom.Credit(1000)
	t.clientTo, err = entity.NewClient("Jane Doe", "jane@domnio.com")
	t.Nil(err)
	t.accountTo = entity.NewAccount(t.clientTo)
	t.accountTo.Credit(1000)
}

func (t *TransactionDBTestSuite) TearDownSuite() {
	defer t.db.Close()
	t.db.Exec("DROP TABLE transactions")
	t.db.Exec("DROP TABLE accounts")
	t.db.Exec("DROP TABLE clients")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (t *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(t.accountFrom, t.accountTo, 100)
	t.Nil(err)
	err = t.transactionDB.Create(transaction)
	t.Nil(err)
}
