package acaseSts

import (
	"encoding/xml"
)

type CitizenshipListRequestType struct {
	Credentials
	XMLName		xml.Name			`xml:"CitizenshipListRequest"`
}

type CitizenshipListType struct {
	BaseResponse
	XMLName		xml.Name				`xml:"CitizenshipList"`
	Citizenship	[]SimpleCodeNameType	`xml:"Citizenship"`
}
