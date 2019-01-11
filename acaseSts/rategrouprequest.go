package acaseSts

import "encoding/xml"

type RateGroupRequestType struct {
	Credentials
	XMLName    xml.Name                `xml:"RateGroupRequest"`
	ActionList RateGroupActionListType `xml:"ActionList"`
}

type RateGroupActionListType struct {
	XMLName    xml.Name                 `xml:"ActionList"`
	Parameters []RateGroupParameterType `xml:"Parameters"`
}

type RateGroupParameterType struct {
	XMLName xml.Name `xml:"Parameters"`
	Name    string   `xml:"Name,attr"`
}

type RateGroupResponseType struct {
	Credentials
	BaseResponse
	XMLName    xml.Name                `xml:"RateGroupResponse"`
	ActionList RateGroupActionListType `xml:"ActionList"`
	RateGroup  []SimpleCodeNameType    `xml:"RateGroup"`
}
