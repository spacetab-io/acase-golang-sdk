package acaseSts

import "encoding/xml"

type StatusListRequestType struct {
	XMLName		xml.Name			`xml:"StatusListRequest"`
	BuyerId		string				`xml:"BuyerId,attr"`
	UserId		string				`xml:"UserId,attr"`
	Password	string				`xml:"Password,attr"`
	Language	LanguageTypeEnum	`xml:"Language,attr,omitempty"`
}

type StatusListResponseType struct {
	XMLName		xml.Name				`xml:"StatusListResponse"`
	Success		SuccessType				`xml:"Success"`
	Error		ErrorType				`xml:"Error,omitempty"`
	Status		[]SimpleCodeNameType	`xml:"Status"`
}
