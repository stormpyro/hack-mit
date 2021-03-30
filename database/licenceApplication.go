package database

import (
	"context"
	"hack-mit/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LicenceApplicationCollection struct {
	collection *mongo.Collection
}

func GetLicenceApplicationCollection() *LicenceApplicationCollection {
	licenceApplication := HackMIT.Collection("licenceApplication")
	return &LicenceApplicationCollection{collection: licenceApplication}
}

func (la *LicenceApplicationCollection) AddNewLicence(ctx context.Context, request *models.LicenceApplication) (*models.LicenceApplication, error) {
	if len(request.ID) > 0 {
		request.ID = primitive.NewObjectID().Hex()
	}
	_, err := la.collection.InsertOne(ctx, request)
	if err != nil {
		return nil, err
	}
	return request, nil
}
