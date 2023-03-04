package main

import (
	"fmt"
	"github.com/KQW/my_page/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.AboutPage)
	fmt.Println("Starting apllication on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
