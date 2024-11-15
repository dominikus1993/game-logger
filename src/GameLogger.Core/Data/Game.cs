namespace GameLogger.Core.Data;

public sealed class Game
{
    public Guid Id { get; set; }
    public string Title { get; set; }
    public DateOnly StartDate { get; set; }
    public DateOnly? FinishDate { get; set; }
    public string? Platform { get; set; }
    public ushort? HoursPlayed { get; set; }
    public ushort? Rating { get; set; }
    public string? Notes { get; set; }
}