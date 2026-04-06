"use client";

import { useDeferredValue, useEffect, useState, useTransition } from "react";
import {
  createCoffeeItem,
  deleteCoffeeItem,
  getCoffeeItems,
  getDashboardSummary,
  updateCoffeeItem
} from "../lib/api";
import type {
  CoffeeItem,
  CoffeeItemPayload,
  DashboardSummary,
  StockFilter
} from "../lib/types";

function getToday() {
  const now = new Date();
  const localDate = new Date(now.getTime() - now.getTimezoneOffset() * 60_000);
  return localDate.toISOString().slice(0, 10);
}

const EMPTY_FORM: CoffeeItemPayload = {
  name: "",
  sku: "",
  origin: "",
  roastLevel: "Medium",
  category: "Cafe House",
  supplier: "",
  quantity: 0,
  minimumStock: 8,
  unitPrice: 0,
  lastRestocked: getToday(),
  notes: ""
};

const moneyFormatter = new Intl.NumberFormat("zh-CN", {
  style: "currency",
  currency: "CNY",
  minimumFractionDigits: 2
});

const numberFormatter = new Intl.NumberFormat("zh-CN");

function getErrorMessage(error: unknown) {
  if (error instanceof Error) {
    return error.message;
  }

  return "请求失败，请稍后再试。";
}

function toFormState(item: CoffeeItem): CoffeeItemPayload {
  return {
    name: item.name,
    sku: item.sku,
    origin: item.origin ?? "",
    roastLevel: item.roastLevel,
    category: item.category,
    supplier: item.supplier ?? "",
    quantity: item.quantity,
    minimumStock: item.minimumStock,
    unitPrice: item.unitPrice,
    lastRestocked: item.lastRestocked.slice(0, 10),
    notes: item.notes ?? ""
  };
}

function getStatus(item: CoffeeItem) {
  if (item.quantity === 0) {
    return { label: "缺货", tone: "empty" };
  }

  if (item.quantity <= item.minimumStock) {
    return { label: "低库存", tone: "low" };
  }

  return { label: "充足", tone: "good" };
}

export function InventoryDashboard() {
  const [items, setItems] = useState<CoffeeItem[]>([]);
  const [summary, setSummary] = useState<DashboardSummary | null>(null);
  const [searchTerm, setSearchTerm] = useState("");
  const [stockFilter, setStockFilter] = useState<StockFilter>("all");
  const [form, setForm] = useState<CoffeeItemPayload>(EMPTY_FORM);
  const [editingId, setEditingId] = useState<number | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [isPending, startTransition] = useTransition();
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [successMessage, setSuccessMessage] = useState<string | null>(null);

  const deferredSearchTerm = useDeferredValue(searchTerm);

  async function hydrate(nextFilter: StockFilter) {
    setIsLoading(true);
    setErrorMessage(null);

    try {
      const [inventory, dashboard] = await Promise.all([
        getCoffeeItems(nextFilter),
        getDashboardSummary()
      ]);

      setItems(inventory);
      setSummary(dashboard);
    } catch (error) {
      setErrorMessage(getErrorMessage(error));
    } finally {
      setIsLoading(false);
    }
  }

  useEffect(() => {
    void hydrate(stockFilter);
  }, [stockFilter]);

  const visibleItems = items.filter((item) => {
    const query = deferredSearchTerm.trim().toLowerCase();
    if (!query) {
      return true;
    }

    return [
      item.name,
      item.sku,
      item.origin ?? "",
      item.category,
      item.supplier ?? ""
    ].some((value) => value.toLowerCase().includes(query));
  });

  const lowStockSpotlight = items.filter((item) => item.quantity <= item.minimumStock).slice(0, 3);

  function clearForm() {
    setEditingId(null);
    setForm({ ...EMPTY_FORM, lastRestocked: getToday() });
  }

  function resetForm() {
    clearForm();
    setSuccessMessage(null);
    setErrorMessage(null);
  }

  function handleFieldChange(
    event: React.ChangeEvent<HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement>
  ) {
    const { name, value } = event.target;
    const numericFields = new Set(["quantity", "minimumStock", "unitPrice"]);

    setForm((current) => ({
      ...current,
      [name]: numericFields.has(name) ? Number(value) : value
    }));
  }

  function handleEdit(item: CoffeeItem) {
    setEditingId(item.id);
    setForm(toFormState(item));
    setSuccessMessage(null);
    setErrorMessage(null);
  }

  function handleDelete(id: number) {
    const confirmed = window.confirm("确认删除这条咖啡库存记录吗？");
    if (!confirmed) {
      return;
    }

    setSuccessMessage(null);
    setErrorMessage(null);

    startTransition(() => {
      void (async () => {
        try {
          await deleteCoffeeItem(id);
          if (editingId === id) {
            clearForm();
          }
          await hydrate(stockFilter);
          setSuccessMessage("库存记录已删除。");
        } catch (error) {
          setErrorMessage(getErrorMessage(error));
        }
      })();
    });
  }

  function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setSuccessMessage(null);
    setErrorMessage(null);

    startTransition(() => {
      void (async () => {
        try {
          const nextFilter = stockFilter === "low" ? "all" : stockFilter;

          if (editingId === null) {
            await createCoffeeItem(form);
            setSuccessMessage("已新增咖啡库存。");
          } else {
            await updateCoffeeItem(editingId, form);
            setSuccessMessage("库存资料已更新。");
          }

          clearForm();
          if (stockFilter !== nextFilter) {
            setStockFilter(nextFilter);
          }
          await hydrate(nextFilter);
        } catch (error) {
          setErrorMessage(getErrorMessage(error));
        }
      })();
    });
  }

  return (
    <main className="inventory-shell">
      <section className="hero-card">
        <div className="hero-copy">
          <div className="eyebrow">Coffee Inventory System</div>
          <h1 className="hero-title">
            咖啡库存中枢
            <span>把进货、预警与价值盘点放到一个面板。</span>
          </h1>
          <p className="hero-text">
            前端使用 Next.js + React，后端使用 .NET 8 Minimal API，数据通过 EF Core 持久化到 MySQL。
            当前界面聚焦门店最常见的三类动作：录入库存、处理低库存、查看库存价值。
          </p>
          <div className="hero-actions">
            <button className="button" type="button" onClick={resetForm}>
              新建库存项目
            </button>
            <button className="button-ghost" type="button" onClick={() => void hydrate(stockFilter)}>
              刷新数据
            </button>
          </div>
        </div>

        <div className="hero-highlights">
          <article className="insight-card">
            <span className="muted-label">关注点</span>
            <strong>{summary?.lowStockItems ?? 0}</strong>
            <span>个品项处于补货预警区间</span>
          </article>
          <article className="insight-card">
            <span className="muted-label">焦点货品</span>
            <strong>{items[0]?.name ?? "等待录入"}</strong>
            <span>优先查看库存紧张的咖啡豆</span>
          </article>
          <article className="insight-card">
            <span className="muted-label">风味覆盖</span>
            <strong>{summary?.originsTracked ?? 0}</strong>
            <span>个产地已纳入仓储管理</span>
          </article>
        </div>
      </section>

      <section className="metrics-grid">
        <article className="metric-card">
          <span className="muted-label">库存品项</span>
          <strong>{numberFormatter.format(summary?.totalItems ?? 0)}</strong>
          <div className="metric-footnote">不同 SKU 的在库数量</div>
        </article>
        <article className="metric-card">
          <span className="muted-label">总库存量</span>
          <strong>{numberFormatter.format(summary?.totalUnits ?? 0)}</strong>
          <div className="metric-footnote">全部咖啡豆库存袋数</div>
        </article>
        <article className="metric-card">
          <span className="muted-label">库存货值</span>
          <strong>{moneyFormatter.format(summary?.totalValue ?? 0)}</strong>
          <div className="metric-footnote">按单价和现存数量计算</div>
        </article>
        <article className="metric-card">
          <span className="muted-label">低库存</span>
          <strong>{numberFormatter.format(summary?.lowStockItems ?? 0)}</strong>
          <div className="metric-footnote">数量低于或等于安全库存</div>
        </article>
        <article className="metric-card">
          <span className="muted-label">数据状态</span>
          <strong>{isLoading || isPending ? "同步中" : "已更新"}</strong>
          <div className="metric-footnote">可随时刷新并继续编辑</div>
        </article>
      </section>

      <section className="content-grid">
        <section className="panel">
          <div className="panel-title-row">
            <div>
              <h2 className="panel-title">库存清单</h2>
              <p className="panel-subtitle">
                支持按关键词搜索，或只查看低库存项目。SKU、产地、供应商都会参与检索。
              </p>
            </div>
          </div>

          <div className="toolbar">
            <div className="search-wrap">
              <input
                className="search-input"
                placeholder="搜索名称、SKU、产地或供应商"
                value={searchTerm}
                onChange={(event) => setSearchTerm(event.target.value)}
              />
            </div>

            <div className="stock-filter">
              <button
                className={`filter-button ${stockFilter === "all" ? "active" : ""}`}
                type="button"
                onClick={() => setStockFilter("all")}
              >
                全部库存
              </button>
              <button
                className={`filter-button ${stockFilter === "low" ? "active" : ""}`}
                type="button"
                onClick={() => setStockFilter("low")}
              >
                低库存优先
              </button>
            </div>
          </div>

          {errorMessage ? <div className="message error">{errorMessage}</div> : null}
          {successMessage ? <div className="message success">{successMessage}</div> : null}

          {visibleItems.length === 0 && !isLoading ? (
            <div className="empty-state">当前条件下没有匹配的库存项目。</div>
          ) : null}

          <div className="table-wrap">
            <table className="inventory-table">
              <thead>
                <tr>
                  <th>咖啡豆</th>
                  <th>库存</th>
                  <th>单价</th>
                  <th>产地 / 烘焙</th>
                  <th>状态</th>
                  <th>最近补货</th>
                  <th>操作</th>
                </tr>
              </thead>
              <tbody>
                {visibleItems.map((item) => {
                  const status = getStatus(item);

                  return (
                    <tr key={item.id}>
                      <td>
                        <div className="inventory-name">
                          <strong>{item.name}</strong>
                          <span className="sku-pill">{item.sku}</span>
                          <span>{item.category}</span>
                        </div>
                      </td>
                      <td>
                        {numberFormatter.format(item.quantity)} / 安全库存 {numberFormatter.format(item.minimumStock)}
                      </td>
                      <td>{moneyFormatter.format(item.unitPrice)}</td>
                      <td>
                        {(item.origin ?? "未设置")} / {item.roastLevel}
                        <br />
                        <span className="muted-label">{item.supplier ?? "未设置供应商"}</span>
                      </td>
                      <td>
                        <span className={`status-badge ${status.tone}`}>{status.label}</span>
                      </td>
                      <td>{item.lastRestocked.slice(0, 10)}</td>
                      <td>
                        <div className="row-actions">
                          <button className="text-button" type="button" onClick={() => handleEdit(item)}>
                            编辑
                          </button>
                          <button
                            className="text-button danger"
                            type="button"
                            onClick={() => handleDelete(item.id)}
                          >
                            删除
                          </button>
                        </div>
                      </td>
                    </tr>
                  );
                })}
              </tbody>
            </table>
          </div>

          <p className="footer-note">
            当前视图会优先展示低库存项目，方便门店先处理缺口，再查看完整库存。
          </p>
        </section>

        <aside className="panel">
          <div className="panel-title-row">
            <div>
              <h2 className="panel-title">{editingId === null ? "新增库存" : "编辑库存"}</h2>
              <p className="panel-subtitle">
                使用同一张表单处理新增与更新，录入后会同步刷新面板统计。
              </p>
            </div>
          </div>

          <form className="form-grid" onSubmit={handleSubmit}>
            <div className="form-row">
              <div className="form-field">
                <label htmlFor="name">咖啡名称</label>
                <input id="name" name="name" value={form.name} onChange={handleFieldChange} required />
              </div>
              <div className="form-field">
                <label htmlFor="sku">SKU</label>
                <input id="sku" name="sku" value={form.sku} onChange={handleFieldChange} required />
              </div>
            </div>

            <div className="form-row">
              <div className="form-field">
                <label htmlFor="origin">产地</label>
                <input id="origin" name="origin" value={form.origin} onChange={handleFieldChange} />
              </div>
              <div className="form-field">
                <label htmlFor="supplier">供应商</label>
                <input id="supplier" name="supplier" value={form.supplier} onChange={handleFieldChange} />
              </div>
            </div>

            <div className="form-row">
              <div className="form-field">
                <label htmlFor="category">分类</label>
                <select id="category" name="category" value={form.category} onChange={handleFieldChange}>
                  <option value="Cafe House">Cafe House</option>
                  <option value="Single Origin">Single Origin</option>
                  <option value="Espresso Base">Espresso Base</option>
                  <option value="Seasonal Reserve">Seasonal Reserve</option>
                  <option value="Cold Brew">Cold Brew</option>
                </select>
              </div>
              <div className="form-field">
                <label htmlFor="roastLevel">烘焙度</label>
                <select
                  id="roastLevel"
                  name="roastLevel"
                  value={form.roastLevel}
                  onChange={handleFieldChange}
                >
                  <option value="Light">Light</option>
                  <option value="Medium">Medium</option>
                  <option value="Medium Dark">Medium Dark</option>
                  <option value="Dark">Dark</option>
                </select>
              </div>
            </div>

            <div className="form-row">
              <div className="form-field">
                <label htmlFor="quantity">当前库存</label>
                <input
                  id="quantity"
                  name="quantity"
                  type="number"
                  min="0"
                  value={form.quantity}
                  onChange={handleFieldChange}
                  required
                />
              </div>
              <div className="form-field">
                <label htmlFor="minimumStock">安全库存</label>
                <input
                  id="minimumStock"
                  name="minimumStock"
                  type="number"
                  min="0"
                  value={form.minimumStock}
                  onChange={handleFieldChange}
                  required
                />
              </div>
            </div>

            <div className="form-row">
              <div className="form-field">
                <label htmlFor="unitPrice">单价</label>
                <input
                  id="unitPrice"
                  name="unitPrice"
                  type="number"
                  min="0"
                  step="0.01"
                  value={form.unitPrice}
                  onChange={handleFieldChange}
                  required
                />
              </div>
              <div className="form-field">
                <label htmlFor="lastRestocked">最近补货</label>
                <input
                  id="lastRestocked"
                  name="lastRestocked"
                  type="date"
                  value={form.lastRestocked}
                  onChange={handleFieldChange}
                  required
                />
              </div>
            </div>

            <div className="form-field">
              <label htmlFor="notes">备注</label>
              <textarea id="notes" name="notes" value={form.notes} onChange={handleFieldChange} />
            </div>

            <div className="hero-actions">
              <button className="button" type="submit" disabled={isPending}>
                {editingId === null ? "保存库存" : "更新库存"}
              </button>
              <button className="button-ghost" type="button" onClick={resetForm}>
                清空表单
              </button>
            </div>
          </form>

          <div className="spotlight">
            {lowStockSpotlight.length === 0 ? (
              <div className="spotlight-card">
                <strong>当前没有低库存项目</strong>
                <p>如果继续录入新品，这里会自动显示最需要补货的咖啡豆。</p>
              </div>
            ) : (
              lowStockSpotlight.map((item) => (
                <div className="spotlight-card" key={item.id}>
                  <strong>{item.name}</strong>
                  <p>
                    当前 {item.quantity}，安全库存 {item.minimumStock}。建议优先联系
                    {item.supplier ? ` ${item.supplier}` : " 供应商"} 补货。
                  </p>
                </div>
              ))
            )}
          </div>
        </aside>
      </section>
    </main>
  );
}
