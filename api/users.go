package main

/*
All endpoints relevant to users
Prefix: /users/v1
*/

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

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

type GetUserComapniesCompany struct {
	ID   uuid.UUID
	Name string
}

type GetUserComapniesResponse struct {
	Companies []GetUserComapniesCompany
	Count     int
}

func GetUserCompanies(w http.ResponseWriter, r *http.Request) {
	code, errMsg, claims := AuthorizeAndRefresh(w, r)
	if code != 0 {
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	params := mux.Vars(r)
	uid, err := uuid.Parse(params["userId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{1603, "userId must be valid uuid"})
		return
	}
	if uid != claims.UID {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// for now mocked to 1 company
	var company GetUserComapniesCompany
	var companies GetUserComapniesResponse
	err = db.QueryRow(
		fmt.Sprintf(
			qGetUserCompanies, uid,
		),
	).Scan(&company.ID, &company.Name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{1601, "User is not owner of any company object"})
		return
	}

	// mock
	companies.Companies = append(companies.Companies, company)
	companies.Count = 1

	json.NewEncoder(w).Encode(companies)
}
