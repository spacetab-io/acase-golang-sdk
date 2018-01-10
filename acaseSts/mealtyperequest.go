package acaseSts

import "encoding/xml"

type MealTypeRequestType struct {
	Credentials
	XMLName		xml.Name			`xml:"MealTypeRequest"`
	Action		[]ActionType		`xml:"Action"`
}

type MealTypeResponseType struct {
	Credentials
	BaseResponse
	XMLName			xml.Name			`xml:"MealTypeResponse"`
	Action			[]ActionType		`xml:"Action"`
	MealTypeList	MealTypeListType	`xml:"MealTypeList"`
}

type MealTypeListType struct {
	XMLName	xml.Name	`xml:"MealTypeList"`
	Items	[]Meal2Type	`xml:"MealType"`
}
