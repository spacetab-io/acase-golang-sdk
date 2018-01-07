package acaseSts

import "encoding/xml"

type ObjTypeListRequestType struct {
	XMLName		xml.Name				`xml:"ObjTypeListRequest"`
	BuyerId		string					`xml:"BuyerId,attr"`
	UserId		string					`xml:"UserId,attr"`
	Password	string					`xml:"Password,attr"`
	Language	LanguageTypeEnum		`xml:"Language,attr,omitempty"`
	Code		string					`xml:"Code,attr"`
	Name		string					`xml:"Name,attr"`
}

type ObjTypeListResponseType struct {
	XMLName			xml.Name			`xml:"ObjTypeListResponse"`
	BuyerId			string				`xml:"BuyerId,attr"`
	UserId			string				`xml:"UserId,attr"`
	Password		string				`xml:"Password,attr"`
	Language		string				`xml:"Language,attr,omitempty"`
	Success			string				`xml:"Success"`
	Error			ErrorType			`xml:"Error,omitempty"`
	ObjType			[]ObjTypeType		`xml:"ObjType"`
}
