package acaseSts

import "encoding/xml"

type CountryListRequestType struct {
	XMLName		xml.Name			`xml:"CountryListRequest"`
	BuyerId		string				`xml:"BuyerId,attr"`
	UserId		string				`xml:"UserId,attr"`
	Password	string				`xml:"Password,attr"`
	Language	LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	CountryCode	int					`xml:"CountryCode,attr,omitempty"`
	CountryName	string				`xml:"CountryName,attr,omitempty"`
}

type CountryListType struct {
	XMLName		xml.Name		`xml:"CountryList"`
	Success		string			`xml:"Success"`
	Error		ErrorType		`xml:"Error,omitempty"`
	Countries	[]CountryType	`xml:"Country"`
}
