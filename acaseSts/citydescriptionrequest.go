package acaseSts

import "encoding/xml"

type CityDescriptionRequestType struct {
	XMLName		xml.Name			`xml:"CityDescriptionRequest"`
	BuyerId		string				`xml:"BuyerId,attr"`
	UserId		string				`xml:"UserId,attr"`
	Password	string				`xml:"Password,attr"`
	Language	LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	CityCode	string				`xml:"CityCode,attr"`
}

type CityDescriptionType struct {
	XMLName		xml.Name		`xml:"CityDescription"`
	Code		string			`xml:"Code,attr"`
	Name		string			`xml:"Name,attr"`
	Description	string			`xml:"Description,attr"`
	Url			string			`xml:"Url,attr"`
	Success		string			`xml:"Success,omitempty"`
	Error		ErrorType		`xml:"Error,omitempty"`
	Country		CountryType		`xml:"Country"`
	AdmUnit1	AdmUnit1Type	`xml:"AdmUnit1"`
	AdmUnit2	AdmUnit2Type	`xml:"AdmUnit2"`
	TypeOfPlace	TypeOfPlaceType	`xml:"TypeOfPlace"`
	Position	PositionType	`xml:"Position"`
}
