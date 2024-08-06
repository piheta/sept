package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	weberrors "github.com/piheta/sept/errors"
	"github.com/piheta/sept/models"
	"github.com/piheta/sept/services"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

// @Summary Get all users
// @Description Get all users
// @Tags User
// @Produce json
// @Failure 401
// @Failure 404
// @Failure 500
// @Success 200 {array} User
// @Router /api/users [get]
func (ec *UserController) GetUsers(c *fiber.Ctx) error {
	jwt := mapReqToJWT(c)
	users, err := ec.userService.GetUsers(jwt)
	if err != nil {
		return err
	}
	return c.JSON(users)
}

// @Summary Get an user by ID
// @Description Get an user by its ID
// @Tags User
// @Produce json
// @Param id path int true "User ID"
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Success 200 {object} User
// @Router /api/users/{id} [get]
func (ec *UserController) GetUser(c *fiber.Ctx) error {
	jwt := mapReqToJWT(c)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return weberrors.NewError(400, "invalid id")
	}

	user, err := ec.userService.GetUser(jwt, id)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

// @Summary Create a new user
// @Description Create a new user
// @Tags User
// @Produce json
// @Param user body User true "User object"
// @Failure 400
// @Failure 401
// @Failure 409
// @Failure 500
// @Success 201 {object} User
// @Router /api/users [post]
func (ec *UserController) CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return weberrors.NewError(400, err.Error())
	}

	if err := ec.userService.CreateUser(&user); err != nil {
		return err
	}

	return c.Status(201).JSON(user)
}

// @Summary Update an existing user
// @Description Update an existing user by its ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body UpdateUser true "UpdateUser object"
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 409
// @Failure 500
// @Success 204 "No Content"
// @Router /api/users/{id} [put]
func (ec *UserController) UpdateUser(c *fiber.Ctx) error {
	jwt := mapReqToJWT(c)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return weberrors.NewError(400, "invalid id")
	}

	var user models.UpdateUser
	if err := c.BodyParser(&user); err != nil {
		return weberrors.NewError(400, err.Error())
	}

	if err := ec.userService.UpdateUser(jwt, id, &user); err != nil {
		return err
	}

	return c.SendStatus(204)
}

// @Summary Delete an user by ID
// @Description Delete an user by its ID
// @Tags User
// @Produce json
// @Param id path int true "User ID"
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Success 204 "No Content"
// @Router /api/users/{id} [delete]
func (ec *UserController) DeleteUser(c *fiber.Ctx) error {
	jwt := mapReqToJWT(c)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return weberrors.NewError(400, "invalid id")
	}

	if err := ec.userService.DeleteUser(jwt, id); err != nil {
		return err
	}

	return c.SendStatus(204)
}
