package acaseSts

import "encoding/xml"

type HotelProductRequestType struct {
	Credentials
	XMLName					xml.Name			`xml:"HotelProductRequest"`
	Hotel					int					`xml:"Hotel,attr"`
	Currency				int					`xml:"Currency,attr"`
	WhereToPay				int					`xml:"WhereToPay,attr"`
	ArrivalDate				string				`xml:"ArrivalDate,attr"`
	DepartureDate			string				`xml:"DepartureDate,attr"`
	Id						string				`xml:"Id,attr"`
	AccommodationId			string				`xml:"AccommodationId,attr"`
	NumberOfGuests			int					`xml:"NumberOfGuests,attr"`
	NumberOfExtraBedsAdult	int					`xml:"NumberOfExtraBedsAdult,attr,omitempty"`
	NumberOfExtraBedsChild	int					`xml:"NumberOfExtraBedsChild,attr,omitempty"`
	NumberOfExtraBedsInfant	int					`xml:"NumberOfExtraBedsInfant,attr,omitempty"`
}

type HotelProductResponseType struct {
	Credentials
	BaseResponse
	SimpleCodeNameType
	XMLName					xml.Name					`xml:"HotelProductResponse"`
	StartDate				string						`xml:"StartDate,attr"`
	EndDate					string						`xml:"EndDate,attr"`
	NumberOfGuests			int							`xml:"NumberOfGuests,attr"`
	NumberOfExtraBedsAdult	int							`xml:"NumberOfExtraBedsAdult,attr"`
	NumberOfExtraBedsChild	int							`xml:"NumberOfExtraBedsChild,attr,omitempty"`
	NumberOfExtraBedsInfant	int							`xml:"NumberOfExtraBedsInfant,attr,omitempty"`
	SpecialRequirements		string						`xml:"SpecialRequirements,attr"`
	Country					CountryType					`xml:"Country"`
	AdmUnit1				SimpleCodeNameType			`xml:"AdmUnit1"`
	AdmUnit2				AdmUnit2Type				`xml:"AdmUnit2"`
	TypeOfPlace				SimpleCodeNameType			`xml:"TypeOfPlace"`
	City					CityType					`xml:"City"`
	Position				PositionType				`xml:"Position"`
	ObjType					ObjTypeType					`xml:"ObjType"`
	Currency				SimpleCodeNameType			`xml:"Currency"`
	Infants					AgeRestrictionType			`xml:"Infants"`
	Children				AgeRestrictionType			`xml:"Children"`
	SpecialRequirementList	SpecialRequirementListType	`xml:"SpecialRequirementList"`
	HotelProductList		HotelProductListType		`xml:"HotelProductList"`
}

type HotelProductListType struct {
	XMLName	xml.Name					`xml:"HotelProductList"`
	Items	[]HotelProductDetailType	`xml:"HotelProductDetail"`
}

type HotelProductDetailType struct {
	XMLName						xml.Name				`xml:"HotelProductDetail"`
	Code 						string					`xml:"Code,attr"`
	RoomName 					string					`xml:"RoomName,attr"`
	RoomCode 					string					`xml:"RoomCode,attr"`
	RateCode 					string					`xml:"RateCode,attr"`
	RateName 					string					`xml:"RateName,attr"`
	RateDescription 			string					`xml:"RateDescription,attr"`
	IsHourRate 					string					`xml:"IsHourRate,attr"`
	MaxHours 					int						`xml:"MaxHours,attr"`
	StartTime 					string					`xml:"StartTime,attr"`
	EndTime 					string					`xml:"EndTime,attr"`
	RateGroupCode 				string					`xml:"RateGroupCode,attr"`
	RateGroupName 				string					`xml:"RateGroupName,attr"`
	MinimumRetailPrice			float64					`xml:"MinimumRetailPrice,attr"`
	Price						float64					`xml:"Price,attr"`
	TravelAgencyCommission		float64					`xml:"TravelAgencyCommission,attr"`
	FullVATInPrice 				string					`xml:"FullVATInPrice,attr"`
	VATIncludedInPrice 			float64					`xml:"VATIncludedInPrice,attr"`
	DeadlineDate 				string					`xml:"DeadlineDate,attr"`
	DeadlineTimeLoc 			string					`xml:"DeadlineTimeLoc,attr"`
	DeadlineTimeSys 			string					`xml:"DeadlineTimeSys,attr"`
	DeadlineTimeUTC 			string					`xml:"DeadlineTimeUTC,attr"`
	PenaltySize 				float64					`xml:"PenaltySize,attr"`
	MaxNumberOfGuests 			int						`xml:"MaxNumberOfGuests,attr"`
	MaxNumberOfExtraBeds		int						`xml:"MaxNumberOfExtraBeds,attr"`
	MaxNumberOfExtraBedsAdult	int						`xml:"MaxNumberOfExtraBedsAdult,attr"`
	MaxNumberOfExtraBedsChild	int						`xml:"MaxNumberOfExtraBedsChild,attr"`
	MaxNumberOfExtraBedsInfant	int						`xml:"MaxNumberOfExtraBedsInfant,attr"`
	SpecialOfferList			SpecialOfferListType	`xml:"SpecialOfferList"`
	Meals						MealsType				`xml:"Meals"`
	Availability				AvailabilityType		`xml:"Availability"`
	WhereToPayList				WhereToPayListType		`xml:"WhereToPayList"`
}

type AvailabilityType struct {
	SimpleCodeNameType
	XMLName		xml.Name		`xml:"Availability"`
	Allotment	AllotmentType	`xml:"Allotment"`
}

type AllotmentType struct {
	XMLName		xml.Name	`xml:"Allotment"`
	Code		string		`xml:"Code,attr"`
	Name		string		`xml:"Name,attr"`
	Rooms		int			`xml:"Rooms,attr,omitempty"`
	Quantity	string		`xml:"Quantity,attr,omitempty"`
}
