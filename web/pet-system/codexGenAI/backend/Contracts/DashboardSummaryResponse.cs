namespace CodexGenAI.CoffeeInventory.Contracts;

public class DashboardSummaryResponse
{
    public int TotalItems { get; set; }
    public int TotalUnits { get; set; }
    public decimal TotalValue { get; set; }
    public int LowStockItems { get; set; }
    public int OriginsTracked { get; set; }
}
