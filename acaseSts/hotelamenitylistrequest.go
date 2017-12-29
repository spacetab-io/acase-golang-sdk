package acaseSts

import "encoding/xml"

type HotelAmenityListRequestType struct {
	XMLName				xml.Name			`xml:"HotelAmenityList"`
	BuyerId				string				`xml:"BuyerId,attr"`
	UserId				string				`xml:"UserId,attr"`
	Password			string				`xml:"Password,attr"`
	Language			LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	HotelAmenityCode	int					`xml:"HotelAmenityCode,attr,omitempty"`
	HotelAmenityName	string				`xml:"HotelAmenityName,attr,omitempty"`
	Options				string				`xml:"Options,attr"`
}

type HotelAmenityListResponseType struct {
	XMLName			xml.Name		`xml:"HotelAmenityList"`
	Success			string			`xml:"Success"`
	Error			ErrorType		`xml:"Error,omitempty"`
	HotelAmenity	[]HotelAmenity	`xml:"HotelAmenity"`
}

type HotelAmenity struct {
	XMLName	xml.Name	`xml:"HotelAmenity"`
	Code	int			`xml:"Code,attr"`
	Name	string		`xml:"Name,attr"`
	Url		string		`xml:"Url,attr,omitempty"`
}

