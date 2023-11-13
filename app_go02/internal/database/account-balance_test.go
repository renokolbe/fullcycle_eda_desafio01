package database

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type AccountBalanceDBTestSuite struct {
	suite.Suite
	db               *sql.DB
	accountBalanceDB *AccountBalanceDB
}

func (s *AccountBalanceDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE account_balance (id varchar(255), balance float)")
	s.accountBalanceDB = NewAccountBalanceDB(db)
}

func (s *AccountBalanceDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE account_balance")
}
func TestAccountBalanceDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountBalanceDBTestSuite))
}

func (s *AccountBalanceDBTestSuite) TestInsertAccountBalance() {
	err := s.accountBalanceDB.InsertAccountBalance("123", 100)
	s.Nil(err)
}

func (s *AccountBalanceDBTestSuite) TestFindAccountBalance() {
	err := s.accountBalanceDB.InsertAccountBalance("123", 100)
	s.Nil(err)
	balance, err := s.accountBalanceDB.GetAccountBalance("123")
	s.Nil(err)
	s.Equal(100.0, balance)
}

func (s *AccountBalanceDBTestSuite) TestUpdateAccountBalance() {
	err := s.accountBalanceDB.InsertAccountBalance("123", 100)
	s.Nil(err)
	err = s.accountBalanceDB.UpdateAccountBalance("123", 200)
	s.Nil(err)
	balance, err := s.accountBalanceDB.GetAccountBalance("123")
	s.Nil(err)
	s.Equal(200.0, balance)
}
