package acaseSts

import "encoding/xml"

type CityDescriptionRequestType struct {
	Credentials
	XMLName		xml.Name			`xml:"CityDescriptionRequest"`
	CityCode	string				`xml:"CityCode,attr"`
}

type CityDescriptionType struct {
	BaseResponse
	XMLName		xml.Name			`xml:"CityDescription"`
	Code		string				`xml:"Code,attr"`
	Name		string				`xml:"Name,attr"`
	Description	string				`xml:"Description,attr"`
	Url			string				`xml:"Url,attr"`
	Country		CountryType			`xml:"Country"`
	AdmUnit1	SimpleCodeNameType	`xml:"AdmUnit1"`
	AdmUnit2	AdmUnit2Type		`xml:"AdmUnit2"`
	TypeOfPlace	SimpleCodeNameType	`xml:"TypeOfPlace"`
	Position	PositionType		`xml:"Position"`
}
