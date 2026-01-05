export interface DashboardSummary {
    total_reconciled: number;
    pending_movements_amount: number;
    pending_invoices_amount: number;
    currency: string;
}

export interface AICategorySuggestion {
    Category: string;
    Confidence: number;
    reason: string;
}