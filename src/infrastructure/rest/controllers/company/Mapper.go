package company

import (
	"pausalac/src/domain"
	"time"
)

// ToResponse maps Company to CompanyResponse
func ToResponse(company *domain.Company) *CompanyResponse {
	return &CompanyResponse{
		ID:                           company.ID.Hex(),
		UserID:                       company.UserID,
		AgencyID:                     company.AgencyID,
		Name:                         company.Name,
		FullName:                     company.FullName,
		PIB:                          company.PIB,
		IdentificationNumber:         company.IdentificationNumber,
		FirstAccountNumber:           company.FirstAccountNumber,
		SecondAccountNumber:          company.SecondAccountNumber,
		ForeignExchangeAccountNumber: company.ForeignExchangeAccountNumber,
		CallNumber:                   company.CallNumber,
		DateOfRegistration:           company.DateOfRegistration,
		City:                         company.City,
		ActivityCodeID:               company.ActivityCodeID,
		MunicipalityID:               company.MunicipalityID,
		EmployedByOtherFirm:          company.EmployedByOtherFirm,
		EmploymentChanged:            company.EmploymentChanged,
		Logo:                         company.Logo,
		StreetAddress:                company.StreetAddress,
		StreetNumber:                 company.StreetNumber,
		Phone:                        company.Phone,
		AgencyEmail:                  company.AgencyEmail,
		SWIFT:                        company.SWIFT,
		IBAN:                         company.IBAN,
		Signature:                    company.Signature,
		EmploymentType:               company.EmploymentType,
		InvoiceDescription:           company.InvoiceDescription,
		CreatedAt:                    company.CreatedAt,
		UpdatedAt:                    company.UpdatedAt,
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
	return &domain.NewCompany{
		UserID:                       req.UserID,
		AgencyID:                     req.AgencyID,
		Name:                         req.Name,
		FullName:                     req.FullName,
		PIB:                          req.PIB,
		IdentificationNumber:         req.IdentificationNumber,
		FirstAccountNumber:           req.FirstAccountNumber,
		SecondAccountNumber:          req.SecondAccountNumber,
		ForeignExchangeAccountNumber: req.ForeignExchangeAccountNumber,
		CallNumber:                   req.CallNumber,
		DateOfRegistration:           req.DateOfRegistration,
		City:                         req.City,
		ActivityCodeID:               req.ActivityCodeID,
		MunicipalityID:               req.MunicipalityID,
		EmployedByOtherFirm:          req.EmployedByOtherFirm,
		EmploymentChanged:            req.EmploymentChanged,
		Logo:                         req.Logo,
		StreetAddress:                req.StreetAddress,
		StreetNumber:                 req.StreetNumber,
		Phone:                        req.Phone,
		AgencyEmail:                  req.AgencyEmail,
		SWIFT:                        req.SWIFT,
		IBAN:                         req.IBAN,
		Signature:                    req.Signature,
		EmploymentType:               req.EmploymentType,
		InvoiceDescription:           req.InvoiceDescription,
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
	if req.SecondAccountNumber != "" {
		companyMap["second_account_number"] = req.SecondAccountNumber
	}
	if req.ForeignExchangeAccountNumber != "" {
		companyMap["foreign_exchange_account_number"] = req.ForeignExchangeAccountNumber
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
	if req.ActivityCodeID != "" {
		companyMap["activity_code_id"] = req.ActivityCodeID
	}
	if req.MunicipalityID != "" {
		companyMap["municipality_id"] = req.MunicipalityID
	}
	if req.EmployedByOtherFirm != "" {
		companyMap["employed_by_other_firm"] = req.EmployedByOtherFirm
	}
	if req.EmploymentChanged != "" {
		companyMap["employment_changed"] = req.EmploymentChanged
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
	if req.SWIFT != "" {
		companyMap["swift"] = req.SWIFT
	}
	if req.IBAN != "" {
		companyMap["iban"] = req.IBAN
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
