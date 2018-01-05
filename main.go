package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joyent/triton-service-groups/templates"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/v1/tsg").Subrouter()

	sub.Methods("POST").Path("/templates").HandlerFunc(templates_v1.Create)
	sub.Methods("GET").Path("/templates/{name}").HandlerFunc(templates_v1.Get)
	sub.Methods("PUT").Path("/templates/{name}").HandlerFunc(templates_v1.Update)
	sub.Methods("DELETE").Path("/templates/{name}").HandlerFunc(templates_v1.Delete)

	log.Fatal(http.ListenAndServe(":3000", router))
}