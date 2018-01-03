package acaseSts

import "encoding/xml"

type ObjectTypeRequestType struct {
	XMLName		xml.Name				`xml:"ObjectTypeRequest"`
	BuyerId		string					`xml:"BuyerId,attr"`
	UserId		string					`xml:"UserId,attr"`
	Password	string					`xml:"Password,attr"`
	Language	LanguageTypeEnum		`xml:"Language,attr,omitempty"`
	Action		[]ObjectTypeActionType	`xml:"Action"`
}

type ObjectTypeActionType struct {
	XMLName		xml.Name			`xml:"Action"`
	Name		string				`xml:"Name,attr"`
	Parameters	ParametersObjType	`xml:"Parameters"`
}

type ParametersObjType struct {
	XMLName			xml.Name	`xml:"Parameters"`
	ObjectTypeCode	int			`xml:"ObjectTypeCode,attr"`
}

type ObjectTypeResponseType struct {
	XMLName			xml.Name			`xml:"ObjectTypeResponse"`
	BuyerId			string				`xml:"BuyerId,attr"`
	UserId			string				`xml:"UserId,attr"`
	Password		string				`xml:"Password,attr"`
	Language		string				`xml:"Language,attr,omitempty"`
	Success			string				`xml:"Success"`
	Error			ErrorType			`xml:"Error,omitempty"`
	Action			[]ObjectActionType	`xml:"Action"`
	ObjectTypeList	ObjectTypeListType	`xml:"ObjectTypeList"`
}

type ObjectTypeListType struct {
	XMLName	xml.Name				`xml:"ObjectTypeList"`
	Items	[]SimpleCodeNameType	`xml:"ObjectType"`
}
