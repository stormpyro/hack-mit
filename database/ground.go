package database

import (
	"context"
	"hack-mit/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type GroundCollection struct {
	collection *mongo.Collection
}

func GetGroundCollection() *GroundCollection {
	ground := HackMIT.Collection("ground")
	return &GroundCollection{collection: ground}
}

func (g *GroundCollection) AddNewGround(ctx context.Context, newGround *models.Ground) (*models.Ground, error) {
	if len(newGround.ID) == 0 {
		newGround.ID = primitive.NewObjectID().Hex()
	}
	_, err := g.collection.InsertOne(ctx, newGround)
	if err != nil {
		return nil, err
	}

	return newGround, nil

}
