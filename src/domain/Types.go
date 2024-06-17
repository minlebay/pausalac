package domain

type DateRangeFilter struct {
	Field string
	Start string
	End   string
}

type Types interface {
	Company | Customer | Service | BankAccount | Invoice
}

type NewTypes interface {
	NewCompany | NewCustomer | NewService | NewBankAccount | NewInvoice
}
