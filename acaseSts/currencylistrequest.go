package acaseSts

import "encoding/xml"

type CurrencyListRequestType struct {
	XMLName		xml.Name			`xml:"CurrencyListRequest"`
	BuyerId		string				`xml:"BuyerId,attr"`
	UserId		string				`xml:"UserId,attr"`
	Password	string				`xml:"Password,attr"`
	Language	LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	Code		int					`xml:"Code,attr,omitempty"`
	Name		string				`xml:"Name,attr,omitempty"`
	Options		string				`xml:"Options,attr"`
}

type CurrencyListResponseType struct {
	XMLName		xml.Name		`xml:"CurrencyListResponse"`
	Success		string			`xml:"Success"`
	Error		ErrorType		`xml:"Error,omitempty"`
	Currency	[]CurrencyType	`xml:"Currency"`
}

type CurrencyType struct {
	XMLName	xml.Name	`xml:"Currency"`
	Code	int			`xml:"Code,attr"`
	Name	string		`xml:"Name,attr,omitempty"`
}
