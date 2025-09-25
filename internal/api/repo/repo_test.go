package repo

import (
	"context"
	"fmt"
	"testing"
	"time"

	writer "github.com/dominikus1993/game-logger/internal/load/repo"
	"github.com/dominikus1993/game-logger/internal/mongo"
	"github.com/dominikus1993/game-logger/pkg/api/repo"
	"github.com/dominikus1993/game-logger/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
)

func TestCount(t *testing.T) {
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

	reader, err := NewMongoGamesReader(client)

	if err != nil {
		t.Fatal(err)
	}

	t.Run("Read when no articles exist", func(t *testing.T) {
		// Act
		count, err := reader.Count(ctx)
		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 0, count)
	})

	t.Run("Read when one article exists", func(t *testing.T) {
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
			}}
		err := writer.WriteGame(ctx, &article)
		assert.NoError(t, err)

		count, err := reader.Count(ctx)
		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 1, count)
	})
}

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

	reader, err := NewMongoGamesReader(client)

	if err != nil {
		t.Fatal(err)
	}

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
		now := time.Now()
		article := model.Game{
			Id:    "testArticle",
			Title: "testArticle",
			Playthroughs: []model.Playthrough{{
				StartDate:  now,
				FinishDate: &now, // 1 day later
				Platform:   "Switch",
				Notes:      "test notes",
			}},
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
