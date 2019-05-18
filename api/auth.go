package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("2eab7d68-fbec-4335-bdfc-0f6ecd44230b")

const jwtTTL = 3600

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

type LoginCredentialsRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginCredentialsResponse struct {
	TTL   int    `json:"ttl"`
	Token string `json:"token"`
}

type Token struct {
	Name     string    `json:"name"`
	JWT      string    `json:"jwt"`
	TTL      time.Time `json:"ttl"`
	Redirect string    `json:"redirect"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Logging")
	var credentials LoginCredentialsRequest
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Println(err)
		log.Println(r.Body)
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
	expirationTime := time.Now().Add(jwtTTL * time.Second)
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

	log.Println("Logged in")
	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	var creds Token
	creds.Name = "token"
	creds.JWT = tokenString
	creds.TTL = expirationTime
	creds.Redirect = fmt.Sprintf("/user/%v", uid)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(creds)
	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "token",
	// 	Value:   tokenString,
	// 	Expires: expirationTime,
	// })
}

func AuthorizeAndRefresh(w http.ResponseWriter, r *http.Request) (code int, errMsg *ErrorResponse, claims *Claims) {
	errMsg = new(ErrorResponse)
	claims = &Claims{}
	// We can obtain the session token from the requests cookies, which come with every request
	// cookie, err := r.Cookie("token")
	// if err != nil {
	// 	if err == http.ErrNoCookie {
	// 		errMsg.Code = 2001
	// 		errMsg.Message = "No cookie in request"
	// 		return http.StatusUnauthorized, errMsg, claims
	// 	}
	// 	log.Println(err)
	// 	errMsg.Code = 2003
	// 	errMsg.Message = "Internal server error"
	// 	return http.StatusBadRequest, errMsg, claims
	// }

	// // Get the JWT string from the cookie
	// tknStr := cookie.Value
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		errMsg.Code = 2001
		errMsg.Message = "No authorization in the request"
		return http.StatusUnauthorized, errMsg, claims
	}

	tknStr := strings.Split(authHeader, " ")[1]

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !tkn.Valid {
		errMsg.Code = 2004
		errMsg.Message = "Token not valid"
		return http.StatusUnauthorized, errMsg, claims
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			errMsg.Code = 2005
			errMsg.Message = "JWT signature invalid"
			return http.StatusUnauthorized, errMsg, claims
		}
		log.Println(err)
		errMsg.Code = 2006
		errMsg.Message = "Error parsing JWT"
		return http.StatusBadRequest, errMsg, claims
	}

	// We ensure that a new token is not issued until enough time has elapsed
	// In this case, a new token will only be issued if the old token is within
	// 30 seconds of expiry. Otherwise, return a bad request status
	// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 60*time.Second {

	// 	// Now, create a new token for the current use, with a renewed expiration time
	// 	expirationTime := time.Now().Add(jwtTTL * time.Second)
	// 	claims.ExpiresAt = expirationTime.Unix()
	// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 	tokenString, err := token.SignedString(jwtKey)
	// 	if err != nil {
	// 		errMsg.Code = 2007
	// 		errMsg.Message = "Internal server error"
	// 		return http.StatusBadRequest, errMsg, claims
	// 	}

	// 	// Set the new token as the users `token` cookie
	// 	http.SetCookie(w, &http.Cookie{
	// 		Name:    "token",
	// 		Value:   tokenString,
	// 		Expires: expirationTime,
	// 	})

	// }

	return 0, errMsg, claims
}
