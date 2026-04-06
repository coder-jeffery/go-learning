using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Design;

namespace CodexGenAI.CoffeeInventory.Data;

public class DesignTimeDbContextFactory : IDesignTimeDbContextFactory<CoffeeDbContext>
{
    public CoffeeDbContext CreateDbContext(string[] args)
    {
        var optionsBuilder = new DbContextOptionsBuilder<CoffeeDbContext>();
        var connectionString = Environment.GetEnvironmentVariable("COFFEE_DB_CONNECTION")
            ?? "Server=localhost;Port=3306;Database=coffee_db;User=root;Password=password;";

        optionsBuilder.UseMySql(connectionString, ServerVersion.AutoDetect(connectionString));
        return new CoffeeDbContext(optionsBuilder.Options);
    }
}
