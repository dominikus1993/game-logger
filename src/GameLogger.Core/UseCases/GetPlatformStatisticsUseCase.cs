using GameLogger.Core.Providers;
using GameLogger.Core.Types;

namespace GameLogger.Core.UseCases;

public sealed class GetPlatformStatisticsUseCase
{
    private readonly IGamesStatisticsProvider _gamesStatisticsProvider;

    public GetPlatformStatisticsUseCase(IGamesStatisticsProvider gamesStatisticsProvider)
    {
        _gamesStatisticsProvider = gamesStatisticsProvider;
    }

    public async Task<Result<IReadOnlyList<PlatformStatistics>>> Execute(CancellationToken cancellationToken = default)
    {
        return await _gamesStatisticsProvider.GetPlatformsStatistics(cancellationToken);
    }   
}