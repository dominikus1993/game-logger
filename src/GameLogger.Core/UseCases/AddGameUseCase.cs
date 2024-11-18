using GameLogger.Core.Data;
using GameLogger.Core.Repositories;
using GameLogger.Core.Types;
using GameLogger.Core.Validators.Abstractions;

namespace GameLogger.Core.UseCases;

public sealed record AddGameRequest(string Title, DateOnly StartDate, DateOnly? FinishDate, string? Platform, ushort? HoursPlayed, ushort? Rating, string? Notes);

public sealed class InvalidaAddGameRequestException : Exception
{
    private const string DefaultMessage = "Invalid AddGameRequest";
    
    public string? AdditonalMessage { get; init; }
    
    public InvalidaAddGameRequestException(string? message) : base(DefaultMessage)
    {
        AdditonalMessage = message;
    }

    public override string ToString()
    {
        return $"{base.ToString()}, {nameof(AdditonalMessage)}: {AdditonalMessage}";
    }
}

public sealed class AddGameUseCase
{
    private readonly IGamesLogsRepository _repository;
    private readonly IValidator<AddGameRequest> _validator;
    
    public AddGameUseCase(IGamesLogsRepository repository, IValidator<AddGameRequest> validator)
    {
        _repository = repository;
        _validator = validator;
    }

    public async Task<Result<Game>> Execute(AddGameRequest game, CancellationToken cancellationToken = default)
    {
        
        var validationResult = await _validator.Validate(game);
        
        if (!validationResult.IsValid)
        {
            return Result.Failure<Game>(new InvalidaAddGameRequestException(validationResult.Message));
        }
        
        var newGame = new Game
        {
            Id = Guid.CreateVersion7(),
            Title = game.Title,
            StartDate = game.StartDate,
            FinishDate = game.FinishDate,
            Platform = game.Platform,
            HoursPlayed = game.HoursPlayed,
            Rating = game.Rating,
            Notes = game.Notes
        };

        var addGameResult = await _repository.WriteGame(newGame, cancellationToken);


        if (addGameResult.IsSuccess)
        {
            return Result.Ok(newGame);
        }
        
        return Result.Failure<Game>(addGameResult.ErrorValue);
    }
}