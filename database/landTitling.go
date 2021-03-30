package database

import (
	"context"
	"hack-mit/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LandTitlingCollection struct {
	collection *mongo.Collection
}

func GetLandTitlingCollection() *LandTitlingCollection {
	landTitling := HackMIT.Collection("landTitling")
	return &LandTitlingCollection{collection: landTitling}
}

func (lt *LandTitlingCollection) AddNewLandTitling(ctx context.Context, newLandTitling *models.LandTitling) (*models.LandTitling, error) {
	if len(newLandTitling.ID) > 0 {
		newLandTitling.ID = primitive.NewObjectID().Hex()
	}
	_, err := lt.collection.InsertOne(ctx, newLandTitling)
	if err != nil {
		return nil, err
	}
	return newLandTitling, nil
}
