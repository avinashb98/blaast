package blaast

import (
	"github.com/avinashb98/blaast/domain/blaast"
)

type Repository interface {
	CreateBlaast(receiverID string, senderID string, text string) (*blaast.Blaast, error)
	GetByID(blaastId string) (*blaast.Blaast, error)
}
