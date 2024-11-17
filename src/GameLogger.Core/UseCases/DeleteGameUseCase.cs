using GameLogger.Core.Repositories;
using GameLogger.Core.Types;

namespace GameLogger.Core.UseCases;

public sealed class DeleteGameUseCase
{
    private readonly IGamesLogsRepository _gamesLogs;

    public DeleteGameUseCase(IGamesLogsRepository gamesLogs)
    {
        _gamesLogs = gamesLogs;
    }

    public async Task<Result<Unit>> Execute(Guid gameId, CancellationToken cancellationToken = default)
    {
        var result = await _gamesLogs.DeleteGame(gameId, cancellationToken);
        return result;
    }
}