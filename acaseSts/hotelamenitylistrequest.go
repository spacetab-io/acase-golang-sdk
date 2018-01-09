package acaseSts

import "encoding/xml"

type HotelAmenityListRequestType struct {
	Credentials
	XMLName				xml.Name			`xml:"HotelAmenityList"`
	HotelAmenityCode	int					`xml:"HotelAmenityCode,attr,omitempty"`
	HotelAmenityName	string				`xml:"HotelAmenityName,attr,omitempty"`
	Options				string				`xml:"Opt,attr,omitempty"`
}

type HotelAmenityListResponseType struct {
	BaseResponse
	XMLName			xml.Name		`xml:"HotelAmenityList"`
	HotelAmenity	[]HotelAmenity	`xml:"HotelAmenity"`
}

type HotelAmenity struct {
	SimpleCodeNameType
	XMLName	xml.Name	`xml:"HotelAmenity"`
	Url		string		`xml:"Url,attr,omitempty"`
}

