package repo

import (
	"context"

	"github.com/dominikus1993/game-logger/internal/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type playedHoursPerPlatformStatsProvider struct {
	client *mongo.MongoClient
}

func NewPlayedHoursPerPlatformStatsProvider(client *mongo.MongoClient) *playedHoursPerPlatformStatsProvider {
	return &playedHoursPerPlatformStatsProvider{client: client}
}

func (p *playedHoursPerPlatformStatsProvider) PlayedHoursPerPlatform(ctx context.Context) (map[string]int, error) {
	collection := p.client.GetCollection()
	pipeline := []bson.M{
		{"$group": bson.M{
			"_id":   "$platform",
			"total": bson.M{"$sum": "$hours_played"},
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
	pipeline := []bson.M{
		{"$group": bson.M{
			"_id":   bson.M{"$year": "$playthroughs.start_date"},
			"total": bson.M{"$sum": "$playthroughs.hours_played"},
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

type ratingStatsProvider struct {
	client *mongo.MongoClient
}

func NewRatingStatsProvider(client *mongo.MongoClient) *ratingStatsProvider {
	return &ratingStatsProvider{client: client}
}

func (r *ratingStatsProvider) AvgRatingPerPlatform(ctx context.Context) (map[string]float64, error) {
	collection := r.client.GetCollection()
	pipeline := []bson.M{
		{"$match": bson.M{"playthroughs.rating": bson.M{"$ne": nil}}}, // Only include games with ratings
		{"$group": bson.M{
			"_id":       "$playthroughs.platform",
			"avgRating": bson.M{"$avg": "$playthroughs.rating"},
		}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	stats := make(map[string]float64)
	for cursor.Next(ctx) {
		var result struct {
			ID        string  `bson:"_id"`
			AvgRating float64 `bson:"avgRating"`
		}
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		stats[result.ID] = result.AvgRating
	}

	return stats, nil
}

func (r *ratingStatsProvider) AvgRatingPerYear(ctx context.Context) (map[int]float64, error) {
	collection := r.client.GetCollection()
	pipeline := []bson.M{
		{"$match": bson.M{"playthroughs.rating": bson.M{"$ne": nil}}}, // Only include games with ratings
		{"$group": bson.M{
			"_id":       bson.M{"$year": "$playthroughs.start_date"},
			"avgRating": bson.M{"$avg": "$playthroughs.rating"},
		}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	stats := make(map[int]float64)
	for cursor.Next(ctx) {
		var result struct {
			ID        int     `bson:"_id"`
			AvgRating float64 `bson:"avgRating"`
		}
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		stats[result.ID] = result.AvgRating
	}

	return stats, nil
}
