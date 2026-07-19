package model

type FoodEvent struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Ingredients string  `json:"ingredients"`
	Price       int     `json:"price"`
	Rate        float64 `json:"rate"`
	Types       string  `json:"types"`
	PicturePath string  `json:"picture_path"`
	CreatedAt   int64   `json:"created_at"`
	UpdatedAt   int64   `json:"updated_at"`
}

func (c *FoodEvent) GetId() string {
	return c.ID
}
