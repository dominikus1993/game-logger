package repo

import (
	"context"
	"time"

	"github.com/dominikus1993/game-logger/internal/mongo"
	"github.com/dominikus1993/game-logger/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	mong "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoGamesWriter struct {
	client *mongo.MongoClient
}

func NewMongoGamesWriter(client *mongo.MongoClient) *MongoGamesWriter {
	return &MongoGamesWriter{client: client}
}

func (w *MongoGamesWriter) WriteGame(ctx context.Context, game *model.Game) error {
	filter := bson.M{"_id": game.Id}
	model := newMongoGame(game)
	col := w.client.GetCollection()
	dbGame := col.FindOne(ctx, filter, options.FindOne())

	if dbGame.Err() == mong.ErrNoDocuments {

	}

	var gameFromDb mongoGame
	err := dbGame.Decode(&gameFromDb)
	if err == nil {

	}

	_, err := col.ReplaceOne(ctx, filter, model, options.Replace().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}

type mongoGame struct {
	Id           string             `bson:"_id"`
	Title        string             `bson:"title"`
	Playthroughs []mongoPlaythrough `bson:"playthroughs,omitempty"`
}

type mongoPlaythrough struct {
	StartDate   time.Time  `json:"start_date"`
	FinishDate  *time.Time `json:"finish_date,omitempty"`
	Platform    string     `json:"platform,omitempty"`
	HoursPlayed *int       `json:"hours_played,omitempty"`
	Rating      *int       `json:"rating,omitempty"`
	Notes       string     `json:"notes,omitempty"`
}

func newMongoGame(game *model.Game) *mongoGame {
	mongoGame := mongoGame{
		Id:    game.Id,
		Title: game.Title,
	}

	for _, pt := range game.Playthroughs {
		mongoPt := mongoPlaythrough{
			StartDate:   pt.StartDate,
			FinishDate:  pt.FinishDate,
			Platform:    pt.Platform,
			HoursPlayed: pt.HoursPlayed,
			Rating:      pt.Rating,
			Notes:       pt.Notes,
		}
		mongoGame.Playthroughs = append(mongoGame.Playthroughs, mongoPt)
	}

	return &mongoGame
}
