package acaseSts

import "encoding/xml"

type HotelSearchRequestType struct {
	Credentials
	XMLName         xml.Name             `xml:"HotelSearchRequest"`
	ArrivalDate     string               `xml:"ArrivalDate,attr"`
	DepartureDate   string               `xml:"DepartureDate,attr"`
	City            int                  `xml:"City,attr,omitempty"`
	PriceFrom       string               `xml:"PriceFrom,attr"`
	PriceTo         string               `xml:"PriceTo,attr"`
	FreeSaleOnly    int                  `xml:"FreeSaleOnly,attr"`
	HotelCategory   int                  `xml:"HotelCategory,attr"`
	Currency        int                  `xml:"Currency,attr"`
	WhereToPay      int                  `xml:"WhereToPay,attr"`
	NumberOfGuests  int                  `xml:"NumberOfGuests,attr"`
	Options         string               `xml:"Options,attr"`
	HotelCode       int                  `xml:"HotelCode,attr,omitempty"`
	HotelName       string               `xml:"HotelName,attr,omitempty"`
	StarList        StarListType         `xml:"StarList"`
	AmenityList     AmenityListType      `xml:"AmenityList"`
	ObjectList      ObjectListType       `xml:"ObjectList"`
	ObjTypeList     *ObjTypeListType     `xml:"ObjTypeList,omitempty"`
	Guests          *GuestsType          `xml:"Guests,omitempty"`
	Destination     *DestinationType     `xml:"Destination,omitempty"`
	DestinationList *DestinationListType `xml:"DestinationList,omitempty"`
}

type DestinationListType struct {
	XMLName  xml.Name            `xml:"DestinationList"`
	Distance int                 `xml:"Distance,attr"`
	TypeCode DestinationTypeEnum `xml:"TypeCode,attr"`
	Code     string              `xml:"Code,attr"`
}

type DestinationType struct {
	XMLName  xml.Name            `xml:"Destination"`
	Distance int                 `xml:"Distance,attr"`
	TypeCode DestinationTypeEnum `xml:"TypeCode,attr"`
	Code     int                 `xml:"Code,attr"`
}

type GuestsType struct {
	XMLName        xml.Name         `xml:"Guests"`
	NumberOfAdults int              `xml:"NumberOfAdults,attr"`
	NumberOfMinors int              `xml:"NumberOfMinors,attr"`
	MinorAgeList   MinorAgeListType `xml:"MinorAgeList"`
}

type MinorAgeListType struct {
	XMLName xml.Name    `xml:"MinorAgeList"`
	Items   []MinorType `xml:"Minor"`
}

type MinorType struct {
	XMLName xml.Name `xml:"Minor"`
	Age     int      `xml:"Age,attr"`
}

type ObjTypeListType struct {
	XMLName xml.Name          `xml:"ObjTypeList"`
	Items   *[]SimpleCodeType `xml:"ObjType,omitempty"`
}

type ObjectListType struct {
	XMLName xml.Name         `xml:"ObjectList"`
	Items   []SimpleCodeType `xml:"Object"`
}

type AmenityListType struct {
	XMLName xml.Name         `xml:"AmenityList"`
	Items   []SimpleCodeType `xml:"Amenity"`
}

type StarListType struct {
	XMLName xml.Name         `xml:"StarList"`
	Items   []SimpleCodeType `xml:"Star"`
}

type HotelSearchResponseType struct {
	BaseResponse
	XMLName                         xml.Name                `xml:"HotelSearchResponse"`
	Language                        string                  `xml:"Language,attr"`
	ArrivalDate                     string                  `xml:"ArrivalDate,attr"`
	DepartureDate                   string                  `xml:"DepartureDate,attr"`
	NumberOfGuests                  int                     `xml:"NumberOfGuests,attr"`
	DestinationList                 DestListFullType        `xml:"DestinationList"`
	AvailableHotelsList             AvailableHotelsListType `xml:"AvailableHotelsList"`
	NoContractAvailableHotelsList   HotelOfferListType      `xml:"NoContractAvailableHotelsList"`
	AlternativeHotelsList           HotelOfferListType      `xml:"AlternativeHotelsList"`
	NoContractAlternativeHotelsList HotelOfferListType      `xml:"NoContractAlternativeHotelsList"`
}

type AvailableHotelsListType struct {
	XMLName xml.Name        `xml:"AvailableHotelsList"`
	Items   []HotelFullType `xml:"Hotel"`
}

type HotelFullType struct {
	SimpleCodeNameType
	XMLName               xml.Name               `xml:"Hotel"`
	Zip                   string                 `xml:"Zip,attr"`
	Address               string                 `xml:"Address,attr"`
	CityCentre            string                 `xml:"CityCentre,attr"`
	Airport               string                 `xml:"Airport,attr"`
	RailwayStation        string                 `xml:"RailwayStation,attr"`
	Underground           string                 `xml:"Underground,attr"`
	Description           string                 `xml:"Description,attr"`
	Url                   string                 `xml:"Url,attr"`
	Ref                   string                 `xml:"Ref,attr"`
	DistanceToDestination int                    `xml:"DistanceToDestination,attr"`
	DefaultCheckOutTime   string                 `xml:"DefaultCheckOutTime,attr"`
	DefaultCheckInTime    string                 `xml:"DefaultCheckInTime,attr"`
	ObjType               ObjTypeType            `xml:"ObjType"`
	City                  CityType               `xml:"City"`
	AdmUnit1              SimpleCodeNameType     `xml:"AdmUnit1"`
	AdmUnit2              AdmUnit2Type           `xml:"AdmUnit2"`
	TypeOfPlace           SimpleCodeNameType     `xml:"TypeOfPlace"`
	Country               CountryType            `xml:"Country"`
	Position              PositionType           `xml:"Position"`
	Rating                SimpleCodeNameType     `xml:"Rating"`
	Stars                 StarsType              `xml:"Stars"`
	Amenities             HotelAmenitiesListType `xml:"Amenities"`
	Currency              SimpleCodeNameType     `xml:"Currency"`
	HotelOfferList        HotelOfferListType     `xml:"HotelOfferList"`
}

type HotelOfferListType struct {
	Items []HotelOfferDetailType `xml:"HotelOfferDetail"`
}

type HotelOfferDetailType struct {
	XMLName                xml.Name             `xml:"HotelOfferDetail"`
	Code                   int                  `xml:"Code,attr"`
	NumberOfGuests         int                  `xml:"NumberOfGuests,attr"`
	NumberOfExtraBedAdult  int                  `xml:"NumberOfExtraBedAdult,attr"`
	NumberOfExtraBedChild  int                  `xml:"NumberOfExtraBedChild,attr"`
	NumberOfExtraBedInfant int                  `xml:"NumberOfExtraBedInfant,attr"`
	IsHourRate             int                  `xml:"IsHourRate,attr"`
	MaxHours               int                  `xml:"MaxHours,attr"`
	TotalPrice             float64              `xml:"TotalPrice,attr"`
	VATIncludedInPrice     float64              `xml:"VATIncludedInPrice,attr"`
	FullVATInPrice         int                  `xml:"FullVATInPrice,attr"`
	TravelAgencyCommission float64              `xml:"TravelAgencyCommission,attr"`
	RoomCode               int                  `xml:"RoomCode,attr"`
	RoomName               string               `xml:"RoomName,attr"`
	RateCode               int                  `xml:"RateCode,attr"`
	RateName               string               `xml:"RateName,attr"`
	RateDescription        string               `xml:"RateDescription,attr"`
	RateGroupCode          int                  `xml:"RateGroupCode,attr"`
	RateGroupName          string               `xml:"RateGroupName,attr"`
	DeadlineDate           string               `xml:"DeadlineDate,attr"`
	DeadlineTimeLoc        string               `xml:"DeadlineTimeLoc,attr"`
	DeadlineTimeSys        string               `xml:"DeadlineTimeSys,attr"`
	DeadlineTimeUTC        string               `xml:"DeadlineTimeUTC,attr"`
	PenaltySize            float64              `xml:"PenaltySize,attr"`
	SpecialOfferList       SpecialOfferListType `xml:"SpecialOfferList"`
	MealIncludedInPrice    SimpleCodeNameType   `xml:"MealIncludedInPrice"`
	Availability           AvailabilityType     `xml:"Availability"`
	WhereToPay             SimpleCodeNameType   `xml:"WhereToPay"`
}

type DestListFullType struct {
	XMLName xml.Name       `xml:"DestinationList"`
	Items   []DestFullType `xml:"Destination"`
}

type DestFullType struct {
	SimpleCodeNameType
	XMLName         xml.Name           `xml:"Destination"`
	DestinationType SimpleCodeNameType `xml:"DestinationType"`
	ObjType         SimpleCodeNameType `xml:"ObjType"`
	HotelType       SimpleCodeNameType `xml:"HotelType"`
	Position        PositionType       `xml:"Position"`
}
