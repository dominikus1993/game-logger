using GameLogger.Core.Data;

namespace GameLogger.Core.UseCases;

public sealed class GetGamesUseCase
{
    public IAsyncEnumerable<Game> Execute(CancellationToken cancellationToken = default)
    {
        return AsyncEnumerable.Empty<Game>();
    }
}