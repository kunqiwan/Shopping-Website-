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

// NoSurf provides a middleware that prevents Cross-Site Request Forgery (CSRF) attacks by verifying
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad save and load session
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
