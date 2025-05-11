package service

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/dominikus1993/game-logger/pkg/model"
	"github.com/xuri/excelize/v2"
)

type ExcelLoadGamesService struct {
	path string
}

func NewExcelLoadGamesService(path string) *ExcelLoadGamesService {
	return &ExcelLoadGamesService{
		path: path,
	}
}

func (s *ExcelLoadGamesService) Load(ctx context.Context) <-chan *model.Game {
	games := make(chan *model.Game, 10)

	go func(ctx context.Context, service *ExcelLoadGamesService) {
		defer close(games)
		f, err := excelize.OpenFile(service.path)
		if err != nil {
			slog.ErrorContext(ctx, "failed to open file", "file", service.path, "error", err)
			return
		}
		defer func() {
			// Close the spreadsheet.
			if err := f.Close(); err != nil {
				slog.ErrorContext(ctx, "failed to close file", "file", service.path, "error", err)
			}
		}()

		rows, err := f.GetRows("Arkusz1")
		if err != nil {
			slog.ErrorContext(ctx, "failed to get rows", "file", service.path, "error", err)
			return
		}

		for i, row := range rows {
			if i < 2 {
				continue
			}

			if len(row) < 6 {
				slog.ErrorContext(ctx, "invalid row", "file", service.path, "row", i, "data", row)
				continue
			}
			// []string len: 6, cap: 8, ["Ori and the Will of the Wisps","5","Switch","2023-10-01","2023-11-01","25"]
			game := &model.Game{
				Id:          generateId(),
				Title:       row[0],
				Rating:      parseRating(row[1]),
				Platform:    row[2],
				StartDate:   row[3],
				FinishDate:  row[4],
				HoursPlayed: parseRating(row[5]),
			}
			games <- game
		}

	}(ctx, s)

	return games
}

func generateId() string {
	return "test"
}

func stringToPointer(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func parseRating(rating string) int {
	if rating == "" {
		return 0
	}
	r, err := strconv.Atoi(rating)
	if err != nil {
		return 0
	}
	return r
}
