package blaast

import (
	"github.com/avinashb98/blaast/domain/blaast"
	"github.com/google/uuid"
	"github.com/kamva/mgm/v3"
)

type Blaast struct {
	mgm.DefaultModel `bson:",inline"`
	BlaastID         string `json:"blaastId" bson:"blaastId"`
	ReceiverID       string `json:"receiverID" bson:"receiverID"`
	SenderID         string `json:"senderID" bson:"senderID"`
	Text             string `json:"text" bson:"text"`
	Status           Status `json:"status" bson:"status"`
}

func (model *Blaast) CollectionName() string {
	return "blaast"
}

func (model *Blaast) ConvertForDomain() *blaast.Blaast {
	return &blaast.Blaast{
		BlaastID:   model.BlaastID,
		ReceiverID: model.ReceiverID,
		SenderID:   model.SenderID,
		Text:       model.Text,
		Status:     string(model.Status),
		CreatedAt:  model.CreatedAt.String(),
	}
}

func NewBlaast(receiverID string, senderID string, text string) *Blaast {
	return &Blaast{
		BlaastID:   uuid.New().String(),
		ReceiverID: receiverID,
		SenderID:   senderID,
		Text:       text,
		Status:     Active,
	}
}
