package repo

import (
	"context"

	"github.com/dominikus1993/game-logger/pkg/model"
)

type LoadGamesQuery struct {
	Page int
	Size int
}

type GamesReader interface {
	LoadGames(ctx context.Context, query LoadGamesQuery) ([]*model.Game, error)
	Count(ctx context.Context) (int, error)
}
