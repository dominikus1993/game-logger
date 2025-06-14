package repo

import "context"

type PlayedHoursStatsProvider interface {
	PlayedHoursPerPlatform(ctx context.Context) (map[string]int, error)
	PlayedHoursPerYear(ctx context.Context) (map[int]int, error)
}
