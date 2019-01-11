package acaseSts

import "encoding/xml"

type OrderListRequestType struct {
	Credentials
	XMLName               xml.Name `xml:"OrderListRequest"`
	ArrivalDateFrom       string   `xml:"ArrivalDateFrom,attr"`
	ArrivalDateTo         string   `xml:"ArrivalDateTo,attr"`
	DepartureDateFrom     string   `xml:"DepartureDateFrom,attr"`
	DepartureDateTo       string   `xml:"DepartureDateTo,attr"`
	DeadlineDateFrom      string   `xml:"DeadlineDateFrom,attr"`
	DeadlineDateTo        string   `xml:"DeadlineDateTo,attr"`
	RegistrationDateFrom  string   `xml:"RegistrationDateFrom,attr"`
	RegistrationDateTo    string   `xml:"RegistrationDateTo,attr"`
	AccommodationDateFrom string   `xml:"AccommodationDateFrom,attr"`
	AccommodationDateTo   string   `xml:"AccommodationDateTo,attr"`
	Hotel                 int      `xml:"Hotel,attr"`
	HotelName             string   `xml:"HotelName,attr"`
	LastName              string   `xml:"LastName,attr"`
}

type OrderListResponseType struct {
	BaseResponse
	XMLName   xml.Name      `xml:"OrderListResponse"`
	IsThatAll YesNoCodeType `xml:"IsThatAll"`
	Orders    OrdersType    `xml:"Orders"`
}

type OrdersType struct {
	XMLName xml.Name    `xml:"Orders"`
	Items   []OrderType `xml:"Order"`
}

type OrderType struct {
	XMLName                xml.Name           `xml:"Order"`
	Id                     int                `xml:"Id,attr"`
	ReferenceNumber        string             `xml:"ReferenceNumber,attr"`
	RegistrationDate       string             `xml:"RegistrationDate,attr"`
	DeadlineDate           string             `xml:"DeadlineDate,attr"`
	Price                  float64            `xml:"Price,attr"`
	TravelAgencyCommission float64            `xml:"TravelAgencyCommission,attr"`
	Penalty                float64            `xml:"Penalty,attr"`
	StartDate              string             `xml:"StartDate,attr"`
	EndDate                string             `xml:"EndDate,attr"`
	Currency               SimpleCodeNameType `xml:"Currency"`
	Status                 SimpleCodeNameType `xml:"Status"`
	ContactPerson          ContactPersonType  `xml:"ContactPerson"`
}
