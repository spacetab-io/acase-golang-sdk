package acaseSts

import "encoding/xml"

type ClientCategoryListRequestType struct {
	Credentials
	XMLName				xml.Name			`xml:"ClientCategoryListRequest"`
	ClientCategoryCode	int					`xml:"ClientCategoryCode,attr,omitempty"`
	ClientCategoryName	string				`xml:"ClientCategoryName,attr,omitempty"`
}

type ClientCategoryListType struct {
	BaseResponse
	XMLName				xml.Name				`xml:"ClientCategoryList"`
	ClientCategories	[]SimpleCodeNameType	`xml:"ClientCategory"`
}
