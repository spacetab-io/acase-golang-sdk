package acaseSts

import "encoding/xml"

type ObjectSubTypeRequestType struct {
	XMLName		xml.Name				`xml:"ObjectRequest"`
	BuyerId		string					`xml:"BuyerId,attr"`
	UserId		string					`xml:"UserId,attr"`
	Password	string					`xml:"Password,attr"`
	Language	LanguageTypeEnum		`xml:"Language,attr,omitempty"`
	Action		[]ObjSubTypeActionType	`xml:"Action"`
}

type ObjSubTypeActionType struct {
	XMLName		xml.Name			`xml:"Action"`
	Name		string				`xml:"Name,attr"`
	Parameters	ObjSubTypeParamType	`xml:"Parameters"`
}

type ObjSubTypeParamType struct {
	XMLName			xml.Name	`xml:"Parameters"`
	ObjectTypeCode	int			`xml:"ObjectTypeCode,attr"`
}

type ObjectSubTypeResponseType struct {
	XMLName		xml.Name				`xml:"ObjectSubTypeResponse"`
	BuyerId		string					`xml:"BuyerId,attr"`
	UserId		string					`xml:"UserId,attr"`
	Password	string					`xml:"Password,attr"`
	Language	string					`xml:"Language,attr,omitempty"`
	Success		string					`xml:"Success"`
	Error		ErrorType				`xml:"Error,omitempty"`
	Action		[]ObjSubTypeActionType	`xml:"Action"`
	ObjectType	[]ObjectTypeSubType		`xml:"ObjectType"`
}

type ObjectTypeSubType struct {
	XMLName			xml.Name			`xml:"ObjectType"`
	Code			int					`xml:"Code,attr"`
	Name			string				`xml:"Name,attr"`
	ObjectSubType	[]ObjectSubTypeType	`xml:"ObjectSubType"`
}

