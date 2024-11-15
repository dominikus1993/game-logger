using GameLogger.Core.Data;
using GameLogger.Core.Repositories;
using GameLogger.Infrastructure.Extensions;
using GameLogger.Infrastructure.Repositories;
using GameLogger.Infrastructure.Tests.Fixture;

namespace GameLogger.Infrastructure.Tests.Repositories;

public class MongoGamesLogsRepositoryTests : IClassFixture<MongoDbFixture>, IAsyncLifetime
{
    private readonly IGamesLogsRepository _repository;
    private readonly MongoDbFixture _fixture;
    
    public MongoGamesLogsRepositoryTests(MongoDbFixture fixture)
    {
        _fixture = fixture;
        _repository = new MongoGamesLogsRepository(fixture.Database);
    }
    
    [Fact]
    public async Task GetGamesWhenCollectionIsEmpty()
    {
        var games = await _repository.GetGames(new GetGamesQuery(1, 10));
        
        Assert.Empty(games);
    }
    
    [Fact]
    public async Task GetGamesWhenCollectionIsNotEmpty()
    {
        // Arrange

        var game = new Game()
        {
            Id = Guid.NewGuid(), Title = "Cyberpunk 2077", Rating = 8, Platform = "PC",
            StartDate = DateOnly.FromDateTime(DateTime.Now),
            FinishDate = DateOnly.FromDateTime(DateTime.Now.AddDays(7)), HoursPlayed = 40
        };
        
        var writeResult = await _repository.WriteGame(game);
        Assert.True(writeResult.IsSuccess);
        
        var games = await _repository.GetGames(new GetGamesQuery(1, 10));
        
        Assert.NotEmpty(games);
        Assert.Single(games);
        var gameFromDb = games[0];
        Assert.Equivalent(game, gameFromDb);
    }
    
    [Fact]
    public async Task DeleteGameWhenExists()
    {
        // Arrange

        var game = new Game()
        {
            Id = Guid.NewGuid(), Title = "Cyberpunk 2077", Rating = 8, Platform = "PC",
            StartDate = DateOnly.FromDateTime(DateTime.Now),
            FinishDate = DateOnly.FromDateTime(DateTime.Now.AddDays(7)), HoursPlayed = 40
        };
        
        var writeResult = await _repository.WriteGame(game);
        Assert.True(writeResult.IsSuccess);
        
        // Act
        
        var deleteResult = await _repository.DeleteGame(game.Id);
        Assert.True(deleteResult.IsSuccess);
        
        var games = await _repository.GetGames(new GetGamesQuery(1, 10));
        
        Assert.Empty(games);
        
    }
    
    [Fact]
    public async Task DeleteGameWhenNoExists()
    {
        // Arrange

        var game = new Game()
        {
            Id = Guid.NewGuid(), Title = "Cyberpunk 2077", Rating = 8, Platform = "PC",
            StartDate = DateOnly.FromDateTime(DateTime.Now),
            FinishDate = DateOnly.FromDateTime(DateTime.Now.AddDays(7)), HoursPlayed = 40
        };
        
        var writeResult = await _repository.WriteGame(game);
        Assert.True(writeResult.IsSuccess);
        
        // Act
        
        var deleteResult = await _repository.DeleteGame(Guid.NewGuid());
        Assert.False(deleteResult.IsSuccess);
        Assert.IsType<InvalidOperationException>(deleteResult.ErrorValue);
        
        var games = await _repository.GetGames(new GetGamesQuery(1, 10));
        
        Assert.NotEmpty(games);
        Assert.Single(games);
        
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