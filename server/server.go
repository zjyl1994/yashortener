package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/template/html/v2"
	"github.com/zjyl1994/yashortener/infra/vars"
)

func Run(listen string) error {
	engine := html.New("./web/template", ".html")
	if vars.DebugMode {
		engine.Reload(true)
	}
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Views:                 engine,
		ProxyHeader:           fiber.HeaderXForwardedFor,
		TrustedProxies:        []string{"127.0.0.1", "::1"},
	})

	admin := app.Group("/admin")
	if vars.AdminUser != "" || vars.AdminPass != "" {
		admin.Use(basicauth.New(basicauth.Config{
			Users: map[string]string{
				vars.AdminUser: vars.AdminPass,
			},
		}))
	}
	admin.Get("/", adminIndexHandler)
	admin.Post("/create", createHandler(true))
	admin.Get("/:code", adminDetailHandler)
	admin.Delete("/:code", adminDeleteHandler)
	admin.Put("/:code", adminUpdateHandler)

	app.Get("/", indexHandler)
	app.Get("/:code", processHandler)
	app.Post("/create", createHandler(false))
	return app.Listen(listen)
}
