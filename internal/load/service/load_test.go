package service

import (
	"context"
	"testing"

	"github.com/dominikus1993/go-toolkit/channels"
	"github.com/stretchr/testify/assert"
)

func TestLoadGamesService(t *testing.T) {
	path := "Games.xlsx"
	service := NewExcelLoadGamesService(path, "Arkusz1")

	assert.NotNil(t, service)

	channel := service.Load(context.Background())

	games := channels.ToSlice(channel)

	assert.NotEmpty(t, games)
	assert.Len(t, games, 52)
	// []string len: 6, cap: 8, ["Ori and the Will of the Wisps","5","Switch","2023-10-01","2023-11-01","25"]
	firstGame := games[0]
	assert.Equal(t, 5, firstGame.Rating)
	assert.Equal(t, "Switch", firstGame.Platform)
	assert.Equal(t, "2023-10-01", firstGame.StartDate)
	assert.Equal(t, "2023-11-01", firstGame.FinishDate)
	assert.Equal(t, 25, firstGame.HoursPlayed)
	assert.Equal(t, "Ori and the Will of the Wisps", firstGame.Title)
}
