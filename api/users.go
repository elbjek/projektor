package main

/*
All endpoints relevant to users
Prefix: /v1/users
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Username string
	email    string
	Role     uuid.UUID
}

type LoginCredentialsRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginCredentialsResponse struct {
	TTL   int    `json:"ttl"`
	Token string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials LoginCredentialsRequest
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{101, "Cannot unmarshal body, bad request"})
		return
	}

	// Attempt to retreive hash and uuid from the DB
	var uid uuid.UUID
	var hash string

	err = db.QueryRow(
		fmt.Sprintf(
			qGetUserUuidAndHash, credentials.Username,
		),
	).Scan(&uid, &hash)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(ErrorResponse{102, "User not found"})
		return
	}

	// Check password and if the hashses match
	valid := CheckPassword(credentials.Password, hash)
	if !valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// JWT expiration time
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		UID:      uid,
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		log.Printf("1003: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

// func GetUser(w http.ResponseWriter, r *http.Request)
