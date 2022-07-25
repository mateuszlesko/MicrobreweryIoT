using Microsoft.Extensions.Configuration;
using Microsoft.EntityFrameworkCore;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.

builder.Services.AddControllers();
builder.Services.AddDbContext<MicroBreweryRecipesAPI.Models.MicroBreweryContext>(options=>
options.UseNpgsql(builder.Configuration.GetSection("ConnectionStrings").Get<string>())
.UseSnakeCaseNamingConvention()
.UseLoggerFactory(LoggerFactory.Create(builder => builder.AddConsole()))
.EnableSensitiveDataLogging()
);
builder.Services.AddDatabaseDeveloperPageExceptionFilter();
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI(c => {
        c.SwaggerEndpoint("/swagger/v1/swagger.json","Beer Recipes API");
        c.RoutePrefix="";
    });
}

app.UseHttpsRedirection();

app.UseAuthorization();

app.MapControllers();

app.Run();
