package blaast

type Service interface {
	CreateBlaast(*CreateBlaastInput) (*Blaast, error)
	GetByID(blaastId string) (*Blaast, error)
}

type Repository interface {
	CreateBlaast(receiverID string, senderID string, text string) (*Blaast, error)
	GetByID(blaastId string) (*Blaast, error)
}
