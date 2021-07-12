package handlers

import (
	"net/http"

	"github.com/itsaadarsh/golang/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTmpl(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTmpl(w, "about.page.html")
}
