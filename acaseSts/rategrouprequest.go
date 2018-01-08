package acaseSts

import "encoding/xml"

type RateGroupRequestType struct {
	XMLName		xml.Name				`xml:"RateGroupRequest"`
	BuyerId		string					`xml:"BuyerId,attr"`
	UserId		string					`xml:"UserId,attr"`
	Password	string					`xml:"Password,attr"`
	Language	LanguageTypeEnum		`xml:"Language,attr,omitempty"`
	ActionList	RateGroupActionListType	`xml:"ActionList"`
}

type RateGroupActionListType struct {
	XMLName		xml.Name					`xml:"ActionList"`
	Parameters	[]RateGroupParameterType	`xml:"Parameters"`
}

type RateGroupParameterType struct {
	XMLName	xml.Name	`xml:"Parameters"`
	Name	string		`xml:"Name,attr"`
}

type RateGroupResponseType struct {
	XMLName		xml.Name				`xml:"RateGroupResponse"`
	BuyerId		string					`xml:"BuyerId,attr"`
	UserId		string					`xml:"UserId,attr"`
	Password	string					`xml:"Password,attr"`
	Language	string					`xml:"Language,attr,omitempty"`
	Success		SuccessType				`xml:"Success"`
	Error		ErrorType				`xml:"Error,omitempty"`
	ActionList	RateGroupActionListType	`xml:"ActionList"`
	RateGroup	[]SimpleCodeNameType	`xml:"RateGroup"`
}
