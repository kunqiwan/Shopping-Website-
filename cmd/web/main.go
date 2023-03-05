package main

import (
	"fmt"
	"github.com/KQW/my_page/pkg/config"
	"github.com/KQW/my_page/pkg/handlers"
	"github.com/KQW/my_page/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	//create the cache of template
	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("template cache error")
	}
	app.TemplateCache = tc
	//use app configure we get to create new repository instance
	repo := handlers.NewRepository(&app)
	//update the repository variable we set in handler class
	handlers.NewHandler(repo)
	render.NewTemplates(&app)
	// the http package will automatically pass the http.ResponseWriter and *http.Request arguments
	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.AboutPage)
	fmt.Println("Starting apllication on port %s", portNumber)
	//_ = http.ListenAndServe(portNumber, nil)
	fmt.Println("paole srv")
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	fmt.Println("paole handler")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
