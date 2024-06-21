package domain

type DateRangeFilter struct {
	Field string
	Start string
	End   string
}

type Types interface {
	User | Company | Customer | Service | BankAccount | Invoice
}
