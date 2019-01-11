package acaseSts

import "encoding/xml"

type StatusListRequestType struct {
	Credentials
	XMLName xml.Name `xml:"StatusListRequest"`
}

type StatusListResponseType struct {
	BaseResponse
	XMLName xml.Name             `xml:"StatusListResponse"`
	Status  []SimpleCodeNameType `xml:"Status"`
}
