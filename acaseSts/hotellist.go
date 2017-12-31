package acaseSts

import "encoding/xml"

type HotelListRequestType struct {
	XMLName			xml.Name			`xml:"HotelList"`
	BuyerId			string				`xml:"BuyerId,attr"`
	UserId			string				`xml:"UserId,attr"`
	Password		string				`xml:"Password,attr"`
	Language		LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	HotelCode		int					`xml:"HotelCode,attr,omitempty"`
	HotelName		string				`xml:"HotelName,attr,omitempty"`
	CountryCode		int					`xml:"CountryCode,attr,omitempty"`
	CityCode		int					`xml:"CityCode,attr,omitempty"`
	HotelRatingCode	int					`xml:"HotelRatingCode,attr,omitempty"`
	Options			string				`xml:"Opt,attr,omitempty"`
}

type HotelListResponseType struct {
	XMLName				xml.Name					`xml:"HotelList"`
	Success				string						`xml:"Success"`
	Error				ErrorType					`xml:"Error,omitempty"`
	Country				[]CountryType				`xml:"Country"`
	AlternativeCountry	[]AlternativeCountryType	`xml:"AlternativeCountry"`
}
