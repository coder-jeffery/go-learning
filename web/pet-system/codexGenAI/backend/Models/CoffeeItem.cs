namespace CodexGenAI.CoffeeInventory.Models;

public class CoffeeItem
{
    public int Id { get; set; }
    public string Name { get; set; } = string.Empty;
    public string Sku { get; set; } = string.Empty;
    public string? Origin { get; set; }
    public string RoastLevel { get; set; } = string.Empty;
    public string Category { get; set; } = string.Empty;
    public string? Supplier { get; set; }
    public int Quantity { get; set; }
    public int MinimumStock { get; set; }
    public decimal UnitPrice { get; set; }
    public DateTime LastRestocked { get; set; }
    public string? Notes { get; set; }
}
