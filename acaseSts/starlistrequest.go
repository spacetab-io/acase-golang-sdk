package acaseSts

import "encoding/xml"

type StarListRequestType struct {
	XMLName		xml.Name			`xml:"StarListRequest"`
	BuyerId		string				`xml:"BuyerId,attr"`
	UserId		string				`xml:"UserId,attr"`
	Password	string				`xml:"Password,attr"`
	Language	LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	Code		int					`xml:"Code,attr"`
	Name		string				`xml:"Name,attr"`
	Options		string				`xml:"Options,attr"`
}

type StarListResponseType struct {
	XMLName		xml.Name	`xml:"StarListResponse"`
	BuyerId		string		`xml:"BuyerId,attr"`
	UserId		string		`xml:"UserId,attr"`
	Password	string		`xml:"Password,attr"`
	Language	string		`xml:"Language,attr,omitempty"`
	Success		SuccessType	`xml:"Success"`
	Error		ErrorType	`xml:"Error,omitempty"`
	Star		[]StarType	`xml:"Star"`
}

type StarType struct {
	XMLName	xml.Name	`xml:"Star"`
	Code	int			`xml:"Code,attr"`
	Name	string		`xml:"Name,attr"`
	Value	string		`xml:"Value,attr"`
}
