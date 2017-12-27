package acaseSts

import (
	"encoding/xml"
)

type LanguageTypeEnum string
const (
	Ru LanguageTypeEnum = "ru"
	En LanguageTypeEnum = "en"
)

type AdmUnitActionNameEnum string
const (
	List AdmUnitActionNameEnum = "LIST"
)

type ErrorType struct {
	XMLName		xml.Name		`xml:"Error"`
	Type		ErrorTypeType	`xml:"Type,omitempty"`
	Pointer		string			`xml:"Pointer"`
	Code		string			`xml:"Code,attr"`
	Description	string			`xml:"Description,attr"`
}

type ErrorTypeType struct {
	XMLName		xml.Name	`xml:"Type"`
	Code		string		`xml:"Code,attr,omitempty"`
	Name		string		`xml:"Name,attr,omitempty"`
}

type CountryType struct {
	XMLName		xml.Name		`xml:"Country"`
	Code		int				`xml:"Code,attr"`
	Name		string			`xml:"Name,attr"`
	Position	*PositionType	`xml:"Position,omitempty"`
	Cities		[]CityType		`xml:"City,omitempty"`
}

type CityType struct {
	XMLName		xml.Name			`xml:"City"`
	Code		int					`xml:"Code,attr"`
	Name		string				`xml:"Name,attr"`
	Genitive	string				`xml:"Genitive,attr,omitempty"`
	Dative		string				`xml:"Dative,attr,omitempty"`
	Accusative	string				`xml:"Accusative,attr,omitempty"`
	Ablative	string				`xml:"Ablative,attr,omitempty"`
	Prepositive	string				`xml:"Prepositive,attr,omitempty"`
	Ref			string				`xml:"Ref,attr,omitempty"`
	CityType	*CityTypeType		`xml:"CityType,omitempty"`
	AdmUnit1	*AdmUnit1Type		`xml:"AdmUnit1,omitempty"`
	AdmUnit2	*AdmUnit2Type		`xml:"AdmUnit2,omitempty"`
	TypeOfPlace	*TypeOfPlaceType	`xml:"TypeOfPlace,omitempty"`
	Position	*PositionType		`xml:"Position,omitempty"`
}

type CityTypeType struct {
	XMLName	xml.Name	`xml:"CityType"`
	Code	int			`xml:"Code,attr,omitempty"`
	Name	string		`xml:"Name,attr,omitempty"`
}

type TypeOfPlaceType struct {
	XMLName	xml.Name	`xml:"TypeOfPlace"`
	Code	int			`xml:"Code,attr,omitempty"`
	Name	string		`xml:"Name,attr,omitempty"`
}

type PositionType struct {
	XMLName		xml.Name	`xml:"Position"`
	Latitude	string		`xml:"Latitude,attr,omitempty"`
	Longitude	string		`xml:"Longitude,attr,omitempty"`
}
