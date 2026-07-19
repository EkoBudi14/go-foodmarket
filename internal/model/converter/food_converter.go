package converter

import (
	"golang-clean-architecture/internal/entity"
	"golang-clean-architecture/internal/model"
)

func FoodToResponse(food *entity.Food) *model.FoodResponse {
	return &model.FoodResponse{
		ID:          food.ID,
		Name:        food.Name,
		Description: food.Description,
		Ingredients: food.Ingredients,
		Price:       food.Price,
		Rate:        food.Rate,
		Types:       food.Types,
		PicturePath: food.PicturePath,
	}
}

func FoodToEvent(food *entity.Food) *model.FoodEvent {
	return &model.FoodEvent{
		ID:          food.ID,
		Name:        food.Name,
		Description: food.Description,
		Ingredients: food.Ingredients,
		Price:       food.Price,
		Rate:        food.Rate,
		Types:       food.Types,
		PicturePath: food.PicturePath,
		CreatedAt:   food.CreatedAt,
		UpdatedAt:   food.UpdatedAt,
	}
}
