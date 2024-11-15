using GameLogger.Core.Data;
using MongoDB.Driver;

namespace GameLogger.Infrastructure.Extensions;

public static class MongoDatabaseExtensions
{
    public static string GetGamesCollectionName(this IMongoDatabase _) => "Games";
    
    public static IMongoCollection<Game> Games(this IMongoDatabase database)
    {
        return database.GetCollection<Game>(database.GetGamesCollectionName());
    }
}