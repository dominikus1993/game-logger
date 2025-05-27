package main

import (
	"context"
	"log"
	"os"
	"sort"

	"github.com/dominikus1993/game-logger/cmd"
	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "mongo-connection-string",
				Value:    "",
				Usage:    "mongo-connection-string",
				Sources:  cli.EnvVars("MONGO_CONNECTION"),
				Required: true,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "parse",
				Usage: "Parse articles from excel file and save them to MongoDB",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "file-path",
						Value:    "",
						Usage:    "file-path",
						Sources:  cli.EnvVars("FILE_PATH"),
						Required: true,
					},
				},
				Action: cmd.Parse,
			},
			{
				Name:   "api",
				Usage:  "Start API server",
				Flags:  []cli.Flag{},
				Action: cmd.Api,
			},
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
