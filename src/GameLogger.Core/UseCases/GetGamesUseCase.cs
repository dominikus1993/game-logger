using GameLogger.Core.Data;

namespace GameLogger.Core.UseCases;

public sealed class GetGamesUseCase
{
    public Task<IReadOnlyList<Game>> Execute(CancellationToken cancellationToken = default)
    {
        return Task.FromResult<IReadOnlyList<Game>>([]);
    }
}