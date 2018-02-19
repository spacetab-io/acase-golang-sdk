package acaseSts

import "encoding/xml"

type OrderInfoRequestType struct {
	Credentials
	XMLName			xml.Name	`xml:"OrderInfoRequest"`
	Id				string		`xml:"Id"`
}

type OrderInfoResponseType struct {
	BaseResponse
	XMLName					xml.Name				`xml:"OrderInfoResponse"`
	Language				string					`xml:"Language,attr"`
	Password				string					`xml:"Password,attr"`
	Id						string					`xml:"Id,attr"`
	ReferenceNumber			string					`xml:"ReferenceNumber,attr"`
	ReferenceNumberName		string					`xml:"ReferenceNumberName,attr"`
	RegistrationDate		string					`xml:"RegistrationDate,attr"`
	DeadlineDate			string					`xml:"DeadlineDate,attr"`
	Price					float64					`xml:"Price,attr"`
	TravelAgencyCommission	float64					`xml:"TravelAgencyCommission,attr"`
	Penalty					float64					`xml:"Penalty,attr"`
	StartDate				string					`xml:"StartDate,attr"`
	EndDate					string					`xml:"EndDate,attr"`
	InvoiceRule				int						`xml:"InvoiceRule,attr"`
	InvoiceId				string					`xml:"InvoiceId,attr"`
	ConfirmId				string					`xml:"ConfirmId,attr"`
	ContractDate			string					`xml:"ContractDate,attr"`
	ContractNumber			string					`xml:"ContractNumber,attr"`
	GuaranteeAmount			float64					`xml:"GuaranteeAmount,attr"`
	Currency				SimpleCodeNameType		`xml:"Currency"`
	Autonom					SimpleCodeNameType		`xml:"Autonom"`
	VatLimit				SimpleCodeNameType		`xml:"VatLimit"`
	TypeContract			SimpleCodeNameType		`xml:"TypeContract"`
	WhereToPay				SimpleCodeNameType		`xml:"WhereToPay"`
	SubjectToProcessing		YesNoCodeType			`xml:"SubjectToProcessing"`
	FinancialCondition		SimpleCodeNameType		`xml:"FinancialCondition"`
	Status					SimpleCodeNameType		`xml:"Status"`
	ContactPerson			ContactPersonType		`xml:"ContactPerson"`
	AccountDetails			AccountDetailsType		`xml:"AccountDetails"`
	Customer				CustomerType			`xml:"Customer"`
	IsCustomer				YesNoCodeType			`xml:"IsCustomer"`
	AccommodationList		AccommodationListType	`xml:"AccommodationList"`
}
