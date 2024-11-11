using GameLogger.Core.Repositories;
using Microsoft.Extensions.DependencyInjection;

namespace GameLogger.Core;

public static class Setup
{
    public static IServiceCollection AddCore(this IServiceCollection services)
    {
        services.AddSingleton<UseCases.GetGamesUseCase>();
        services.AddSingleton<IGamesLogsRepository, FakeGamesLogsRepository>();
        
        return services;
    }
}