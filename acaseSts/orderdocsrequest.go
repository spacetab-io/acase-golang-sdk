package acaseSts

import "encoding/xml"

type OrderDocActionName string
const (
	TASKADD OrderDocActionName = "TaskAdd"
	TASKSTATUS OrderDocActionName = "TaskStatus"
	TASKRESPONSE OrderDocActionName = "TaskResponse"
)

type OrderDocsRequestType struct {
	Credentials
	XMLName		xml.Name			`xml:"OrderDocsRequest"`
	Action		OrderDocActionType	`xml:"Action"`
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
	Credentials
	BaseResponse
	XMLName		xml.Name				`xml:"OrderDocsResponse"`
	Action		OrderDocActionType		`xml:"Action"`
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
