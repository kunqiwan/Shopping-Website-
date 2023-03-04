package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}
func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	//get all files which named *.page.tmpl from ./templates
	//files.Glob will base on the filepath,find the file that has specific name, it will return a slice
	pages, err := filepath.Glob("./templates/*.pages.tmpl")
	if err != nil {
		return myCache, err
	}
	//range through all files ending with*.page.tmpl
	for _, page := range pages {
		//Base function will return the page name after ./templates/
		name := filepath.Base(page)
		// firstly, allocate the specific name to the html file we find,parse the file and allocate it to ts
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		//find the layout file
		matches, err := filepath.Glob("./templates/base.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			//parse all the base.layout.tmpl files and add them to the template set
			ts, err = ts.ParseGlob("./templates/base.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
