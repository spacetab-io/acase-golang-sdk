package acaseSts

import "encoding/xml"

type ObjectRequestType struct {
	Credentials
	XMLName		xml.Name			`xml:"ObjectRequest"`
	Action		[]ObjectActionType	`xml:"Action"`
}

type ObjectActionType struct {
	XMLName		xml.Name				`xml:"Action"`
	Name		string					`xml:"Name,attr"`
	Parameters	ObjectParametersType	`xml:"Parameters"`
}

type ObjectParametersType struct {
	XMLName				xml.Name	`xml:"Parameters"`
	ObjectTypeCode		int			`xml:"ObjectTypeCode,attr"`
	ObjectSubTypeCode	int			`xml:"ObjectSubTypeCode,attr"`
	CityCode			int			`xml:"CityCode,attr"`
}

type ObjectResponseType struct {
	Credentials
	BaseResponse
	XMLName		xml.Name				`xml:"ObjectResponse"`
	Action		[]ObjectActionType		`xml:"Action"`
	Object		[]ObjRespType			`xml:"Object"`
}

type ObjRespType struct {
	SimpleCodeNameType
	XMLName			xml.Name			`xml:"Object"`
	ObjectSubType	SimpleCodeNameType	`xml:"ObjectSubType"`
}