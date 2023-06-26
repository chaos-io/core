package jwtauth_test

import (
	"fmt"
	"net/http"

	jwtauth2 "github.com/chaos-io/core/httputil/middleware/jwtauth"
)

func ExampleVerifyToken() {
	// Create HTTP router.
	r := chi.NewRouter()

	// Plug middleware in for any handler
	// Note: signing method and key func are required
	r.Use(jwtauth2.VerifyToken(
		jwtauth2.WithSigningMethod(jwt.SigningMethodHS512),
		jwtauth2.WithKeyFunc(func(_ *jwt.Token) (interface{}, error) {
			return []byte("my_symmetric_secret"), nil
		}),
	))

	// add handlers and catch parsed and verified JWT
	r.Handle("/user/info", http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		token := r.Context().Value(jwtauth2.TokenCtxKey)
		fmt.Printf("user token: %+v", token)
	}))
}
