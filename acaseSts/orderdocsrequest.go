package acaseSts

import "encoding/xml"

type OrderDocActionName string
const (
	TASKADD OrderDocActionName = "TaskAdd"
	TASKSTATUS OrderDocActionName = "TaskStatus"
	TASKRESPONSE OrderDocActionName = "TaskResponse"
)

type OrderDocsRequestType struct {
	XMLName		xml.Name				`xml:"OrderDocsRequest"`
	BuyerId		string					`xml:"BuyerId,attr"`
	UserId		string					`xml:"UserId,attr"`
	Password	string					`xml:"Password,attr"`
	Language	LanguageTypeEnum		`xml:"Language,attr,omitempty"`
	Action		[]OrderDocActionType	`xml:"Action"`
}

type OrderDocActionType struct {
	XMLName		xml.Name				`xml:"Action"`
	Name		string					`xml:"Name,attr"`
	Parameters	OrderDocParametersType	`xml:"Parameters"`
}

type OrderDocParametersType struct {
	XMLName	xml.Name				`xml:"Parameters"`
	DocId	int						`xml:"DocId,attr,omitempty"`
	TaskId	int						`xml:"TaskId,attr,omitempty"`
	DocType	*OrderDocParamsDocType	`xml:"DocType,omitempty"`
}

type OrderDocParamsDocType struct {
	XMLName	xml.Name	`xml:"DocType"`
	Code	int			`xml:"Code,attr,omitempty"`
}

type OrderDocsResponseType struct {
	XMLName		xml.Name				`xml:"OrderDocsResponse"`
	BuyerId		string					`xml:"BuyerId,attr"`
	UserId		string					`xml:"UserId,attr"`
	Password	string					`xml:"Password,attr"`
	Language	string					`xml:"Language,attr,omitempty"`
	Success		string					`xml:"Success"`
	Error		ErrorType				`xml:"Error,omitempty"`
	Action		[]OrderDocActionType	`xml:"Action"`
	TaskInfo	[]TaskInfoType			`xml:"TaskInfo"`
	File		*OrderDocFileType		`xml:"File,omitempty"`
}

type OrderDocFileType struct {
	XMLName		xml.Name	`xml:"File"`
	ContentType	string		`xml:"ContentType,attr"`
	Type		string		`xml:"type,attr"`
	Value 		string 		`xml:",chardata"`
}

type TaskInfoType struct {
	XMLName		xml.Name			`xml:"TaskInfo"`
	Code		int					`xml:"Code,attr"`
	TaskStatus	SimpleCodeNameType	`xml:"TaskStatus"`
}
