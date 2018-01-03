package acaseSts

import "encoding/xml"

type MealTypeRequestType struct {
	XMLName		xml.Name			`xml:"MealTypeRequest"`
	BuyerId		string				`xml:"BuyerId,attr"`
	UserId		string				`xml:"UserId,attr"`
	Password	string				`xml:"Password,attr"`
	Language	LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	Action		[]ActionType		`xml:"Action"`
}

type MealTypeResponseType struct {
	XMLName			xml.Name			`xml:"MealTypeResponse"`
	BuyerId			string				`xml:"BuyerId,attr"`
	UserId			string				`xml:"UserId,attr"`
	Password		string				`xml:"Password,attr"`
	Language		string				`xml:"Language,attr,omitempty"`
	Success			string				`xml:"Success"`
	Error			ErrorType			`xml:"Error,omitempty"`
	Action			[]ActionType		`xml:"Action"`
	MealTypeList	MealTypeListType	`xml:"MealTypeList"`
}

type MealTypeListType struct {
	XMLName	xml.Name	`xml:"MealTypeList"`
	Items	[]Meal2Type	`xml:"MealType"`
}
