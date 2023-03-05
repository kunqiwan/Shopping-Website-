package render

import (
	"bytes"
	"fmt"
	"github.com/KQW/my_page/pkg/config"
	"github.com/KQW/my_page/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

// RenderTemplate renders templates
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	//in main method, we create cache,and save it in appConfig Struct,just call it there
	tc := app.TemplateCache
	//get requested template from the cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("err in template parsing")
	}
	//create buffer to capture the output of the executed template
	buf := new(bytes.Buffer)
	//causes the template to be executed with no additional data, and the output is written to the buffer
	//pass the template data to the buffer
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}
	//render the template
	//the contents of the buffer are written to the HTTP response using the WriteTo() method
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	//create a map,key is the name of page,value is the template
	myCache := map[string]*template.Template{}
	//get all files which named *.page.tmpl from ./templates
	//files.Glob will base on the filepath,find the file that has specific name, it will return a slice
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	//range through all files ending with*.page.tmpl
	fmt.Println("pages:", pages)
	for _, page := range pages {
		//Base function will return the page name after ./templates/
		name := filepath.Base(page)
		// firstly, allocate the specific name to the html file we find,parse the file and allocate it to ts
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		//find the layout file
		fmt.Println("this is pass 2")
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			//parse all the base.layout.tmpl files and add them to the template set
			//a template set is a collection of related templates
			//a text-based format that can include placeholders for dynamic data
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
