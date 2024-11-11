using GameLogger.Core.Data;
using GameLogger.Core.Repositories;
using Microsoft.Extensions.Logging;

namespace GameLogger.Core.UseCases;

internal static partial class GetGamesUseCaseLogger
{
    [LoggerMessage(0, LogLevel.Information, "Getting games with query")]
    public static partial void LogGettingGames(this ILogger logger, [LogProperties(SkipNullProperties = true)]GetGamesQuery query);
    
    [LoggerMessage(1, LogLevel.Warning, "No games found with query")]
    public static partial void LogNoGamesFound(this ILogger logger, [LogProperties(SkipNullProperties = true)]GetGamesQuery query);
}

public sealed class GetGamesUseCase
{
    private readonly IGamesLogsRepository _repository;
    private readonly ILogger<GetGamesUseCase> _logger;
    
    public GetGamesUseCase(IGamesLogsRepository repository, ILogger<GetGamesUseCase> logger)
    {
        _repository = repository;
        _logger = logger;
    }

    public async Task<IReadOnlyList<Game>> Execute(GetGamesQuery query, CancellationToken cancellationToken = default)
    {
        _logger.LogGettingGames(query);
        var result =  await _repository.GetGames(query, cancellationToken);

        if (result is null or {Count: 0})
        {
            _logger.LogNoGamesFound(query);
            return [];
        }
        
        return result;
    }
}