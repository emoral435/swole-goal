package db

import (
	"database/sql"
)

// CreateUpdatePasswordParam returns a new UpdatePasswordParam struct
func CreateUpdatePasswordParam(ID int64, password string) *UpdatePasswordParams {
	return &UpdatePasswordParams{
		ID:       ID,
		Password: password,
	}
}

// CreateUpdateUsernameParam returns a new UpdateUsernameParams struct
func CreateUpdateUsernameParam(ID int64, username string) *UpdateUsernameParams {
	return &UpdateUsernameParams{
		ID:       ID,
		Username: username,
	}
}

// CreateUpdateEmailParam returns a new UpdateEmailParams struct
func CreateUpdateEmailParam(ID int64, email string) *UpdateEmailParams {
	return &UpdateEmailParams{
		ID:    ID,
		Email: email,
	}
}

// NewUserInfo is a struct that only contains non-sensitive information of a user.
//This is all user information, outside of their password and when they created their account
type NewUserInfo struct {
	// email to sign in - also to send reminders
	Email    string       `json:"email"`
	Username string       `json:"username"`
	Birthday sql.NullTime `json:"birthday"`
}

// CreateNewUserInfo returns a new NewUserInfo struct
func CreateNewUserInfo(email string, username string, birthday sql.NullTime) *NewUserInfo {
	return &NewUserInfo{
		Email:    email,
		Username: username,
		Birthday: birthday,
	}
}

// CreateUpdateBirthdayParams returns a new UpdateBirthdayParams struct
func CreateUpdateBirthdayParams(ID int64, bday sql.NullTime) *UpdateBirthdayParams {
	return &UpdateBirthdayParams{
		ID:       ID,
		Birthday: bday,
	}
}
