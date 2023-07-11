package config

import (
	"github.com/gofiber/fiber/v2"
	"clean-architecture/exception"
)
func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}