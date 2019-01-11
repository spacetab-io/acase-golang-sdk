package acaseSts

import "encoding/xml"

type MealTypeRequestType struct {
	Credentials
	XMLName xml.Name           `xml:"MealTypeRequest"`
	Action  MealTypeActionType `xml:"Action"`
}

type MealTypeActionType struct {
	XMLName    xml.Name               `xml:"Action"`
	Name       string                 `xml:"Name,attr"`
	Parameters MealTypeParametersType `xml:"Parameters"`
}

type MealTypeParametersType struct {
	XMLName      xml.Name `xml:"Parameters"`
	MealTypeCode string   `xml:"MealTypeCode,attr"`
	MealTypeName string   `xml:"MealName,attr"`
}

type MealTypeResponseType struct {
	Credentials
	BaseResponse
	XMLName      xml.Name           `xml:"MealTypeResponse"`
	Action       MealTypeActionType `xml:"Action"`
	MealTypeList MealTypeListType   `xml:"MealTypeList"`
}

type MealTypeListType struct {
	XMLName xml.Name             `xml:"MealTypeList"`
	Items   []SimpleCodeNameType `xml:"MealType"`
}
