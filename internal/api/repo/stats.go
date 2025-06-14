package repo

import (
	"context"

	"github.com/dominikus1993/game-logger/internal/mongo"
)

type playedHoursPerPlatformStatsProvider struct {
	client *mongo.MongoClient
}

func NewPlayedHoursPerPlatformStatsProvider(client *mongo.MongoClient) *playedHoursPerPlatformStatsProvider {
	return &playedHoursPerPlatformStatsProvider{client: client}
}

func (p *playedHoursPerPlatformStatsProvider) PlayedHoursPerPlatform(ctx context.Context) (map[string]int, error) {
	collection := p.client.GetCollection()
	pipeline := []map[string]any{
		{"$group": map[string]any{
			"_id":   "$platform",
			"total": map[string]string{"$sum": "$hours_played"},
		}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	stats := make(map[string]int)
	for cursor.Next(ctx) {
		var result struct {
			ID    string `bson:"_id"`
			Total int    `bson:"total"`
		}
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		stats[result.ID] = result.Total
	}

	return stats, nil
}

func (p *playedHoursPerPlatformStatsProvider) PlayedHoursPerYear(ctx context.Context) (map[int]int, error) {
	collection := p.client.GetCollection()
	pipeline := []map[string]any{
		{"$group": map[string]any{
			"_id":   map[string]string{"$year": "$start_date"},
			"total": map[string]string{"$sum": "$hours_played"},
		}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	stats := make(map[int]int)
	for cursor.Next(ctx) {
		var result struct {
			ID    int `bson:"_id"`
			Total int `bson:"total"`
		}
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		stats[result.ID] = result.Total
	}

	return stats, nil
}
