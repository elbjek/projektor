package main

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("2eab7d68-fbec-4335-bdfc-0f6ecd44230b")

// Claims JWT specs
type Claims struct {
	UID      uuid.UUID `json:"user_id"`
	Username string    `json:"username"`
	jwt.StandardClaims
}

func HashAndSalt(pwd string) string {
	bpwd := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(bpwd, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func CheckPassword(pwd string, hash string) bool {
	bpwd := []byte(pwd)
	bhash := []byte(hash)
	err := bcrypt.CompareHashAndPassword(bhash, bpwd)
	if err != nil {
		return false
	}
	return true
}

// func CheckJwtAndRefresh(cookie *http.Cookie) {
// 	// We can obtain the session token from the requests cookies, which come with every request
// 	cookie, err := c.Cookie("token")
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			// If the cookie is not set, return an unauthorized status
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		// For any other type of error, return a bad request status
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// Get the JWT string from the cookie
// 	tknStr := cookie.Value

// 	// Initialize a new instance of `Claims`
// 	claims := &Claims{}

// 	// Parse the JWT string and store the result in `claims`.
// 	// Note that we are passing the key in this method as well. This method will return an error
// 	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
// 	// or if the signature does not match
// 	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
// 		return jwtKey, nil
// 	})
// 	if !tkn.Valid {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// }
