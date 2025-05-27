package cmd

import (
	"context"
	"log/slog"

	"github.com/dominikus1993/game-logger/internal/load/repo"
	"github.com/dominikus1993/game-logger/internal/load/service"
	"github.com/dominikus1993/game-logger/internal/mongo"
	"github.com/dominikus1993/game-logger/pkg/load/usecase"
	"github.com/urfave/cli/v3"
)

type ParseArgs struct {
	filePath              string
	mongoConnectionString string
}

func NewParseArgs(context *cli.Command) *ParseArgs {
	filePath := context.String("file-path")
	mongo := context.String("mongo-connection-string")
	return &ParseArgs{filePath: filePath, mongoConnectionString: mongo}
}

func Parse(ctx context.Context, cmd *cli.Command) error {
	p := NewParseArgs(cmd)
	slog.InfoContext(ctx, "Parse Articles And Send It")
	mongodbClient, err := mongo.NewClient(ctx, p.mongoConnectionString, "Games", "games")
	if err != nil {
		slog.ErrorContext(ctx, "can't create mongodb client", "error", err)
		return cli.Exit("can't create mongodb client", 1)
	}
	defer mongodbClient.Close(ctx)
	repo := repo.NewMongoGamesWriter(mongodbClient)
	articlesProvider := service.NewExcelLoadGamesService(p.filePath, "Sheet1")

	usecase := usecase.NewLoadGamesUseCase(articlesProvider, repo)

	err = usecase.Execute(ctx)

	if err != nil {
		slog.ErrorContext(ctx, "Error while parsing articles", "error", err)
		return cli.Exit("Error while parsing articles", 0)
	}
	slog.InfoContext(ctx, "Parsing articles finished")
	return nil
}
