package main

import (
	"github.com/gofiber/template/html/v2"

	"github.com/gofiber/fiber/v2"
)

func main() {
	engine := html.New("./views", ".html")
	// Reload the templates on each render, good for development
	engine.Reload(true) // Optional. Default: false

	// Debug will print each template that is parsed, good for debugging
	engine.Debug(true) // Optional. Default: false
	app := fiber.New(fiber.Config{
		// Pass in Views Template Engine
		Views: engine,
		// Enables/Disables access to `ctx.Locals()` entries in rendered views
		// (defaults to false)
		PassLocalsToViews: false,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Listen(":3000")
}
