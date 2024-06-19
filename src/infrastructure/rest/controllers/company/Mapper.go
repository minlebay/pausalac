package company

import (
	"pausalac/src/domain"
	"pausalac/src/infrastructure/rest/controllers/bankaccount"
	"time"
)

// ToResponse maps Company to CompanyResponse
func ToResponse(company *domain.Company) *CompanyResponse {

	var bankAccounts []bankaccount.BankAccountResponse
	for _, bankAccount := range company.BankAccounts {
		bankAccountResponse := bankaccount.ToResponse(&bankAccount)
		bankAccounts = append(bankAccounts, bankAccountResponse)
	}

	return &CompanyResponse{
		Id:                           company.Id.Hex(),
		Author:                       company.Author,
		AgencyId:                     company.AgencyId,
		Name:                         company.Name,
		FullName:                     company.FullName,
		PIB:                          company.PIB,
		IdentificationNumber:         company.IdentificationNumber,
		ForeignExchangeAccountNumber: company.ForeignExchangeAccountNumber,
		CallNumber:                   company.CallNumber,
		DateOfRegistration:           company.DateOfRegistration,
		City:                         company.City,
		ActivityCodeId:               company.ActivityCodeId,
		MunicipalityId:               company.MunicipalityId,
		Logo:                         company.Logo,
		StreetAddress:                company.StreetAddress,
		StreetNumber:                 company.StreetNumber,
		Phone:                        company.Phone,
		AgencyEmail:                  company.AgencyEmail,
		Signature:                    company.Signature,
		EmploymentType:               company.EmploymentType,
		InvoiceDescription:           company.InvoiceDescription,
		CreatedAt:                    company.CreatedAt,
		UpdatedAt:                    company.UpdatedAt,
		BankAccounts:                 bankAccounts,
	}
}

// ToDomainArray maps an array of Company to CompanyResponse array
func ToResponseArray(companies *[]domain.Company) []CompanyResponse {
	var companyResponses []CompanyResponse
	for _, company := range *companies {
		companyResponses = append(companyResponses, *ToResponse(&company))
	}
	return companyResponses
}

// ToDomain converts CreateCompanyRequest to domain Company
func ToDomain(req *CreateCompanyRequest) *domain.NewCompany {

	var bankAccounts []domain.NewBankAccount
	for _, bankAccount := range req.BankAccounts {
		bankAccountDomain := bankaccount.ToDomain(&bankAccount)
		bankAccounts = append(bankAccounts, *bankAccountDomain)
	}

	return &domain.NewCompany{
		Author:                       req.Author,
		AgencyId:                     req.AgencyId,
		Name:                         req.Name,
		FullName:                     req.FullName,
		PIB:                          req.PIB,
		IdentificationNumber:         req.IdentificationNumber,
		ForeignExchangeAccountNumber: req.ForeignExchangeAccountNumber,
		CallNumber:                   req.CallNumber,
		DateOfRegistration:           req.DateOfRegistration,
		City:                         req.City,
		ActivityCodeId:               req.ActivityCodeId,
		MunicipalityId:               req.MunicipalityId,
		Logo:                         req.Logo,
		StreetAddress:                req.StreetAddress,
		StreetNumber:                 req.StreetNumber,
		Phone:                        req.Phone,
		AgencyEmail:                  req.AgencyEmail,
		Signature:                    req.Signature,
		EmploymentType:               req.EmploymentType,
		InvoiceDescription:           req.InvoiceDescription,
		BankAccounts:                 bankAccounts,
	}
}

// ToMap maps UpdateCompanyRequest to a map for updating the company
func ToDomainUpdate(req *UpdateCompanyRequest) map[string]interface{} {
	companyMap := make(map[string]interface{})
	if req.Name != "" {
		companyMap["name"] = req.Name
	}
	if req.FullName != "" {
		companyMap["full_name"] = req.FullName
	}
	if req.PIB != "" {
		companyMap["pib"] = req.PIB
	}
	if req.IdentificationNumber != "" {
		companyMap["identification_number"] = req.IdentificationNumber
	}
	if req.FirstAccountNumber != "" {
		companyMap["first_account_number"] = req.FirstAccountNumber
	}
	if req.CallNumber != "" {
		companyMap["call_number"] = req.CallNumber
	}
	if req.DateOfRegistration != "" {
		companyMap["date_of_registration"] = req.DateOfRegistration
	}
	if req.City != "" {
		companyMap["city"] = req.City
	}
	if req.ActivityCodeId != "" {
		companyMap["activity_code_id"] = req.ActivityCodeId
	}
	if req.MunicipalityId != "" {
		companyMap["municipality_id"] = req.MunicipalityId
	}
	if req.Logo != "" {
		companyMap["logo"] = req.Logo
	}
	if req.StreetAddress != "" {
		companyMap["street_address"] = req.StreetAddress
	}
	if req.StreetNumber != "" {
		companyMap["street_number"] = req.StreetNumber
	}
	if req.Phone != "" {
		companyMap["phone"] = req.Phone
	}
	if req.AgencyEmail != "" {
		companyMap["agency_email"] = req.AgencyEmail
	}
	if req.Signature != "" {
		companyMap["signature"] = req.Signature
	}
	if req.EmploymentType != "" {
		companyMap["employment_type"] = req.EmploymentType
	}
	if req.InvoiceDescription != "" {
		companyMap["invoice_description"] = req.InvoiceDescription
	}
	companyMap["updated_at"] = time.Now()
	return companyMap
}
