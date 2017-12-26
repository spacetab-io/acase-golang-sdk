package acaseSts

import "encoding/xml"

type AdmUnit2RequestType struct {
	XMLName		xml.Name			`xml:"AdmUnit2Request"`
	BuyerId		string				`xml:"BuyerId,attr"`
	UserId		string				`xml:"UserId,attr"`
	Password	string				`xml:"Password,attr"`
	Language	LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	Action		AdmUnit2ActionType	`xml:"Action"`
}

type AdmUnit2ActionType struct {
	XMLName		xml.Name						`xml:"Action"`
	Name		AdmUnitActionNameEnum			`xml:"Name,attr"`
	Parameters	AdmUnit2ActionTypeParameters	`xml:"Parameters"`
}

type AdmUnit2ActionTypeParameters struct {
	CountryCode		int		`xml:"CountryCode,attr"`
	AdmUnit1Code	string	`xml:"AdmUnit1Code,attr"`
	AdmUnit1Name	string	`xml:"AdmUnit1Name,attr"`
	AdmUnit2Code	string	`xml:"AdmUnit2Code,attr"`
	AdmUnit2Name	string	`xml:"AdmUnit2Name,attr"`
}

type AdmUnit2ResponseType struct {
	XMLName			xml.Name			`xml:"AdmUnit2Response"`
	BuyerId			string				`xml:"BuyerId,attr"`
	UserId			string				`xml:"UserId,attr"`
	Password		string				`xml:"Password,attr"`
	Language		LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	Action			AdmUnit2ActionType	`xml:"Action,omitempty"`
	Success			string				`xml:"Success"`
	AdmUnit2List	AdmUnit2ListType	`xml:"AdmUnit2List,omitempty"`
	Error			ErrorType			`xml:"Error,omitempty"`
}

type AdmUnit2ListType struct {
	XMLName		xml.Name		`xml:"AdmUnit2List",json:"-"`
	AdmUnit2	[]AdmUnit2Type	`xml:"AdmUnit2",json:"adm_unit_2"`
}

type AdmUnit2Type struct {
	AdmUnit1	AdmUnit1Type	`xml:"AdmUnit1",json:"adm_unit_1"`
	Code		int				`xml:"Code,attr",json:"code"`
	Name		string			`xml:"Name,attr",json:"name"`
}