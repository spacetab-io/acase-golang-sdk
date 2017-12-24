package acase_sdk

import (
	"encoding/xml"
)

type LanguageTypeEnum string
const (
	Ru LanguageTypeEnum = "ru"
	En LanguageTypeEnum = "en"
)

type AdmUnitActionNameEnum string
const (
	List AdmUnitActionNameEnum = "LIST"
)

type ErrorType struct {
	XMLName		xml.Name		`xml:"Error"`
	Type		ErrorTypeType	`xml:"Type"`
	Pointer		string			`xml:"Pointer"`
	Code		int				`xml:"Code,attr"`
	Description	string			`xml:"Description"`
}

type ErrorTypeType struct {
	XMLName		xml.Name	`xml:"Type"`
	Code		int			`xml:"Code,attr"`
	Name		string		`xml:"Code,attr"`
}
