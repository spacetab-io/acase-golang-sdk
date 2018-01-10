package acaseSts

import "encoding/xml"

type ObjectSubTypeRequestType struct {
	Credentials
	XMLName		xml.Name				`xml:"ObjectRequest"`
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
	Credentials
	BaseResponse
	XMLName		xml.Name				`xml:"ObjectSubTypeResponse"`
	Action		[]ObjSubTypeActionType	`xml:"Action"`
	ObjectType	[]ObjectTypeSubType		`xml:"ObjectType"`
}

type ObjectTypeSubType struct {
	SimpleCodeNameType
	XMLName			xml.Name				`xml:"ObjectType"`
	ObjectSubType	[]SimpleCodeNameType	`xml:"ObjectSubType"`
}

