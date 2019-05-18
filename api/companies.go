package main

/*
All endpoints relevant to companies
Prefix: /companies/v1
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lib/pq"

	"github.com/google/uuid"
)

type GetCompanyResponse struct {
	ID      uuid.UUID   `json:"id"`
	Name    string      `json:"name"`
	FieldId uuid.UUID   `json:"fieldId"`
	Tags    []uuid.UUID `json:"tags"`
}

func GetCompany(w http.ResponseWriter, r *http.Request) {
	code, errMsg, _ := AuthorizeAndRefresh(w, r)
	if code != 0 {
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	params := mux.Vars(r)
	companyId, err := uuid.Parse(params["companyId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{1701, "companyId must be valid uuid"})
		return
	}

	var company GetCompanyResponse
	company.ID = companyId
	err = db.QueryRow(
		fmt.Sprintf(
			qGetCompany, companyId,
		),
	).Scan(&company.Name, &company.FieldId, pq.Array(&company.Tags))
	if err != nil {
		log.Printf("1702: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{1501, "Internal Server Error"})
		return
	}

	json.NewEncoder(w).Encode(company)
}
