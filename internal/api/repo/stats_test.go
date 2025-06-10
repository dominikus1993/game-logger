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

func TestPlayedHoursPerPlatform(t *testing.T) {
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

	reader := NewPlayedHoursPerPlatformStatsProvider(client)

	t.Run("Read when no games exist", func(t *testing.T) {
		// Act
		games, err := reader.PlayedHoursPerPlatform(ctx)
		// Assert
		assert.NoError(t, err)
		assert.Empty(t, games)
	})

	t.Run("Read when one game exists", func(t *testing.T) {
		// Act
		now := time.Now()
		hoursPlayed := 25
		platform := "Switch"
		article := model.Game{
			Id:          "testArticle",
			Title:       "testArticle",
			StartDate:   now,
			FinishDate:  &now, // 1 day later
			Platform:    platform,
			Notes:       "test notes",
			HoursPlayed: &hoursPlayed,
		}
		err := writer.WriteGame(ctx, &article)
		assert.NoError(t, err)

		games, err := reader.PlayedHoursPerPlatform(ctx)
		assert.NoError(t, err)
		assert.NotEmpty(t, games)
		assert.Len(t, games, 1)
		assert.Equal(t, hoursPlayed, games[platform])
	})
}
