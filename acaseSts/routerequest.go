package acaseSts

import "encoding/xml"

type RouteRequestType struct {
	XMLName		xml.Name			`xml:"RouteRequest"`
	BuyerId		string				`xml:"BuyerId,attr"`
	UserId		string				`xml:"UserId,attr"`
	Password	string				`xml:"Password,attr"`
	Language	LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	Action		ActionRouteType		`xml:"Action"`
}

type ActionRouteType struct {
	XMLName		xml.Name			`xml:"Action"`
	Name		string				`xml:"Name,attr"`
	Parameters	RouteParametersType	`xml:"Parameters"`
}

type RouteParametersType struct {
	XMLName		xml.Name	`xml:"Parameters"`
	StartPoint	PointType	`xml:"StartPoint"`
	EndPoint	PointType	`xml:"EndPoint"`
}

type PointType struct {
	Name	string	`xml:"Name,attr"`
	Type	int		`xml:"Type,attr"`
	Code	int		`xml:"Code,attr"`
}

type RouteResponseType struct {
	XMLName				xml.Name				`xml:"RouteResponse"`
	BuyerId				string					`xml:"BuyerId,attr"`
	UserId				string					`xml:"UserId,attr"`
	Password			string					`xml:"Password,attr"`
	Language			string					`xml:"Language,attr,omitempty"`
	Success				SuccessType				`xml:"Success"`
	Error				ErrorType				`xml:"Error,omitempty"`
	Action				ActionRouteType			`xml:"Action"`
	TransferPointList	TransferPointListType	`xml:"TransferPointList"`
}

type TransferPointListType struct {
	XMLName			xml.Name			`xml:"TransferPointList"`
	TransferPoint	[]TransferPointType	`xml:"TransferPoint"`
}

type TransferPointType struct {
	XMLName				xml.Name			`xml:"TransferPoint"`
	Code				int					`xml:"Code,attr"`
	TransferPointType	SimpleCodeNameType	`xml:"TransferPointType"`
	Country				CountryType			`xml:"Country"`
	AdmUnit1			AdmUnit1Type		`xml:"AdmUnit1"`
	AdmUnit2			AdmUnit2Type		`xml:"AdmUnit2"`
	TypeOfPlace			SimpleCodeNameType	`xml:"TypeOfPlace"`
	City				CityType			`xml:"City"`
	IsAddressDetailed	*YesNoCodeType		`xml:"IsAddressDetailed,omitempty"`
	HasRaceNo			*YesNoCodeType		`xml:"HasRaceNo,omitempty"`
	Hotel				*HotelType			`xml:"Hotel,omitempty"`
	Object				*ObjectType			`xml:"Object,omitempty"`
	ObjectSubType		*ObjectSubTypeType	`xml:"ObjectSubType,omitempty"`
	ObjectType			*ObjectTypeType		`xml:"ObjectType,omitempty"`
}

