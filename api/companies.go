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

	"github.com/google/uuid"
)

type CompanyTags struct {
	TagSize       string `json:"size"`
	TagLocation   string `json:"location"`
	TagInvestment string `json:"investment"`
}

type GetCompanyResponse struct {
	ID        uuid.UUID   `json:"id"`
	PIB       int         `json:"pib"`
	MIB       int         `json:"mib"`
	Name      string      `json:"name"`
	CreatedAt string      `json:"established"`
	Employees int         `json:"employees"`
	FieldId   string      `json:"fieldId"`
	Tags      CompanyTags `json:"tags"`
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
	).Scan(&company.Name, &company.CreatedAt, &company.Employees, &company.FieldId, &company.Tags.TagSize, &company.Tags.TagLocation, &company.Tags.TagInvestment, &company.PIB, &company.MIB)
	if err != nil {
		log.Printf("1702: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{1501, "Internal Server Error"})
		return
	}

	company.FieldId = FieldIDToName(company.FieldId)
	company.Tags.TagInvestment = TagInvestmentIDToName(company.Tags.TagInvestment)
	company.Tags.TagLocation = TagLocationIDToName(company.Tags.TagLocation)
	company.Tags.TagSize = TagSizeIDToName(company.Tags.TagSize)

	json.NewEncoder(w).Encode(company)
}
