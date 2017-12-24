package acaseSts

import (
	"encoding/xml"
)

type CitizenshipListRequestType struct {
	XMLName		xml.Name			`xml:"CitizenshipListRequest"`
	BuyerId		string				`xml:"BuyerId,attr"`
	UserId		string				`xml:"UserId,attr"`
	Password	string				`xml:"Password,attr"`
	Language	LanguageTypeEnum	`xml:"Language,attr,omitempty"`
}

type CitizenshipListType struct {
	XMLName		xml.Name			`xml:"CitizenshipList"`
	Success		string				`xml:"Success,omitempty"`
	Error		ErrorType			`xml:"Error,omitempty"`
	Citizenship	[]CitizenshipType	`xml:"Citizenship"`
}

type CitizenshipType struct {
	XMLName	xml.Name	`xml:"Citizenship"`
	Code	string		`xml:"Code,attr"`
	Name	string		`xml:"Name,attr"`
}
