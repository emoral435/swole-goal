package middleware

import (
	"encoding/json"
	"mime"
	"net/http"

	util "github.com/emoral435/swole-goal/utils"
)

func EnforceJSONHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		contentType := req.Header.Get("Content-Type")

		if contentType != "" {
			mt, _, err := mime.ParseMediaType(contentType)
			if err != nil {
				http.Error(res, "Malformed Content-Type header", http.StatusBadRequest)
				gotEncoded := json.NewEncoder(res).Encode(util.CreateErrorResponse(err.Error(), http.StatusBadRequest))
				if gotEncoded != nil {
					http.Error(res, "Also could not format the response lol.", http.StatusInternalServerError)
				}
				return
			}

			if mt != "application/json" {
				http.Error(res, "Content-Type header must be application/json", http.StatusUnsupportedMediaType)
				json.NewEncoder(res).Encode(util.CreateErrorResponse(mt, http.StatusBadRequest))
				return
			}
		}

		next.ServeHTTP(res, req)
	})
}
