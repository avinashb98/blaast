package blaast

import (
	"github.com/avinashb98/blaast/config"
	"github.com/avinashb98/blaast/domain/blaast"
	"github.com/avinashb98/blaast/errors"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

type mongoRepository struct {
	mongoConfig *config.Mongo
}

func New(config *config.Mongo) Repository {
	return mongoRepository{
		mongoConfig: config,
	}
}

func (m mongoRepository) CreateBlaast(receiverID string, senderID string, text string) (*blaast.Blaast, error) {
	newBlaast := NewBlaast(receiverID, senderID, text)
	err := mgm.Coll(newBlaast).Create(newBlaast)
	return newBlaast.ConvertForDomain(), err
}

func (m mongoRepository) GetByID(blaastId string) (*blaast.Blaast, error) {
	ctx := mgm.Ctx()
	foundBlaast := &Blaast{}
	err := mgm.Coll(foundBlaast).FirstWithCtx(ctx, bson.M{"blaastId": blaastId}, foundBlaast)
	if err != nil {
		if strings.Contains(err.Error(), "no documents") {
			return nil, &errors.NotFoundError{Message: "blaast not found"}
		}
		return nil, err
	}
	return foundBlaast.ConvertForDomain(), nil
}
