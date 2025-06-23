package usecases

import (
	"context"

	"github.com/dominikus1993/game-logger/pkg/api/repo"
	"github.com/dominikus1993/game-logger/pkg/model"
)

type LoadGamesQuery struct {
	Page int // Page number for pagination
	Size int // Number of games per page
}

type LoadGamesUseCase struct {
	provider repo.GamesReader
}

type LoadGamesResponse struct {
	Games []*model.Game // List of games loaded
	Total int           // Total number of games available
}

func NewLoadGamesUseCase(provider repo.GamesReader) (*LoadGamesUseCase, error) {
	return &LoadGamesUseCase{provider: provider}, nil
}

func (uc *LoadGamesUseCase) Execute(ctx context.Context, query LoadGamesQuery) (*LoadGamesResponse, error) {
	games, err := uc.provider.LoadGames(ctx, repo.LoadGamesQuery{Page: query.Page, Size: query.Size})
	if err != nil {
		return nil, err
	}

	total, err := uc.provider.Count(ctx)
	if err != nil {
		return nil, err
	}
	response := &LoadGamesResponse{
		Games: games,
		Total: total,
	}
	return response, nil
}
