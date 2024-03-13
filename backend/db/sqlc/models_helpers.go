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

func CreateUpdateUsernameParam(ID int64, username string) *UpdateUsernameParams {
	return &UpdateUsernameParams{
		ID:       ID,
		Username: username,
	}
}

func CreateUpdateEmailParam(ID int64, email string) *UpdateEmailParams {
	return &UpdateEmailParams{
		ID:    ID,
		Email: email,
	}
}

type NewUserInfo struct {
	// email to sign in - also to send reminders
	Email    string       `json:"email"`
	Username string       `json:"username"`
	Birthday sql.NullTime `json:"birthday"`
}

func CreateNewUserInfo(email string, username string, birthday sql.NullTime) *NewUserInfo {
	return &NewUserInfo{
		Email:    email,
		Username: username,
		Birthday: birthday,
	}
}

func CreateUpdateBrithdayParams(ID int64, bday sql.NullTime) *UpdateBirthdayParams {
	return &UpdateBirthdayParams{
		ID:       ID,
		Birthday: bday,
	}
}
