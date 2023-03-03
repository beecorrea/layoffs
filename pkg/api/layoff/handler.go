package layoff

import "github.com/gofiber/fiber/v2"

func ListLayoffs(c *fiber.Ctx) error {
	comp := c.Params("company", "")
	lfs := GetLayoffs(comp)

	resp := LayoffResponse{lfs}

	return c.JSON(resp)
}

type LayoffResponse struct {
	Layoffs []Layoff `json:"layoffs"`
}
