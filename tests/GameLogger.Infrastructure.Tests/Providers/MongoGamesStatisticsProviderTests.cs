using GameLogger.Core.Data;
using GameLogger.Core.Providers;
using GameLogger.Core.Repositories;
using GameLogger.Infrastructure.Extensions;
using GameLogger.Infrastructure.Providers;
using GameLogger.Infrastructure.Repositories;
using GameLogger.Infrastructure.Tests.Fixture;

namespace GameLogger.Infrastructure.Tests.Providers;

public class MongoGamesStatisticsProviderTests: IClassFixture<MongoDbFixture>, IAsyncLifetime
{
    private readonly IGamesStatisticsProvider _repository;
    private readonly IGamesLogsRepository _gamesLogsRepository;
    private readonly MongoDbFixture _fixture;
    
    public MongoGamesStatisticsProviderTests(MongoDbFixture fixture)
    {
        _fixture = fixture;
        _repository = new MongoGamesStatisticsProvider(fixture.Database);
        _gamesLogsRepository = new MongoGamesLogsRepository(fixture.Database);
    }
    
    [Fact]
    public async Task GetPlatformsStatisticsWhenCollectionIsEmpty()
    {
        var platformsStatistics = await _repository.GetPlatformsStatistics();
        
        Assert.True(platformsStatistics.IsSuccess);
        Assert.Empty(platformsStatistics.Value);
    }

    [Fact]
    public async Task GetPlatformsStatistics()
    {
        // Arrange
        Game[] games =
        [
            new Game()
            {
                Id = Guid.CreateVersion7(), Title = "Cyberpunk 2077", Rating = 8, Platform = "PC",
                StartDate = DateOnly.FromDateTime(DateTime.Now),
                FinishDate = DateOnly.FromDateTime(DateTime.Now.AddDays(7)), HoursPlayed = 40
            },
            new Game()
            {
                Id = Guid.CreateVersion7(), Title = "Wiedzmin 3", Rating = 8, Platform = "XSX",
                StartDate = DateOnly.FromDateTime(DateTime.Now),
                FinishDate = DateOnly.FromDateTime(DateTime.Now.AddDays(7)), HoursPlayed = 40
            },
            new Game()
            {
                Id = Guid.CreateVersion7(), Title = "Mario Odyssey", Rating = 8, Platform = "Switch",
                StartDate = DateOnly.FromDateTime(DateTime.Now),
                FinishDate = DateOnly.FromDateTime(DateTime.Now.AddDays(7)), HoursPlayed = 40
            }
        ];
        
        foreach (var game in games)
        {
            var writeResult = await _gamesLogsRepository.WriteGame(game);
            Assert.True(writeResult.IsSuccess);
        }
        
        // Act
        var platformsStatistics = await _repository.GetPlatformsStatistics();
        
        // Assert
        Assert.True(platformsStatistics.IsSuccess);
        Assert.NotEmpty(platformsStatistics.Value);
        
        await Verify(platformsStatistics.Value);
    }
    

    public Task InitializeAsync()
    {
        return Task.CompletedTask;
    }

    public async Task DisposeAsync()
    {
        await _fixture.Database.DropCollectionAsync(_fixture.Database.GetGamesCollectionName());
    }
}