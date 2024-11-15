using MongoDB.Driver;
using Testcontainers.MongoDb;

namespace GameLogger.Infrastructure.Tests.Fixture;

public sealed class MongoDbFixture : IAsyncLifetime
{
    private readonly MongoDbContainer _container = new MongoDbBuilder().Build();
    
    public IMongoClient Client { get; private set; }
    public IMongoDatabase Database { get; private set; }
    
    public async Task InitializeAsync()
    {
        await _container.StartAsync();
        Client = new MongoClient(_container.GetConnectionString());
        Database = Client.GetDatabase("GamesLogger");
    }

    public Task DisposeAsync()
    {
        return _container.StopAsync();
    }
}