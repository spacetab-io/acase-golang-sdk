package acaseSts

import "encoding/xml"

type OrderInfoNotifyRequestType struct {
	XMLName                xml.Name              `xml:"OrderInfoNotifyRequest"`
	Password               string                `xml:"Password,attr"`
	Language               LanguageTypeEnum      `xml:"Language,attr,omitempty"`
	ContractNumber         string                `xml:"ContractNumber,attr"`
	ContractDate           string                `xml:"ContractDate,attr"`
	GuaranteeAmount        float64               `xml:"GuaranteeAmount,attr"`
	NotifierId             string                `xml:"NotifierId,attr"`
	Id                     int                   `xml:"Id,attr"`
	ReferenceNumber        string                `xml:"ReferenceNumber,attr"`
	ReferenceNumberName    string                `xml:"ReferenceNumberName,attr"`
	RegistrationDate       string                `xml:"RegistrationDate,attr"`
	DeadlineDate           string                `xml:"DeadlineDate,attr"`
	Price                  float64               `xml:"Price,attr"`
	TravelAgencyCommission float64               `xml:"TravelAgencyCommission,attr"`
	Penalty                float64               `xml:"Penalty,attr"`
	StartDate              string                `xml:"StartDate,attr"`
	EndDate                string                `xml:"EndDate,attr"`
	InvoiceRule            int                   `xml:"InvoiceRule,attr"`
	InvoiceId              int                   `xml:"InvoiceId,attr"`
	ConfirmId              int                   `xml:"ConfirmId,attr"`
	Currency               SimpleCodeNameType    `xml:"Currency"`
	Autonom                SimpleCodeNameType    `xml:"Autonom"`
	VatLimit               SimpleCodeNameType    `xml:"VatLimit"`
	TypeContract           SimpleCodeNameType    `xml:"TypeContract"`
	WhereToPay             SimpleCodeNameType    `xml:"WhereToPay"`
	SubjectToProcessing    YesNoCodeType         `xml:"SubjectToProcessing"`
	FinancialCondition     SimpleCodeNameType    `xml:"FinancialCondition"`
	Status                 SimpleCodeNameType    `xml:"Status"`
	ContactPerson          ContactPersonType     `xml:"ContactPerson"`
	AccountDetails         AccountDetailsType    `xml:"AccountDetails"`
	Customer               CustomerType          `xml:"Customer"`
	IsCustomer             SimpleCodeNameType    `xml:"IsCustomer"`
	AccommodationList      AccommodationListType `xml:"AccommodationList"`
}

type AccommodationListType struct {
	XMLName       xml.Name            `xml:"AccommodationList"`
	Accommodation []AccommodationType `xml:"Accommodation"`
}

type AccommodationType struct {
	XMLName                        xml.Name            `xml:"Accommodation"`
	Id                             int                 `xml:"Id,attr"`
	ArrivalDate                    string              `xml:"ArrivalDate,attr"`
	ArrivalTime                    string              `xml:"ArrivalTime,attr"`
	DepartureDate                  string              `xml:"DepartureDate,attr"`
	DepartureTime                  string              `xml:"DepartureTime,attr"`
	NumberOfNights                 int                 `xml:"NumberOfNights,attr"`
	NumberOfRooms                  int                 `xml:"NumberOfRooms,attr"`
	NumberOfGuests                 int                 `xml:"NumberOfGuests,attr"`
	NumberOfExtraBedsAdult         int                 `xml:"NumberOfExtraBedsAdult,attr,omitempty"`
	NumberOfExtraBedsChild         int                 `xml:"NumberOfExtraBedsChild,attr,omitempty"`
	NumberOfExtraBedsInfant        int                 `xml:"NumberOfExtraBedsInfant,attr,omitempty"`
	AdditionalInfo                 string              `xml:"AdditionalInfo,attr"`
	ReferenceNumber                string              `xml:"ReferenceNumber,attr,omitempty"`
	SupplierInfo                   string              `xml:"SupplierInfo,attr"`
	ConfirmationNumber             string              `xml:"ConfirmationNumber,attr"`
	Price                          float64             `xml:"Price,attr"`
	VATIncludedInPrice             float64             `xml:"VATIncludedInPrice,attr"`
	TravelAgencyCommission         float64             `xml:"TravelAgencyCommission,attr"`
	Penalty                        float64             `xml:"Penalty,attr"`
	DeadlineDate                   string              `xml:"DeadlineDate,attr"`
	DeadlineTimeLoc                string              `xml:"DeadlineTimeLoc,attr"`
	DeadlineTimeSys                string              `xml:"DeadlineTimeSys,attr"`
	DeadlineTimeUTC                string              `xml:"DeadlineTimeUTC,attr"`
	PossiblePenaltySize            float64             `xml:"PossiblePenaltySize,attr"`
	VoucherId                      int                 `xml:"VoucherId,attr,omitempty"`
	Gain                           float64             `xml:"Gain,attr,omitempty"`
	PenaltyGain                    float64             `xml:"PenaltyGain,attr,omitempty"`
	ToBeCancelled                  YesNoCodeType       `xml:"ToBeCancelled"`
	ToBeChanged                    YesNoCodeType       `xml:"ToBeChanged"`
	Hotel                          SimpleCodeNameType  `xml:"Hotel"`
	ObjType                        ObjTypeType         `xml:"ObjType"`
	Country                        SimpleCodeNameType  `xml:"Country"`
	AdmUnit1                       *SimpleCodeNameType `xml:"AdmUnit1,omitempty"`
	AdmUnit2                       *SimpleCodeNameType `xml:"AdmUnit2,omitempty"`
	TypeOfPlace                    *SimpleCodeNameType `xml:"TypeOfPlace,omitempty"`
	City                           SimpleCodeNameType  `xml:"City"`
	Status                         SimpleCodeNameType  `xml:"Status"`
	AmendmentRejected              YesNoCodeType       `xml:"AmendmentRejected"`
	AllowCancel                    YesNoCodeType       `xml:"AllowCancel"`
	AllowModify                    YesNoCodeType       `xml:"AllowModify"`
	AllowSetGain                   *YesNoCodeType      `xml:"AllowSetGain,omitempty"`
	Product                        ProductType         `xml:"Product"`
	Meal                           SimpleCodeNameType  `xml:"Meal"`
	Persons                        PersonsType         `xml:"Persons"`
	Penalties                      PenaltiesType       `xml:"Penalties"`
	IsEarlyCheckIn                 YesNoCodeType       `xml:"IsEarlyCheckIn"`
	CriticalEarlyCheckIn           SimpleCodeNameType  `xml:"CriticalEarlyCheckIn"`
	EarlyCheckInConfirmationStatus SimpleCodeNameType  `xml:"EarlyCheckInConfirmationStatus"`
	IsHour                         YesNoCodeType       `xml:"IsHour"`
	IsLateCheckOut                 YesNoCodeType       `xml:"IsLateCheckOut"`
	CriticalLateCheckOut           SimpleCodeNameType  `xml:"CriticalLateCheckOut"`
	LateCheckOutConfirmationStatus SimpleCodeNameType  `xml:"LateCheckOutConfirmationStatus"`
}

type PenaltiesType struct {
	XMLName xml.Name      `xml:"Penalties"`
	Items   []PenaltyType `xml:"Penalty"`
}

type PenaltyType struct {
	SimpleCodeNameType
	XMLName            xml.Name           `xml:"Penalty"`
	Gain               float64            `xml:"Gain,attr"`
	FullVATInPrice     int                `xml:"FullVATInPrice,attr"`
	Id                 int                `xml:"Id,attr"`
	VATIncludedInPrice float64            `xml:"VATIncludedInPrice,attr"`
	Price              float64            `xml:"Price,attr"`
	AllowSetGain       YesNoCodeType      `xml:"AllowSetGain"`
	Reason             SimpleCodeNameType `xml:"Reason"`
}

type PersonsType struct {
	XMLName xml.Name     `xml:"Persons"`
	Items   []PersonType `xml:"Person"`
}

type PersonType struct {
	XMLName     xml.Name           `xml:"Person"`
	LastName    string             `xml:"LastName,attr"`
	FirstName   string             `xml:"FirstName,attr"`
	Company     string             `xml:"Company,attr,omitempty"`
	Position    string             `xml:"Position,attr,omitempty"`
	Category    SimpleCodeNameType `xml:"Category"`
	Citizenship SimpleCodeNameType `xml:"Citizenship"`
}

type ProductType struct {
	XMLName         xml.Name      `xml:"Product"`
	Code            int           `xml:"Code,attr"`
	RoomName        string        `xml:"RoomName,attr"`
	RoomCode        int           `xml:"RoomCode,attr"`
	RateCode        int           `xml:"RateCode,attr"`
	RateName        string        `xml:"RateName,attr"`
	RateDescription string        `xml:"RateDescription,attr"`
	RateGroupCode   int           `xml:"RateGroupCode,attr"`
	RateGroupName   string        `xml:"RateGroupName,attr"`
	Allotment       AllotmentType `xml:"Allotment"`
}

type AccountDetailsType struct {
	XMLName xml.Name           `xml:"AccountDetails"`
	IsGain  SimpleCodeNameType `xml:"IsGain"`
}

type ContactPersonType struct {
	XMLName xml.Name `xml:"ContactPerson"`
	Name    string   `xml:"Name,attr"`
	Phone   string   `xml:"Phone,attr"`
	Fax     string   `xml:"Fax,attr"`
	Email   string   `xml:"Email,attr"`
}

type OrderInfoNotifyResponseType struct {
	BaseResponse
	XMLName xml.Name `xml:"OrderInfoNotifyResponse"`
}
