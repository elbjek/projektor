package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type ErrorResponse struct {
	Code    int
	Message string
}

func RunAPI() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(JsonMiddleware)

	handlers.AllowedOrigins([]string{"*"})

	// c := cors.New(cors.Options{
	// 	AllowedOrigins: []string{
	// 		"http://localhost:8080/*",
	// 		"http://localhost:8080",
	// 	},
	// 	// AllowedHeaders: []string{"location"},
	// })
	// handler := c.Handler(router)

	router.HandleFunc("/login", Login).Methods("POST")
	router.HandleFunc("/users/v1/user", GetUser).Methods("GET")
	router.HandleFunc("/users/v1/user/{userId}/companies", GetUserCompanies).Methods("GET")
	router.HandleFunc("/companies/v1/company/{companyId}", GetCompany).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}

func JsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// func CORSMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Add("Access-Control-Allow-Headers", "*")
// 		w.Header().Add("Access-Control-Allow-Methods", "*")
// 		w.Header().Add("Access-Control-Allow-Origin", "*")
// 		next.ServeHTTP(w, r)
// 	})
// }
