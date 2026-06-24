package http

import (
	"golang-clean-architecture/internal/model"
	"golang-clean-architecture/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type FoodController struct {
	useCase *usecase.FoodUseCase
	Log     *logrus.Logger
}

func NewFoodController(usecase *usecase.FoodUseCase, logrus *logrus.Logger) *FoodController {
	return &FoodController{
		useCase: usecase,
		Log:     logrus,
	}
}

func (c *FoodController) Get(ctx *fiber.Ctx) error {

	foodId := ctx.Params("foodId")

	request := &model.FoodRequest{
		ID: foodId,
	}

	response, err := c.useCase.Get(ctx.UserContext(), request)

	if err != nil {
		c.Log.WithError(err).Error("error getting food")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.FoodResponse]{Data: response})
}
