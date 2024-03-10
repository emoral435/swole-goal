package util

import (
	"encoding/json"
	"net/http"
)

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

func CheckError(err error, res http.ResponseWriter, req *http.Request) error {
	// deal with bad request (params for creating user not satisfied)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(CreateErrorResponse(err.Error(), http.StatusBadRequest))
		return err
	}
	return nil
}
