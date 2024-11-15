using GameLogger.Core.Repositories;
using GameLogger.Core.UseCases;
using Microsoft.AspNetCore.Builder;
using Microsoft.Extensions.DependencyInjection;

namespace GameLogger.Core;

public static class Setup
{
    public static IServiceCollection AddCore(this IServiceCollection services)
    {
        services.AddScoped<UseCases.GetGamesUseCase>();
        services.AddScoped<InitDatabaseUseCase>();
        return services;
    }

    public static async Task Init(this WebApplication app)
    {
        await using var scope = app.Services.CreateAsyncScope();
        var useCase = scope.ServiceProvider.GetRequiredService<InitDatabaseUseCase>();
        await useCase.Execute();
    }
    
}