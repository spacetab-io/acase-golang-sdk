package acaseSts

import "encoding/xml"

type CountryDescriptionRequestType struct {
	XMLName		xml.Name			`xml:"CountryDescriptionRequest"`
	BuyerId		string				`xml:"BuyerId,attr"`
	UserId		string				`xml:"UserId,attr"`
	Password	string				`xml:"Password,attr"`
	Language	LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	CountryCode	int					`xml:"CountryCode,attr"`
}

type CountryDescriptionType struct {
	XMLName		xml.Name		`xml:"CountryDescription"`
	Code		int				`xml:"Code,attr"`
	Name		string			`xml:"Name,attr"`
	Description	string			`xml:"Description,attr"`
	Success		string			`xml:"Success,omitempty"`
	Error		ErrorType		`xml:"Error,omitempty"`
	Position	PositionType	`xml:"Position"`
}
