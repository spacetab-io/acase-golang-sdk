package acaseSts

import "encoding/xml"

type StarListRequestType struct {
	Credentials
	SimpleCodeNameType
	XMLName		xml.Name			`xml:"StarListRequest"`
	Options		string				`xml:"Options,attr"`
}

type StarListResponseType struct {
	Credentials
	BaseResponse
	XMLName		xml.Name	`xml:"StarListResponse"`
	Star		[]StarType	`xml:"Star"`
}

type StarType struct {
	SimpleCodeNameType
	XMLName	xml.Name	`xml:"Star"`
	Value	string		`xml:"Value,attr"`
}
