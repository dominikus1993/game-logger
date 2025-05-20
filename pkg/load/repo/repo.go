package repo

import (
	"context"

	"github.com/dominikus1993/game-logger/pkg/model"
)

type GamesWriter interface {
	WriteGame(ctx context.Context, games *model.Game) error
}
