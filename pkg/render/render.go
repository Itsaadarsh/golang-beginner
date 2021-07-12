package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTmpl(w http.ResponseWriter, tmpl string) {

	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error while parsing template : ", err)
		return
	}
}
