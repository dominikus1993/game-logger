using GameLogger.Core;
using GameLogger.Core.Repositories;
using GameLogger.Core.UseCases;
using GameLogger.Infrastructure;
using Microsoft.AspNetCore.Mvc;
using Scalar.AspNetCore;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddAntiforgery();
builder.Services.AddOpenApi();
builder.Services.AddCore();
builder.Services.AddInfrastructure(builder.Configuration);
var app = builder.Build();
// await app.Init();
// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.MapOpenApi();
    app.MapScalarApiReference();
}

app.UseAntiforgery();


app.MapGet("/games", async ([FromServices] GetGamesUseCase useCase, CancellationToken cancellationToken, [FromQuery]uint page = 1, [FromQuery]ushort pageSize = 12) =>
{
    if (page < 1)
    {
        return Results.BadRequest("Page must be greater than 0");
    }
    
    if (pageSize < 1)
    {
        return Results.BadRequest("Page size must be greater than 0");
    }
    if (pageSize > 100)
    {
        return Results.BadRequest("Page size must be less than or equal to 100");
    }
    
    var result = await useCase.Execute(new GetGamesQuery(page, pageSize), cancellationToken);
    return Results.Ok(result);
});

await app.RunAsync();

