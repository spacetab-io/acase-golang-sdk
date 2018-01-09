package acaseSts

import "encoding/xml"

type SpecialOfferTypeRequestType struct {
	XMLName		xml.Name					`xml:"SpecialOfferTypeRequest"`
	BuyerId		string						`xml:"BuyerId,attr"`
	UserId		string						`xml:"UserId,attr"`
	Password	string						`xml:"Password,attr"`
	Language	LanguageTypeEnum			`xml:"Language,attr,omitempty"`
	ActionList	SpecialOfferActionListType	`xml:"ActionList"`
}

type SpecialOfferActionListType struct {
	XMLName		xml.Name			`xml:"ActionList"`
	Parameters	SimpleCodeNameType	`xml:"Parameters"`
}

type SpecialOfferTypeResponseType struct {
	XMLName					xml.Name					`xml:"SpecialOfferTypeResponse"`
	BuyerId					string						`xml:"BuyerId,attr"`
	UserId					string						`xml:"UserId,attr"`
	Password				string						`xml:"Password,attr"`
	Language				string						`xml:"Language,attr,omitempty"`
	Success					SuccessType					`xml:"Success"`
	Error					ErrorType					`xml:"Error,omitempty"`
	ActionList				SpecialOfferActionListType	`xml:"ActionList"`
	SpecialOfferTypeList	SpecialOfferTypeListType	`xml:"SpecialOfferTypeList"`
}

type SpecialOfferTypeListType struct {
	XMLName				xml.Name			`xml:"SpecialOfferTypeList"`
	SpecialOfferType	[]SpecOfferTypeType	`xml:"SpecialOfferType"`
}

type SpecOfferTypeType struct {
	XMLName		xml.Name		`xml:"SpecialOfferType"`
	Code		int				`xml:"Code,attr"`
	Name		string			`xml:"Name,attr"`
	IsMultiple	YesNoCodeType	`xml:"IsMultiple"`
}