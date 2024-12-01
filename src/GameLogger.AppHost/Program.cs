using Projects;

var builder = DistributedApplication.CreateBuilder(args);

var database = builder.AddMongoDB("database");

builder.AddProject<GameLogger_Web>("web")
    .WithHttpEndpoint()
    .WithReference(database)
    .WaitFor(database);

builder.Build().Run();