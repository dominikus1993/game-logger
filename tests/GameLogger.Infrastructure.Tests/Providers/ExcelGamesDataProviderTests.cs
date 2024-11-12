using GameLogger.Infrastructure.Providers;

namespace GameLogger.Infrastructure.Tests.Providers;

public sealed class ExcelGamesDataProviderTests
{
    [Fact]
    public async Task TestWhenFileIsNotEmpty()
    {
        // Arrange
        var configuration = new ExcelConfiguration
        {
            Path = "games.xlsx"
        };
        var provider = new ExcelGamesDataProvider(configuration);
        
        // Act
        var games = await provider.Provide().ToListAsync();
        
        // Assert
        Assert.NotEmpty(games);
    }
}