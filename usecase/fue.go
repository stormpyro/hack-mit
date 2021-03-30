package usecase

import (
	"context"
	"hack-mit/constants"
	"hack-mit/database"
	"hack-mit/models"
)

type FueUsecase struct {
	applicant          *database.ApplicantCollection
	ground             *database.GroundCollection
	landTitling        *database.LandTitlingCollection
	licenceApplication *database.LicenceApplicationCollection
}

func NewFueUsecase() *FueUsecase {
	applicant := database.GetApplicantCollection()
	ground := database.GetGroundCollection()
	landTitling := database.GetLandTitlingCollection()
	licenceApplication := database.GetLicenceApplicationCollection()
	return &FueUsecase{
		applicant:          applicant,
		ground:             ground,
		landTitling:        landTitling,
		licenceApplication: licenceApplication}
}

func (h *FueUsecase) AddNewApplicant(ctx context.Context, newApplicant *models.Applicant) (*models.ResponseSuccess, error) {
	applicant, err := h.applicant.AddNewApplicant(nil, newApplicant)
	if err != nil {
		return nil, constants.CustomError(500, err.Error())
	}
	if applicant == nil {
		return nil, constants.ErrorDatabase()
	}

	return &models.ResponseSuccess{Message: "Applicant added"}, nil

}

func (h *FueUsecase) AddNewGround(ctx context.Context, newGround *models.Ground) (*models.ResponseSuccess, error) {
	ground, err := h.ground.AddNewGround(nil, newGround)
	if err != nil {
		return nil, constants.CustomError(500, err.Error())
	}
	if ground == nil {
		return nil, constants.ErrorDatabase()
	}
	return &models.ResponseSuccess{Message: "Ground added"}, nil
}

func (h *FueUsecase) AddNewLandTitling(ctx context.Context, newLandTitling *models.LandTitling) (*models.ResponseSuccess, error) {
	landTitling, err := h.landTitling.AddNewLandTitling(nil, newLandTitling)
	if err != nil {
		return nil, constants.CustomError(500, err.Error())
	}
	if landTitling == nil {
		return nil, constants.ErrorDatabase()
	}
	return &models.ResponseSuccess{Message: "Land titling added"}, nil
}

func (h *FueUsecase) AddNewLicence(ctx context.Context, newLicenceApplication *models.LicenceApplication) (*models.ResponseSuccess, error) {
	licenceApplication, err := h.licenceApplication.AddNewLicence(nil, newLicenceApplication)
	if err != nil {
		return nil, constants.CustomError(500, err.Error())
	}
	if licenceApplication == nil {
		return nil, constants.ErrorDatabase()
	}
	return &models.ResponseSuccess{Message: "Licence application added"}, nil
}
