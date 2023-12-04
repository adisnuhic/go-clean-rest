package repositories

// "sqlmock" is a mock library implementing sql/driver.
// Which has one and only purpose — to simulate any sql driver behavior in tests, without needing a real database connection.
// It helps to maintain correct TDD workflow.

// "suite" with it, you can build a testing suite as a struct, build setup/teardown methods and testing methods on your struct,
// and run them with 'go test' as per normal.

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/adisnuhic/go-clean/internal/models"
	"github.com/adisnuhic/go-clean/internal/repositories"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
)

// AccountSuite -
type AccountSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repo repositories.IAccountRepository
	user *models.User
}

// SetupSuite -
func (s *AccountSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	if err != nil {
		s.T().Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	s.DB, err = gorm.Open("mysql", db)
	if err != nil {
		s.T().Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	s.repo = NewMySQLAccountRepository(s.DB)
}

func (s *AccountSuite) Test_Register() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec("INSERT INTO `users`").WithArgs(s.user).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	_, err := s.repo.Register(s.user)
	s.Nil(err)
}
