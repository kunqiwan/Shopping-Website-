package main

import (
	"fmt"
	"github.com/KQW/my_page/pkg/config"
	"github.com/KQW/my_page/pkg/handlers"
	"github.com/KQW/my_page/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	//InProduction - Production Model - false PRODUCTION Model -true
	app.InProduction = false

	//set the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	//create the cache of template
	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("template cache error")
	}
	app.TemplateCache = tc
	app.UseCache = false

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
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
