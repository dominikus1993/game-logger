namespace GameLogger.Core.Validators.Abstractions;

public sealed class ValidationResult
{
    public bool IsValid { get; init; }
    public string? Message { get; init; }
}

public interface IValidator<in T>
{
    ValueTask<ValidationResult> Validate(T entity);
}