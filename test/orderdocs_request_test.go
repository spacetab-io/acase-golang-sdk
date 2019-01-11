package test

import (
	"context"
	"testing"

	"github.com/nbio/st"
	"github.com/tmconsulting/acase-golang-sdk/acaseSts"
)

func TestOrderDocRequest_Create_Ok(t *testing.T) {
	testRequest("orderdoc_create_response_example.xml", false)
	defer gockOff()

	data, err := acApi.OrderDocRequest(context.Background(), acaseSts.TASKADD, 0, 2835883, 7)
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.Action.Name, "TaskAdd")
	st.Expect(t, data.Action.Parameters.DocId, 2835883)
	st.Expect(t, data.Action.Parameters.DocType.Code, 7)
	st.Expect(t, len(data.TaskInfo), 1)
	st.Expect(t, data.TaskInfo[0].Code, 25749)
	st.Expect(t, data.TaskInfo[0].TaskStatus.Code, 1)
	st.Expect(t, data.TaskInfo[0].TaskStatus.Name, "New")
}

func TestOrderDocRequest_CheckStatus_Ok(t *testing.T) {
	testRequest("orderdoc_checkstatus_response_example.xml", false)
	defer gockOff()

	data, err := acApi.OrderDocRequest(context.Background(), acaseSts.TASKSTATUS, 25729, 0, 0)
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.Action.Name, "TaskStatus")
	st.Expect(t, data.Action.Parameters.TaskId, 25729)
	st.Expect(t, len(data.TaskInfo), 1)
	st.Expect(t, data.TaskInfo[0].Code, 25729)
	st.Expect(t, data.TaskInfo[0].TaskStatus.Code, 5)
	st.Expect(t, data.TaskInfo[0].TaskStatus.Name, "Success")
}

func TestOrderDocRequest_GetResult_Ok(t *testing.T) {
	testRequest("orderdoc_getresult_response_example.xml", false)
	defer gockOff()

	data, err := acApi.OrderDocRequest(context.Background(), acaseSts.TASKRESPONSE, 25729, 0, 0)
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.Action.Name, "TaskResponse")
	st.Expect(t, data.Action.Parameters.TaskId, 25729)
	st.Expect(t, len(data.TaskInfo), 1)
	st.Expect(t, data.TaskInfo[0].Code, 25729)
	st.Expect(t, data.TaskInfo[0].TaskStatus.Code, 5)
	st.Expect(t, data.TaskInfo[0].TaskStatus.Name, "Success")
	var tstFile *acaseSts.OrderDocFileType = nil
	st.Reject(t, data.File, tstFile)
	st.Expect(t, data.File.Type, "xsd:base64Binary")
	st.Expect(t, data.File.ContentType, "PDF")
	st.Expect(t, data.File.Value, "ЙЦУКЕН")
}

func TestOrderDocRequest_Error(t *testing.T) {
	testRequest("orderdoc_error_example.xml", true)
	defer gockOff()

	_, err := acApi.OrderDocRequest(context.Background(), acaseSts.TASKADD, 0, 2835883, 7)

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}
