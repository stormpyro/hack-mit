package fue

import (
	"context"
	"hack-mit/models"
)

type Usecase interface {
	AddNewApplicant(ctx context.Context, newApplicant *models.Applicant) (*models.ResponseSuccess, error)
	AddNewGround(ctx context.Context, newGround *models.Ground) (*models.ResponseSuccess, error)
	AddNewLandTitling(ctx context.Context, newLandTitling *models.LandTitling) (*models.ResponseSuccess, error)
	AddNewLicence(ctx context.Context, newLicenceApplication *models.LicenceApplication) (*models.ResponseSuccess, error)
}
