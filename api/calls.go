package main

/*
All endpoints relevant to companies
Prefix: /calls/v1
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/lib/pq"

	"github.com/gorilla/mux"

	"github.com/google/uuid"
)

type Check struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Assist      string    `json:"assist_text"`
}

type Call struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Deadline    string    `json:"deadline"`
	Program     string    `json:"program"`
	Description string    `json:"description"`
}

type GetCallResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Deadline    string    `json:"deadline"`
	Program     string    `json:"program"`
	Description []string  `json:"description"`
	Checks      []Check   `json:"checks"`
}

func GetCall(w http.ResponseWriter, r *http.Request) {
	code, errMsg, _ := AuthorizeAndRefresh(w, r)
	if code != 0 {
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(errMsg)
		return
	}

	params := mux.Vars(r)
	callId, err := uuid.Parse(params["callId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{1701, "companyId must be valid uuid"})
		return
	}

	var call GetCallResponse
	var checkList []uuid.UUID
	call.ID = callId
	err = db.QueryRow(
		fmt.Sprintf(
			qGetCall, callId,
		),
	).Scan(&call.Title, &call.Deadline, &call.Program, pq.Array(&call.Description), pq.Array(&checkList))
	if err != nil {
		log.Printf("1802: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{1801, "Internal Server Error"})
		return
	}

	for _, check := range checkList {
		call.Checks = append(call.Checks, GetCheck(check))
	}

	json.NewEncoder(w).Encode(call)
}

type GetCallsResponse struct {
	Count int    `json:"count"`
	Calls []Call `json:"calls"`
}

func GetCalls(w http.ResponseWriter, r *http.Request) {
	var response GetCallsResponse

	rows, err := db.Query(
		qGetCalls,
	)
	if err != nil {
		log.Panicln(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{1805, "Internal Server Error"})
		return
	}

	var call Call
	for rows.Next() {
		call = Call{}
		rows.Scan(
			&call.ID, &call.Title, &call.Deadline, &call.Program, &call.Description,
		)
		response.Calls = append(response.Calls, call)
	}
	rows.Close()

	response.Count = len(response.Calls)

	json.NewEncoder(w).Encode(response)
}
