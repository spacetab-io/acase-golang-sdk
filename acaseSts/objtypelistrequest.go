package acaseSts

import "encoding/xml"

type ObjTypeListRequestType struct {
	Credentials
	XMLName xml.Name `xml:"ObjTypeListRequest"`
	Code    string   `xml:"Code,attr"`
	Name    string   `xml:"Name,attr"`
}

type ObjTypeListResponseType struct {
	Credentials
	BaseResponse
	XMLName xml.Name      `xml:"ObjTypeListResponse"`
	ObjType []ObjTypeType `xml:"ObjType"`
}
