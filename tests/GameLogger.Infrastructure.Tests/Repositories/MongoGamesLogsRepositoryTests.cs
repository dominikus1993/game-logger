using GameLogger.Core.Data;
using GameLogger.Core.Repositories;
using GameLogger.Infrastructure.Extensions;
using GameLogger.Infrastructure.Repositories;
using GameLogger.Infrastructure.Tests.Fixture;

namespace GameLogger.Infrastructure.Tests.Repositories;

public class MongoGamesLogsRepositoryTests : IClassFixture<MongoDbFixture>, IDisposable
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
        var games = await _repository.GetGames(new GetGamesQuery(1, 10), TestContext.Current.CancellationToken);
        
        Assert.Empty(games);
    }
    
    [Fact]
    public async Task GetGamesWhenCollectionIsNotEmpty()
    {
        // Arrange

        var game = new Game()
        {
            Id = Guid.CreateVersion7(), Title = "Cyberpunk 2077", Rating = 8, Platform = "PC",
            StartDate = DateOnly.FromDateTime(DateTime.Now),
            FinishDate = DateOnly.FromDateTime(DateTime.Now.AddDays(7)), HoursPlayed = 40
        };
        
        var writeResult = await _repository.WriteGame(game, TestContext.Current.CancellationToken);
        Assert.True(writeResult.IsSuccess);
        
        var games = await _repository.GetGames(new GetGamesQuery(1, 10), TestContext.Current.CancellationToken);
        
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
            Id = Guid.CreateVersion7(), Title = "Cyberpunk 2077", Rating = 8, Platform = "PC",
            StartDate = DateOnly.FromDateTime(DateTime.Now),
            FinishDate = DateOnly.FromDateTime(DateTime.Now.AddDays(7)), HoursPlayed = 40
        };
        
        var writeResult = await _repository.WriteGame(game, TestContext.Current.CancellationToken);
        Assert.True(writeResult.IsSuccess);
        
        // Act
        
        var deleteResult = await _repository.DeleteGame(game.Id, TestContext.Current.CancellationToken);
        Assert.True(deleteResult.IsSuccess);
        
        var games = await _repository.GetGames(new GetGamesQuery(1, 10), TestContext.Current.CancellationToken);
        
        Assert.Empty(games);
        
    }
    
    [Fact]
    public async Task DeleteGameWhenNoExists()
    {
        // Arrange

        var game = new Game()
        {
            Id = Guid.CreateVersion7(), Title = "Cyberpunk 2077", Rating = 8, Platform = "PC",
            StartDate = DateOnly.FromDateTime(DateTime.Now),
            FinishDate = DateOnly.FromDateTime(DateTime.Now.AddDays(7)), HoursPlayed = 40
        };
        
        var writeResult = await _repository.WriteGame(game, TestContext.Current.CancellationToken);
        Assert.True(writeResult.IsSuccess);
        
        // Act
        
        var deleteResult = await _repository.DeleteGame(Guid.CreateVersion7(), TestContext.Current.CancellationToken);
        Assert.False(deleteResult.IsSuccess);
        Assert.IsType<InvalidOperationException>(deleteResult.ErrorValue);
        
        var games = await _repository.GetGames(new GetGamesQuery(1, 10), TestContext.Current.CancellationToken);
        
        Assert.NotEmpty(games);
        Assert.Single(games);
        
    }
    
    
    [Fact]
    public async Task UpdateGameWhenExists()
    {
        // Arrange

        var game = new Game()
        {
            Id = Guid.CreateVersion7(), Title = "Cyberpunk 2077", Rating = 8, Platform = "PC",
            StartDate = DateOnly.FromDateTime(DateTime.Now),
            FinishDate = DateOnly.FromDateTime(DateTime.Now.AddDays(7)), HoursPlayed = 40
        };
        
        var writeResult = await _repository.WriteGame(game, TestContext.Current.CancellationToken);
        Assert.True(writeResult.IsSuccess);
        
        // Act

        var newGame = new Game()
        {
            Id = game.Id,
            HoursPlayed = 2112,
            Rating = 555,
            StartDate = game.StartDate,
            FinishDate = game.FinishDate,
            Platform = game.Platform,
            Title = game.Title,
        };
        
        var deleteResult = await _repository.UpdateGame(newGame, TestContext.Current.CancellationToken);
        Assert.True(deleteResult.IsSuccess);
        
        var games = await _repository.GetGames(new GetGamesQuery(1, 10), TestContext.Current.CancellationToken);
        
        Assert.NotEmpty(games);
        Assert.Single(games);
        var gameFromDb = games[0];
        Assert.Equivalent(newGame, gameFromDb);
    }
    
    [Fact]
    public async Task UpdateGameWhenNoExists()
    {
        // Arrange

        var game = new Game()
        {
            Id = Guid.CreateVersion7(), Title = "Cyberpunk 2077", Rating = 8, Platform = "PC",
            StartDate = DateOnly.FromDateTime(DateTime.Now),
            FinishDate = DateOnly.FromDateTime(DateTime.Now.AddDays(7)), HoursPlayed = 40
        };
        
        var writeResult = await _repository.UpdateGame(game, TestContext.Current.CancellationToken);
        Assert.True(writeResult.IsSuccess);
        
        // Act
        
        var deleteResult = await _repository.DeleteGame(Guid.CreateVersion7(), TestContext.Current.CancellationToken);
        Assert.False(deleteResult.IsSuccess);
        Assert.IsType<InvalidOperationException>(deleteResult.ErrorValue);
        
        var games = await _repository.GetGames(new GetGamesQuery(1, 10), TestContext.Current.CancellationToken);
        
        Assert.NotEmpty(games);
        Assert.Single(games);
        
    }

    public void Dispose()
    {
        _fixture.CleanDatabase();
    }
}