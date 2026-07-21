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

func (c *FoodController) List(ctx *fiber.Ctx) error {
	request := &model.FoodRequest{
		Limit: 10,
	}

	response, err := c.useCase.List(ctx.UserContext(), request)

	if err != nil {
		c.Log.WithError(err).Error("error getting food")
		return err
	}

	return ctx.JSON(model.WebResponse[[]*model.FoodResponse]{Data: response})

}

func (c *FoodController) Create(ctx *fiber.Ctx) error {

	request := &model.FoodRequest{}
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	response, err := c.useCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error creating food")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.FoodResponse]{Data: response})
}

func (c *FoodController) Update(ctx *fiber.Ctx) error {
	request := &model.FoodRequest{}

	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	request.ID = ctx.Params("foodId")

	response, err := c.useCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error updating food")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.FoodResponse]{Data: response})
}

func (c *FoodController) Delete(ctx *fiber.Ctx) error {
	request := &model.FoodRequest{
		ID: ctx.Params("foodId"),
	}

	err := c.useCase.Delete(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error deleting food")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
