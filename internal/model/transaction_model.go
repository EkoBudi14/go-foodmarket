package model

type TransactionResponse struct {
	ID         string        `json:"id"`
	User       *UserResponse `json:"user"`
	Food       *FoodResponse `json:"food"`
	Quantity   int           `json:"quantity"`
	Total      int           `json:"total"`
	Status     string        `json:"status"`
	PaymentUrl string        `json:"payment_url"`
	CreatedAt  int64         `json:"created_at"`
	UpdatedAt  int64         `json:"updated_at"`
}
