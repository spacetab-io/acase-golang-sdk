package acaseStructs

import (
	"encoding/xml"
)

type AdmUnit1RequestType struct {
	XMLName		xml.Name			`xml:"AdmUnit1Request"`
	BuyerId		string				`xml:"BuyerId,attr"`
	UserId		string				`xml:"UserId,attr"`
	Password	string				`xml:"Password,attr"`
	Language	LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	Action		ActionType			`xml:"Action"`
}

type ActionType struct {
	XMLName		xml.Name						`xml:"Action"`
	Name		AdmUnitActionNameEnum			`xml:"Name,attr"`
	Parameters	AdmUnit1ActionTypeParameters	`xml:"Parameters"`
}

type AdmUnit1ActionTypeParameters struct {
	CountryCode		int		`xml:"CountryCode,attr"`
	AdmUnit1Code	string	`xml:"AdmUnit1Code,attr"`
	AdmUnit1Name	string	`xml:"AdmUnit1Name,attr"`
}

type AdmUnit1ResponseType struct {
	XMLName			xml.Name			`xml:"AdmUnit1Response"`
	BuyerId			string				`xml:"BuyerId,attr"`
	UserId			string				`xml:"UserId,attr"`
	Password		string				`xml:"Password,attr"`
	Language		LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	Action			ActionType			`xml:"Action"`
	Success			string				`xml:"Success,omitempty"`
	AdmUnit1List	AdmUnit1ListType	`xml:"AdmUnit1List"`
	Error			ErrorType			`xml:"Error,omitempty"`
}

type AdmUnit1ListType struct {
	XMLName		xml.Name		`xml:"AdmUnit1List",json:"-"`
	AdmUnit1	[]AdmUnit1Type	`xml:"AdmUnit1",json:"adm_unit_1"`
}

type AdmUnit1Type struct {
	Code	int		`xml:"Code,attr",json:"code"`
	Name	string	`xml:"Name,attr",json:"name"`
}

type ErrorType struct {
	XMLName		xml.Name		`xml:"Error"`
	Code		int				`xml:"Code,attr"`
	Description	string			`xml:"Description"`
	Type		ErrorTypeType	`xml:"Type"`
}

type ErrorTypeType struct {
	XMLName		xml.Name	`xml:"Type"`
	Code		int			`xml:"Code,attr"`
	Name		string		`xml:"Code,attr"`
}
