package model

type FoodResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Ingredients string  `json:"ingredients"`
	Price       int     `json:"price"`
	Rate        float64 `json:"rate"`
	Types       string  `json:"types"`
	PicturePath string  `json:"picture_path"`
}
