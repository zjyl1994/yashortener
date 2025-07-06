package server

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/zjyl1994/yashortener/infra/vars"
	"github.com/zjyl1994/yashortener/service"
)

func indexHandler(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func createHandler(fromAdmin bool) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if !vars.AnonymousCreate && !fromAdmin {
			return c.Status(fiber.StatusForbidden).SendString("anonymous create disabled")
		}

		if c.FormValue("url") == "" {
			return c.Status(fiber.StatusBadRequest).SendString("url is required")
		}

		url := strings.TrimSpace(c.FormValue("url"))
		code := strings.TrimSpace(c.FormValue("code"))
		short, err := service.CreateLink(code, url)
		if err != nil {
			return err
		}

		baseUrl := vars.BaseURL
		if baseUrl == "" {
			baseUrl = c.BaseURL()
		}
		link := strings.TrimSuffix(baseUrl, "/") + "/" + short
		return c.SendString(link)
	}
}

func processHandler(c *fiber.Ctx) error {
	code := c.Params("code")
	link, err := service.GetLink(code)
	if err != nil {
		return err
	}
	if link == nil {
		return c.Status(fiber.StatusNotFound).Render("status", fiber.Map{"code": "404", "message": "link not found"})
	}
	err = service.RecordAccess(link.ID, c.IP(), c.Get(fiber.HeaderUserAgent))
	if err != nil {
		logrus.Errorln("write access record faild", link.ID)
	}
	return c.Redirect(link.Link)
}
