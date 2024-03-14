package util

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// ErrorResponse is a generic errror response struct that we can return back as a JSON error response
type ErrorResponse struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

// CreateErrorResponse is a generic error response struct maker
func CreateErrorResponse(errorMsg string, status int) *ErrorResponse {
	return &ErrorResponse{
		Error:  errorMsg,
		Status: status,
	}
}

// SuccessResponse is a generic success response struct that we can return back as a JSON error response
type SuccessResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// CreateSuccessResponse is a generic success response struct maker
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
			gotEncoded := json.NewEncoder(res).Encode(CreateErrorResponse(err.Error(), http.StatusBadRequest))
			if gotEncoded != nil {
				return gotEncoded
			}
			return err
		}

		// server-side error
		res.WriteHeader(http.StatusInternalServerError)
		gotEncoded := json.NewEncoder(res).Encode(CreateErrorResponse(err.Error(), http.StatusInternalServerError))
		if gotEncoded != nil {
			return gotEncoded
		}
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

// ReturnValidJSONResponse sends back a http 200 status code indicating the request was successful and valid.
//
// This function accepts any valid JSON response to send back as a response as well
func ReturnValidJSONResponse(res http.ResponseWriter, jsonObject any) {
	// send back the correct response
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(jsonObject)
}
