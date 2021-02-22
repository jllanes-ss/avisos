package handler

import (

	//"github.com/labstack/echo"
	"time"

	"github.com/gofiber/fiber/v2"

	//validator "gopkg.in/go-playground/validator.v9"

	//"github.com/bxcodec/go-clean-arch/domain"

	"github.com/jllanes-ss/avisos/account/domain"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// ArticleHandler  represent the httphandler for article
type Handler struct {
	UserUseCase  domain.UserUseCase
	TokenUseCase domain.TokenUseCase
}

type Config struct {
	F               *fiber.App
	UserUseCase     domain.UserUseCase
	TokenUseCase    domain.TokenUseCase
	BaseURL         string
	TimeoutDuration time.Duration
	MaxBodyBytes    int64
}

// NewArticleHandler will initialize the articles/ resources endpoint
//func NewUserHandler(app *fiber.App, us domain.UserUseCase) {
func NewUserHandler(c *Config) {
	h := &Handler{
		UserUseCase:  c.UserUseCase,
		TokenUseCase: c.TokenUseCase,
	}

	user := c.F.Group(c.BaseURL)
	user.Get("/:id", h.GetByID)
	user.Post("/login", h.Login)

	c.F.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
}
