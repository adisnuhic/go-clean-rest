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

// AuthProviderSuite -
type AuthProviderSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repo         repositories.IAuthProviderRepository
	authProvider *models.AuthProvider
}

// SetupSuite -
func (s *AuthProviderSuite) SetupSuite() {
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

	s.authProvider = &models.AuthProvider{
		Provider:  "local",
		UserID:    1,
		UID:       "test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.repo = NewMySQLAuthProviderRepository(s.DB)
}

func TestInitAuthProvider(t *testing.T) {
	suite.Run(t, new(AuthProviderSuite))
}

func (s *AuthProviderSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *AuthProviderSuite) TestGetByUserID() {
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `auth_providers` WHERE (user_id = ?)")).
		WithArgs(s.authProvider.UserID).
		WillReturnRows(sqlmock.NewRows([]string{"provider", "user_id", "uid"}).
			AddRow(s.authProvider.Provider, s.authProvider.UserID, s.authProvider.UID))

	_, err := s.repo.GetByUserID(s.authProvider.UserID)
	s.Nil(err)
}

func (s *AuthProviderSuite) TestGetByUserIDProviderID() {
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `auth_providers` WHERE (user_id = ?) AND (provider = ?)")).
		WithArgs(s.authProvider.UserID, s.authProvider.Provider).
		WillReturnRows(sqlmock.NewRows([]string{"provider", "user_id", "uid"}).
			AddRow(s.authProvider.Provider, s.authProvider.UserID, s.authProvider.UID))

	_, err := s.repo.GetByUserIDProviderID(s.authProvider.UserID, s.authProvider.Provider)
	s.Nil(err)
}

func (s *AuthProviderSuite) TestUpdate() {
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE auth_providers SET uid = ? WHERE user_id = ? AND provider = ?")).
		WithArgs(s.authProvider.UID, s.authProvider.UserID, s.authProvider.Provider).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.repo.Update(s.authProvider)
	s.Nil(err)
}
