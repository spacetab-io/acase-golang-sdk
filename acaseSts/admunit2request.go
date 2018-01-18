package acaseSts

import "encoding/xml"

type AdmUnit2RequestType struct {
	Credentials
	XMLName		xml.Name			`xml:"AdmUnit2Request"`
	Action		AdmUnit2ActionType	`xml:"Action"`
}

type AdmUnit2ActionType struct {
	XMLName		xml.Name						`xml:"Action"`
	Name		AdmUnitActionNameEnum			`xml:"Name,attr"`
	Parameters	AdmUnit2ActionTypeParameters	`xml:"Parameters"`
}

type AdmUnit2ActionTypeParameters struct {
	CountryCode		int		`xml:"CountryCode,attr"`
	AdmUnit1Code	int		`xml:"AdmUnit1Code,attr,omitempty"`
	AdmUnit1Name	string	`xml:"AdmUnit1Name,attr,omitempty"`
	AdmUnit2Code	int		`xml:"AdmUnit2Code,attr,omitempty"`
	AdmUnit2Name	string	`xml:"AdmUnit2Name,attr,omitempty"`
}

type AdmUnit2ResponseType struct {
	Credentials
	BaseResponse
	XMLName			xml.Name			`xml:"AdmUnit2Response"`
	Action			AdmUnit2ActionType	`xml:"Action,omitempty"`
	AdmUnit2List	AdmUnit2ListType	`xml:"AdmUnit2List,omitempty"`
}

type AdmUnit2ListType struct {
	XMLName		xml.Name		`xml:"AdmUnit2List",json:"-"`
	AdmUnit2	[]AdmUnit2Type	`xml:"AdmUnit2",json:"adm_unit_2"`
}

type AdmUnit2Type struct {
	AdmUnit1	SimpleCodeNameType	`xml:"AdmUnit1",json:"adm_unit_1"`
	Code		int					`xml:"Code,attr",json:"code"`
	Name		string				`xml:"Name,attr",json:"name"`
}