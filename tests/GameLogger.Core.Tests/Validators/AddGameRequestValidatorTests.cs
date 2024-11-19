using GameLogger.Core.UseCases;
using GameLogger.Core.Validators;

namespace GameLogger.Core.Tests.Validators;

public class AddGameRequestValidatorTests
{
    [Fact]
    public async Task Validate_WhenIsNull_ReturnsInvalidResult()
    {
        // Arrange
        var validator = new AddGameRequestValidator();

        // Act
        await Assert.ThrowsAsync<ArgumentNullException>(async () => await validator.Validate(null!));
    }
    
    [Fact]
    public async Task Validate_WhenTitleIsEmpty_ReturnsInvalidResult()
    {
        // Arrange
        var validator = new AddGameRequestValidator();
        var request = new AddGameRequest(string.Empty, DateOnly.FromDateTime(DateTime.Now), null, null, null, null, null);

        // Act
        var result = await validator.Validate(request);

        // Assert
        Assert.False(result.IsValid);
        Assert.Equal("Title is required", result.Message);
    }
    
    [Fact]
    public async Task Validate_WhenStartDateIsDefault_ReturnsInvalidResult()
    {
        // Arrange
        var validator = new AddGameRequestValidator();
        var request = new AddGameRequest("Title", default, null, null, null, null, null);

        // Act
        var result = await validator.Validate(request);

        // Assert
        Assert.False(result.IsValid);
        Assert.Equal("Start date is required", result.Message);
    }
    
    [Fact]
    public async Task Validate_WhenRequestIsValid_ReturnsValidResult()
    {
        // Arrange
        var validator = new AddGameRequestValidator();
        var request = new AddGameRequest("Title", DateOnly.FromDateTime(DateTime.Now), null, null, null, null, null);

        // Act
        var result = await validator.Validate(request);

        // Assert
        Assert.True(result.IsValid);
        Assert.Null(result.Message);
    }
}