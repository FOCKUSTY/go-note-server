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
	fmt.Println(client)
	collection := client.Database.Collection("note")
	fmt.Println(collection)
	fmt.Println("Collection create")
}

func CreateNote(note NoteCreate) *Note {
	createData := CreateTable()

	return &Note{
		Id:         createData.id,
		CreatedAt:  createData.createdAt,
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
