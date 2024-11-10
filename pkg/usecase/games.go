package usecase

import (
	"context"

	"github.com/dominikus1993/game-logger/pkg/data"
	"github.com/dominikus1993/game-logger/pkg/repositories"
)

type GamesUsecase struct {
	gamesRepository repositories.GamesRepository
}

func NewGamesUsecase(gamesRepository repositories.GamesRepository) *GamesUsecase {
	return &GamesUsecase{
		gamesRepository: gamesRepository,
	}
}

func (uc *GamesUsecase) GetGames(ctx context.Context, req *repositories.GetGamesRequest) ([]data.Game, error) {
	return uc.gamesRepository.GetGames(ctx, req)
}
