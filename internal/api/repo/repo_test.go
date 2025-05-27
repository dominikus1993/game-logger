package repo

import (
	"context"
	"fmt"
	"testing"

	writer "github.com/dominikus1993/game-logger/internal/load/repo"
	"github.com/dominikus1993/game-logger/internal/mongo"
	"github.com/dominikus1993/game-logger/pkg/api/repo"
	"github.com/dominikus1993/game-logger/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
)

func TestLoadGame(t *testing.T) {
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
	writer := writer.NewMongoGamesWriter(client)

	reader := NewMongoGamesReader(client)

	t.Run("Read when no articles exist", func(t *testing.T) {
		// Act
		query := repo.LoadGamesQuery{
			Page: 1,
			Size: 10,
		}
		games, err := reader.LoadGames(ctx, query)
		// Assert
		assert.NoError(t, err)
		assert.Empty(t, games)
	})

	t.Run("Read when one article exists", func(t *testing.T) {
		// Act
		article := model.Game{
			Id:          "testArticle",
			Title:       "testArticle",
			StartDate:   "2023-10-01",
			FinishDate:  "2023-11-01",
			Platform:    "Switch",
			HoursPlayed: 25,
			Rating:      5,
			Notes:       "test notes",
		}
		err := writer.WriteGame(ctx, &article)
		assert.NoError(t, err)

		query := repo.LoadGamesQuery{
			Page: 1,
			Size: 10,
		}

		games, err := reader.LoadGames(ctx, query)
		assert.NoError(t, err)
		assert.NotEmpty(t, games)
		assert.Len(t, games, 1)
	})
}
