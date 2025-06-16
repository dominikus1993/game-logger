package usecases

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRatingStatsProvider struct {
	mock.Mock
}

func (m *mockRatingStatsProvider) AvgRatingPerPlatform(ctx context.Context) (map[string]float64, error) {
	args := m.Called(ctx)
	return args.Get(0).(map[string]float64), args.Error(1)
}

func (m *mockRatingStatsProvider) AvgRatingPerYear(ctx context.Context) (map[int]float64, error) {
	args := m.Called(ctx)
	return args.Get(0).(map[int]float64), args.Error(1)
}

func TestRatingStatsUseCase_AvgRatingPerPlatform_Success(t *testing.T) {
	ctx := context.Background()
	mockProvider := new(mockRatingStatsProvider)
	expectedStats := map[string]float64{
		"PlayStation": 4.5,
		"Xbox":        4.2,
		"PC":          4.7,
	}
	mockProvider.On("AvgRatingPerPlatform", ctx).Return(expectedStats, nil)

	uc := NewRatingStatsUseCase(mockProvider)
	result, err := uc.AvgRatingPerPlatform(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedStats, result)
	mockProvider.AssertExpectations(t)
}

func TestRatingStatsUseCase_AvgRatingPerPlatform_Error(t *testing.T) {
	ctx := context.Background()
	mockProvider := new(mockRatingStatsProvider)
	mockProvider.On("AvgRatingPerPlatform", ctx).Return(map[string]float64(nil), assert.AnError)

	uc := NewRatingStatsUseCase(mockProvider)
	result, err := uc.AvgRatingPerPlatform(ctx)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockProvider.AssertExpectations(t)
}

func TestRatingStatsUseCase_AvgRatingPerYear_Success(t *testing.T) {
	ctx := context.Background()
	mockProvider := new(mockRatingStatsProvider)
	expectedStats := map[int]float64{
		2023: 4.3,
		2024: 4.6,
		2025: 4.8,
	}
	mockProvider.On("AvgRatingPerYear", ctx).Return(expectedStats, nil)

	uc := NewRatingStatsUseCase(mockProvider)
	result, err := uc.AvgRatingPerYear(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedStats, result)
	mockProvider.AssertExpectations(t)
}

func TestRatingStatsUseCase_AvgRatingPerYear_Error(t *testing.T) {
	ctx := context.Background()
	mockProvider := new(mockRatingStatsProvider)
	mockProvider.On("AvgRatingPerYear", ctx).Return(map[int]float64(nil), assert.AnError)

	uc := NewRatingStatsUseCase(mockProvider)
	result, err := uc.AvgRatingPerYear(ctx)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockProvider.AssertExpectations(t)
}
