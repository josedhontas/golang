package main

import (
	"fmt"
	"net/http"
	"golang-crud/helper"
	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("Start server")

	routes := httprouter.New()
	routes.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome Home")
	})

	server := http.Server{Addr: "localhost:8000", Handler: routes}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
