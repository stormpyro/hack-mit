package fue

import (
	"encoding/json"
	"hack-mit/constants"
	"hack-mit/models"

	"github.com/labstack/echo/v4"
)

type FueHandler struct {
	FueUsecase Usecase
}

func NewFueHandler(r *echo.Group, us Usecase) {
	handler := &FueHandler{
		FueUsecase: us,
	}
	r.POST("/applicant", handler.AddNewApplicant)
	r.POST("/ground", handler.AddNewGround)
	r.POST("/landtitling", handler.AddNewLandTitling)
	r.POST("/licenceapplication", handler.AddNewLicence)
}

// @Summary Add a new applicant
// @Description	Add a new applicant
// @Tags	Del solicitante
// @Accept	json
// @Produce	json
// @Param applicant	body	models.Applicant	true	"New applicant"
// @Success	200	{object}	models.ResponseSuccess
// @Failure	500	{object}	models.ResponseError
// @Router	/api/v1/applicant	[post]
func (h *FueHandler) AddNewApplicant(ctx echo.Context) error {
	applicant := models.Applicant{}
	errDecode := json.NewDecoder(ctx.Request().Body).Decode(&applicant)
	if errDecode != nil {
		return constants.CustomError(500, errDecode.Error())
	}
	success, errAdd := h.FueUsecase.AddNewApplicant(ctx.Request().Context(), &applicant)
	if errAdd != nil {
		return errAdd
	}
	return ctx.JSON(200, success)
}

// @Summary Add a new ground
// @Description	Add a new ground
// @Tags	Del terreno
// @Accept	json
// @Produce	json
// @Param applicant	body	models.Ground	true	"New ground"
// @Success	200	{object}	models.ResponseSuccess
// @Failure	500	{object}	models.ResponseError
// @Router	/api/v1/ground	[post]
func (h *FueHandler) AddNewGround(ctx echo.Context) error {
	ground := models.Ground{}
	errDecode := json.NewDecoder(ctx.Request().Body).Decode(&ground)
	if errDecode != nil {
		return constants.CustomError(500, errDecode.Error())
	}
	success, errAdd := h.FueUsecase.AddNewGround(ctx.Request().Context(), &ground)
	if errAdd != nil {
		return errAdd
	}
	return ctx.JSON(200, success)
}

// @Summary Add a new land titling
// @Description	Add a new land titling
// @Tags	De la titulación del predio
// @Accept	json
// @Produce	json
// @Param applicant	body	models.Ground	true	"New land titling"
// @Success	200	{object}	models.ResponseSuccess
// @Failure	500	{object}	models.ResponseError
// @Router	/api/v1/landtitling	[post]
func (h *FueHandler) AddNewLandTitling(ctx echo.Context) error {
	landTitling := models.LandTitling{}
	errDecode := json.NewDecoder(ctx.Request().Body).Decode(&landTitling)
	if errDecode != nil {
		return constants.CustomError(500, errDecode.Error())
	}
	success, errAdd := h.FueUsecase.AddNewLandTitling(ctx.Request().Context(), &landTitling)
	if errAdd != nil {
		return errAdd
	}
	return ctx.JSON(200, success)
}

// @Summary Add a new licence application
// @Description	Add a new request to get a licence application
// @Tags	Solicitud de licencia de edificación
// @Accept	json
// @Produce	json
// @Param applicant	body	models.LicenceApplication	true	"New licence application"
// @Success	200	{object}	models.ResponseSuccess
// @Failure	500	{object}	models.ResponseError
// @Router	/api/v1/licenceapplication	[post]
func (h *FueHandler) AddNewLicence(ctx echo.Context) error {
	licence := models.LicenceApplication{}
	errDecode := json.NewDecoder(ctx.Request().Body).Decode(&licence)
	if errDecode != nil {
		return constants.CustomError(500, errDecode.Error())
	}
	success, errAdd := h.FueUsecase.AddNewLicence(ctx.Request().Context(), &licence)
	if errAdd != nil {
		return errAdd
	}
	return ctx.JSON(200, success)
}
