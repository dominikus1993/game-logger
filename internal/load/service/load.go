package service

import (
	"context"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/dominikus1993/game-logger/pkg/model"
	"github.com/google/uuid"
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
			startDate, err := time.Parse("2006-01-02", row.GetCell(3).String())
			if err != nil {
				return err
			}
			finishDate := row.GetCell(4).String()
			var finishDateTime *time.Time
			if finishDate != "" {
				finishDateTimeParsed, err := time.Parse("2006-01-02", finishDate)
				if err != nil {
					return err
				}
				finishDateTime = &finishDateTimeParsed
			}
			game := &model.Game{
				Id:    generateId(title),
				Title: title,

				Playthroughs: []model.Playthrough{{
					Rating:      parseRating(row.GetCell(1).String()),
					Platform:    row.GetCell(2).String(),
					StartDate:   startDate,
					FinishDate:  finishDateTime,
					HoursPlayed: parseRating(row.GetCell(5).String()),
				}},
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

func generateId(title string) string {
	if title == "" {
		return uuid.NewString()
	}
	normalizedTitle := lowercaseAndTirm(title)

	if normalizedTitle == "" {
		return uuid.NewString()
	}
	data := normalizedTitle
	return uuid.NewSHA1(uuid.NameSpaceDNS, []byte(data)).String()
}

func lowercaseAndTirm(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func parseRating(rating string) *int {
	if rating == "" {
		return nil
	}
	r, err := strconv.Atoi(rating)
	if err != nil {
		return nil
	}
	return &r
}

type ExcelGame struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	StartDate   time.Time  `json:"start_date"`
	FinishDate  *time.Time `json:"finish_date,omitempty"`
	Platform    string     `json:"platform,omitempty"`
	HoursPlayed *int       `json:"hours_played,omitempty"`
	Rating      *int       `json:"rating,omitempty"`
	Notes       string     `json:"notes,omitempty"`
}
