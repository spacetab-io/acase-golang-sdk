package acaseSts

import "encoding/xml"

type ObjectRequestType struct {
	XMLName		xml.Name			`xml:"ObjectRequest"`
	BuyerId		string				`xml:"BuyerId,attr"`
	UserId		string				`xml:"UserId,attr"`
	Password	string				`xml:"Password,attr"`
	Language	LanguageTypeEnum	`xml:"Language,attr,omitempty"`
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
	XMLName		xml.Name				`xml:"ObjectResponse"`
	BuyerId		string					`xml:"BuyerId,attr"`
	UserId		string					`xml:"UserId,attr"`
	Password	string					`xml:"Password,attr"`
	Language	LanguageTypeEnum		`xml:"Language,attr,omitempty"`
	Success		string					`xml:"Success"`
	Error		ErrorType				`xml:"Error,omitempty"`
	Action		[]ObjectActionType		`xml:"Action"`
	Object		[]ObjRespType			`xml:"Object"`
}

type ObjRespType struct {
	XMLName			xml.Name			`xml:"Object"`
	Code			int					`xml:"Code,attr"`
	Name			string				`xml:"Name,attr"`
	ObjectSubType	SimpleCodeNameType	`xml:"ObjectSubType"`
}