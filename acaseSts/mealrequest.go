package acaseSts

import "encoding/xml"

type MealRequestType struct {
	Credentials
	XMLName		xml.Name			`xml:"MealRequest"`
	Action		[]ActionType		`xml:"Action"`
}

type ActionType struct {
	XMLName		xml.Name			`xml:"Action"`
	Name		string				`xml:"Name,attr"`
	Parameters	ParametersType	`xml:"Parameters"`
}

type ParametersType struct {
	XMLName			xml.Name	`xml:"Parameters"`
	MealCode		int			`xml:"MealCode,attr,omitempty"`
	MealTypeCode	int			`xml:"MealTypeCode,attr"`
	MealName		string		`xml:"MealName,attr"`
}

type MealResponseType struct {
	Credentials
	BaseResponse
	XMLName		xml.Name		`xml:"MealResponse"`
	Action		[]ActionType	`xml:"Action"`
	MealList	MealListType	`xml:"MealList"`
}

type MealListType struct {
	XMLName	xml.Name	`xml:"MealList"`
	Items	[]Meal2Type	`xml:"Meal"`
}

type Meal2Type struct {
	SimpleCodeNameType
	TypeCode	int			`xml:"TypeCode,attr,omitempty"`
}
