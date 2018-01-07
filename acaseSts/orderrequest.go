package acaseSts

import "encoding/xml"

type OrderRequestType struct {
	XMLName				xml.Name				`xml:"OrderListRequest"`
	BuyerId				string					`xml:"BuyerId,attr"`
	UserId				string					`xml:"UserId,attr"`
	Password			string					`xml:"Password,attr"`
	Language			LanguageTypeEnum		`xml:"Language,attr,omitempty"`
	Id					int						`xml:"Id,attr"`
	ReferenceNumber		string					`xml:"String20Type,attr"`
	Currency			CurrencyType			`xml:"Currency"`
	WhereToPay			SimpleCodeNameType		`xml:"WhereToPay"`
	Customer			CustomerType			`xml:"Customer"`
	AccommodationList	AccommodationListType	`xml:"AccommodationList"`
}

type OrderResponseType struct {
	XMLName	xml.Name	`xml:"OrderResponse"`
	Success	SuccessType	`xml:"Success"`
	Error	ErrorType	`xml:"Error,omitempty"`
}

type SuccessType struct {
	XMLName	xml.Name	`xml:"Success"`
	Id		int			`xml:"Id,attr"`
}
