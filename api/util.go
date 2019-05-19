package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

func GetProductID(name string) (id uuid.UUID) {
	err := db.QueryRow(fmt.Sprintf(qGetRoleId, name)).Scan(&id)
	if err != nil {
		log.Println(err)
		return
	}
	return id
}

func FieldIDToName(id string) (name string) {
	err := db.QueryRow(fmt.Sprintf(qFieldIDToName, id)).Scan(&name)
	if err != nil {
		log.Println(err)
		return
	}
	return name
}

func TagLocationIDToName(id string) (name string) {
	err := db.QueryRow(fmt.Sprintf(qLocationIDToName, id)).Scan(&name)
	if err != nil {
		log.Println(err)
		return
	}
	return name
}

func TagSizeIDToName(id string) (name string) {
	err := db.QueryRow(fmt.Sprintf(qSizeIDToName, id)).Scan(&name)
	if err != nil {
		log.Println(err)
		return
	}
	return name
}

func TagInvestmentIDToName(id string) (name string) {
	err := db.QueryRow(fmt.Sprintf(qInvestmentIDToName, id)).Scan(&name)
	if err != nil {
		log.Println(err)
		return
	}
	return name
}

func GetCheck(id uuid.UUID) (check Check) {
	check.ID = id
	err := db.QueryRow(fmt.Sprintf(qGetCheck, id)).Scan(&check.PreChecked, &check.Description, &check.Assist)
	if err != nil {
		log.Println(err)
		return
	}
	return check
}
