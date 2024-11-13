using GameLogger.Core.Providers;
using GameLogger.Infrastructure.Providers;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;

namespace GameLogger.Infrastructure;

public static class Setup
{
    public static IServiceCollection AddInfrastructure(this IServiceCollection services, IConfiguration configuration)
    {
        services.AddSingleton(configuration.GetRequiredSection("Excel").Get<ExcelConfiguration>()!);
        services.AddScoped<IGamesDataProvider, ExcelGamesDataProvider>();

        return services;
    }
}