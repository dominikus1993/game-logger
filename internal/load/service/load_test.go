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
	// []string len: 6, cap: 8, ["Ori and the Will of the Wisps","5","Switch2","2023-10-01","2023-11-01","25"]
	firstGame := games[0]
	assert.NotNil(t, firstGame)
	assert.NotNil(t, firstGame.Rating)
	assert.Equal(t, "bb2dc65f-b61d-58a9-86d4-1767afba61a1", firstGame.Id)
	assert.Equal(t, 5, *firstGame.Rating)
	assert.Equal(t, "Switch", firstGame.Platform)
	assert.Equal(t, "2023-10-01", firstGame.StartDate.Format("2006-01-02"))
	assert.Equal(t, "2023-11-01", firstGame.FinishDate.Format("2006-01-02"))
	assert.NotNil(t, firstGame.HoursPlayed)
	assert.Equal(t, 25, *firstGame.HoursPlayed)
	assert.Equal(t, "Ori and the Will of the Wisps", firstGame.Title)
}

func TestParseRating(t *testing.T) {
	tests := []struct {
		input    string
		expected *int
	}{
		{"", nil},
		{"5", intPtr(5)},
		{"10", intPtr(10)},
		{"-1", intPtr(-1)},
		{"11", intPtr(11)},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := parseRating(test.input)
			assert.Equal(t, test.expected, result)
		})
	}

}

func intPtr(i int) *int {
	return &i
}
