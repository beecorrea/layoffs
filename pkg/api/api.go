package api

import (
	"context"
	"log"

	"github.com/beecorrea/layoffs/pkg/api/layoff"
	"github.com/beecorrea/layoffs/pkg/persistence"
	"github.com/gofiber/fiber/v2"
)

func NewLayoffApp() {
	ctx := context.Background()

	handler := GetComponents(ctx)

	app := fiber.New()

	app.Get("/api/layoff/:company?", handler.ListLayoffs)
	app.Post("/api/layoff", handler.SaveLayoff)
	log.Fatal(app.Listen(":3000"))
}

func GetComponents(ctx context.Context) layoff.LayoffHandler {
	db, err := persistence.Open(ctx)
	if err != nil {
		panic(err)
	}
	repo := layoff.LayoffRepository{Database: db}
	service := layoff.NewLayoffService(repo)
	handler := layoff.NewLayoffHandler(service)
	return handler
}
