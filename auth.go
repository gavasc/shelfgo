package main

import "net/http"

func isAuthorized(r *http.Request, expectedUser, expectedPass string) bool {
	user, pass, ok := r.BasicAuth()

	return ok && user == expectedUser && pass == expectedPass
}
