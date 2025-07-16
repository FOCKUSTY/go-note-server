package note

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	Database *mongo.Client
}

func (service *Service) Get() string {
	err := service.Database.Database("note").CreateCollection(context.Background(), "note")

	if err != nil {
		fmt.Println(err)
		return "database error #1"
	}

	collections, collectionErr := service.Database.Database("note").ListCollectionNames(context.Background(), struct{}{})

	if collectionErr != nil {
		fmt.Println(err)
		return "database error #1"
	}

	return strings.Join(collections, "\n")
}

func CreateService(Database *mongo.Client) *Service {
	return &Service{Database}
}
