package systemmodels

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	Name   string `json:"name"`
	Client *mongo.Client
}
