using System.Runtime.CompilerServices;
using ClosedXML.Excel;
using GameLogger.Core.Data;
using GameLogger.Core.Providers;

namespace GameLogger.Infrastructure.Providers;

public sealed class ExcelConfiguration
{
    public string Path { get; set; }
}

internal sealed class ExcelGamesDataProvider : IGamesDataProvider
{
    private readonly ExcelConfiguration _configuration;
    
    public ExcelGamesDataProvider(ExcelConfiguration configuration)
    {
        _configuration = configuration;
    }
    
    public async IAsyncEnumerable<Game> Provide([EnumeratorCancellation] CancellationToken cancellationToken = default)
    {
        await using var file = File.OpenRead(_configuration.Path);
        using var excel = new XLWorkbook(file);
        
        var sheet = excel.Worksheet(1);
        var rows = sheet.RowsUsed();
        foreach (var row in rows)
        {
            if (row.RowNumber() < 3)
            {
                continue;
            }

            var title = row.Cell(1);
            var rating = row.Cell(2);
            var platform = row.Cell(3);
            var startDate = row.Cell(4);
            var endDate = row.Cell(5);
            var playTime = row.Cell(6);
            
            yield return new Game
            {
                Title = title.GetString(),
                Rating = rating.GetValue<ushort>(),
                Platform = platform.GetString(),
                StartDate = startDate.GetDateTime(),
                FinishDate = endDate.GetDateTime(),
                HoursPlayed = playTime.GetValue<ushort>()
            };
        }
        
    }
}