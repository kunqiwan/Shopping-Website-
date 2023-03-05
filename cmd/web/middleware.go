package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

// WriteToConsole log a message whenever a request is received
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("the page loaded")
		//calls the ServeHTTP method of the next handler to continue processing the request
		next.ServeHTTP(w, r)
	})
}

//NoSurf provides a middleware that prevents Cross-Site Request Forgery (CSRF) attacks by verifying
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{HttpOnly: true, Path: "/", Secure: false, SameSite: http.SameSiteLaxMode})
	return csrfHandler
}
