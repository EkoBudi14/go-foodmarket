package usecase

import (
	"context"
	"golang-clean-architecture/internal/entity"
	"golang-clean-architecture/internal/gateway/messaging"
	"golang-clean-architecture/internal/model"
	"golang-clean-architecture/internal/model/converter"
	"golang-clean-architecture/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type FoodUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	FoodRepository *repository.FoodRepository
	FoodProducer   *messaging.FoodProducer
}

func NewFoodUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate,
	foodRepository *repository.FoodRepository, foodProducer *messaging.FoodProducer) *FoodUseCase {
	return &FoodUseCase{
		DB:             db,
		Log:            log,
		Validate:       validate,
		FoodRepository: foodRepository,
		FoodProducer:   foodProducer,
	}
}

func (c *FoodUseCase) Get(ctx context.Context, request *model.FoodRequest) (*model.FoodResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	food := new(entity.Food)
	if err := c.FoodRepository.FindByid(tx, food, request.ID); err != nil {
		c.Log.WithError(err).Error("error getting food")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting food")
		return nil, fiber.ErrInternalServerError
	}

	return converter.FoodToResponse(food), nil

}
