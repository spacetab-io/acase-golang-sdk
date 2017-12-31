package acaseSts

import "encoding/xml"

type HotelProductRequestType struct {
	XMLName					xml.Name			`xml:"HotelProductRequest"`
	BuyerId					string				`xml:"BuyerId,attr"`
	UserId					string				`xml:"UserId,attr"`
	Password				string				`xml:"Password,attr"`
	Language				LanguageTypeEnum	`xml:"Language,attr,omitempty"`
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
	XMLName				xml.Name					`xml:"HotelList"`
	Success				string						`xml:"Success"`
	Error				ErrorType					`xml:"Error,omitempty"`
	Country				[]CountryType				`xml:"Country"`
	AlternativeCountry	[]AlternativeCountryType	`xml:"AlternativeCountry"`
}

