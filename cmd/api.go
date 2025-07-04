package cmd

import (
	"context"
	"log"
	"log/slog"
	"strconv"

	"github.com/dominikus1993/game-logger/internal/api/repo"
	"github.com/dominikus1993/game-logger/internal/mongo"
	domainRepo "github.com/dominikus1993/game-logger/pkg/api/repo"
	"github.com/dominikus1993/game-logger/pkg/api/usecases"
	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
	"github.com/samber/do"
	"github.com/urfave/cli/v3"
)

type ApiParseArgs struct {
	filePath              string
	mongoConnectionString string
}

func NewApiParseArgs(context *cli.Command) *ApiParseArgs {
	filePath := context.String("file-path")
	mongo := context.String("mongo-connection-string")
	return &ApiParseArgs{filePath: filePath, mongoConnectionString: mongo}
}

func Api(ctx context.Context, cmd *cli.Command) error {
	slog.InfoContext(ctx, "Run Api Server")
	p := NewApiParseArgs(cmd)
	mongodbClient, err := mongo.NewClient(ctx, p.mongoConnectionString, "Games", "games")
	if err != nil {
		slog.ErrorContext(ctx, "can't create mongodb client", "error", err)
		return cli.Exit("can't create mongodb client", 1)
	}
	defer mongodbClient.Close(ctx)

	statsProvider := repo.NewPlayedHoursPerPlatformStatsProvider(mongodbClient)
	injector := do.New()

	do.Provide(injector, func(i *do.Injector) (*mongo.MongoClient, error) {
		return mongodbClient, nil
	})

	do.Provide(injector, func(i *do.Injector) (domainRepo.GamesReader, error) {
		client := do.MustInvoke[*mongo.MongoClient](i)
		return repo.NewMongoGamesReader(client)
	})

	do.Provide(injector, func(i *do.Injector) (*usecases.LoadGamesUseCase, error) {
		gamesReader := do.MustInvoke[domainRepo.GamesReader](i)
		return usecases.NewLoadGamesUseCase(gamesReader)
	})

	engine := html.New("./public", ".html")

	// Add helper functions to template engine
	engine.AddFunc("sub", func(a, b int) int { return a - b })
	engine.AddFunc("add", func(a, b int) int { return a + b })
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		// Pass in Views Template Engine
		Views: engine,

		// Enables/Disables access to `ctx.Locals()` entries in rendered views
		// (defaults to false)
		PassLocalsToViews: false,
	})

	// Define a route for the GET method on the root path '/'
	app.Get("/ping", func(c fiber.Ctx) error {
		return c.SendString("pong")
	})

	app.Get("/", func(c fiber.Ctx) error {
		pageInt := 1
		limitInt := 10
		loadGamesUseCase, err := do.Invoke[*usecases.LoadGamesUseCase](injector)
		if err != nil {
			slog.ErrorContext(ctx, "Error while invoking LoadGamesUseCase", slog.Any("error", err))
			return c.Status(fiber.StatusInternalServerError).SendString("Error")
		}
		res, err := loadGamesUseCase.Execute(ctx, usecases.LoadGamesQuery{Page: pageInt, Size: limitInt})

		if err != nil {
			slog.ErrorContext(ctx, "Error while loading games", slog.Any("error", err))
			return c.Status(fiber.StatusInternalServerError).SendString("Error while loading games")
		}
		if len(res.Games) == 0 {
			slog.WarnContext(ctx, "No games found")
			return c.Status(fiber.StatusNotFound).SendString("No games found")
		}

		return c.Render("index", fiber.Map{
			"Games": res.Games,
			"Page":  pageInt,
			"Limit": limitInt,
			"Count": len(res.Games),
			"Total": res.Total,
		})
	})

	app.Get("/games", func(c fiber.Ctx) error {
		page := c.Query("page", "1")
		limit := c.Query("limit", "10")
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			slog.ErrorContext(ctx, "Invalid page number", slog.String("page", page), slog.Any("error", err))
			return c.Status(fiber.StatusBadRequest).SendString("Invalid page number")
		}
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			slog.ErrorContext(ctx, "Invalid limit number", slog.String("limit", limit), slog.Any("error", err))
			return c.Status(fiber.StatusBadRequest).SendString("Invalid limit number")
		}

		loadGamesUseCase, err := do.Invoke[*usecases.LoadGamesUseCase](injector)
		if err != nil {
			slog.ErrorContext(ctx, "Error while invoking LoadGamesUseCase", slog.Any("error", err))
			return c.Status(fiber.StatusInternalServerError).SendString("Error")
		}
		res, err := loadGamesUseCase.Execute(ctx, usecases.LoadGamesQuery{Page: pageInt, Size: limitInt})

		if err != nil {
			slog.ErrorContext(ctx, "Error while loading games", slog.Any("error", err))
			return c.Status(fiber.StatusInternalServerError).SendString("Error while loading games")
		}
		if len(res.Games) == 0 {
			slog.WarnContext(ctx, "No games found")
		}

		return c.Render("games", fiber.Map{
			"Games": res.Games,
			"Page":  pageInt,
			"Limit": limitInt,
			"Count": len(res.Games),
			"Total": res.Total,
		})
	})

	app.Get("/gamesperplatform", func(c fiber.Ctx) error {

		useCase := usecases.NewPlayedHoursPerPlatformUseCase(statsProvider)
		res, err := useCase.Execute(ctx)

		if err != nil {
			slog.ErrorContext(ctx, "Error while loading games", slog.Any("error", err))
			return c.Status(fiber.StatusInternalServerError).SendString("Error while loading games")
		}
		if len(res) == 0 {
			slog.WarnContext(ctx, "No stats found")
		}

		return c.Render("gamesperplatform", fiber.Map{
			"Stats": res,
		})
	})

	app.Get("/gamesperyear", func(c fiber.Ctx) error {

		useCase := usecases.NewPlayedHoursPerYearUseCase(statsProvider)
		res, err := useCase.Execute(ctx)

		if err != nil {
			slog.ErrorContext(ctx, "Error while loading games", slog.Any("error", err))
			return c.Status(fiber.StatusInternalServerError).SendString("Error while loading games")
		}
		if len(res) == 0 {
			slog.WarnContext(ctx, "No stats found")
		}

		return c.Render("gamesperyear", fiber.Map{
			"Stats": res,
		})
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
	slog.InfoContext(ctx, "Parsing articles finished")
	return nil
}
