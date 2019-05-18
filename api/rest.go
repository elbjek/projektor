package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ErrorResponse struct {
	Code    int
	Message string
}

func RunAPI() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(JsonMiddleware)

	router.HandleFunc("/login", Login).Methods("POST")

	log.Fatal(http.ListenAndServe(":80", router))
}

func JsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
