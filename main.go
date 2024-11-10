package main

import (
	"strconv"

	"github.com/dominikus1993/game-logger/pkg/repositories"
	"github.com/dominikus1993/game-logger/pkg/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/template/html/v2"
)

var getGamesUseCase *usecase.GamesUsecase

func init() {
	getGamesUseCase = usecase.NewGamesUsecase(repositories.NewFakeGamesRepository())
}

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
		page := ParseQueryInt(c.Query("page", "1"), 1)
		pageSize := ParseQueryInt(c.Query("page_size", "10"), 10)
		games, err := getGamesUseCase.GetGames(c.Context(), &repositories.GetGamesRequest{Page: page, PageSize: pageSize})
		if err != nil {
			log.WithContext(c.Context()).Errorf("Failed to get games: %v", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.Render("games", fiber.Map{
			"Games": games,
		})
	})

	app.Listen(":3000")
}

func ParseQueryInt(value string, defaultValue int) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return result
}
