package repo

import (
	"context"

	"github.com/dominikus1993/game-logger/internal/mongo"
	"github.com/dominikus1993/game-logger/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoGamesWriter struct {
	client *mongo.MongoClient
}

func NewMongoGamesWriter(client *mongo.MongoClient) *MongoGamesWriter {
	return &MongoGamesWriter{client: client}
}

func (w *MongoGamesWriter) WriteGame(ctx context.Context, game *model.Game) error {
	filter := bson.M{"id": game.Id}
	model := newMongoGame(game)
	col := w.client.GetCollection()
	_, err := col.ReplaceOne(ctx, filter, model, options.Replace().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
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

func newMongoGame(game *model.Game) *mongoGame {
	return &mongoGame{
		Id:          game.Id,
		Title:       game.Title,
		StartDate:   game.StartDate,
		FinishDate:  game.FinishDate,
		Platform:    game.Platform,
		HoursPlayed: game.HoursPlayed,
		Rating:      game.Rating,
		Notes:       game.Notes,
	}
}
