package acaseSts

import "encoding/xml"

type ClientCategoryListRequestType struct {
	XMLName				xml.Name			`xml:"ClientCategoryListRequest"`
	BuyerId				string				`xml:"BuyerId,attr"`
	UserId				string				`xml:"UserId,attr"`
	Password			string				`xml:"Password,attr"`
	Language			LanguageTypeEnum	`xml:"Language,attr,omitempty"`
	ClientCategoryCode	int					`xml:"ClientCategoryCode,attr,omitempty"`
	ClientCategoryName	string				`xml:"ClientCategoryName,attr,omitempty"`
}

type ClientCategoryListType struct {
	XMLName				xml.Name				`xml:"ClientCategoryList"`
	Success				string					`xml:"Success"`
	Error				ErrorType				`xml:"Error,omitempty"`
	ClientCategories	[]ClientCategoryType	`xml:"ClientCategory"`
}

type ClientCategoryType struct {
	XMLName	xml.Name	`xml:"ClientCategory"`
	Code	int			`xml:"Code,attr"`
	Name	string		`xml:"Name,attr"`
}
