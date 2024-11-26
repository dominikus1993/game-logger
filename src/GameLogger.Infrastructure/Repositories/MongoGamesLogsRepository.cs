using GameLogger.Core.Data;
using GameLogger.Core.Repositories;
using GameLogger.Core.Types;
using GameLogger.Infrastructure.Extensions;
using MongoDB.Driver;

namespace GameLogger.Infrastructure.Repositories;

public sealed class MongoGamesLogsRepository : IGamesLogsRepository
{
    private readonly IMongoDatabase _mongoDatabase;
    private readonly IMongoCollection<Game> _games;

    public MongoGamesLogsRepository(IMongoDatabase mongoDatabase)
    {
        _mongoDatabase = mongoDatabase;
        _games = _mongoDatabase.Games();
    }

    public async Task<IReadOnlyList<Game>> GetGames(GetGamesQuery query, CancellationToken cancellationToken = default)
    {
        var limit = query.PageSize;
        var skip = query.PageSize * (query.Page - 1);
        var result = await _games.Find(FilterDefinition<Game>.Empty)
            .SortByDescending(x => x.StartDate)
            .Skip((int)skip)
            .Limit(limit)
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

    public async Task<Result<Unit>> UpdateGame(Game game, CancellationToken cancellationToken = default)
    {
        var filter = Builders<Game>.Filter.Eq(g => g.Id, game.Id);
        await _games.FindOneAndReplaceAsync(filter, game, new FindOneAndReplaceOptions<Game, Game>()
        {
            IsUpsert = true,
        }, cancellationToken);
        
        return Result.UnitResult;
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