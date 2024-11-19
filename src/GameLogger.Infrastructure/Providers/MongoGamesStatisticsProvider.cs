using GameLogger.Core.Data;
using GameLogger.Core.Providers;
using GameLogger.Core.Types;
using GameLogger.Infrastructure.Extensions;
using MongoDB.Driver;
using MongoDB.Driver.Linq;

namespace GameLogger.Infrastructure.Providers;

public sealed class MongoGamesStatisticsProvider : IGamesStatisticsProvider
{
    private readonly IMongoDatabase _mongoDatabase;
    private readonly IMongoCollection<Game> _games;

    public MongoGamesStatisticsProvider(IMongoDatabase mongoDatabase)
    {
        _mongoDatabase = mongoDatabase;
        _games = _mongoDatabase.Games();
    }
    
    public async Task<Result<IReadOnlyList<PlatformStatistics>>> GetPlatformsStatistics(CancellationToken cancellationToken = default)
    {
        var query = _games.AsQueryable().GroupBy(x => x.Platform).Select(x => new PlatformStatistics() { Platform = x.Key, GamesCount = x.Count() });

        var result = await query.ToListAsync(cancellationToken: cancellationToken);
        
        return Result.Ok<IReadOnlyList<PlatformStatistics>>(result);
    }
}