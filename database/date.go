package database

import (
	env "main/env"
	"strconv"
	"time"
)

func GenerateId() string {
	timeStamp := env.Get("TIME_STAMP")

	integer, err := strconv.Atoi(timeStamp)

	if err != nil {
		return strconv.Itoa(int(time.Now().Unix()))
	}

	id := int(time.Now().Unix()) - integer

	return strconv.Itoa(id)
}

type CreateTableData struct {
	id        string
	createdAt string
}

func CreateTable() *CreateTableData {
	return &CreateTableData{
		id:        GenerateId(),
		createdAt: time.Now().Format(time.RFC3339),
	}
}
