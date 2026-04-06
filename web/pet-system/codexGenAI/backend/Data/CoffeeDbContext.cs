using CodexGenAI.CoffeeInventory.Models;
using Microsoft.EntityFrameworkCore;

namespace CodexGenAI.CoffeeInventory.Data;

public class CoffeeDbContext : DbContext
{
    public CoffeeDbContext(DbContextOptions<CoffeeDbContext> options) : base(options)
    {
    }

    public DbSet<CoffeeItem> CoffeeItems => Set<CoffeeItem>();

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.Entity<CoffeeItem>(entity =>
        {
            entity.ToTable("coffee_items");
            entity.HasIndex(e => e.Sku).IsUnique();
            entity.Property(e => e.Name).HasMaxLength(120).IsRequired();
            entity.Property(e => e.Sku).HasMaxLength(40).IsRequired();
            entity.Property(e => e.Origin).HasMaxLength(80);
            entity.Property(e => e.RoastLevel).HasMaxLength(50).IsRequired();
            entity.Property(e => e.Category).HasMaxLength(60).IsRequired();
            entity.Property(e => e.Supplier).HasMaxLength(100);
            entity.Property(e => e.UnitPrice).HasPrecision(10, 2);
            entity.Property(e => e.LastRestocked).HasColumnType("datetime(6)");
            entity.Property(e => e.Notes).HasMaxLength(500);
        });
    }
}
