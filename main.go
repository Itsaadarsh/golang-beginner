package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":3000"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTmpl(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTmpl(w, "about.page.html")
}

func renderTmpl(w http.ResponseWriter, tmpl string) {

	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error while parsing template : ", err)
		return
	}
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	http.ListenAndServe(portNumber, nil)
}
