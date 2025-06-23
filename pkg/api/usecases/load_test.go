package usecases

import (
	"context"
	"errors"
	"testing"

	"github.com/dominikus1993/game-logger/pkg/api/repo"
	"github.com/dominikus1993/game-logger/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockGamesReader struct {
	mock.Mock
}

func (m *mockGamesReader) LoadGames(ctx context.Context, query repo.LoadGamesQuery) ([]*model.Game, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]*model.Game), args.Error(1)
}

func (m *mockGamesReader) Count(ctx context.Context) (int, error) {
	args := m.Called(ctx)
	return args.Int(0), args.Error(1)
}

func TestLoadGamesUseCase_Execute_Success(t *testing.T) {
	ctx := context.Background()
	mockReader := new(mockGamesReader)
	games := []*model.Game{{Id: "1", Title: "Game1"}, {Id: "2", Title: "Game2"}}
	mockReader.On("LoadGames", ctx, repo.LoadGamesQuery{Page: 1, Size: 2}).Return(games, nil)
	mockReader.On("Count", ctx).Return(2, nil)

	uc, err := NewLoadGamesUseCase(mockReader)
	assert.NoError(t, err)
	resp, err := uc.Execute(ctx, LoadGamesQuery{Page: 1, Size: 2})

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 2, resp.Total)
	assert.Len(t, resp.Games, 2)
	assert.Equal(t, "Game1", resp.Games[0].Title)
	mockReader.AssertExpectations(t)
}

func TestLoadGamesUseCase_Execute_LoadGamesError(t *testing.T) {
	ctx := context.Background()
	var games []*model.Game = nil
	mockReader := new(mockGamesReader)
	mockReader.On("LoadGames", ctx, repo.LoadGamesQuery{Page: 1, Size: 2}).Return(games, errors.New("failed"))

	uc, err := NewLoadGamesUseCase(mockReader)
	assert.NoError(t, err)
	resp, err := uc.Execute(ctx, LoadGamesQuery{Page: 1, Size: 2})

	assert.Error(t, err)
	assert.Nil(t, resp)
	mockReader.AssertExpectations(t)
}

func TestLoadGamesUseCase_Execute_CountError(t *testing.T) {
	ctx := context.Background()
	mockReader := new(mockGamesReader)
	games := []*model.Game{{Id: "1", Title: "Game1"}}
	mockReader.On("LoadGames", ctx, repo.LoadGamesQuery{Page: 1, Size: 1}).Return(games, nil)
	mockReader.On("Count", ctx).Return(0, assert.AnError)

	uc, err := NewLoadGamesUseCase(mockReader)
	assert.NoError(t, err)
	resp, err := uc.Execute(ctx, LoadGamesQuery{Page: 1, Size: 1})

	assert.Error(t, err)
	assert.Nil(t, resp)
	mockReader.AssertExpectations(t)
}
