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

type UserFlags struct {
	FieldName string `json:"name"`
	Set       bool   `json:"set"`
}

type GetUserResponse struct {
	ID              uuid.UUID
	Username        string                `json:"username"`
	Email           string                `json:"email"`
	Firstname       string                `json:"firstname"`
	Lastname        string                `json:"lastname"`
	EducationFlags  map[string]*UserFlags `json:"education_flags"`
	ExperienceFlags map[string]*UserFlags `json:"experience_flags"`
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

	var e_eu_applied_before, e_eu_funded_before, e_eu_subsidy_before, e_wrote_projects, e_enlish_fluent, e_edu_writing_projects, e_edu_project_mngmt, e_edu_financial_reporting, e_edu_eu_standards, e_edu_market_emerge bool

	err := db.QueryRow(
		fmt.Sprintf(
			qGetUser, user.ID,
		),
	).Scan(&user.Email, &user.Firstname, &user.Lastname, &e_eu_applied_before, &e_eu_funded_before, &e_eu_subsidy_before, &e_wrote_projects, &e_enlish_fluent, &e_edu_writing_projects, &e_edu_project_mngmt, &e_edu_financial_reporting, &e_edu_eu_standards, &e_edu_market_emerge)

	user.ExperienceFlags = make(map[string]*UserFlags)
	user.EducationFlags = make(map[string]*UserFlags)
	user.ExperienceFlags["e_eu_applied_before"] = new(UserFlags)
	user.ExperienceFlags["e_eu_applied_before"].FieldName = "Da li ste aplicirali za EU sredstva?"
	user.ExperienceFlags["e_eu_applied_before"].Set = e_eu_applied_before

	user.ExperienceFlags["e_eu_funded_before"] = new(UserFlags)
	user.ExperienceFlags["e_eu_funded_before"].FieldName = "Da li ste nekad dobili sredstva iz EU fondova?"
	user.ExperienceFlags["e_eu_funded_before"].Set = e_eu_funded_before

	user.ExperienceFlags["e_eu_subsidy_before"] = new(UserFlags)
	user.ExperienceFlags["e_eu_subsidy_before"].FieldName = "Da li ste koristili subvencionisane kredite iz EU fondova?"
	user.ExperienceFlags["e_eu_subsidy_before"].Set = e_eu_subsidy_before

	user.ExperienceFlags["e_wrote_projects"] = new(UserFlags)
	user.ExperienceFlags["e_wrote_projects"].FieldName = "Da li znate da pišete projektne aplikacije?"
	user.ExperienceFlags["e_wrote_projects"].Set = e_wrote_projects

	user.ExperienceFlags["e_enlish_fluent"] = new(UserFlags)
	user.ExperienceFlags["e_enlish_fluent"].FieldName = "Da li neko u firmi priča engleski?"
	user.ExperienceFlags["e_enlish_fluent"].Set = e_enlish_fluent

	user.EducationFlags["e_edu_writing_projects"] = new(UserFlags)
	user.EducationFlags["e_edu_writing_projects"].FieldName = "Pisanje projekata"
	user.EducationFlags["e_edu_writing_projects"].Set = e_edu_writing_projects

	user.EducationFlags["e_edu_project_mngmt"] = new(UserFlags)
	user.EducationFlags["e_edu_project_mngmt"].FieldName = "Vođenje projekata"
	user.EducationFlags["e_edu_project_mngmt"].Set = e_edu_project_mngmt

	user.EducationFlags["e_edu_financial_reporting"] = new(UserFlags)
	user.EducationFlags["e_edu_financial_reporting"].FieldName = "Finansijsko izveštavanje projekata"
	user.EducationFlags["e_edu_financial_reporting"].Set = e_edu_financial_reporting

	user.EducationFlags["e_edu_eu_standards"] = new(UserFlags)
	user.EducationFlags["e_edu_eu_standards"].FieldName = "Usklađivanje sa standardima EU"
	user.EducationFlags["e_edu_eu_standards"].Set = e_edu_eu_standards

	user.EducationFlags["e_edu_market_emerge"] = new(UserFlags)
	user.EducationFlags["e_edu_market_emerge"].FieldName = "Izlazak na nova tržišta"
	user.EducationFlags["e_edu_market_emerge"].Set = e_edu_market_emerge

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{1501, "Internal Server Error"})
		return
	}

	json.NewEncoder(w).Encode(user)
}

// type CreateUserRequest struct {
// 	Username  string `json:"username"`
// 	Email     string `json:"email"`
// 	Password  string `json:"password"`
// 	Firstname string `json:"firstname"`
// 	Lastname  string `json:"lastname"`
// }

// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	var request CreateUserRequest
// 	err := json.NewDecoder(r.Body).Decode(&request)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	err := db.QueryRow(
// 		fmt.Sprintf(
// 			qCreateUser, user.ID,
// 		),
// 	).Scan(&user.Email, &user.Firstname, &user.Lastname)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(ErrorResponse{1501, "Internal Server Error"})
// 		return
// 	}

// 	json.NewEncoder(w).Encode(user)
// }

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
