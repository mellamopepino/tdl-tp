package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type server struct{}

func home(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(`{"message": "hello world"}`))
}

func main() {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	homeRouter := apiRouter.PathPrefix("/home").Subrouter()
	homeRouter.HandleFunc("", home).Methods(http.MethodGet) // /api/home

	http.ListenAndServe(":8080", router)
}