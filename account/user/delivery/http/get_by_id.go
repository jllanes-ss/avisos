package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jllanes-ss/avisos/account/domain/apperrors"
)

// GetByID will get article by given id
func (uh *Handler) GetByID(c *fiber.Ctx) error {
	uid, err := uuid.Parse(c.Params("id"))

	// This shouldn't happen, as our middleware ought to throw an error.
	// This is an extra safety measure
	// We'll extract this logic later as it will be common to all handler
	// methods which require a valid user
	if err != nil {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
		err := apperrors.NewInternal()
		return c.Status(err.Status()).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})

	}

	user, err := uh.UserUseCase.Get(c.Context(), uid)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
		//return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

// func (a *UserHandler) GetByID(c *fiber.Ctx) error {
// 	uid, _ := uuid.Parse(c.Params("id"))

// 	user, err := a.UserService.GetByID(c.Context(), uid)
// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
// 		//return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
// 	}

// 	return c.JSON(fiber.Map{"status": "success", "message": "Product found", "data": user})
// }
