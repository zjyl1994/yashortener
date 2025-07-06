package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/template/html/v2"
	"github.com/zjyl1994/yashortener/infra/vars"
)

func Run(listen string) error {
	engine := html.New("./web/template", ".html")
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Views:                 engine,
	})
	app.Get("/", indexHandler)
	app.Get("/:code", processHandler)
	app.Post("/create", createHandler)
	admin := app.Group("/admin")
	if vars.AdminUser != "" || vars.AdminPass != "" {
		admin.Use(basicauth.New(basicauth.Config{
			Users: map[string]string{
				vars.AdminUser: vars.AdminPass,
			},
		}))
	}
	admin.Get("/", adminIndexHandler)
	admin.Get("/:code", adminDetailHandler)
	admin.Delete("/:code", adminDeleteHandler)
	return app.Listen(listen)
}
