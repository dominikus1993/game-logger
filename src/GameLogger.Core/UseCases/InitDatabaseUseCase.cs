using GameLogger.Core.Providers;
using GameLogger.Core.Repositories;

namespace GameLogger.Core.UseCases;

public sealed class InitDatabaseUseCase
{
    private readonly IGamesDataProvider _gamesDataProvider;
    private readonly IGamesLogsRepository _gamesLogsRepository;

    public InitDatabaseUseCase(IGamesDataProvider gamesDataProvider, IGamesLogsRepository gamesLogsRepository)
    {
        _gamesDataProvider = gamesDataProvider;
        _gamesLogsRepository = gamesLogsRepository;
    }
    
    public async Task Execute(CancellationToken cancellationToken = default)
    {
        var gamesFromDb = await _gamesLogsRepository.GetGames(new GetGamesQuery(1, 1), cancellationToken);
        if (gamesFromDb is { Count: > 0 })
        {
            return;
        }
        
        var games = await _gamesDataProvider.Provide(cancellationToken).ToArrayAsync(cancellationToken: cancellationToken);
        foreach (var game in games)
        {
            var result = await _gamesLogsRepository.WriteGame(game, cancellationToken);
            if (!result.IsSuccess)
            {
                throw new InvalidOperationException("Failed to write game to database", result.ErrorValue);
            }
        }
    }
}