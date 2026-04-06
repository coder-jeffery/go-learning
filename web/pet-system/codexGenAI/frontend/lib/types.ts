export type StockFilter = "all" | "low";

export interface CoffeeItem {
  id: number;
  name: string;
  sku: string;
  origin: string | null;
  roastLevel: string;
  category: string;
  supplier: string | null;
  quantity: number;
  minimumStock: number;
  unitPrice: number;
  lastRestocked: string;
  notes: string | null;
}

export interface CoffeeItemPayload {
  name: string;
  sku: string;
  origin: string;
  roastLevel: string;
  category: string;
  supplier: string;
  quantity: number;
  minimumStock: number;
  unitPrice: number;
  lastRestocked: string;
  notes: string;
}

export interface DashboardSummary {
  totalItems: number;
  totalUnits: number;
  totalValue: number;
  lowStockItems: number;
  originsTracked: number;
}
