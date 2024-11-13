using GameLogger.Core.Repositories;
using Microsoft.Extensions.DependencyInjection;

namespace GameLogger.Core;

public static class Setup
{
    public static IServiceCollection AddCore(this IServiceCollection services)
    {
        services.AddScoped<UseCases.GetGamesUseCase>();
        services.AddScoped<IGamesLogsRepository, FakeGamesLogsRepository>();
        
        return services;
    }
}