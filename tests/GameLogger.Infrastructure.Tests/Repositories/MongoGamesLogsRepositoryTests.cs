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

    public Task InitializeAsync()
    {
        return Task.CompletedTask;
    }

    public async Task DisposeAsync()
    {
        await _fixture.Database.DropCollectionAsync(_fixture.Database.GetGamesCollectionName());
    }
}