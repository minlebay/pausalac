package domain

type DateRangeFilter struct {
	Field string
	Start string
	End   string
}

type Types interface {
	User | Company | Customer | Service | BankAccount | Invoice
}

type NewTypes interface {
	NewUser | NewCompany | NewCustomer | NewService | NewBankAccount | NewInvoice
}
