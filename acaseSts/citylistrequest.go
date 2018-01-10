package acaseSts

import "encoding/xml"

type CityListRequestType struct {
	Credentials
	XMLName		xml.Name			`xml:"CityListRequest"`
	CountryCode	string				`xml:"CountryCode,attr,omitempty"`
	CountryName	string				`xml:"CountryName,attr,omitempty"`
	CityCode	int					`xml:"CityCode,attr,omitempty"`
	CityName	string				`xml:"CityName,attr,omitempty"`
}

type CityListType struct {
	BaseResponse
	XMLName		xml.Name		`xml:"CityList"`
	Countries	[]CountryType	`xml:"Country"`
}

