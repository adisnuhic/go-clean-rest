package repositories

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/adisnuhic/go-clean/internal/models"
	"github.com/adisnuhic/go-clean/internal/repositories"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// UserSuite -
type UserSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repo repositories.IUserRepository
	user *models.User
}

// SetupSuite -
func (s *UserSuite) SetupSuite() {
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

	s.user = &models.User{
		ID:          1,
		FirstName:   "test",
		LastName:    "test",
		Email:       "test@test.com",
		AcceptedTos: true,
		IsConfirmed: true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.repo = NewMySQLUserRepository(s.DB)
}

func (s *UserSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

// TestInitUser -
func TestInitUser(t *testing.T) {
	suite.Run(t, new(UserSuite))
}

// TestGetByID -
func (s *UserSuite) TestGetByID() {
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL AND ((id = ?))")).
		WithArgs(s.user.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "accepted_tos", "is_confirmed", "created_at", "updated_at"}).
			AddRow(s.user.ID, s.user.FirstName, s.user.LastName, s.user.Email, s.user.AcceptedTos, s.user.IsConfirmed, s.user.CreatedAt, s.user.UpdatedAt))

	user, err := s.repo.GetByID(s.user.ID)
	s.Nil(err)

	require.Nil(s.T(), err)
	require.NotNil(s.T(), user)
	require.Equal(s.T(), s.user.ID, user.ID)
	require.Equal(s.T(), s.user.FirstName, user.FirstName)
	require.Equal(s.T(), s.user.LastName, user.LastName)
	require.Equal(s.T(), s.user.Email, user.Email)
	require.Equal(s.T(), s.user.AcceptedTos, user.AcceptedTos)
	require.Equal(s.T(), s.user.IsConfirmed, user.IsConfirmed)
}

func (s *UserSuite) TestGetByEmail() {
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL AND ((email = ?))")).
		WithArgs(s.user.Email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "accepted_tos", "is_confirmed", "created_at", "updated_at"}).
			AddRow(s.user.ID, s.user.FirstName, s.user.LastName, s.user.Email, s.user.AcceptedTos, s.user.IsConfirmed, s.user.CreatedAt, s.user.UpdatedAt))

	user, err := s.repo.GetByEmail(s.user.Email)
	s.Nil(err)

	require.Nil(s.T(), err)
	require.NotNil(s.T(), user)
	require.Equal(s.T(), s.user.ID, user.ID)
	require.Equal(s.T(), s.user.FirstName, user.FirstName)
	require.Equal(s.T(), s.user.LastName, user.LastName)
	require.Equal(s.T(), s.user.Email, user.Email)
	require.Equal(s.T(), s.user.AcceptedTos, user.AcceptedTos)
	require.Equal(s.T(), s.user.IsConfirmed, user.IsConfirmed)
}
