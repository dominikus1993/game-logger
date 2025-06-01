package repo

import (
	"context"

	"github.com/dominikus1993/game-logger/internal/mongo"
	"github.com/dominikus1993/game-logger/pkg/api/repo"
	"github.com/dominikus1993/game-logger/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoGamesReader struct {
	client *mongo.MongoClient
}

func NewMongoGamesReader(client *mongo.MongoClient) *MongoGamesReader {
	return &MongoGamesReader{client: client}
}

func (w *MongoGamesReader) LoadGames(ctx context.Context, query repo.LoadGamesQuery) ([]*model.Game, error) {
	col := w.client.GetCollection()
	filter := bson.M{}

	if query.Page < 1 {
		query.Page = 1
	}
	if query.Size < 1 {
		query.Size = 10
	}
	opts := options.Find().
		SetSkip(int64((query.Page - 1) * query.Size)).
		SetLimit(int64(query.Size)).
		SetSort(bson.M{"start_date": 1})

	cursor, err := col.Find(ctx, filter, opts)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var games []*model.Game
	for cursor.Next(ctx) {
		var mongoGame mongoGame
		if err := cursor.Decode(&mongoGame); err != nil {
			return nil, err
		}
		games = append(games, &model.Game{
			Id:          mongoGame.Id,
			Title:       mongoGame.Title,
			StartDate:   mongoGame.StartDate,
			FinishDate:  mongoGame.FinishDate,
			Platform:    mongoGame.Platform,
			HoursPlayed: mongoGame.HoursPlayed,
			Rating:      mongoGame.Rating,
			Notes:       mongoGame.Notes,
		})
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return games, nil
}

func (w *MongoGamesReader) Count(ctx context.Context) (int, error) {
	col := w.client.GetCollection()
	count, err := col.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

type mongoGame struct {
	Id          string `bson:"id"`
	Title       string `bson:"title"`
	StartDate   string `bson:"start_date"`
	FinishDate  string `bson:"finish_date,omitempty"`
	Platform    string `bson:"platform,omitempty"`
	HoursPlayed int    `bson:"hours_played,omitempty"`
	Rating      int    `bson:"rating,omitempty"`
	Notes       string `bson:"notes,omitempty"`
}
