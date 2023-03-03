package api

import (
	"log"

	"github.com/beecorrea/layoffs/pkg/api/layoff"
	"github.com/gofiber/fiber/v2"
)

func NewLayoffApp() {
	app := fiber.New()

	app.Get("*", func(c *fiber.Ctx) error {
		return c.SendString("TODO implement api")
	})

	app.Get("/api/:company?", layoff.ListLayoffs)

	log.Fatal(app.Listen(":3000"))
}
