package main

/*
All endpoints relevant to users
Prefix: /v1/users
*/

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type GetUserResponse struct {
	ID        uuid.UUID
	Username  string
	Email     string
	Firstname string
	Lastname  string
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	code, errMsg, claims := AuthorizeAndRefresh(w, r)
	if code != 0 {
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(errMsg)
		return
	}
	var user GetUserResponse
	user.ID = claims.UID
	user.Username = claims.Username

	err := db.QueryRow(
		fmt.Sprintf(
			qGetUser, user.ID,
		),
	).Scan(&user.Email, &user.Firstname, &user.Lastname)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{1501, "Internal Server Error"})
		return
	}

	json.NewEncoder(w).Encode(user)
}
