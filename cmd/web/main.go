package main

import (
	"net/http"

	"github.com/itsaadarsh/golang/pkg/handlers"
)

const portNumber = ":3000"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	http.ListenAndServe(portNumber, nil)
}
