using GameLogger.Core.Data;

namespace GameLogger.Core.Providers;

public interface IGamesDataProvider
{
    IAsyncEnumerable<Game> Provide(CancellationToken cancellationToken = default);
}