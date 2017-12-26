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
	Type		ErrorTypeType	`xml:"Type"`
	Pointer		string			`xml:"Pointer"`
	Code		string			`xml:"Code,attr"`
	Description	string			`xml:"Description,attr"`
}

type ErrorTypeType struct {
	XMLName		xml.Name	`xml:"Type"`
	Code		string		`xml:"Code,attr"`
	Name		string		`xml:"Name,attr"`
}

type CountryType struct {
	XMLName		xml.Name		`xml:"Country"`
	Code		int				`xml:"Code,attr"`
	Name		string			`xml:"Name,attr"`
	Position	PositionType	`xml:"Position"`
	Cities		[]CityType		`xml:"City,omitempty"`
}

type CityType struct {
	XMLName		xml.Name		`xml:"City"`
	Code		int				`xml:"Code,attr"`
	Name		string			`xml:"Name,attr"`
	Genitive	string			`xml:"Genitive,attr"`
	Dative		string			`xml:"Dative,attr"`
	Accusative	string			`xml:"Accusative,attr"`
	Ablative	string			`xml:"Ablative,attr"`
	Prepositive	string			`xml:"Prepositive,attr"`
	Ref			string			`xml:"Ref,arre"`
	AdmUnit1	AdmUnit1Type	`xml:"AdmUnit1"`
	AdmUnit2	AdmUnit2Type	`xml:"AdmUnit2"`
	TypeOfPlace	TypeOfPlaceType	`xml:"TypeOfPlace"`
	Position	PositionType	`xml:"Position"`
}

type TypeOfPlaceType struct {
	XMLName	xml.Name	`xml:"TypeOfPlace"`
	Code	int			`xml:"Code,attr"`
	Name	string		`xml:"Name,attr"`
}

type PositionType struct {
	XMLName		xml.Name	`xml:"Position"`
	Latitude	string		`xml:"Latitude,attr"`
	Longitude	string		`xml:"Longitude,attr"`
}
