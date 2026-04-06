using CodexGenAI.CoffeeInventory.Contracts;
using CodexGenAI.CoffeeInventory.Data;
using CodexGenAI.CoffeeInventory.Models;
using Microsoft.EntityFrameworkCore;

var builder = WebApplication.CreateBuilder(args);

var connectionString = builder.Configuration.GetConnectionString("Default")
    ?? "Server=localhost;Port=3306;Database=coffee_db;User=root;Password=password;";

builder.Services.AddDbContext<CoffeeDbContext>(options =>
    options.UseMySql(connectionString, ServerVersion.AutoDetect(connectionString)));

builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

builder.Services.AddCors(options =>
{
    var allowedOrigins = builder.Configuration.GetSection("AllowedOrigins").Get<string[]>()
        ?? new[] { "http://localhost:3000" };

    options.AddDefaultPolicy(policy =>
        policy.WithOrigins(allowedOrigins)
              .AllowAnyHeader()
              .AllowAnyMethod());
});

var app = builder.Build();

await using (var scope = app.Services.CreateAsyncScope())
{
    var db = scope.ServiceProvider.GetRequiredService<CoffeeDbContext>();
    await db.Database.EnsureCreatedAsync();
    await CoffeeDbSeeder.SeedAsync(db);
}

if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseCors();

app.MapGet("/api/coffee", async (string? search, string? stock, CoffeeDbContext db) =>
{
    var query = db.CoffeeItems.AsNoTracking().AsQueryable();

    if (!string.IsNullOrWhiteSpace(search))
    {
        var pattern = $"%{search.Trim()}%";
        query = query.Where(item =>
            EF.Functions.Like(item.Name, pattern) ||
            EF.Functions.Like(item.Sku, pattern) ||
            (item.Origin != null && EF.Functions.Like(item.Origin, pattern)) ||
            EF.Functions.Like(item.Category, pattern) ||
            (item.Supplier != null && EF.Functions.Like(item.Supplier, pattern)));
    }

    if (string.Equals(stock, "low", StringComparison.OrdinalIgnoreCase))
    {
        query = query.Where(item => item.Quantity <= item.MinimumStock);
    }

    var items = await query
        .OrderBy(item => item.Quantity <= item.MinimumStock ? 0 : 1)
        .ThenBy(item => item.Name)
        .ToListAsync();

    return Results.Ok(items);
});

app.MapGet("/api/coffee/{id:int}", async (int id, CoffeeDbContext db) =>
    await db.CoffeeItems.FindAsync(id) is CoffeeItem coffee
        ? Results.Ok(coffee)
        : Results.NotFound());

app.MapGet("/api/dashboard", async (CoffeeDbContext db) =>
{
    var query = db.CoffeeItems.AsNoTracking();
    var totalItems = await query.CountAsync();
    var totalUnits = totalItems == 0 ? 0 : await query.SumAsync(item => item.Quantity);
    var totalValue = totalItems == 0 ? 0m : await query.SumAsync(item => item.UnitPrice * item.Quantity);
    var lowStockItems = totalItems == 0 ? 0 : await query.CountAsync(item => item.Quantity <= item.MinimumStock);
    var originsTracked = totalItems == 0
        ? 0
        : await query
            .Where(item => item.Origin != null && item.Origin != "")
            .Select(item => item.Origin)
            .Distinct()
            .CountAsync();

    return Results.Ok(new DashboardSummaryResponse
    {
        TotalItems = totalItems,
        TotalUnits = totalUnits,
        TotalValue = totalValue,
        LowStockItems = lowStockItems,
        OriginsTracked = originsTracked
    });
});

app.MapPost("/api/coffee", async (CoffeeItemRequest request, CoffeeDbContext db) =>
{
    var errors = request.Validate();
    if (errors.Count > 0)
    {
        return Results.ValidationProblem(errors);
    }

    var normalizedSku = request.Sku.Trim().ToUpperInvariant();
    var skuExists = await db.CoffeeItems.AnyAsync(item => item.Sku == normalizedSku);
    if (skuExists)
    {
        return Results.Conflict(new { message = "SKU already exists." });
    }

    var coffee = new CoffeeItem();
    request.ApplyTo(coffee);
    db.CoffeeItems.Add(coffee);
    await db.SaveChangesAsync();

    return Results.Created($"/api/coffee/{coffee.Id}", coffee);
});

app.MapPut("/api/coffee/{id:int}", async (int id, CoffeeItemRequest request, CoffeeDbContext db) =>
{
    var errors = request.Validate();
    if (errors.Count > 0)
    {
        return Results.ValidationProblem(errors);
    }

    var coffee = await db.CoffeeItems.FindAsync(id);
    if (coffee is null)
    {
        return Results.NotFound();
    }

    var normalizedSku = request.Sku.Trim().ToUpperInvariant();
    var skuExists = await db.CoffeeItems.AnyAsync(item => item.Sku == normalizedSku && item.Id != id);
    if (skuExists)
    {
        return Results.Conflict(new { message = "SKU already exists." });
    }

    request.ApplyTo(coffee);

    await db.SaveChangesAsync();
    return Results.NoContent();
});

app.MapDelete("/api/coffee/{id:int}", async (int id, CoffeeDbContext db) =>
{
    var coffee = await db.CoffeeItems.FindAsync(id);
    if (coffee is null) return Results.NotFound();

    db.CoffeeItems.Remove(coffee);
    await db.SaveChangesAsync();
    return Results.NoContent();
});

app.MapGet("/health", () => Results.Ok("ok"));

await app.RunAsync();
