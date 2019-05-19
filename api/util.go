package main

import (
	"fmt"

	"github.com/google/uuid"
)

func GetProductID(name string) (id uuid.UUID) {
	err := db.QueryRow(fmt.Sprintf(qGetRoleId, name)).Scan(&id)
	if err != nil {

		return
	}
	return id
}
