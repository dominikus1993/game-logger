using GameLogger.Core.Providers;
using GameLogger.Core.Repositories;
using GameLogger.Infrastructure.Extensions;
using GameLogger.Infrastructure.Providers;
using GameLogger.Infrastructure.Repositories;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using MongoDB.Driver;

namespace GameLogger.Infrastructure;

public static class Setup
{
    public static IServiceCollection AddInfrastructure(this IServiceCollection services, IConfiguration configuration)
    {
        services.AddSingleton(configuration.GetRequiredSection("Excel").Get<ExcelConfiguration>()!);
        services.AddScoped<IGamesDataProvider, ExcelGamesDataProvider>();
        services.AddMongoDb(configuration);
        services.AddScoped<IGamesLogsRepository, FakeGamesLogsRepository>();
        services.AddScoped<IGamesStatisticsProvider, FakeGamesStatisticsProvider>();
        return services;
    }
    
    private static IServiceCollection AddMongoDb(this IServiceCollection services, IConfiguration configuration)
    {
        services.AddSingleton(MongoClientSettings.FromConnectionString(configuration.GetConnectionString("GamesDb")));
        services.AddSingleton<IMongoClient>(sp =>
        {
            var settings = sp.GetRequiredService<MongoClientSettings>();
            return new MongoClient(settings);
        });
        services.AddSingleton<IMongoDatabase>(sp =>
        {
            var client = sp.GetRequiredService<IMongoClient>();
            var db = client.GetDatabase("GamesLogger");
            db.MapGamesCollection();
            return db;
        });
    
        return services;
    }
}