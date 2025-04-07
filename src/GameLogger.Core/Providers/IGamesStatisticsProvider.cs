using GameLogger.Core.Types;

namespace GameLogger.Core.Providers;

public sealed class PlatformStatistics
{
    public string Platform { get; init; }
    public int GamesCount { get; init; }
    public double? AvgHoursPlayed { get; init; }
    public double? AvgRating { get; set; }
}

public interface IGamesStatisticsProvider
{
    Task<Result<IReadOnlyList<PlatformStatistics>>> GetPlatformsStatistics(CancellationToken cancellationToken = default);
}

public sealed class FakeGamesStatisticsProvider : IGamesStatisticsProvider
{
    public Task<Result<IReadOnlyList<PlatformStatistics>>> GetPlatformsStatistics(CancellationToken cancellationToken = default)
    {
        var result = new List<PlatformStatistics>
        {
            new()
            {
                Platform = "PC",
                GamesCount = 10,
                AvgHoursPlayed = 20,
                AvgRating = 4.5
            },
            new()
            {
                Platform = "Xbox",
                GamesCount = 5,
                AvgHoursPlayed = 15,
                AvgRating = 4.0
            }
        };
        
        return Task.FromResult(Result.Ok<IReadOnlyList<PlatformStatistics>>(result));
    }
}