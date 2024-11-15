using GameLogger.Core.Repositories;
using Microsoft.Extensions.DependencyInjection;

namespace GameLogger.Core;

public static class Setup
{
    public static IServiceCollection AddCore(this IServiceCollection services)
    {
        services.AddScoped<UseCases.GetGamesUseCase>();
        return services;
    }
}