package usecases

import (
	"context"

	"github.com/dominikus1993/game-logger/pkg/api/repo"
)

type PlayedHoursPerPlatofmUseCase struct {
	provider repo.PlayedHoursStatsProvider
}

func NewPlayedHoursPerPlatformUseCase(provider repo.PlayedHoursStatsProvider) *PlayedHoursPerPlatofmUseCase {
	return &PlayedHoursPerPlatofmUseCase{provider: provider}
}

func (uc *PlayedHoursPerPlatofmUseCase) Execute(ctx context.Context) (map[string]int, error) {
	stats, err := uc.provider.PlayedHoursPerPlatform(ctx)
	if err != nil {
		return nil, err
	}
	return stats, nil
}
