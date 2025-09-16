package repo

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/dominikus1993/game-logger/internal/mongo"
	"github.com/dominikus1993/game-logger/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
)

func TestWriteGame(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	// Arrange
	ctx := context.Background()
	mongodbContainer, err := mongodb.Run(ctx, "mongo:6")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := mongodbContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	connectionString, err := mongodbContainer.ConnectionString(ctx)
	if err != nil {
		t.Fatal(fmt.Errorf("can't download mongo conectionstring, %w", err))
	}
	client, err := mongo.NewClient(ctx, connectionString, "Games", "games")
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close(ctx)
	repo := NewMongoGamesWriter(client)

	t.Run("Write article once", func(t *testing.T) {
		// Act
		now := time.Now()
		hoursPlayed := 25
		rating := 5
		article := model.Game{
			Id:    "testArticle",
			Title: "testArticle",
			Playthroughs: []model.Playthrough{
				{
					StartDate:   now,
					FinishDate:  &now,
					Platform:    "Switch",
					HoursPlayed: &hoursPlayed,
					Rating:      &rating,
					Notes:       "test notes",
				},
			},
		}
		err := repo.WriteGame(ctx, &article)
		assert.NoError(t, err)
	})
}
