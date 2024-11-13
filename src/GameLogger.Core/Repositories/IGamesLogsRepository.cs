using GameLogger.Core.Data;
using GameLogger.Core.Providers;

namespace GameLogger.Core.Repositories;

public sealed record GetGamesQuery(int Page, int PageSize);
public interface IGamesLogsRepository
{
    Task<IReadOnlyList<Game>> GetGames(GetGamesQuery query, CancellationToken cancellationToken = default);
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
}