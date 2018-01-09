package acaseSts

import "encoding/xml"

type PenaltyReasonRequestType struct {
	XMLName		xml.Name				`xml:"PenaltyReasonRequest"`
	BuyerId		string					`xml:"BuyerId,attr"`
	UserId		string					`xml:"UserId,attr"`
	Password	string					`xml:"Password,attr"`
	Language	LanguageTypeEnum		`xml:"Language,attr,omitempty"`
	Action		PenaltyReasonActionType	`xml:"Action"`
}

type PenaltyReasonActionType struct {
	XMLName	xml.Name	`xml:"Action"`
	Name	string		`xml:"Name,attr"`
}

type PenaltyReasonResponseType struct {
	XMLName			xml.Name				`xml:"PenaltyReasonResponse"`
	BuyerId			string					`xml:"BuyerId,attr"`
	UserId			string					`xml:"UserId,attr"`
	Password		string					`xml:"Password,attr"`
	Language		string					`xml:"Language,attr,omitempty"`
	Success			SuccessType				`xml:"Success"`
	Error			ErrorType				`xml:"Error,omitempty"`
	Action			PenaltyReasonActionType	`xml:"Action"`
	PenaltyReason	[]SimpleCodeNameType	`xml:"PenaltyReason"`
}
