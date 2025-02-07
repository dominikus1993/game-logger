using GameLogger.Infrastructure.Extensions;
using MongoDB.Driver;
using Testcontainers.MongoDb;

namespace GameLogger.Infrastructure.Tests.Fixture;

public sealed class MongoDbFixture : IAsyncLifetime
{
    private readonly MongoDbContainer _container = new MongoDbBuilder().Build();
    
    public IMongoClient Client { get; private set; }
    public IMongoDatabase Database { get; private set; }
    
    public async ValueTask InitializeAsync()
    {
        await _container.StartAsync();
        Client = new MongoClient(_container.GetConnectionString());
        Database = Client.GetDatabase("GamesLogger");
        Database.MapGamesCollection();
    }
    
    public void CleanDatabase()
    {
        Database.DropCollection(Database.GetGamesCollectionName());
    }

    public async ValueTask DisposeAsync()
    {
        await _container.StopAsync();
    }
}