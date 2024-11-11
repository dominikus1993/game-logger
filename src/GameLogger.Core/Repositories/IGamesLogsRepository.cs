using GameLogger.Core.Data;

namespace GameLogger.Core.Repositories;

public sealed record GetGamesQuery(int Page, int PageSize);
public interface IGamesLogsRepository
{
    Task<IReadOnlyList<Game>> GetGames(GetGamesQuery query, CancellationToken cancellationToken = default);
}

public sealed class FakeGamesLogsRepository : IGamesLogsRepository
{
    public Task<IReadOnlyList<Game>> GetGames(GetGamesQuery query, CancellationToken cancellationToken = default)
    {
        return Task.FromResult<IReadOnlyList<Game>>([]);
    }
}