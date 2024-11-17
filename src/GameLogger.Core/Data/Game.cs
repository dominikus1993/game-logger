namespace GameLogger.Core.Data;

public sealed class Game
{
    public Guid Id { get; init; }
    public string Title { get; init; }
    public DateOnly StartDate { get; init; }
    public DateOnly? FinishDate { get; init; }
    public string? Platform { get; init; }
    public ushort? HoursPlayed { get; init; }
    public ushort? Rating { get; init; }
    public string? Notes { get; init; }

    private bool Equals(Game other)
    {
        return Id.Equals(other.Id);
    }

    public override bool Equals(object? obj)
    {
        return ReferenceEquals(this, obj) || obj is Game other && Equals(other);
    }

    public override int GetHashCode()
    {
        return Id.GetHashCode();
    }
}