package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/zjyl1994/yashortener/infra/utils"
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
	pager := utils.CalcPage(page, pageSize, count)
	logrus.Debugln("pager", utils.ToJson(pager))
	return c.Render("admin", fiber.Map{
		"links": links,
		"pager": pager,
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
	if link == nil {
		return c.Status(fiber.StatusNotFound).Render("status", fiber.Map{"code": "404", "message": "link not found"})
	}
	access, count, err := service.ListAccessRecord(link.ID, page, pageSize)
	if err != nil {
		return err
	}
	pager := utils.CalcPage(page, pageSize, count)
	logrus.Debugln("pager", utils.ToJson(pager))
	return c.Render("detail", fiber.Map{
		"link":   link,
		"access": access,
		"pager":  pager,
	})
}

func adminUpdateHandler(c *fiber.Ctx) error {
	code := c.Params("code")
	link := c.FormValue("link")
	if link == "" {
		return c.Status(fiber.StatusBadRequest).Render("status", fiber.Map{"code": "400", "message": "link is required"})
	}
	err := service.UpdateLink(code, link)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusOK)
}
