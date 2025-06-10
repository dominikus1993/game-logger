package repo

import "context"

type PlayedHoursPerPlatformStatsProvider interface {
	PlayedHoursPerPlatform(ctx context.Context) (map[string]int, error)
}
