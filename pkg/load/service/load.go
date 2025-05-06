package service

import (
	"context"

	"github.com/dominikus1993/game-logger/pkg/model"
)

type LoadGamesService interface {
	Load(ctx context.Context) <-chan *model.Game
}
