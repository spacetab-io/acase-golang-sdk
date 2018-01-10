package acaseSts

import "encoding/xml"

type OrderRequestType struct {
	Credentials
	XMLName				xml.Name				`xml:"OrderListRequest"`
	Id					int						`xml:"Id,attr"`
	ReferenceNumber		string					`xml:"String20Type,attr"`
	Currency			SimpleCodeNameType		`xml:"Currency"`
	WhereToPay			SimpleCodeNameType		`xml:"WhereToPay"`
	Customer			CustomerType			`xml:"Customer"`
	AccommodationList	AccommodationListType	`xml:"AccommodationList"`
}

type OrderAwocRequestType struct {
	Credentials
	XMLName			xml.Name			`xml:"OrderListRequest"`
	Id				int					`xml:"Id,attr"`
	ReferenceNumber	string				`xml:"String20Type,attr"`
	Currency		SimpleCodeNameType	`xml:"Currency"`
	WhereToPay		SimpleCodeNameType	`xml:"WhereToPay"`
	Customer		CustomerType		`xml:"Customer"`
	AwocList		AwocListType		`xml:"AwocList"`
}

type OrderResponseType struct {
	BaseResponse
	XMLName	xml.Name	`xml:"OrderResponse"`
}
