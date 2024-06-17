package invoice

import (
	"log"
	"pausalac/src/domain"
	bankAccountController "pausalac/src/infrastructure/rest/controllers/bankaccount"
	customerController "pausalac/src/infrastructure/rest/controllers/customer"
	serviceController "pausalac/src/infrastructure/rest/controllers/service"
	"pausalac/src/infrastructure/rest/controllers/user"
	"strconv"
	"time"
)

func ToResponseArray(invoices *[]domain.Invoice) *[]domain.Invoice {
	var response []domain.Invoice
	for _, invoice := range *invoices {
		response = append(response, *ToResponse(&invoice))
	}
	return &response

}

func ToResponse(invoice *domain.Invoice) *domain.Invoice {
	return &domain.Invoice{
		Id:            invoice.Id,
		Author:        invoice.Author,
		Created:       invoice.Created,
		BankAccount:   invoice.BankAccount,
		Cancelled:     invoice.Cancelled,
		Client:        invoice.Client,
		CreatedAt:     invoice.CreatedAt,
		Comment:       invoice.Comment,
		Currency:      invoice.Currency,
		IBAN:          invoice.IBAN,
		Date:          invoice.Date,
		Number:        invoice.Number,
		PaidDate:      invoice.PaidDate,
		PaidValue:     invoice.PaidValue,
		SentDate:      invoice.SentDate,
		Services:      invoice.Services,
		Status:        invoice.Status,
		SWIFT:         invoice.SWIFT,
		TradingDate:   invoice.TradingDate,
		TraidingPlace: invoice.TraidingPlace,
		Type:          invoice.Type,
		ValueInRSD:    invoice.ValueInRSD,
		UpdatedAt:     invoice.UpdatedAt,
	}
}

func ToDomain(r *CreateInvoiceRequest) *domain.NewInvoice {

	created, err := time.Parse(time.RFC3339, r.Created)
	if err != nil {
		log.Println("Error parsing time")
		return nil
	}

	canceled, err := strconv.ParseBool(r.Cancelled)
	if err != nil {
		log.Println("Error parsing bool")
		return nil
	}

	date, err := time.Parse(time.RFC3339, r.Date)
	if err != nil {
		log.Println("Error parsing time")
		return nil
	}

	paidDate, err := time.Parse(time.RFC3339, r.PaidDate)
	if err != nil {
		log.Println("Error parsing time")
		return nil
	}

	paidValue, err := strconv.ParseInt(r.PaidValue, 10, 64)
	if err != nil {
		log.Println("Error parsing int")
		return nil
	}

	sentDate, err := time.Parse(time.RFC3339, r.SentDate)
	if err != nil {
		log.Println("Error parsing time")
		return nil
	}

	tradingDate, err := time.Parse(time.RFC3339, r.TradingDate)
	if err != nil {
		log.Println("Error parsing time")
		return nil
	}

	valueInRSD, err := strconv.ParseFloat(r.ValueInRSD, 64)
	if err != nil {
		log.Println("Error parsing int")
		return nil
	}

	return &domain.NewInvoice{
		Author:        *r.Author.ToDomain().ToDomainMapper(),
		Created:       created,
		BankAccount:   *bankAccountController.ToDomain(&r.BankAccount).ToDomainBankAccountMapper(),
		Cancelled:     canceled,
		Client:        *customerController.ToDomain(&r.Client).ToDomainCustomerMapper(),
		Comment:       r.Comment,
		Currency:      r.Currency,
		IBAN:          r.IBAN,
		Date:          date,
		Number:        r.Number,
		PaidDate:      paidDate,
		PaidValue:     paidValue,
		SentDate:      sentDate,
		Services:      serviceController.ToDomainArray(&r.Services).ToDomainServiceArrayMapper(),
		Status:        domain.InvoiceStatus(r.Status),
		SWIFT:         r.SWIFT,
		TradingDate:   tradingDate,
		TraidingPlace: r.TraidingPlace,
		Type:          r.Type,
		ValueInRSD:    valueInRSD,
	}
}

func ToDomainUpdate(r *UpdateInvoiceRequest) map[string]interface{} {

	invoiceMap := make(map[string]interface{})
	if r.Author != (user.UpdateUserRequest{}) {
		invoiceMap["author"] = r.Author.ToDomainUpdate()
	}
	if r.BankAccount != (bankAccountController.UpdateBankAccountRequest{}) {
		invoiceMap["bank_account"] = bankAccountController.ToDomainUpdate(&r.BankAccount)
	}
	if r.Cancelled != "" {
		invoiceMap["cancelled"] = r.Cancelled
	}
	if r.Client != (customerController.UpdateCustomerRequest{}) {
		invoiceMap["client"] = customerController.ToDomainUpdate(&r.Client)
	}
	if r.Comment != "" {
		invoiceMap["comment"] = r.Comment
	}
	if r.Currency != "" {
		invoiceMap["currency"] = r.Currency
	}
	if r.IBAN != "" {
		invoiceMap["iban"] = r.IBAN
	}
	if r.Date != "" {
		invoiceMap["date"] = r.Date
	}
	if r.Number != "" {
		invoiceMap["number"] = r.Number
	}
	if r.PaidDate != "" {
		invoiceMap["paid_date"] = r.PaidDate
	}
	if r.PaidValue != "" {
		invoiceMap["paid_value"] = r.PaidValue
	}
	if r.SentDate != "" {
		invoiceMap["sent_date"] = r.SentDate
	}
	if !r.Services.Equals(serviceController.UpdateServiceArrayRequest{}) {
		invoiceMap["services"] = r.Services
	}
	if r.Status != "" {
		invoiceMap["status"] = r.Status
	}
	if r.SWIFT != "" {
		invoiceMap["swift"] = r.SWIFT
	}
	if r.TradingDate != "" {
		invoiceMap["trading_date"] = r.TradingDate
	}
	if r.TraidingPlace != "" {
		invoiceMap["traiding_place"] = r.TraidingPlace
	}
	if r.Type != "" {
		invoiceMap["type"] = r.Type
	}
	if r.ValueInRSD != "" {
		invoiceMap["value_in_rsd"] = r.ValueInRSD
	}
	invoiceMap["updated_at"] = time.Now()
	return invoiceMap
}
