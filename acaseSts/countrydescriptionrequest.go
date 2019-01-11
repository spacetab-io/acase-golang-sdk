package acaseSts

import "encoding/xml"

type CountryDescriptionRequestType struct {
	Credentials
	XMLName     xml.Name `xml:"CountryDescriptionRequest"`
	CountryCode int      `xml:"CountryCode,attr"`
}

type CountryDescriptionType struct {
	BaseResponse
	XMLName     xml.Name     `xml:"CountryDescription"`
	Code        int          `xml:"Code,attr"`
	Name        string       `xml:"Name,attr"`
	Description string       `xml:"Description,attr"`
	Position    PositionType `xml:"Position"`
}
