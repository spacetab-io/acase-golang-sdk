package acaseSts

import "encoding/xml"

type CurrencyListRequestType struct {
	Credentials
	XMLName xml.Name `xml:"CurrencyListRequest"`
	Code    int      `xml:"Code,attr,omitempty"`
	Name    string   `xml:"Name,attr,omitempty"`
	Options string   `xml:"Options,attr"`
}

type CurrencyListResponseType struct {
	BaseResponse
	XMLName  xml.Name             `xml:"CurrencyListResponse"`
	Currency []SimpleCodeNameType `xml:"Currency"`
}
