package repo

import (
	"context"
	"fmt"
	"testing"
	"time"

	writer "github.com/dominikus1993/game-logger/internal/load/repo"
	"github.com/dominikus1993/game-logger/internal/mongo"
	"github.com/dominikus1993/game-logger/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
)

func TestAvgRatingPerPlatform(t *testing.T) {
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

	reader := NewRatingStatsProvider(client)

	t.Run("Read when no games exist", func(t *testing.T) {
		// Act
		ratings, err := reader.AvgRatingPerPlatform(ctx)
		// Assert
		assert.NoError(t, err)
		assert.Empty(t, ratings)
	})

	t.Run("Read when one game exists", func(t *testing.T) {
		// Act
		now := time.Now()
		rating := 8
		platform := "Switch"
		game := model.Game{
			Id:    "testGame1",
			Title: "testGame1",
			Playthroughs: []model.Playthrough{{
				StartDate: now,
				Platform:  platform,
				Rating:    &rating,
			}},
		}
		err := writer.WriteGame(ctx, &game)
		assert.NoError(t, err)

		ratings, err := reader.AvgRatingPerPlatform(ctx)
		assert.NoError(t, err)
		assert.NotEmpty(t, ratings)
		assert.Len(t, ratings, 1)
		assert.Equal(t, float64(rating), ratings[platform])
	})

	t.Run("Read when multiple games with same platform exist", func(t *testing.T) {
		// Act
		now := time.Now()
		rating1 := 8
		rating2 := 6
		platform := "PC"

		game1 := model.Game{
			Id:    "testGame2",
			Title: "testGame2",
			Playthroughs: []model.Playthrough{{
				StartDate: now,
				Platform:  platform,
				Rating:    &rating1,
			}},
		}
		game2 := model.Game{
			Id:    "testGame3",
			Title: "testGame3",
			Playthroughs: []model.Playthrough{{
				StartDate: now,
				Platform:  platform,
				Rating:    &rating2,
			}},
		}

		err := writer.WriteGame(ctx, &game1)
		assert.NoError(t, err)
		err = writer.WriteGame(ctx, &game2)
		assert.NoError(t, err)

		ratings, err := reader.AvgRatingPerPlatform(ctx)
		assert.NoError(t, err)
		assert.NotEmpty(t, ratings)
		assert.Contains(t, ratings, platform)
		expectedAvg := float64(rating1+rating2) / 2.0
		assert.Equal(t, expectedAvg, ratings[platform])
	})
}

func TestAvgRatingPerYear(t *testing.T) {
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

	reader := NewRatingStatsProvider(client)

	t.Run("Read when no games exist", func(t *testing.T) {
		// Act
		ratings, err := reader.AvgRatingPerYear(ctx)
		// Assert
		assert.NoError(t, err)
		assert.Empty(t, ratings)
	})

	t.Run("Read when one game exists", func(t *testing.T) {
		// Act
		now := time.Now()
		year := now.Year()
		rating := 9
		platform := "PS5"
		game := model.Game{
			Id:    "testGame4",
			Title: "testGame4",
			Playthroughs: []model.Playthrough{{
				StartDate: now,
				Platform:  platform,
				Rating:    &rating,
			}},
		}
		err := writer.WriteGame(ctx, &game)
		assert.NoError(t, err)

		ratings, err := reader.AvgRatingPerYear(ctx)
		assert.NoError(t, err)
		assert.NotEmpty(t, ratings)
		assert.Len(t, ratings, 1)
		assert.Equal(t, float64(rating), ratings[year])
	})

	t.Run("Read when multiple games with same year exist", func(t *testing.T) {
		// Act
		pastYear := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		year := pastYear.Year()
		rating1 := 7
		rating2 := 9

		game1 := model.Game{
			Id:    "testGame5",
			Title: "testGame5",
			Playthroughs: []model.Playthrough{{
				StartDate: pastYear,
				Platform:  "Xbox",
				Rating:    &rating1,
			}},
		}
		game2 := model.Game{
			Id:    "testGame6",
			Title: "testGame6",
			Playthroughs: []model.Playthrough{{
				StartDate: pastYear,
				Platform:  "Nintendo",
				Rating:    &rating2,
			}},
		}

		err := writer.WriteGame(ctx, &game1)
		assert.NoError(t, err)
		err = writer.WriteGame(ctx, &game2)
		assert.NoError(t, err)

		ratings, err := reader.AvgRatingPerYear(ctx)
		assert.NoError(t, err)
		assert.NotEmpty(t, ratings)
		assert.Contains(t, ratings, year)
		expectedAvg := float64(rating1+rating2) / 2.0
		assert.Equal(t, expectedAvg, ratings[year])
	})
}
