package acaseSts

import "encoding/xml"

type TypeOfPlaceRequestType struct {
	XMLName		xml.Name						`xml:"TypeOfPlaceRequest"`
	BuyerId		string							`xml:"BuyerId,attr"`
	UserId		string							`xml:"UserId,attr"`
	Password	string							`xml:"Password,attr"`
	Language	LanguageTypeEnum				`xml:"Language,attr,omitempty"`
	Action		TypeOfPlaceRequestActionType	`xml:"Action"`
}

type TypeOfPlaceRequestActionType struct {
	XMLName		xml.Name				`xml:"Action"`
	Name		string					`xml:"Name,attr"`
	Parameters	TypeOfPlaceParamType	`xml:"Parameters"`
}

type TypeOfPlaceParamType struct {
	XMLName			xml.Name	`xml:"Parameters"`
	TypeOfPlaceCode	int			`xml:"TypeOfPlaceCode,attr"`
	TypeOfPlaceName	string		`xml:"TypeOfPlaceName,attr"`
}

type TypeOfPlaceResponseType struct {
	XMLName			xml.Name						`xml:"TypeOfPlaceResponse"`
	BuyerId			string							`xml:"BuyerId,attr"`
	UserId			string							`xml:"UserId,attr"`
	Password		string							`xml:"Password,attr"`
	Language		string							`xml:"Language,attr,omitempty"`
	Success			SuccessType						`xml:"Success"`
	Error			ErrorType						`xml:"Error,omitempty"`
	Action			TypeOfPlaceRequestActionType	`xml:"Action"`
	TypeOfPlaceList	TypeOfPlaceListType				`xml:"TypeOfPlaceList"`
}

type TypeOfPlaceListType struct {
	XMLName		xml.Name				`xml:"TypeOfPlaceList"`
	TypeOfPlace	[]SimpleCodeNameType	`xml:"TypeOfPlace"`
}
