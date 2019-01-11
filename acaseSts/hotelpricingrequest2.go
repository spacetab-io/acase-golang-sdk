package acaseSts

import "encoding/xml"

type HotelPricingRequest2Type struct {
	Credentials
	XMLName                 xml.Name `xml:"HotelPricingRequest2"`
	Hotel                   int      `xml:"Hotel,attr"`
	ProductCode             int      `xml:"ProductCode,attr"`
	Currency                int      `xml:"Currency,attr"`
	WhereToPay              int      `xml:"WhereToPay,attr"`
	ArrivalDate             string   `xml:"ArrivalDate,attr"`
	DepartureDate           string   `xml:"DepartureDate,attr"`
	ArrivalTime             string   `xml:"ArrivalTime,attr"`
	DepartureTime           string   `xml:"DepartureTime,attr"`
	NumberOfGuests          int      `xml:"NumberOfGuests,attr"`
	NumberOfExtraBedsAdult  int      `xml:"NumberOfExtraBedsAdult,attr"`
	NumberOfExtraBedsChild  int      `xml:"NumberOfExtraBedsChild,attr,omitempty"`
	NumberOfExtraBedsInfant int      `xml:"NumberOfExtraBedsInfant,attr,omitempty"`
	Meal                    int      `xml:"Meal,attr"`
	Id                      string   `xml:"Id,attr"`
	AccommodationId         string   `xml:"AccommodationId,attr"`
}

type HotelPricingResponse2Type struct {
	BaseResponse
	SimpleCodeNameType
	XMLName                 xml.Name                   `xml:"HotelPricingResponse2"`
	Country                 CountryType                `xml:"Country"`
	StartDate               string                     `xml:"StartDate,attr"`
	EndDate                 string                     `xml:"EndDate,attr"`
	NumberOfGuests          int                        `xml:"NumberOfGuests,attr"`
	NumberOfExtraBedsAdult  int                        `xml:"NumberOfExtraBedsAdult,attr"`
	NumberOfExtraBedsChild  int                        `xml:"NumberOfExtraBedsChild,attr,omitempty"`
	NumberOfExtraBedsInfant int                        `xml:"NumberOfExtraBedsInfant,attr,omitempty"`
	SpecialRequirements     string                     `xml:"SpecialRequirements,attr"`
	AdmUnit1                SimpleCodeNameType         `xml:"AdmUnit1"`
	AdmUnit2                AdmUnit2Type               `xml:"AdmUnit2"`
	TypeOfPlace             SimpleCodeNameType         `xml:"TypeOfPlace"`
	City                    CityType                   `xml:"City"`
	ObjType                 ObjTypeType                `xml:"ObjType"`
	Position                PositionType               `xml:"Position"`
	Currency                SimpleCodeNameType         `xml:"Currency"`
	Infants                 AgeRestrictionType         `xml:"Infants"`
	Children                AgeRestrictionType         `xml:"Children"`
	SpecialRequirementList  SpecialRequirementListType `xml:"SpecialRequirementList"`
	HotelPricingDetail      HotelPricingDetailType     `xml:"HotelPricingDetail"`
	EarlyCheckInList        EarlyCheckInListType       `xml:"EarlyCheckInList"`
	LateCheckOutList        LateCheckOutListType       `xml:"LateCheckOutList"`
}

type LateCheckOutListType struct {
	XMLName xml.Name           `xml:"LateCheckOutList"`
	Items   []LateCheckOutType `xml:"LateCheckOut"`
}

type LateCheckOutType struct {
	XMLName                      xml.Name           `xml:"LateCheckOut"`
	Time                         string             `xml:"Time,attr"`
	Price                        float64            `xml:"Price,attr"`
	Text                         string             `xml:"Text,attr"`
	Guaranteed                   SimpleCodeNameType `xml:"Guaranteed"`
	EarlyCheckInLateCheckOutRule SimpleCodeNameType `xml:"EarlyCheckInLateCheckOutRule"`
}

type EarlyCheckInListType struct {
	XMLName xml.Name           `xml:"EarlyCheckInList"`
	Items   []EarlyCheckInType `xml:"EarlyCheckIn"`
}

type EarlyCheckInType struct {
	XMLName                      xml.Name           `xml:"EarlyCheckIn"`
	Time                         string             `xml:"Time,attr"`
	Price                        float64            `xml:"Price,attr"`
	Text                         string             `xml:"Text,attr"`
	Guaranteed                   SimpleCodeNameType `xml:"Guaranteed"`
	EarlyCheckInLateCheckOutRule SimpleCodeNameType `xml:"EarlyCheckInLateCheckOutRule"`
}

type HotelPricingDetailType struct {
	XMLName                    xml.Name             `xml:"HotelPricingDetail"`
	Code                       string               `xml:"Code,attr"`
	RoomName                   string               `xml:"RoomName,attr"`
	RoomCode                   string               `xml:"RoomCode,attr"`
	RateCode                   string               `xml:"RateCode,attr"`
	RateName                   string               `xml:"RateName,attr"`
	RateDescription            string               `xml:"RateDescription,attr"`
	RateGroupCode              string               `xml:"RateGroupCode,attr"`
	RateGroupName              string               `xml:"RateGroupName,attr"`
	IsHourRate                 string               `xml:"IsHourRate,attr"`
	MaxHours                   int                  `xml:"MaxHours,attr"`
	StartTime                  string               `xml:"StartTime,attr"`
	EndTime                    string               `xml:"EndTime,attr"`
	MinimumRetailPrice         float64              `xml:"MinimumRetailPrice,attr"`
	Price                      float64              `xml:"Price,attr"`
	TravelAgencyCommission     float64              `xml:"TravelAgencyCommission,attr"`
	FullVATInPrice             string               `xml:"FullVATInPrice,attr"`
	VATIncludedInPrice         float64              `xml:"VATIncludedInPrice,attr"`
	DeadlineDate               string               `xml:"DeadlineDate,attr"`
	DeadlineTimeLoc            string               `xml:"DeadlineTimeLoc,attr"`
	DeadlineTimeSys            string               `xml:"DeadlineTimeSys,attr"`
	DeadlineTimeUTC            string               `xml:"DeadlineTimeUTC,attr"`
	PenaltySize                float64              `xml:"PenaltySize,attr"`
	DefaultCheckInTime         string               `xml:"DefaultCheckInTime,attr"`
	DefaultCheckOutTime        string               `xml:"DefaultCheckOutTime,attr"`
	CheckInTime                string               `xml:"CheckInTime,attr"`
	CheckOutTime               string               `xml:"CheckOutTime,attr"`
	MaxNumberOfGuests          int                  `xml:"MaxNumberOfGuests,attr"`
	MaxNumberOfExtraBeds       int                  `xml:"MaxNumberOfExtraBeds,attr"`
	MaxNumberOfExtraBedsAdult  int                  `xml:"MaxNumberOfExtraBedsAdult,attr"`
	MaxNumberOfExtraBedsChild  int                  `xml:"MaxNumberOfExtraBedsChild,attr"`
	MaxNumberOfExtraBedsInfant int                  `xml:"MaxNumberOfExtraBedsInfant,attr"`
	SpecialOfferList           SpecialOfferListType `xml:"SpecialOfferList"`
	AllowEarlierCheckIn        YesNoCodeType        `xml:"AllowEarlierCheckIn"`
	AllowLateCheckOut          YesNoCodeType        `xml:"AllowLateCheckOut"`
	AllowToAmendBookings       YesNoCodeType        `xml:"AllowToAmendBookings"`
	AllowToAmendGuestNames     YesNoCodeType        `xml:"AllowToAmendGuestNames"`
	AllGuestNamesRequired      YesNoCodeType        `xml:"AllGuestNamesRequired"`
	GuestCitizenshipRequired   YesNoCodeType        `xml:"GuestCitizenshipRequired"`
	Meals                      MealsType            `xml:"Meals"`
	PenaltyRuleList            PenaltyRuleListType  `xml:"PenaltyRuleList"`
	WhereToPayList             WhereToPayListType   `xml:"WhereToPayList"`
	PriceDescription           PriceDescriptionType `xml:"PriceDescription"`
}

type PriceDescriptionType struct {
	XMLName             xml.Name                `xml:"PriceDescription"`
	EarlierCheckInPrice EarlierCheckInPriceType `xml:"EarlierCheckInPrice"`
	LateCheckOutPrice   LateCheckOutPriceType   `xml:"LateCheckOutPrice"`
	Days                DaysType                `xml:"Days"`
}

type DaysType struct {
	XMLName xml.Name  `xml:"Days"`
	Items   []DayType `xml:"Day"`
}

type DayType struct {
	XMLName                xml.Name `xml:"Day"`
	Date                   string   `xml:"Date,attr"`
	MinimumRetailPrice     float64  `xml:"MinimumRetailPrice,attr"`
	Price                  float64  `xml:"Price,attr"`
	TravelAgencyCommission float64  `xml:"TravelAgencyCommission,attr"`
	AdditionalMeal         float64  `xml:"AdditionalMeal,attr"`
	ExtraBedsAdultPrice    float64  `xml:"ExtraBedsAdultPrice,attr"`
	ExtraBedsChildPrice    float64  `xml:"ExtraBedsChildPrice,attr"`
	ExtraBedsInfantPrice   float64  `xml:"ExtraBedsInfantPrice,attr"`
}

type LateCheckOutPriceType struct {
	XMLName                xml.Name `xml:"LateCheckOutPrice"`
	Price                  float64  `xml:"Price,attr"`
	TravelAgencyCommission float64  `xml:"TravelAgencyCommission,attr"`
}

type EarlierCheckInPriceType struct {
	XMLName                xml.Name `xml:"EarlierCheckInPrice"`
	Price                  float64  `xml:"Price,attr"`
	TravelAgencyCommission float64  `xml:"TravelAgencyCommission,attr"`
}

type WhereToPayListType struct {
	XMLName xml.Name             `xml:"WhereToPayList"`
	Items   []SimpleCodeNameType `xml:"WhereToPay"`
}

type PenaltyRuleListType struct {
	XMLName xml.Name          `xml:"PenaltyRuleList"`
	Items   []PenaltyRuleType `xml:"PenaltyRule"`
}

type PenaltyRuleType struct {
	XMLName             xml.Name `xml:"PenaltyRule"`
	HoursFrom           int      `xml:"HoursFrom,attr"`
	HoursTo             int      `xml:"HoursTo,attr"`
	Value               float64  `xml:"Value,attr"`
	CalculationRuleCode int      `xml:"CalculationRuleCode,attr"`
	CalculationRuleName string   `xml:"CalculationRuleName,attr"`
}

type MealsType struct {
	XMLName xml.Name   `xml:"Meals"`
	Items   []MealType `xml:"Meal"`
}

type MealType struct {
	SimpleCodeNameType
	XMLName            xml.Name      `xml:"Meal"`
	Price              float64       `xml:"Price,attr"`
	MinimumRetailPrice float64       `xml:"MinimumRetailPrice,attr"`
	IncludedInPrice    YesNoCodeType `xml:"IncludedInPrice"`
}

type SpecialOfferListType struct {
	XMLName xml.Name           `xml:"SpecialOfferList"`
	Items   []SpecialOfferType `xml:"SpecialOffer"`
}

type SpecialOfferType struct {
	XMLName          xml.Name             `xml:"SpecialOffer"`
	StayNights       int                  `xml:"StayNights,attr"`
	PayNights        int                  `xml:"PayNights,attr"`
	SpecialOfferType SpecialOfferTypeType `xml:"SpecialOfferType"`
	IsMultiple       YesNoCodeType        `xml:"IsMultiple"`
}

type SpecialOfferTypeType struct {
	SimpleCodeNameType
	XMLName xml.Name `xml:"SpecialOfferType"`
	Id      int      `xml:"Id,attr"`
}
