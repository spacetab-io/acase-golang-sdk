package acaseSts

import "encoding/xml"

type OrderInfoAwocNotifyRequestType struct {
	XMLName					xml.Name			`xml:"OrderInfoAwocNotifyRequest"`
	Id						int					`xml:"Id,attr"`
	ReferenceNumber			string				`xml:"ReferenceNumber,attr"`
	ReferenceNumberName		string				`xml:"ReferenceNumberName,attr"`
	RegistrationDate		string				`xml:"RegistrationDate,attr"`
	DeadlineDate			string				`xml:"DeadlineDate,attr"`
	Price					float64				`xml:"Price,attr"`
	TravelAgencyCommission	float64				`xml:"TravelAgencyCommission,attr"`
	Penalty					float64				`xml:"Penalty,attr"`
	StartDate				string				`xml:"StartDate,attr"`
	EndDate					string				`xml:"EndDate,attr"`
	InvoiceRule				int					`xml:"InvoiceRule,attr"`
	InvoiceId				int					`xml:"InvoiceId,attr"`
	ConfirmId				int					`xml:"ConfirmId,attr"`
	Success					string				`xml:"Success"`
	Currency				CurrencyType		`xml:"Currency"`
	WhereToPay				SimpleCodeNameType	`xml:"WhereToPay"`
	Status					SimpleCodeNameType	`xml:"Status"`
	ContactPerson			ContactPersonType	`xml:"ContactPerson"`
	Customer				CustomerType		`xml:"Customer"`
	IsCustomer				SimpleCodeNameType	`xml:"IsCustomer"`
	AwocList				AwocListType		`xml:"AwocList"`
}

type AwocListType struct {
	XMLName	xml.Name	`xml:"AwocList"`
	Items	[]AwocType	`xml:"Awoc"`
}

type AwocType struct {
	XMLName					xml.Name			`xml:"Awoc"`
	Id						int					`xml:"Id,attr"`
	ParentId				int					`xml:"ParentId,attr"`
	Subject					string				`xml:"Subject,attr"`
	ArrivalDate				string				`xml:"ArrivalDate,attr"`
	DepartureDate			string				`xml:"DepartureDate,attr"`
	NumberOfNights			int					`xml:"NumberOfNights,attr"`
	ArrivalTime				string				`xml:"ArrivalTime,attr"`
	DepartureTime			string				`xml:"DepartureTime,attr"`
	Price					float64				`xml:"Price,attr"`
	VATIncludedInPrice		float64				`xml:"VATIncludedInPrice,attr"`
	TravelAgencyCommission	float64				`xml:"TravelAgencyCommission,attr"`
	DeadlineDate			string				`xml:"DeadlineDate,attr"`
	PossiblePenaltySize		float64				`xml:"PossiblePenaltySize,attr"`
	Penalty					float64				`xml:"Penalty,attr"`
	ConfirmationNumber		string				`xml:"ConfirmationNumber,attr"`
	ReferenceNumber			string				`xml:"ReferenceNumber,attr"`
	SupplierInfo			string				`xml:"SupplierInfo,attr"`
	ToBeCancelled			YesNoCodeType		`xml:"ToBeCancelled"`
	ToBeChanged				YesNoCodeType		`xml:"ToBeChanged"`
	ToBeSelected			YesNoCodeType		`xml:"ToBeSelected"`
	Hotel					SimpleCodeNameType	`xml:"Hotel"`
	Country					SimpleCodeNameType	`xml:"Country"`
	AdmUnit1				AdmUnit1Type		`xml:"AdmUnit1"`
	AdmUnit2				AdmUnit2Type		`xml:"AdmUnit2"`
	TypeOfPlace				SimpleCodeNameType	`xml:"TypeOfPlace"`
	City					SimpleCodeNameType	`xml:"City"`
	Status					SimpleCodeNameType	`xml:"Status"`
	AllowCancel				YesNoCodeType		`xml:"AllowCancel"`
	AllowModify				YesNoCodeType		`xml:"AllowModify"`
	AllowSelect				YesNoCodeType		`xml:"AllowSelect"`
	DetailList				DetailListType		`xml:"DetailList"`
	Persons					PersonsType			`xml:"Persons"`
	Penalties				PenaltiesType		`xml:"Penalties"`
}

type DetailListType struct {
	XMLName	xml.Name		`xml:"DetailList"`
	Items	[]DetailType	`xml:"Detail"`
}

type DetailType struct {
	XMLName					xml.Name	`xml:"Detail"`
	Id						int			`xml:"Id,attr"`
	Subject					string		`xml:"Subject,attr"`
	Quantity				float64		`xml:"Quantity,attr"`
	Duration				float64		`xml:"Duration,attr"`
	Price					float64		`xml:"Price,attr"`
	TravelAgencyCommission	float64		`xml:"TravelAgencyCommission,attr"`
	VATIncludedInPrice		float64		`xml:"VATIncludedInPrice,attr"`
}

