package database

import (
	"context"
	"hack-mit/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApplicantCollection struct {
	collection *mongo.Collection
}

func GetApplicantCollection() *ApplicantCollection {
	applicant := HackMIT.Collection("applicant")
	return &ApplicantCollection{collection: applicant}
}

func (ac *ApplicantCollection) AddNewApplicant(ctx context.Context, newApplicant *models.Applicant) (*models.Applicant, error) {
	if len(newApplicant.ID) == 0 {
		newApplicant.ID = primitive.NewObjectID().Hex()
	}
	_, err := ac.collection.InsertOne(ctx, newApplicant)
	if err != nil {
		return nil, err
	}
	return newApplicant, nil
}
