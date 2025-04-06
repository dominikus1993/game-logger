using GameLogger.Core;
using GameLogger.Infrastructure;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddAntiforgery();
builder.Services.AddCore();
builder.Services.AddInfrastructure(builder.Configuration);
var app = builder.Build();
// await app.Init();
// Configure the HTTP request pipeline.
if (!app.Environment.IsDevelopment())
{
    app.UseExceptionHandler("/Error", createScopeForErrors: true);
    // The default HSTS value is 30 days. You may want to change this for production scenarios, see https://aka.ms/aspnetcore-hsts.
}

app.UseAntiforgery();


app.MapGet("/", () => "Hello World!");
await app.RunAsync();
