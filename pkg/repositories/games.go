package repositories

import (
	"context"

	"github.com/dominikus1993/game-logger/pkg/data"
)

type GetGamesRequest struct {
	Page     int
	PageSize int
}

type GamesRepository interface {
	GetGames(ctx context.Context, req *GetGamesRequest) ([]data.Game, error)
}

type FakeGamesRepository struct {
}

func NewFakeGamesRepository() *FakeGamesRepository {
	return &FakeGamesRepository{}
}

func (r *FakeGamesRepository) GetGames(ctx context.Context, req *GetGamesRequest) ([]data.Game, error) {
	return []data.Game{
		{ID: "1", Name: "Game 1", PlayStart: "2021-01-01", PlayEnd: "2021-01-02", Rating: 5, Platform: "PS4"},
		{ID: "2", Name: "Game 2", PlayStart: "2021-01-02", PlayEnd: "2021-01-03", Rating: 4, Platform: "PS4"},
	}, nil
}
