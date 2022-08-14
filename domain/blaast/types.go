package blaast

type CreateBlaastInput struct {
	ReceiverID string `json:"receiver_id"`
	SenderID   string `json:"sender_id"`
	Text       string `json:"text"`
}

type Blaast struct {
	BlaastID   string `json:"blaast_id"`
	ReceiverID string `json:"receiver_id"`
	SenderID   string `json:"sender_id"`
	Text       string `json:"text"`
	Status     string `json:"status"`
	CreatedAt  string `json:"created_at"`
}
