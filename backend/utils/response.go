package util

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// Generic Error Response
type ErrorResponse struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

func CreateErrorResponse(errorMsg string, status int) *ErrorResponse {
	return &ErrorResponse{
		Error:  errorMsg,
		Status: status,
	}
}

// Generic success response
type SuccessResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func CreateSuccessResponse(successMsg string, status int) *SuccessResponse {
	return &SuccessResponse{
		Message: successMsg,
		Status:  status,
	}
}

// Generic Check Error response that sends generic JSON response
//
// If there was an issue from the database, then we return the error that was from the database.
// If it was a server-side error, we return the error from the server and send a status 500 err
func CheckError(err error, res http.ResponseWriter, req *http.Request) error {
	// deal with bad request (params for creating user not satisfied)
	if err != nil {
		// database error - most likely the users fault
		if err == sql.ErrNoRows {
			res.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(res).Encode(CreateErrorResponse(err.Error(), http.StatusBadRequest))
		}

		// server-side error
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(CreateErrorResponse(err.Error(), http.StatusInternalServerError))
		return err
	}
	return nil
}

// HashPassword hashes our password - 60 bytes long
func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashPassword), nil
}

// CompareHash compares the password that we store in the database against the input password
func CompareHash(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
