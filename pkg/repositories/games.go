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
