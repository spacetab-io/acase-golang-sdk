package acaseSts

import "encoding/xml"

type ObjectTypeRequestType struct {
	Credentials
	XMLName		xml.Name				`xml:"ObjectTypeRequest"`
	Action		ObjectTypeActionType	`xml:"Action"`
}

type ObjectTypeActionType struct {
	XMLName		xml.Name			`xml:"Action"`
	Name		string				`xml:"Name,attr"`
	Parameters	ParametersObjType	`xml:"Parameters"`
}

type ParametersObjType struct {
	XMLName			xml.Name	`xml:"Parameters"`
	ObjectTypeCode	string		`xml:"ObjectTypeCode,attr,omitempty"`
}

type GeoObjectTypeResponseType struct {
	Credentials
	BaseResponse
	XMLName			xml.Name			`xml:"GeoObjectTypeResponse"`
	Action			ObjectActionType	`xml:"Action"`
	ObjectTypeList	ObjectTypeListType	`xml:"ObjectTypeList"`
}

type ObjectTypeListType struct {
	XMLName	xml.Name				`xml:"ObjectTypeList"`
	Items	[]SimpleCodeNameType	`xml:"ObjectType"`
}
