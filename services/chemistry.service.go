package services

import (
	"kietchung/models"
	"kietchung/request"
)

type ChemistryService interface {
	GetMaterialUrl(chemistry *request.GetChemistryReq) ([]*models.Chemistry, error)
}
