package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	//create template cache
	fmt.Println("this is pass 5")
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	//get requested template from the cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}
	//create buffer to capture the output of the executed template
	fmt.Println("this is pass 6")
	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("this is pass 77")
	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("this is pass 8")
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	//get all files which named *.page.tmpl from ./templates
	//files.Glob will base on the filepath,find the file that has specific name, it will return a slice
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	//range through all files ending with*.page.tmpl
	fmt.Println("this is pass 1")
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
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
		fmt.Println("this is pass 3")
	}
	fmt.Println("this is pass 4")
	return myCache, nil
}
