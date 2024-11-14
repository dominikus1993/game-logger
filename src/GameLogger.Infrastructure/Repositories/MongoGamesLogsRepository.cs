using GameLogger.Core.Data;
using GameLogger.Core.Repositories;
using GameLogger.Core.Types;
using MongoDB.Driver;

namespace GameLogger.Infrastructure.Repositories;

public sealed class MongoGamesLogsRepository : IGamesLogsRepository
{
    private const string MongoCollectionName = "games";
    private readonly IMongoDatabase _mongoDatabase;
    private readonly IMongoCollection<Game> _games;

    public MongoGamesLogsRepository(IMongoDatabase mongoDatabase)
    {
        _mongoDatabase = mongoDatabase;
        _games = _mongoDatabase.GetCollection<Game>(MongoCollectionName);
    }

    public async Task<IReadOnlyList<Game>> GetGames(GetGamesQuery query, CancellationToken cancellationToken = default)
    {
        var result = await _games.Find(FilterDefinition<Game>.Empty)
            .Skip(query.PageSize * query.Page)
            .Limit(query.PageSize)
            .ToListAsync(cancellationToken);

        return result;
    }

    public async Task<Result<Unit>> WriteGame(Game game, CancellationToken cancellationToken = default)
    {
        try
        {
            await _games.InsertOneAsync(game, cancellationToken: cancellationToken);
            return Result.UnitResult;
        }
        catch (Exception e)
        {
            return Result.Failure<Unit>(e);
        }
    }

    public async Task<Result<Unit>> DeleteGame(Guid id, CancellationToken cancellationToken = default)
    {
        var filter = Builders<Game>.Filter.Eq(g => g.Id, id);
        var result = await _games.DeleteOneAsync(filter, cancellationToken: cancellationToken);
        
        if (result.DeletedCount == 0)
        {
            return Result.Failure<Unit>(new InvalidOperationException("Game not found"));
        }
        
        return Result.UnitResult;
    }
}