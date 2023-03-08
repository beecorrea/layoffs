package layoff

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

type LayoffHandler struct {
	service LayoffService
}

func NewLayoffHandler(service LayoffService) LayoffHandler {
	return LayoffHandler{service}
}

type ListLayoffResponse struct {
	Msg     string   `json:"message"`
	Layoffs []Layoff `json:"layoffs"`
}

func (lh LayoffHandler) ListLayoffs(c *fiber.Ctx) error {
	comp := c.Params("company", "")
	ctx := c.UserContext()

	lfs, err := lh.service.GetLayoffs(ctx, comp)
	resp := ListLayoffResponse{Msg: "success", Layoffs: lfs}
	if err != nil {
		log.Println(err)
		resp.Msg = fmt.Sprintf("couldn't list layoffs for %s", comp)
	}

	return c.JSON(resp)
}

type SaveLayoffResponse struct {
	Msg    string `json:"message"`
	Layoff Layoff `json:"layoff"`
}

func (lh LayoffHandler) SaveLayoff(c *fiber.Ctx) error {
	var layoff Layoff
	if err := c.BodyParser(&layoff); err != nil {
		return err
	}

	if layoff.ReportedAt == nil {
		now := time.Now()
		layoff.ReportedAt = &now
	}

	ctx := c.UserContext()
	lf, err := lh.service.repo.PersistLayoff(ctx, layoff)
	resp := SaveLayoffResponse{Msg: "success", Layoff: *lf}
	if err != nil {
		log.Println(err)
		resp.Msg = fmt.Sprintf("couldn't save layoff at %s", layoff.CompanyName)
		return c.Status(500).JSON(resp)
	}

	if err := c.Status(201).JSON(resp); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
