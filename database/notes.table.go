package database

import (
	"context"
	"fmt"
	"main/env"

	"go.mongodb.org/mongo-driver/mongo"
)

type Note struct {
	Id         string
	CreatedAt  string
	Title      string
	Body       string
	Attacments []string
}

type NoteCreate struct {
	Title string
	Body  string
}

func init() {
	client := CreateClient(env.Get("MONGO_URL"))
	client.Database.CreateCollection(context.TODO(), "note")
	fmt.Println("Collection create")
}

func CreateNote(note NoteCreate) *Note {
	createData := CreateTable()

	return &Note{
		Id:         createData.Id,
		CreatedAt:  createData.CreatedAt,
		Title:      note.Title,
		Body:       note.Body,
		Attacments: []string{},
	}
}

func (note Note) Insert() (*mongo.InsertOneResult, error) {
	client := CreateClient(env.Get("MONGO_URL"))
	collection := client.Database.Collection("note")

	return collection.InsertOne(context.TODO(), note)
}
