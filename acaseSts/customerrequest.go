package acaseSts

import "encoding/xml"

type CustomerRequestCreateType struct {
	XMLName			xml.Name			`xml:"CustomerRequest"`
	BuyerId			string				`xml:"BuyerId,attr"`
	UserId			string				`xml:"UserId,attr"`
	Password		string				`xml:"Password,attr"`
	Language		LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	ActionCreate	ActionCreateType	`xml:"ActionCreate"`
}

type ActionCreateType struct {
	XMLName		xml.Name						`xml:"ActionCreate"`
	Parameters	CustomerReqCreateParametersType	`xml:"Parameters"`
}

type CustomerReqCreateParametersType struct {
	XMLName			xml.Name		`xml:"Parameters"`
	CustomerCode	string			`xml:"CustomerCode,attr"`
	Customer		CustomerType	`xml:"Customer"`
}

type CustomerType struct {
	XMLName		xml.Name			`xml:"Customer"`
	FullName	string				`xml:"FullName,attr"`
	ZipCode		string				`xml:"ZipCode,attr"`
	Address		string				`xml:"Address,attr"`
	PIAddress	string				`xml:"PIAddress,attr"`
	INN			string				`xml:"INN,attr"`
	KPP			string				`xml:"KPP,attr"`
	Phone		string				`xml:"Phone,attr"`
	Code		int					`xml:"Code,attr,omitempty"`
	Name		string				`xml:"Name,attr"`
	BuyerType	BuyerTypeType		`xml:"BuyerType"`
	Country		CountryType			`xml:"Country"`
	City		CityType			`xml:"City"`
	Actual		*ActualType			`xml:"Actual,omitempty"`
	AllowModify	*AllowModifyType	`xml:"AllowModify,omitempty"`
}

type AllowModifyType struct {
	XMLName	xml.Name	`xml:"AllowModify"`
	Code	int			`xml:"Code,attr,omitempty"`
	Name	string		`xml:"Name,attr,omitempty"`
}

type ActualType struct {
	XMLName	xml.Name	`xml:"Actual"`
	Code	int			`xml:"Code,attr,omitempty"`
	Name	string		`xml:"Name,attr,omitempty"`
}

type BuyerTypeType struct {
	XMLName	xml.Name	`xml:"BuyerType"`
	Code	int			`xml:"Code,attr"`
	Name	string		`xml:"Name,attr"`
}

type CustomerResponseCreateType struct {
	XMLName			xml.Name			`xml:"CustomerResponse"`
	BuyerId			string				`xml:"BuyerId,attr"`
	UserId			string				`xml:"UserId,attr"`
	Password		string				`xml:"Password,attr"`
	Language		LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	ActionCreate	ActionCreateType	`xml:"ActionCreate"`
	Success			string				`xml:"Success"`
	Error			ErrorType			`xml:"Error,omitempty"`
	Customer		CustomerType		`xml:"Customer"`
}