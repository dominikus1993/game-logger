using GameLogger.Core.UseCases;
using GameLogger.Core.Validators.Abstractions;

namespace GameLogger.Core.Validators;

public sealed class AddGameRequestValidator : IValidator<AddGameRequest>
{
    public ValueTask<ValidationResult> Validate(AddGameRequest entity)
    {
        if (string.IsNullOrWhiteSpace(entity.Title))
        {
            return ValueTask.FromResult(new ValidationResult
            {
                IsValid = false,
                Message = "Title is required"
            });
        }

        if (entity.StartDate == default)
        {
            return ValueTask.FromResult(new ValidationResult
            {
                IsValid = false,
                Message = "Start date is required"
            });
        }

        return ValueTask.FromResult(new ValidationResult
        {
            IsValid = true
        });
    }
}