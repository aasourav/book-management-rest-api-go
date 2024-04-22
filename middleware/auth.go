package middleware

import (
	"fmt"
	"gorilla-test/data"
	"net/http"
)

func ValidateUser(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.Header.Get("username")
		password := r.Header.Get("password")

		if data.Users[username] != password || username == "" {
			w.Write([]byte("failed to authenticate"))
			return
		}
		f(w, r)
	}
}

func ValidateOwner(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.Header.Get("username")
		password := r.Header.Get("password")
		IfOwner := r.Header.Get("userType")

		if IfOwner != "owner" {
			w.Write([]byte("you are  not owner"))
		}

		if data.Users[username] != password || username == "" {
			w.Write([]byte("failed to authenticate"))
			return
		}
		f(w, r)
	}
}

func TrackNumberOfRequests(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data.NumberOfRequest += 1
		fmt.Println("Request number: ", data.NumberOfRequest)
	})
}
