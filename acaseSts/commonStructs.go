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

type Credentials struct {
	BuyerId		string				`xml:"BuyerId,attr"`
	UserId		string				`xml:"UserId,attr"`
	Password	string				`xml:"Password,attr"`
	Language	LanguageTypeEnum	`xml:"Language,attr,omitempty"`
}

type BaseResponse struct {
	Success	SuccessType	`xml:"Success"`
	Error	*ErrorType	`xml:"Error,omitempty"`
}

type SuccessType struct {
	XMLName	xml.Name	`xml:"Success"`
	Id		int			`xml:"Id,attr,omitempty"`
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
	Code		int						`xml:"Code,attr"`
	Name		string					`xml:"Name,attr"`
	Genitive	string					`xml:"Genitive,attr,omitempty"`
	Dative		string					`xml:"Dative,attr,omitempty"`
	Accusative	string					`xml:"Accusative,attr,omitempty"`
	Ablative	string					`xml:"Ablative,attr,omitempty"`
	Prepositive	string					`xml:"Prepositive,attr,omitempty"`
	Ref			string					`xml:"Ref,attr,omitempty"`
	CityType	*SimpleCodeNameType		`xml:"CityType,omitempty"`
	AdmUnit1	*SimpleCodeNameType		`xml:"AdmUnit1,omitempty"`
	AdmUnit2	*AdmUnit2Type			`xml:"AdmUnit2,omitempty"`
	TypeOfPlace	*SimpleCodeNameType		`xml:"TypeOfPlace,omitempty"`
	Position	*PositionType			`xml:"Position,omitempty"`
	Hotels		*[]HotelType			`xml:"Hotel,omitempty"`
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
	Rating			SimpleCodeNameType		`xml:"Rating"`
	Type			SimpleCodeNameType		`xml:"Type"`
	Stars			StarsType				`xml:"Stars"`
	VirtualTour		VirtualTourType			`xml:"VirtualTour"`
	FreeSale		SimpleCodeNameType		`xml:"FreeSale"`
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

type VirtualTourType struct {
	XMLName	xml.Name	`xml:"VirtualTour"`
	Code	int			`xml:"Code,attr"`
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
	UseThisAge	YesNoCodeType	`xml:"UseThisAge"`
}

type SimpleCodeNameType struct {
	Code	int		`xml:"Code,attr"`
	Name	string	`xml:"Name,attr"`
}

type SimpleCodeType struct {
	Code	int	`xml:"Code,attr"`
}

