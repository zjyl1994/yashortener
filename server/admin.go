package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zjyl1994/yashortener/service"
)

func adminIndexHandler(c *fiber.Ctx) error {
	keyword := c.Query("keyword")
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("size", 10)
	links, count, err := service.ListLink(keyword, page, pageSize)
	if err != nil {
		return err
	}
	return c.Render("admin", fiber.Map{
		"links": links,
		"count": count,
	})
}

func adminDeleteHandler(c *fiber.Ctx) error {
	code := c.Params("code")
	err := service.DeleteLink(code)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusOK)
}

func adminDetailHandler(c *fiber.Ctx) error {
	code := c.Params("code")
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("size", 10)
	link, err := service.GetLink(code)
	if err != nil {
		return err
	}
	access, count, err := service.ListAccessRecord(link.ID, page, pageSize)
	if err != nil {
		return err
	}
	return c.Render("detail", fiber.Map{
		"link":   link,
		"count":  count,
		"access": access,
	})
}
