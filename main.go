package main

import (
	"github.com/dominikus1993/game-logger/pkg/data"
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

	app.Get("/games", func(c *fiber.Ctx) error {
		return c.Render("games", fiber.Map{
			"Games": []data.Game{
				{ID: "1", Name: "Game 1", PlayStart: "2021-01-01", PlayEnd: "2021-01-02", Rating: 5, Platform: "PS4"},
			},
		})
	})

	app.Listen(":3000")
}
