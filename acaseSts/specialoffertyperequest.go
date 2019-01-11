package acaseSts

import "encoding/xml"

type SpecialOfferTypeRequestType struct {
	Credentials
	XMLName    xml.Name                   `xml:"SpecialOfferTypeRequest"`
	ActionList SpecialOfferActionListType `xml:"ActionList"`
}

type SpecialOfferActionListType struct {
	XMLName    xml.Name           `xml:"ActionList"`
	Parameters SimpleCodeNameType `xml:"Parameters"`
}

type SpecialOfferTypeResponseType struct {
	Credentials
	BaseResponse
	XMLName              xml.Name                   `xml:"SpecialOfferTypeResponse"`
	ActionList           SpecialOfferActionListType `xml:"ActionList"`
	SpecialOfferTypeList SpecialOfferTypeListType   `xml:"SpecialOfferTypeList"`
}

type SpecialOfferTypeListType struct {
	XMLName          xml.Name            `xml:"SpecialOfferTypeList"`
	SpecialOfferType []SpecOfferTypeType `xml:"SpecialOfferType"`
}

type SpecOfferTypeType struct {
	SimpleCodeNameType
	XMLName    xml.Name      `xml:"SpecialOfferType"`
	IsMultiple YesNoCodeType `xml:"IsMultiple"`
}
