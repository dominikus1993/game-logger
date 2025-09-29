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
	assert.NotEmpty(t, firstGame.Playthroughs)
	play := firstGame.Playthroughs[0]
	assert.NotNil(t, play.Rating)
	assert.Equal(t, 5, *play.Rating)
	assert.Equal(t, "Switch", play.Platform)
	assert.Equal(t, "2023-10-01", play.StartDate.Format("2006-01-02"))
	assert.Equal(t, "2023-11-01", play.FinishDate.Format("2006-01-02"))
	assert.NotNil(t, play.HoursPlayed)
	assert.Equal(t, 25, *play.HoursPlayed)
	assert.Equal(t, "Ori and the Will of the Wisps", firstGame.Title)
}

func TestParseRating(t *testing.T) {
	tests := []struct {
		input    string
		expected *int
	}{
		{"", nil},
		{"5", Pointer(5)},
		{"10", Pointer(10)},
		{"-1", Pointer(-1)},
		{"11", Pointer(11)},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := parseRating(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func FuzzLowercaseAndTirm(f *testing.F) {
	f.Add("  Ori and the Will of the Wisps  ")
	f.Add("  Ori and the Will of the Wisps")
	f.Add("Ori and the Will of the Wisps  ")
	f.Add("Ori and the Will of the Wisps")
	f.Fuzz(func(t *testing.T, input string) {
		result := lowercaseAndTirm(input)
		assert.Equal(t, "ori and the will of the wisps", result)
	})
}

func TestGenerateId(t *testing.T) {
	title := "Test Game"
	expectedId := "b1bf8ccb-bf4c-5d84-bd77-3f81da7caacb" // Example UUID, replace with actual expected value

	id := generateId(title)
	assert.Equal(t, expectedId, id)
}

func Pointer[T any](i T) *T {
	return &i
}
