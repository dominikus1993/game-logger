using Microsoft.Extensions.DependencyInjection;

namespace GameLogger.Core;

public static class Setup
{
    public static IServiceCollection AddCore(this IServiceCollection services)
    {
        services.AddSingleton<UseCases.GetGamesUseCase>();
        
        return services;
    }
}