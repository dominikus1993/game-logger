using GameLogger.Core.Data;
using GameLogger.Core.Providers;
using GameLogger.Core.Types;

namespace GameLogger.Core.Repositories;

public sealed record GetGamesQuery(uint Page, ushort PageSize);
public interface IGamesLogsRepository
{
    Task<IReadOnlyList<Game>> GetGames(GetGamesQuery query, CancellationToken cancellationToken = default);
    
    Task<Result<Unit>> WriteGame(Game game, CancellationToken cancellationToken = default);
    Task<Result<Unit>> UpdateGame(Game game, CancellationToken cancellationToken = default);
    
    Task<Result<Unit>> DeleteGame(Guid id, CancellationToken cancellationToken = default);
}

public sealed class FakeGamesLogsRepository : IGamesLogsRepository
{
    private readonly IGamesDataProvider _gamesDataProvider;

    public FakeGamesLogsRepository(IGamesDataProvider gamesDataProvider)
    {
        _gamesDataProvider = gamesDataProvider;
    }

    public async Task<IReadOnlyList<Game>> GetGames(GetGamesQuery query, CancellationToken cancellationToken = default)
    {
        return await _gamesDataProvider.Provide(cancellationToken).ToArrayAsync(cancellationToken: cancellationToken);
    }

    public Task<Result<Unit>> WriteGame(Game game, CancellationToken cancellationToken = default)
    {
        return Task.FromResult(Result.UnitResult);
    }

    public Task<Result<Unit>> UpdateGame(Game game, CancellationToken cancellationToken = default)
    {
        throw new NotImplementedException();
    }

    public Task<Result<Unit>> DeleteGame(Guid id, CancellationToken cancellationToken = default)
    {
        return Task.FromResult(Result.UnitResult);
    }
}