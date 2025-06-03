package usecase

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/dominikus1993/game-logger/pkg/load/repo"
	"github.com/dominikus1993/game-logger/pkg/load/service"
)

type LoadGamesUseCase struct {
	provider service.LoadGamesService
	writer   repo.GamesWriter
}

func NewLoadGamesUseCase(provider service.LoadGamesService, writer repo.GamesWriter) *LoadGamesUseCase {
	return &LoadGamesUseCase{provider: provider, writer: writer}
}

func (uc *LoadGamesUseCase) Execute(ctx context.Context) error {

	stream := uc.provider.Load(ctx)
	var err error
	for game := range stream {
		slog.Info("loaded game", "id", game.Id, "title", game.Title)
		lerr := uc.writer.WriteGame(ctx, game)
		if err != nil {
			err = errors.Join(err, fmt.Errorf("failed to write game %s: %w", game.Id, lerr))
		}
	}
	return err
}
