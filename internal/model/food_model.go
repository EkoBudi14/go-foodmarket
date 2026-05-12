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

type FoodRequest struct {
	ID        string  `json:"id" query:"id" validate:"omitempty,max=100"`
	Limit     int     `json:"limit" query:"limit" validate:"omitempty,min=1"`
	Name      string  `json:"name" query:"name" validate:"omitempty,max=100"`
	Types     string  `json:"types" query:"types" validate:"omitempty,max=100"`
	PriceFrom int     `json:"price_from" query:"price_from" validate:"omitempty,min=0"`
	PriceTo   int     `json:"price_to" query:"price_to" validate:"omitempty,min=0"`
	RateFrom  float64 `json:"rate_from" query:"rate_from" validate:"omitempty,min=0"`
	RateTo    float64 `json:"rate_to" query:"rate_to" validate:"omitempty,min=0"`
}
