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

// TokenSuite -
type TokenSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repo  repositories.ITokenRepository
	token *models.Token
}

// SetupSuite -
func (s *TokenSuite) SetupSuite() {
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

	userID := uint64(1)

	s.token = &models.Token{
		ID:          1,
		UserID:      &userID,
		Token:       "test",
		TokenTypeID: 1,
		Code:        "test",
		ExpiresAt:   time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.repo = NewMySQLTokenRepository(s.DB)
}

func (s *TokenSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInitToken(t *testing.T) {
	suite.Run(t, new(TokenSuite))
}

func (s *TokenSuite) TestCreateToken() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `tokens` (`id`,`user_id`,`token`,`token_type_id`,`code`,`expires_at`,`created_at`,`updated_at`) VALUES (?,?,?,?,?,?,?,?)")).
		WithArgs(s.token.ID, s.token.UserID, s.token.Token, s.token.TokenTypeID, s.token.Code, AnyTime{}, AnyTime{}, AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()

	token, err := s.repo.CreateToken(s.token)
	require.Nil(s.T(), err)
	require.NotNil(s.T(), token)
	require.Equal(s.T(), token, s.token)
}

func (s *TokenSuite) TestGetByToken() {
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `tokens` WHERE (token = ?)")).
		WithArgs(s.token.Token).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "token", "token_type_id", "code"}).
			AddRow(s.token.ID, s.token.UserID, s.token.Token, s.token.TokenTypeID, s.token.Code))

	token, err := s.repo.GetByToken(s.token.Token)
	require.Nil(s.T(), err)
	require.NotNil(s.T(), token)
	require.Equal(s.T(), token.Token, s.token.Token)
}

func (s *TokenSuite) TestUpdate() {
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta("UPDATE `tokens` SET `user_id` = ?, `token` = ?, `token_type_id` = ?, `code` = ?, `expires_at` = ?, `created_at` = ?, `updated_at` = ? WHERE `tokens`.`id` = ?")).
		WithArgs(s.token.UserID, s.token.Token, s.token.TokenTypeID, s.token.Code, AnyTime{}, AnyTime{}, AnyTime{}, s.token.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	s.mock.ExpectCommit()

	err := s.repo.Update(s.token)
	s.Nil(err)
}
