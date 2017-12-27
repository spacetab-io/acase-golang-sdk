package acaseSts

import "encoding/xml"

type CustomerRequestInfoType struct {
	XMLName		xml.Name			`xml:"CustomerRequest"`
	BuyerId		string				`xml:"BuyerId,attr"`
	UserId		string				`xml:"UserId,attr"`
	Password	string				`xml:"Password,attr"`
	Language	LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	ActionInfo	ActionInfoType		`xml:"ActionInfo"`
}

type ActionInfoType struct {
	XMLName		xml.Name						`xml:"ActionInfo"`
	Parameters	CustomerRequestParametersType	`xml:"Parameters"`
}

type CustomerRequestListType struct {
	XMLName		xml.Name			`xml:"CustomerRequest"`
	BuyerId		string				`xml:"BuyerId,attr"`
	UserId		string				`xml:"UserId,attr"`
	Password	string				`xml:"Password,attr"`
	Language	LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	ActionList	ActionListType		`xml:"ActionList"`
}

type ActionListType struct {
	XMLName		xml.Name						`xml:"ActionList"`
	Parameters	CustomerRequestParametersType	`xml:"Parameters"`
}

type CustomerRequestDeleteType struct {
	XMLName			xml.Name			`xml:"CustomerRequest"`
	BuyerId			string				`xml:"BuyerId,attr"`
	UserId			string				`xml:"UserId,attr"`
	Password		string				`xml:"Password,attr"`
	Language		LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	ActionDelete	ActionDeleteType	`xml:"ActionDelete"`
}

type ActionDeleteType struct {
	XMLName		xml.Name						`xml:"ActionDelete"`
	Parameters	CustomerRequestParametersType	`xml:"Parameters"`
}

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
	Parameters	CustomerRequestParametersType	`xml:"Parameters"`
}

type CustomerRequestUpdateType struct {
	XMLName			xml.Name			`xml:"CustomerRequest"`
	BuyerId			string				`xml:"BuyerId,attr"`
	UserId			string				`xml:"UserId,attr"`
	Password		string				`xml:"Password,attr"`
	Language		LanguageTypeEnum	`xml:"Language,attr,omitempty"`
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

type CustomerResponseDeleteType struct {
	XMLName			xml.Name			`xml:"CustomerResponse"`
	BuyerId			string				`xml:"BuyerId,attr"`
	UserId			string				`xml:"UserId,attr"`
	Password		string				`xml:"Password,attr"`
	Language		LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	ActionDelete	ActionDeleteType	`xml:"ActionDelete"`
	Success			string				`xml:"Success"`
	Error			ErrorType			`xml:"Error,omitempty"`
	Customer		*CustomerType		`xml:"Customer,omitempty"`
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

type CustomerResponseUpdateType struct {
	XMLName			xml.Name			`xml:"CustomerResponse"`
	BuyerId			string				`xml:"BuyerId,attr"`
	UserId			string				`xml:"UserId,attr"`
	Password		string				`xml:"Password,attr"`
	Language		LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	ActionUpdate	ActionUpdateType	`xml:"ActionUpdate"`
	Success			string				`xml:"Success"`
	Error			ErrorType			`xml:"Error,omitempty"`
	Customer		CustomerType		`xml:"Customer"`
}

type CustomerResponseListType struct {
	XMLName			xml.Name			`xml:"CustomerResponse"`
	BuyerId			string				`xml:"BuyerId,attr"`
	UserId			string				`xml:"UserId,attr"`
	Password		string				`xml:"Password,attr"`
	Language		LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	ActionList		ActionListType		`xml:"ActionList"`
	Success			string				`xml:"Success"`
	Error			ErrorType			`xml:"Error,omitempty"`
	CustomerList	CustomerListType	`xml:"CustomerList"`
}

type CustomerListType struct {
	XMLName		xml.Name		`xml:"CustomerList"`
	Customers	[]CustomerType	`xml:"Customer"`
}

type CustomerResponseInfoType struct {
	XMLName			xml.Name			`xml:"CustomerResponse"`
	BuyerId			string				`xml:"BuyerId,attr"`
	UserId			string				`xml:"UserId,attr"`
	Password		string				`xml:"Password,attr"`
	Language		LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	ActionInfo		ActionInfoType		`xml:"ActionInfo"`
	Success			string				`xml:"Success"`
	Error			ErrorType			`xml:"Error,omitempty"`
	Customer		CustomerType		`xml:"Customer"`
}
