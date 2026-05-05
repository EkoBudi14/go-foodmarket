package http

import (
	"golang-clean-architecture/internal/delivery/http/middleware"
	"golang-clean-architecture/internal/model"
	"golang-clean-architecture/internal/usecase"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log     *logrus.Logger
	UseCase *usecase.UserUseCase
}

func NewUserController(useCase *usecase.UserUseCase, logger *logrus.Logger) *UserController {
	return &UserController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	request := new(model.RegisterUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to register user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	request := new(model.LoginUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Login(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to login user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

func (c *UserController) Current(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.GetUserRequest{
		ID: auth.ID,
	}

	response, err := c.UseCase.Current(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to get current user")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

func (c *UserController) Logout(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.LogoutUserRequest{
		ID: auth.ID,
	}

	response, err := c.UseCase.Logout(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to logout user")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: response})
}

func (c *UserController) Update(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.UpdateUserRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	request.ID = auth.ID
	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to update user")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

func (c *UserController) UpdatePhoto(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	file, err := ctx.FormFile("file")
	if err != nil {
		c.Log.Warnf("Failed to get file from request : %+v", err)
		return fiber.ErrBadRequest
	}

	if file.Size > 2*1024*1024 {
		return ctx.Status(401).JSON(model.WebResponse[any]{
			Errors: "File size exceeds 2MB limit",
		})
	}

	contentType := file.Header.Get("Content-Type")
	allowedTypes := map[string]bool{
		"image/jpeg": true, "image/png": true,
		"image/gif": true, "image/webp": true,
	}

	if !allowedTypes[contentType] {
		return ctx.Status(401).JSON(model.WebResponse[any]{
			Errors: "File must be an image (jpeg, png, gif, webp)",
		})
	}

	uploadDir := "./assets/user"
	os.MkdirAll(uploadDir, os.ModePerm)

	filename := uuid.New().String() + filepath.Ext(file.Filename)
	savePath := filepath.Join(uploadDir, filename)

	if err := ctx.SaveFile(file, savePath); err != nil {
		c.Log.Warnf("Failed to save file : %+v", err)
		return fiber.ErrInternalServerError
	}

	request := &model.UpdateProfilePhotoRequest{ID: auth.ID}

	storePath := "assets/user/" + filename

	response, err := c.UseCase.UpdatePhoto(ctx.UserContext(), request, storePath)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to update user photo")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})

}
