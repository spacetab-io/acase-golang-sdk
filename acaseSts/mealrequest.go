package acaseSts

import "encoding/xml"

type MealRequestType struct {
	XMLName		xml.Name			`xml:"MealRequest"`
	BuyerId		string				`xml:"BuyerId,attr"`
	UserId		string				`xml:"UserId,attr"`
	Password	string				`xml:"Password,attr"`
	Language	LanguageTypeEnum	`xml:"Language,attr,omitempty"`
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
	XMLName		xml.Name		`xml:"MealResponse"`
	BuyerId		string			`xml:"BuyerId,attr"`
	UserId		string			`xml:"UserId,attr"`
	Password	string			`xml:"Password,attr"`
	Language	string			`xml:"Language,attr,omitempty"`
	Success		string			`xml:"Success"`
	Error		ErrorType		`xml:"Error,omitempty"`
	Action		[]ActionType	`xml:"Action"`
	MealList	MealListType	`xml:"MealList"`
}

type MealListType struct {
	XMLName	xml.Name	`xml:"MealList"`
	Items	[]Meal2Type	`xml:"Meal"`
}

type Meal2Type struct {
	Code		int			`xml:"Code,attr"`
	TypeCode	int			`xml:"TypeCode,attr,omitempty"`
	Name		string		`xml:"Name,attr"`
}
