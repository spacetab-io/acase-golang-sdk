package acaseSts

import "encoding/xml"

type HotelListRequestType struct {
	Credentials
	XMLName			xml.Name			`xml:"HotelListRequest"`
	HotelCode		int					`xml:"HotelCode,attr,omitempty"`
	HotelName		string				`xml:"HotelName,attr,omitempty"`
	CountryCode		int					`xml:"CountryCode,attr,omitempty"`
	CityCode		int					`xml:"CityCode,attr,omitempty"`
	HotelRatingCode	int					`xml:"HotelRatingCode,attr,omitempty"`
	Options			string				`xml:"Opt,attr,omitempty"`
}

type HotelListResponseType struct {
	BaseResponse
	XMLName				xml.Name					`xml:"HotelList"`
	Country				[]CountryType				`xml:"Country"`
	AlternativeCountry	[]AlternativeCountryType	`xml:"AlternativeCountry"`
}
