package acaseSts

import "encoding/xml"

type TypeOfPlaceRequestType struct {
	Credentials
	XMLName xml.Name                     `xml:"TypeOfPlaceRequest"`
	Action  TypeOfPlaceRequestActionType `xml:"Action"`
}

type TypeOfPlaceRequestActionType struct {
	XMLName    xml.Name             `xml:"Action"`
	Name       string               `xml:"Name,attr"`
	Parameters TypeOfPlaceParamType `xml:"Parameters"`
}

type TypeOfPlaceParamType struct {
	XMLName         xml.Name `xml:"Parameters"`
	TypeOfPlaceCode string   `xml:"TypeOfPlaceCode,attr"`
	TypeOfPlaceName string   `xml:"TypeOfPlaceName,attr"`
}

type TypeOfPlaceResponseType struct {
	Credentials
	BaseResponse
	XMLName         xml.Name                     `xml:"TypeOfPlaceResponse"`
	Action          TypeOfPlaceRequestActionType `xml:"Action"`
	TypeOfPlaceList TypeOfPlaceListType          `xml:"TypeOfPlaceList"`
}

type TypeOfPlaceListType struct {
	XMLName     xml.Name             `xml:"TypeOfPlaceList"`
	TypeOfPlace []SimpleCodeNameType `xml:"TypeOfPlace"`
}
