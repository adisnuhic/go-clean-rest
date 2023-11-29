package ecode

const (

	// Status codes
	// ErrStatusNotFound code
	ErrStatusNotFoundCode = 404

	// ErrUnableToParseQueryStringCode code
	ErrUnableToParseQueryStringCode = 1
	// ErrUnableToParseQueryStringCode message
	ErrUnableToParseQueryStringMsg = "unable to parse query string parameter"

	// ErrRequestParamValidationCode code
	ErrRequestParamValidationCode = 10
	// ErrRequestParamValidationMsg message
	ErrRequestParamValidationMsg = "reuqest param validation error"

	// ErrPermissionCode code
	ErrPermissionCode = 20
	// ErrPermissionMsg message
	ErrPermissionMsg = "permission error"

	// ErrCurrentPasswordMismatchedCode code
	ErrCurrentPasswordMismatchedCode = 30
	// ErrCurrentPasswordMismatchedMsg message
	ErrCurrentPasswordMismatchedMsg = "current password does not match"

	// ErrUnableToFetchUserCode code
	ErrUnableToFetchUserCode = 40
	// ErrUnableToFetchUserMsg message
	ErrUnableToFetchUserMsg = "unable to fetch user"

	// ErrUserExistsCode code
	ErrUserExistsCode = 50
	// ErrUserExistsMsg message
	ErrUserExistsMsg = "user already exists"

	// ErrUserDoesNotExistsCode code
	ErrUserDoesNotExistsCode = 60
	// ErrUserDoesNotExistsMsg message
	ErrUserDoesNotExistsMsg = "user does not exists"

	// ErrLoginFailedCode code
	ErrLoginFailedCode = 70
	// ErrLoginFailedMsg message
	ErrLoginFailedMsg = "wrong username and/or password"

	// ErrUnableToCreateAuthCode code
	ErrUnableToCreateAuthCode = 80
	// ErrUnableToCreateAuthMsg message
	ErrUnableToCreateAuthMsg = "unable to create auth"

	// ErrUnableToFetchAuthCode code
	ErrUnableToFetchAuthCode = 90
	// ErrUnableToFetchAuthMsg message
	ErrUnableToFetchAuthMsg = "unable to fetch auth"

	// ErrUnableToSaveAuthCode code
	ErrUnableToSaveAuthCode = 100
	// ErrUnableToSaveAuthMsg message
	ErrUnableToSaveAuthMsg = "unable to save auth"

	// ErrUnableToCreateTokenCode code
	ErrUnableToCreateTokenCode = 110
	// ErrUnableToCreateTokenMsg message
	ErrUnableToCreateTokenMsg = "unable to create token"

	// ErrUnableToGenerateHashCode code
	ErrUnableToGenerateHashCode = 120
	// ErrUnableToGenerateHashMsg message
	ErrUnableToGenerateHashMsg = "unable to generate hash"

	// ErrUnableToGenerateAccessTokenCode code
	ErrUnableToGenerateAccessTokenCode = 130
	// ErrUnableToGenerateAccessTokenMsg message
	ErrUnableToGenerateAccessTokenMsg = "unable to generate access token"

	// ErrUnableToGetTokenCode code
	ErrUnableToGetTokenCode = 140
	// ErrUnableToGetTokenMsg message
	ErrUnableToGetTokenMsg = "unable to fetch token"

	// ErrNotRefreshTokenCode code
	ErrNotRefreshTokenCode = 150
	// ErrNotRefreshTokenMsg message
	ErrNotRefreshTokenMsg = "token is not refresh token"

	// ErrTokenExpiredCode code
	ErrTokenExpiredCode = 160
	// ErrTokenExpiredMsg message
	ErrTokenExpiredMsg = "token is expired"

	// ErrUnableToUpdateTokenCode code
	ErrUnableToUpdateTokenCode = 170
	// ErrUnableToUpdateTokenMsg message
	ErrUnableToUpdateTokenMsg = "unable to update token"

	// ErrGeneratingTokenCode code
	ErrGeneratingTokenCode = 180
	// ErrGeneratingTokenMsg message
	ErrGeneratingTokenMsg = "error generating token"

	// ErrUnableToCreateUserCode code
	ErrUnableToCreateUserCode = 190
	// ErrUnableToCreateUserMsg message
	ErrUnableToCreateUserMsg = "unable to create user"

	// ErrUnableToGetUserCode code
	ErrUnableToGetUserCode = 200
	// ErrUnableToGetUserMsg message
	ErrUnableToGetUserMsg = "unable to get user"
)
