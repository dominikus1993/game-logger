package repo

import "github.com/dominikus1993/game-logger/pkg/model"

type GamesWriter interface {
	WriteGame(games *model.Game) error
}
