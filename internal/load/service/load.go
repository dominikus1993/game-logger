package service

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/dominikus1993/game-logger/pkg/model"
	"github.com/tealeg/xlsx/v3"
)

type ExcelLoadGamesService struct {
	path      string
	sheetName string
}

func NewExcelLoadGamesService(path, sheetName string) *ExcelLoadGamesService {
	return &ExcelLoadGamesService{
		path:      path,
		sheetName: sheetName,
	}
}

func (s *ExcelLoadGamesService) Load(ctx context.Context) <-chan *model.Game {
	games := make(chan *model.Game, 10)

	go func(ctx context.Context, service *ExcelLoadGamesService) {
		defer close(games)
		wb, err := xlsx.OpenFile(service.path)
		if err != nil {
			slog.Error("failed to open file", "path", service.path, "error", err)
			return
		}
		sheet, ok := wb.Sheet[service.sheetName]
		if !ok {
			slog.Error("sheet not found", "sheetName", service.sheetName)
			return
		}
		err = sheet.ForEachRow(func(row *xlsx.Row) error {
			title := row.GetCell(0).String()
			if shouldSkipRow(title) {
				return nil
			}
			game := &model.Game{
				Id:          generateId(),
				Title:       row.GetCell(0).String(),
				Rating:      parseRating(row.GetCell(1).String()),
				Platform:    row.GetCell(2).String(),
				StartDate:   row.GetCell(3).String(),
				FinishDate:  row.GetCell(4).String(),
				HoursPlayed: parseRating(row.GetCell(5).String()),
			}
			games <- game
			return nil
		})

		if err != nil {
			slog.Error("failed to read sheet", "sheetName", service.sheetName, "error", err)
			return
		}
	}(ctx, s)

	return games
}

func shouldSkipRow(title string) bool {
	return title == "" || title == "Lista" || title == "Gra"
}

func generateId() string {
	return "test"
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
