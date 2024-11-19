using GameLogger.Core.Types;

namespace GameLogger.Core.Providers;

public sealed class PlatformStatistics
{
    public string Platform { get; init; }
    public int GamesCount { get; init; }
    public double? AvgHoursPlayed { get; init; }
}

public interface IGamesStatisticsProvider
{
    Task<Result<IReadOnlyList<PlatformStatistics>>> GetPlatformsStatistics(CancellationToken cancellationToken = default);
}