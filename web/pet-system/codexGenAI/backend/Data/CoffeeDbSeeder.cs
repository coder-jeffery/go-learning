using CodexGenAI.CoffeeInventory.Models;
using Microsoft.EntityFrameworkCore;

namespace CodexGenAI.CoffeeInventory.Data;

public static class CoffeeDbSeeder
{
    public static async Task SeedAsync(CoffeeDbContext db)
    {
        if (await db.CoffeeItems.AnyAsync())
        {
            return;
        }

        var now = DateTime.UtcNow;

        var items = new[]
        {
            new CoffeeItem
            {
                Name = "Ethiopia Yirgacheffe",
                Sku = "ETH-YIR-250",
                Origin = "Ethiopia",
                RoastLevel = "Light",
                Category = "Single Origin",
                Supplier = "Blue Harbor Trading",
                Quantity = 18,
                MinimumStock = 12,
                UnitPrice = 58.00m,
                LastRestocked = now.AddDays(-4),
                Notes = "Floral aroma with bergamot and stone fruit."
            },
            new CoffeeItem
            {
                Name = "Colombia Supremo",
                Sku = "COL-SUP-500",
                Origin = "Colombia",
                RoastLevel = "Medium",
                Category = "Cafe House",
                Supplier = "Andes Direct",
                Quantity = 26,
                MinimumStock = 15,
                UnitPrice = 64.50m,
                LastRestocked = now.AddDays(-2),
                Notes = "Balanced caramel profile for all-day service."
            },
            new CoffeeItem
            {
                Name = "Guatemala Antigua",
                Sku = "GUA-ANT-250",
                Origin = "Guatemala",
                RoastLevel = "Medium Dark",
                Category = "Seasonal Reserve",
                Supplier = "Volcano Select",
                Quantity = 9,
                MinimumStock = 10,
                UnitPrice = 72.00m,
                LastRestocked = now.AddDays(-7),
                Notes = "Chocolate body with gentle spice."
            },
            new CoffeeItem
            {
                Name = "Sumatra Mandheling",
                Sku = "SUM-MAN-500",
                Origin = "Indonesia",
                RoastLevel = "Dark",
                Category = "Espresso Base",
                Supplier = "Island Roast Co.",
                Quantity = 14,
                MinimumStock = 8,
                UnitPrice = 69.00m,
                LastRestocked = now.AddDays(-3),
                Notes = "Low acidity and dense body for blends."
            },
            new CoffeeItem
            {
                Name = "Brazil Cerrado",
                Sku = "BRA-CER-1000",
                Origin = "Brazil",
                RoastLevel = "Medium",
                Category = "Cold Brew",
                Supplier = "Terra Verde",
                Quantity = 6,
                MinimumStock = 12,
                UnitPrice = 88.00m,
                LastRestocked = now.AddDays(-10),
                Notes = "Nutty sweetness suited for cold extraction."
            }
        };

        await db.CoffeeItems.AddRangeAsync(items);
        await db.SaveChangesAsync();
    }
}
