package repository

import (
	"golang-clean-architecture/internal/entity"

	"github.com/sirupsen/logrus"
)

type FoodRepository struct {
	Repository[entity.Food]
	Log *logrus.Logger
}

func NewFoodRepository(log *logrus.Logger) *FoodRepository {
	return &FoodRepository{
		Log: log,
	}
}
