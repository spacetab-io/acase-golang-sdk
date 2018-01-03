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

type DestinationTypeEnum int
const (
	CITY  = 2
	HOTEL = 3
	POI   = 4
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
	Ref			string			`xml:"Ref,attr,omitempty"`
	Position	*PositionType	`xml:"Position,omitempty"`
	Cities		*[]CityType		`xml:"City,omitempty"`
}

type AlternativeCountryType struct {
	XMLName		xml.Name		`xml:"AlternativeCountry"`
	Code		int				`xml:"Code,attr"`
	Name		string			`xml:"Name,attr"`
	Position	*PositionType	`xml:"Position,omitempty"`
	Cities		*[]CityType		`xml:"AlternativeCity,omitempty"`
}

type CityType struct {
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
	Hotels		*[]HotelType		`xml:"Hotel,omitempty"`
}

type HotelType struct {
	XMLName			xml.Name				`xml:"Hotel"`
	Code			int						`xml:"Code,attr"`
	Name			string					`xml:"Name,attr"`
	Zip				string					`xml:"Zip,attr"`
	Address			string					`xml:"Address,attr"`
	Underground		string					`xml:"Underground,attr"`
	CityCentre		string					`xml:"CityCentre,attr"`
	Description		string					`xml:"Description,attr"`
	Url				string					`xml:"Url,attr"`
	Position		PositionType			`xml:"Position"`
	Rating			RatingType				`xml:"Rating"`
	Type			TypeType				`xml:"Type"`
	Stars			StarsType				`xml:"Stars"`
	VirtualTour		VirtualTourType			`xml:"VirtualTour"`
	FreeSale		FreeSaleType			`xml:"FreeSale"`
	Amenities		HotelAmenitiesListType	`xml:"Amenities"`
}

type HotelAmenitiesListType struct {
	XMLName	xml.Name	`xml:"Amenities"`
	Amenity	[]HAType	`xml:"HA"`
}

type HAType struct {
	XMLName	xml.Name	`xml:"HA"`
	Code	int			`xml:"Code,attr"`
}

type FreeSaleType struct {
	XMLName	xml.Name	`xml:"FreeSale"`
	Code	int			`xml:"Code,attr"`
	Name	string		`xml:"Name,attr"`
}

type VirtualTourType struct {
	XMLName	xml.Name	`xml:"VirtualTour"`
	Code	int			`xml:"Code,attr"`
}

type TypeType struct {
	XMLName	xml.Name	`xml:"Type"`
	Code	int			`xml:"Code,attr"`
	Name	string		`xml:"Name,attr"`
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
	XMLName			xml.Name	`xml:"Position"`
	Latitude		string		`xml:"Latitude,attr,omitempty"`
	Longitude		string		`xml:"Longitude,attr,omitempty"`
	OptimalGmapZoom	int			`xml:"OptimalGmapZoom,attr,omitempty"`
}

type YesNoCodeType struct {
	Code	int		`xml:"Code,attr"`
	Name	string	`xml:"Name,attr"`
}

type AgeRestrictionType struct {
	AgeTo		int				`xml:"AgeTo,attr"`
	AgeFrom		int				`xml:"AgeFrom"`
	UseThisAge	UseThisAgeType	`xml:"UseThisAge"`
}

type UseThisAgeType struct {
	XMLName	xml.Name	`xml:"UseThisAge"`
	Code	int			`xml:"Code,attr"`
	Name	string		`xml:"Name,attr"`
}

type SimpleCodeNameType struct {
	Code	int		`xml:"Code,attr"`
	Name	string	`xml:"Name,attr"`
}