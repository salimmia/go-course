package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")

	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

var tc = make(map[string]*template.Template);

func RenderTemplate(w http.ResponseWriter, t string){
	var tmpl *template.Template;
	var err  error;

	_, inMap := tc[t];
	// check to see if we already have the template in our cache
	if !inMap{
		// need to create the template
		log.Println("Creating templates and adding in the cache")
		err = createTemplateCache(t);

		if err != nil{
			log.Println(err)
		}
	} else{
		// We have the template in the cache
		log.Println("Using cache template");
	}
	tmpl = tc[t];
	err = tmpl.Execute(w, nil);

	if err != nil{
		log.Println(err)
	}
}

func createTemplateCache(t string) error{
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.html",
	}
	tmpl, err := template.ParseFiles(templates...)

	if err != nil{
		return err
	}
	tc[t] = tmpl

	return nil
}