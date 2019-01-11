package test

import (
	"context"
	"testing"

	"github.com/nbio/st"
)

func TestObjectTypeRequest_Ok(t *testing.T) {
	testRequest("objecttype_response_example.xml", false)
	defer gockOff()

	data, err := acApi.ObjectTypeRequest(context.Background(), 0)
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.Action.Name, "LIST")
	st.Expect(t, data.Action.Parameters.ObjectTypeCode, "")
	st.Expect(t, len(data.ObjectTypeList.Items), 8)
	st.Expect(t, data.ObjectTypeList.Items[0].Code, 9500001)
	st.Expect(t, data.ObjectTypeList.Items[0].Name, "Hе имеет значения")
	st.Expect(t, data.ObjectTypeList.Items[1].Code, 800001)
	st.Expect(t, data.ObjectTypeList.Items[1].Name, "Аэропорт")
	st.Expect(t, data.ObjectTypeList.Items[2].Code, 900001)
	st.Expect(t, data.ObjectTypeList.Items[2].Name, "Достопримечательности")
	st.Expect(t, data.ObjectTypeList.Items[3].Code, 800104)
	st.Expect(t, data.ObjectTypeList.Items[3].Name, "Метро")
	st.Expect(t, data.ObjectTypeList.Items[4].Code, 1000039)
	st.Expect(t, data.ObjectTypeList.Items[4].Name, "Ж/д вокзалы")
	st.Expect(t, data.ObjectTypeList.Items[5].Code, 1000040)
	st.Expect(t, data.ObjectTypeList.Items[5].Name, "Морской порт")
	st.Expect(t, data.ObjectTypeList.Items[6].Code, 1000041)
	st.Expect(t, data.ObjectTypeList.Items[6].Name, "Речной порт")
	st.Expect(t, data.ObjectTypeList.Items[7].Code, 1000042)
	st.Expect(t, data.ObjectTypeList.Items[7].Name, "Центр города")
}

func TestObjectTypeRequest_Error(t *testing.T) {
	testRequest("objecttype_error_example.xml", true)
	defer gockOff()

	_, err := acApi.ObjectTypeRequest(context.Background(), 0)

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}
