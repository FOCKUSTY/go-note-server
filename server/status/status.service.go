package status

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	Database *mongo.Client
}

func (service Service) Get() string {
	return "Not status yet."
}

func (service Service) GetDatabase() string {
	err := service.Database.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println(err)
		return "some error in database."
	}

	return "Status OK!"
}

func CreateService(Database *mongo.Client) *Service {
	return &Service{Database}
}
