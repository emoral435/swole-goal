package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/emoral435/swole-goal/api/token"
	util "github.com/emoral435/swole-goal/utils"
)

type AUTHPAYLOADKEY string

const (
	authKey        = "authentication"
	authTypeBearer = "bearer"
)

var authPayloadKey AUTHPAYLOADKEY = "authorization_payload"

func AuthMiddleware(tokenMaker token.Maker, next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get(authKey)

		if len(authHeader) <= 0 {
			http.Error(res, "Authorization header not provided.", http.StatusBadRequest)
			panic(http.ErrAbortHandler)
		}

		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			http.Error(res, "Invalid authorization header format", http.StatusBadRequest)
			panic(http.ErrAbortHandler)
		}

		authType := strings.ToLower(fields[0])

		if authType != authTypeBearer {
			http.Error(res, "Unsupported authorization type", http.StatusBadRequest)
			panic(http.ErrAbortHandler)
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)

		if err = util.CheckError(err, res, req); err != nil {
			return
		}

		ctx := context.WithValue(req.Context(), authPayloadKey, payload) // https://stackoverflow.com/questions/39946583/how-to-pass-context-in-golang-request-to-middleware to see where this came from
		next.ServeHTTP(res, req.WithContext(ctx))
	})
}
