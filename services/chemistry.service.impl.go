package services

import (
	"context"
	"errors"
	"kietchung/request"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"kietchung/models"
)

/*
@Author: DevProblems(Sarang Kumar)
@YTChannel: https://www.youtube.com/channel/UCVno4tMHEXietE3aUTodaZQ
*/
type ChemistryServiceImpl struct {
	refDocumentCollection *mongo.Collection
	chemistryCollection   *mongo.Collection
	ctx                   context.Context
}

func NewUserService(chemistryCollection *mongo.Collection, refDocument *mongo.Collection, ctx context.Context) ChemistryService {
	return &ChemistryServiceImpl{
		refDocumentCollection: refDocument,
		chemistryCollection:   chemistryCollection,
		ctx:                   ctx,
	}
}

func (c *ChemistryServiceImpl) GetMaterialUrl(chemistry *request.GetChemistryReq) ([]*models.Chemistry, error) {
	filter := bson.M{}
	var res []*models.Chemistry
	if chemistry.TypeChemical != "" {
		filter["type_chemical"] = chemistry.TypeChemical
	}

	if chemistry.GroupName != "" {
		filter["group_name"] = chemistry.GroupName
	}

	if chemistry.TypeSpectrum != "" {
		filter["type_spectrum"] = chemistry.TypeSpectrum
	}

	if chemistry.Chemical != "" {
		filter["chemical"] = chemistry.Chemical
	}

	cursor, err := c.chemistryCollection.Find(c.ctx, filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(c.ctx) {
		var chemistryRes models.Chemistry
		err := cursor.Decode(&chemistryRes)
		if err != nil {
			return nil, err
		}
		res = append(res, &chemistryRes)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(c.ctx)

	if len(res) == 0 {
		return nil, errors.New("documents not found")
	}
	return res, err
}

func (c *ChemistryServiceImpl) GetReferenceDocument(refDoc *request.GetRefDocument) ([]*models.ReferenceDocument, error) {
	filter := bson.M{}
	var res []*models.ReferenceDocument
	if refDoc.Type != "" {
		filter["type"] = refDoc.Type
	}

	cursor, err := c.refDocumentCollection.Find(c.ctx, filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(c.ctx) {
		var chemistryRes models.ReferenceDocument
		err := cursor.Decode(&chemistryRes)
		if err != nil {
			return nil, err
		}
		res = append(res, &chemistryRes)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(c.ctx)

	if len(res) == 0 {
		return nil, errors.New("documents not found")
	}
	return res, err
}

func (c *ChemistryServiceImpl) FixAkan() {
	filter := bson.M{}
	filter["group_name"] = "Amin"

	update := bson.M{}
	update["group_name"] = "Amine"
	_, _ = c.chemistryCollection.UpdateMany(c.ctx, filter, update)
}

type MenuResponse struct {
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Children []*MenuResponse `json:"children"`
}

func (c *ChemistryServiceImpl) CreateChildren(typeChemistry string) []*MenuResponse {
	filter := bson.M{}
	filter["type_chemical"] = typeChemistry
	cursor, err := c.chemistryCollection.Find(c.ctx, filter)
	if err != nil {
		return nil
	}

	var groupNameMap = make(map[string][]string)
	var chemicalMap = make(map[string][]string)

	for cursor.Next(c.ctx) {
		var chemistryRes models.Chemistry
		err := cursor.Decode(&chemistryRes)
		if err != nil {
			return nil
		}

		groupNameMap[chemistryRes.GroupName] = append(groupNameMap[chemistryRes.GroupName], chemistryRes.Chemical)
		chemicalMap[chemistryRes.Chemical] = append(chemicalMap[chemistryRes.Chemical], chemistryRes.TypeSpectrum)
	}

	var res []*MenuResponse

	if err := cursor.Err(); err != nil {
		return nil
	}

	for key := range groupNameMap {
		res = append(res, &MenuResponse{
			ID:   key,
			Name: key,
		})
	}

	for _, value := range res {
		listValue := groupNameMap[value.Name]
		for _, v := range listValue {
			finaleValue := &MenuResponse{
				ID:   v,
				Name: v,
			}

			for _, chemisVaue := range chemicalMap[v] {
				finaleValue.Children = append(finaleValue.Children, &MenuResponse{
					ID:   chemisVaue,
					Name: chemisVaue,
				})
			}
			value.Children = append(value.Children, finaleValue)
		}
	}

	cursor.Close(c.ctx)

	if len(res) == 0 {
		return nil
	}

	return res
}

func (c *ChemistryServiceImpl) GetMenu(req *request.GetMenu) ([]*MenuResponse, error) {

	var sampleRes []*MenuResponse
	sampleRes = append(sampleRes, &MenuResponse{
		ID:       "HYDROCACBON",
		Name:     "Hydrocacbon",
		Children: c.CreateChildren("HYDROCACBON"),
	})

	sampleRes = append(sampleRes, &MenuResponse{
		ID:       "HYDROCACBON_DERIVATIVE",
		Name:     "Dẫn Xuất Hydrocacbon",
		Children: c.CreateChildren("HYDROCACBON_DERIVATIVE"),
	})

	return sampleRes, nil
}
