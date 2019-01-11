package acaseSts

import "encoding/xml"

type CountryListRequestType struct {
	Credentials
	XMLName     xml.Name `xml:"CountryListRequest"`
	CountryCode int      `xml:"CountryCode,attr,omitempty"`
	CountryName string   `xml:"CountryName,attr,omitempty"`
}

type CountryListType struct {
	BaseResponse
	XMLName   xml.Name      `xml:"CountryList"`
	Countries []CountryType `xml:"Country"`
}
