package handler

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jllanes-ss/avisos/account/domain"
	"github.com/jllanes-ss/avisos/account/domain/apperrors"
)

type loginReq struct {
	Email    string `json:"email" xml:"name" form:"name" binding:"required,email"`
	Password string `json:"password" xml:"password" form:"password" binding:"required,gte=6,lte=30"`
}

// Signin used to authenticate extant user
func (h *Handler) Login(c *fiber.Ctx) error {

	req := new(loginReq)

	if err := c.BodyParser(req); err != nil {
		return err
	}

	log.Println(req.Email)    // john
	log.Println(req.Password) // doe

	// email, err := uuid.Parse(c.Params("email"))

	user := &domain.User{
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.UserUseCase.Login(c.Context(), user); err != nil {
		log.Printf("Failed to sign in user: %v\n", err.Error())
		breq := apperrors.NewBadRequest("Failed to sign in user")
		return c.Status(breq.Status()).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}

	tokens, err := h.TokenUseCase.NewPairFromUser(c.Context(), user, "")

	if err != nil {
		log.Printf("Failed to create tokens for user: %v\n", err.Error())

		breq := apperrors.NewBadRequest("Failed to sign in user")
		return c.Status(breq.Status()).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"tokens": tokens,
	// })

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": "ok", "message": "Login information", "data": tokens}) //user
}
