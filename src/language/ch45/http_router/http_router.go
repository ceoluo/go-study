package main

import (
	"fmt"
	httprouter "github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main()  {
	router := httprouter.New()

	router.GET("/", index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Hello(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	_, _ = fmt.Fprintf(writer, "Hello, %s", params.ByName("name"))
}

func index(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	_, _ = fmt.Fprint(writer, "Welcome!\n")
}