package acaseSts

import "encoding/xml"

type RouteRequestType struct {
	Credentials
	XMLName		xml.Name			`xml:"RouteRequest"`
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
	SimpleCodeNameType
	Type	int		`xml:"Type,attr"`
}

type RouteResponseType struct {
	Credentials
	BaseResponse
	XMLName				xml.Name				`xml:"RouteResponse"`
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
	AdmUnit1			SimpleCodeNameType	`xml:"AdmUnit1"`
	AdmUnit2			AdmUnit2Type		`xml:"AdmUnit2"`
	TypeOfPlace			SimpleCodeNameType	`xml:"TypeOfPlace"`
	City				CityType			`xml:"City"`
	IsAddressDetailed	*YesNoCodeType		`xml:"IsAddressDetailed,omitempty"`
	HasRaceNo			*YesNoCodeType		`xml:"HasRaceNo,omitempty"`
	Hotel				*HotelType			`xml:"Hotel,omitempty"`
	Object				*ObjectType			`xml:"Object,omitempty"`
	ObjectSubType		*SimpleCodeNameType	`xml:"ObjectSubType,omitempty"`
	ObjectType			*SimpleCodeNameType	`xml:"ObjectType,omitempty"`
}

