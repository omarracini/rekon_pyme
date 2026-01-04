package domain

type DashboardSummary struct {
    TotalReconciled float64 `json:"total_reconciled"`
    PendingMovements float64 `json:"pending_movements_amount"`
    PendingInvoices  float64 `json:"pending_invoices_amount"`
    Currency         string  `json:"currency"`
}