package services

import (
	"kietchung/models"
	"kietchung/request"
)

type ChemistryService interface {
	GetMaterialUrl(chemistry *request.GetChemistryReq) ([]*models.Chemistry, error)
	GetReferenceDocument(chemistry *request.GetRefDocument) ([]*models.ReferenceDocument, error)
	CreateChildren(typeChemistry string) []*MenuResponse
	GetMenu(req *request.GetMenu) ([]*MenuResponse, error)
	FixAkan()
}
