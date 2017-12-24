package acaseSts

import "encoding/xml"

type CityListRequestType struct {
	XMLName		xml.Name			`xml:"CityListRequest"`
	BuyerId		string				`xml:"BuyerId,attr"`
	UserId		string				`xml:"UserId,attr"`
	Password	string				`xml:"Password,attr"`
	Language	LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	CountryCode	string				`xml:"CountryCode,attr,omitempty"`
	CountryName	string				`xml:"CountryName,attr,omitempty"`
	CityCode	string				`xml:"CityCode,attr,omitempty"`
	CityName	string				`xml:"CityName,attr,omitempty"`
}

type CityListType struct {
	XMLName		xml.Name		`xml:"CityList"`
	Success		string			`xml:"Success,omitempty"`
	Error		ErrorType		`xml:"Error,omitempty"`
	Countries	[]CountryType	`xml:"Country"`
}

