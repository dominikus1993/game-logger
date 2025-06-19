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

type playedHoursPerYearUseCase struct {
	provider repo.PlayedHoursStatsProvider
}

func NewPlayedHoursPerYearUseCase(provider repo.PlayedHoursStatsProvider) *playedHoursPerYearUseCase {
	return &playedHoursPerYearUseCase{provider: provider}
}

func (uc *playedHoursPerYearUseCase) Execute(ctx context.Context) (map[int]int, error) {
	stats, err := uc.provider.PlayedHoursPerYear(ctx)
	if err != nil {
		return nil, err
	}
	return stats, nil
}

type RatingStatsUseCase struct {
	provider repo.RatingStatsProvider
}

func NewRatingStatsUseCase(provider repo.RatingStatsProvider) *RatingStatsUseCase {
	return &RatingStatsUseCase{provider: provider}
}

func (uc *RatingStatsUseCase) AvgRatingPerPlatform(ctx context.Context) (map[string]float64, error) {
	stats, err := uc.provider.AvgRatingPerPlatform(ctx)
	if err != nil {
		return nil, err
	}
	return stats, nil
}

func (uc *RatingStatsUseCase) AvgRatingPerYear(ctx context.Context) (map[int]float64, error) {
	stats, err := uc.provider.AvgRatingPerYear(ctx)
	if err != nil {
		return nil, err
	}
	return stats, nil
}
