package acaseSts

import "encoding/xml"

type CustomerRequestInfoType struct {
	Credentials
	XMLName		xml.Name			`xml:"CustomerRequest"`
	ActionInfo	ActionInfoType		`xml:"ActionInfo"`
}

type ActionInfoType struct {
	XMLName		xml.Name						`xml:"ActionInfo"`
	Parameters	CustomerRequestParametersType	`xml:"Parameters"`
}

type CustomerRequestListType struct {
	Credentials
	XMLName		xml.Name			`xml:"CustomerRequest"`
	ActionList	ActionListType		`xml:"ActionList"`
}

type ActionListType struct {
	XMLName		xml.Name						`xml:"ActionList"`
	Parameters	CustomerRequestParametersType	`xml:"Parameters"`
}

type CustomerRequestDeleteType struct {
	Credentials
	XMLName			xml.Name			`xml:"CustomerRequest"`
	ActionDelete	ActionDeleteType	`xml:"ActionDelete"`
}

type ActionDeleteType struct {
	XMLName		xml.Name						`xml:"ActionDelete"`
	Parameters	CustomerRequestParametersType	`xml:"Parameters"`
}

type CustomerRequestCreateType struct {
	Credentials
	XMLName			xml.Name			`xml:"CustomerRequest"`
	ActionCreate	ActionCreateType	`xml:"ActionCreate"`
}

type ActionCreateType struct {
	XMLName		xml.Name						`xml:"ActionCreate"`
	Parameters	CustomerRequestParametersType	`xml:"Parameters"`
}

type CustomerRequestUpdateType struct {
	Credentials
	XMLName			xml.Name			`xml:"CustomerRequest"`
	ActionUpdate	ActionUpdateType	`xml:"ActionUpdate"`
}

type ActionUpdateType struct {
	XMLName		xml.Name						`xml:"ActionUpdate"`
	Parameters	CustomerRequestParametersType	`xml:"Parameters"`
}

type CustomerRequestParametersType struct {
	XMLName			xml.Name		`xml:"Parameters"`
	Sort			int				`xml:"Sort,attr,omitempty"`
	ShowActualOnly	int				`xml:"ShowActualOnly,attr,omitempty"`
	CustomerCode	int				`xml:"CustomerCode,attr,omitempty"`
	Customer		*CustomerType	`xml:"Customer,omitempty"`
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
	BuyerType	SimpleCodeNameType	`xml:"BuyerType"`
	Country		CountryType			`xml:"Country"`
	City		CityType			`xml:"City"`
	Actual		*YesNoCodeType		`xml:"Actual,omitempty"`
	AllowModify	*YesNoCodeType		`xml:"AllowModify,omitempty"`
}

type CustomerResponseDeleteType struct {
	Credentials
	BaseResponse
	XMLName			xml.Name			`xml:"CustomerResponse"`
	ActionDelete	ActionDeleteType	`xml:"ActionDelete"`
	Customer		*CustomerType		`xml:"Customer,omitempty"`
}

type CustomerResponseCreateType struct {
	Credentials
	BaseResponse
	XMLName			xml.Name			`xml:"CustomerResponse"`
	ActionCreate	ActionCreateType	`xml:"ActionCreate"`
	Customer		CustomerType		`xml:"Customer"`
}

type CustomerResponseUpdateType struct {
	Credentials
	BaseResponse
	XMLName			xml.Name			`xml:"CustomerResponse"`
	ActionUpdate	ActionUpdateType	`xml:"ActionUpdate"`
	Customer		CustomerType		`xml:"Customer"`
}

type CustomerResponseListType struct {
	Credentials
	BaseResponse
	XMLName			xml.Name			`xml:"CustomerResponse"`
	ActionList		ActionListType		`xml:"ActionList"`
	CustomerList	CustomerListType	`xml:"CustomerList"`
}

type CustomerListType struct {
	XMLName		xml.Name		`xml:"CustomerList"`
	Customers	[]CustomerType	`xml:"Customer"`
}

type CustomerResponseInfoType struct {
	Credentials
	BaseResponse
	XMLName			xml.Name			`xml:"CustomerResponse"`
	ActionInfo		ActionInfoType		`xml:"ActionInfo"`
	Customer		CustomerType		`xml:"Customer"`
}
