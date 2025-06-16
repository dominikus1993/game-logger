package repo

import "context"

type PlayedHoursStatsProvider interface {
	PlayedHoursPerPlatform(ctx context.Context) (map[string]int, error)
	PlayedHoursPerYear(ctx context.Context) (map[int]int, error)
}

type RatingStatsProvider interface {
	AvgRatingPerPlatform(ctx context.Context) (map[string]float64, error)
	AvgRatingPerYear(ctx context.Context) (map[int]float64, error)
}
