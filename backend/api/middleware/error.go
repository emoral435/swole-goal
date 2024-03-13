package middleware

import "net/http"

// Use this special type for your handler funcs
type MyHandlerFunc func(w http.ResponseWriter, r *http.Request) error

// Pattern for endpoint on middleware chain, not takes a diff signature.
func ErrorHandler(h MyHandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Execute the final handler, and deal with errors
		err := h(w, r)
		if err != nil {
			// Deal with error here, show user error template, log etc
			panic(http.StateActive)
		}
	})
}
