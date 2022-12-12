package controllers

import (
	"github.com/gin-gonic/gin"
	"kietchung/request"
	"kietchung/services"
	"net/http"
)

type ChemistryController struct {
	ChemistryService services.ChemistryService
}

func New(chemistryService services.ChemistryService) ChemistryController {
	return ChemistryController{
		ChemistryService: chemistryService,
	}
}

func (uc *ChemistryController) GetMaterialUrl(ctx *gin.Context) {
	var getChemistryReq request.GetChemistryReq
	err := ctx.ShouldBindQuery(&getChemistryReq)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	chemistries, err := uc.ChemistryService.GetMaterialUrl(&getChemistryReq)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, chemistries)
}

func (uc *ChemistryController) GetReferenceDocument(ctx *gin.Context) {
	var getRefDocument request.GetRefDocument
	err := ctx.ShouldBindQuery(&getRefDocument)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	refDocument, err := uc.ChemistryService.GetReferenceDocument(&getRefDocument)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, refDocument)
}

func (uc *ChemistryController) RegisterUserRoutes(rg *gin.RouterGroup) {
	chemistryRoute := rg.Group("/chemistry")
	chemistryRoute.GET("/get-material", uc.GetMaterialUrl)
}
