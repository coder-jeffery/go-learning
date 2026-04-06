using CodexGenAI.CoffeeInventory.Models;

namespace CodexGenAI.CoffeeInventory.Contracts;

public class CoffeeItemRequest
{
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

    public Dictionary<string, string[]> Validate()
    {
        var errors = new Dictionary<string, List<string>>();

        void Add(string key, string message)
        {
            if (!errors.TryGetValue(key, out var messages))
            {
                messages = new List<string>();
                errors[key] = messages;
            }

            messages.Add(message);
        }

        if (string.IsNullOrWhiteSpace(Name))
        {
            Add("name", "Name is required.");
        }

        if (string.IsNullOrWhiteSpace(Sku))
        {
            Add("sku", "SKU is required.");
        }

        if (string.IsNullOrWhiteSpace(RoastLevel))
        {
            Add("roastLevel", "Roast level is required.");
        }

        if (string.IsNullOrWhiteSpace(Category))
        {
            Add("category", "Category is required.");
        }

        if (Quantity < 0)
        {
            Add("quantity", "Quantity cannot be negative.");
        }

        if (MinimumStock < 0)
        {
            Add("minimumStock", "Minimum stock cannot be negative.");
        }

        if (UnitPrice < 0)
        {
            Add("unitPrice", "Unit price cannot be negative.");
        }

        return errors.ToDictionary(pair => pair.Key, pair => pair.Value.ToArray());
    }

    public void ApplyTo(CoffeeItem coffee)
    {
        coffee.Name = Name.Trim();
        coffee.Sku = Sku.Trim().ToUpperInvariant();
        coffee.Origin = string.IsNullOrWhiteSpace(Origin) ? null : Origin.Trim();
        coffee.RoastLevel = RoastLevel.Trim();
        coffee.Category = Category.Trim();
        coffee.Supplier = string.IsNullOrWhiteSpace(Supplier) ? null : Supplier.Trim();
        coffee.Quantity = Quantity;
        coffee.MinimumStock = MinimumStock;
        coffee.UnitPrice = UnitPrice;
        coffee.LastRestocked = LastRestocked == default ? DateTime.UtcNow : LastRestocked;
        coffee.Notes = string.IsNullOrWhiteSpace(Notes) ? null : Notes.Trim();
    }
}
