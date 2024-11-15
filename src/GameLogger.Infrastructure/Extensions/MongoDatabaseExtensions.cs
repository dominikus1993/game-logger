using GameLogger.Core.Data;
using MongoDB.Bson;
using MongoDB.Bson.Serialization;
using MongoDB.Bson.Serialization.Serializers;
using MongoDB.Driver;

namespace GameLogger.Infrastructure.Extensions;

public static class MongoDatabaseExtensions
{
    public static string GetGamesCollectionName(this IMongoDatabase _) => "Games";
    
    public static IMongoCollection<Game> Games(this IMongoDatabase database)
    {
        return database.GetCollection<Game>(database.GetGamesCollectionName());
    }
    
    public static void MapGamesCollection(this IMongoDatabase database)
    {
        BsonClassMap.RegisterClassMap<Game>(classMap =>
        {
            classMap.MapIdField(g => g.Id);
            var guidSerializer = new GuidSerializer(GuidRepresentation.Standard);
            classMap.MapField(g => g.Id).SetSerializer(guidSerializer);
            classMap.AutoMap();
        });
    }
}