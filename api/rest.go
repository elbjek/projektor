package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type ErrorResponse struct {
	Code    int
	Message string
}

func RunAPI() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(JsonMiddleware)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"},
	})
	handler := c.Handler(router)

	router.HandleFunc("/login", Login).Methods("POST")
	router.HandleFunc("/users/v1/user", GetUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":80", handler))
}

func JsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
