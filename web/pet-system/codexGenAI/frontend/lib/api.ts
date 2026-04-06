import type { CoffeeItem, CoffeeItemPayload, DashboardSummary, StockFilter } from "./types";

const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL ?? "http://localhost:5216";

async function request<T>(path: string, init?: RequestInit): Promise<T> {
  const response = await fetch(`${API_BASE_URL}${path}`, {
    ...init,
    headers: {
      "Content-Type": "application/json",
      ...(init?.headers ?? {})
    },
    cache: "no-store"
  });

  if (!response.ok) {
    let message = "请求失败，请检查后端服务是否可用。";

    try {
      const payload = await response.json();

      if (payload?.message) {
        message = payload.message;
      } else if (payload?.errors) {
        const firstMessage = Object.values(payload.errors).flat()[0];
        if (typeof firstMessage === "string") {
          message = firstMessage;
        }
      }
    } catch {
      message = response.statusText || message;
    }

    throw new Error(message);
  }

  if (response.status === 204) {
    return undefined as T;
  }

  return (await response.json()) as T;
}

export function getCoffeeItems(filter: StockFilter) {
  const searchParams = new URLSearchParams();

  if (filter === "low") {
    searchParams.set("stock", "low");
  }

  const suffix = searchParams.toString();
  return request<CoffeeItem[]>(`/api/coffee${suffix ? `?${suffix}` : ""}`);
}

export function getDashboardSummary() {
  return request<DashboardSummary>("/api/dashboard");
}

export function createCoffeeItem(payload: CoffeeItemPayload) {
  return request<CoffeeItem>("/api/coffee", {
    method: "POST",
    body: JSON.stringify(payload)
  });
}

export function updateCoffeeItem(id: number, payload: CoffeeItemPayload) {
  return request<void>(`/api/coffee/${id}`, {
    method: "PUT",
    body: JSON.stringify(payload)
  });
}

export function deleteCoffeeItem(id: number) {
  return request<void>(`/api/coffee/${id}`, {
    method: "DELETE"
  });
}
