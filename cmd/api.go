package cmd

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"github.com/gofiber/fiber/v3"
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
	slog.InfoContext(ctx, "Parse Articles And Send It")
	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		msg := c.Query("msg", "Nobody")
		return c.SendString(fmt.Sprintf("Hello, World 👋! %s", msg))
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
	slog.InfoContext(ctx, "Parsing articles finished")
	return nil
}
