package acaseSts

import (
	"encoding/xml"
)

type AdmUnit1RequestType struct {
	Credentials
	XMLName xml.Name           `xml:"AdmUnit1Request"`
	Action  AdmUnit1ActionType `xml:"Action"`
}

type AdmUnit1ActionTypeParameters struct {
	CountryCode  int    `xml:"CountryCode,attr"`
	AdmUnit1Code string `xml:"AdmUnit1Code,attr"`
	AdmUnit1Name string `xml:"AdmUnit1Name,attr"`
}

type AdmUnit1ResponseType struct {
	Credentials
	BaseResponse
	XMLName      xml.Name           `xml:"AdmUnit1Response"`
	Action       AdmUnit1ActionType `xml:"Action,omitempty"`
	AdmUnit1List AdmUnit1ListType   `xml:"AdmUnit1List,omitempty"`
}

type AdmUnit1ActionType struct {
	XMLName    xml.Name                     `xml:"Action"`
	Name       AdmUnitActionNameEnum        `xml:"Name,attr"`
	Parameters AdmUnit1ActionTypeParameters `xml:"Parameters"`
}

type AdmUnit1ListType struct {
	XMLName  xml.Name             `xml:"AdmUnit1List"`
	AdmUnit1 []SimpleCodeNameType `xml:"AdmUnit1"`
}
