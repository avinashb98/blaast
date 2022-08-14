package blaast

type service struct {
	r Repository
}

func New(r Repository) Service {
	return service{r: r}
}

func (s service) CreateBlaast(blaast *CreateBlaastInput) (*Blaast, error) {
	return s.r.CreateBlaast(blaast.ReceiverID, blaast.SenderID, blaast.Text)
}

func (s service) GetByID(blaastId string) (*Blaast, error) {
	return s.r.GetByID(blaastId)
}
