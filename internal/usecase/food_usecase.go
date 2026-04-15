package usecase

import (
	"golang-clean-architecture/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type FoodUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	FoodRepository *repository.FoodRepository
}

func NewFoodUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate,
	foodRepository *repository.FoodRepository) *FoodUseCase {
	return &FoodUseCase{
		DB:             db,
		Log:            log,
		Validate:       validate,
		FoodRepository: foodRepository,
	}
}
