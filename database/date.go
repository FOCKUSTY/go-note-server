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
	Id        string
	CreatedAt string
}

func CreateTable() *CreateTableData {
	return &CreateTableData{
		Id:        GenerateId(),
		CreatedAt: time.Now().Format(time.RFC3339),
	}
}
