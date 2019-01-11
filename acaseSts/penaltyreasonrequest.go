package acaseSts

import "encoding/xml"

type PenaltyReasonRequestType struct {
	Credentials
	XMLName xml.Name                `xml:"PenaltyReasonRequest"`
	Action  PenaltyReasonActionType `xml:"Action"`
}

type PenaltyReasonActionType struct {
	XMLName xml.Name `xml:"Action"`
	Name    string   `xml:"Name,attr"`
}

type PenaltyReasonResponseType struct {
	Credentials
	BaseResponse
	XMLName       xml.Name                `xml:"PenaltyReasonResponse"`
	Action        PenaltyReasonActionType `xml:"Action"`
	PenaltyReason []SimpleCodeNameType    `xml:"PenaltyReason"`
}
