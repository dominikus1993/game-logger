using GameLogger.Core.Data;
using GameLogger.Core.Repositories;
using GameLogger.Core.Types;

namespace GameLogger.Core.UseCases;

public sealed class UpdateGameUseCase
{
    private readonly IGamesLogsRepository _gamesLogsRepository;

    public UpdateGameUseCase(IGamesLogsRepository gamesLogsRepository)
    {
        _gamesLogsRepository = gamesLogsRepository;
    }

    public async Task<Result<Unit>> Execute(Game game, CancellationToken cancellationToken = default)
    {
        ArgumentNullException.ThrowIfNull(game);
        return await _gamesLogsRepository.UpdateGame(game, cancellationToken);
    }
}