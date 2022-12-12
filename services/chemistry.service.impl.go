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
	filter["group_name"] = "Akan"

	update := bson.M{}
	update["group_name"] = "Ankan"
	_, _ = c.chemistryCollection.UpdateMany(c.ctx, filter, update)
}

func (c *ChemistryServiceImpl) GetMenu(req *request.GetMenu) ([]string, error) {
	if req.TypeChemical == "" {
		return []string{"Hydro Cacbon", "Dẫn xuất hydro các bon"}, nil
	}
	check := 1
	filter := bson.M{}
	filter["type_chemical"] = req.TypeChemical
	if req.GroupName != "" {
		filter["group_name"] = req.GroupName
		check = 2
		if req.Chemical != "" {
			filter["chemical"] = req.Chemical
			check = 3
		}
	}

	var res []string
	cursor, err := c.chemistryCollection.Find(c.ctx, filter)
	if err != nil {
		return nil, err
	}
	var checkMap = make(map[string]string)
	for cursor.Next(c.ctx) {
		var chemistryRes models.Chemistry
		err := cursor.Decode(&chemistryRes)
		if err != nil {
			return nil, err
		}

		if check == 1 {
			_, ok := checkMap[chemistryRes.GroupName]
			if !ok {
				res = append(res, chemistryRes.GroupName)
				checkMap[chemistryRes.GroupName] = ""
			}
		} else if check == 2 {
			_, ok := checkMap[chemistryRes.Chemical]
			if !ok {
				res = append(res, chemistryRes.Chemical)
				checkMap[chemistryRes.Chemical] = ""
			}
		} else {
			_, ok := checkMap[chemistryRes.TypeSpectrum]
			if !ok {
				res = append(res, chemistryRes.TypeSpectrum)
				checkMap[chemistryRes.TypeSpectrum] = ""
			}
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(c.ctx)

	if len(res) == 0 {
		return nil, errors.New("documents not found")
	}

	return res, nil
}
