package test

import (
	"testing"
	"github.com/nbio/st"
)

func TestObjectSubTypeRequest_Ok(t *testing.T) {
	testRequest("objectsubtype_response_example.xml", false)
	defer gockOff()

	data, err := acApi.ObjectSubTypeRequest(0)
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.Action.Name, "LIST")
	st.Expect(t, data.Action.Parameters.ObjectTypeCode, "")
	st.Expect(t, len(data.ObjectType), 3)
	st.Expect(t, data.ObjectType[0].Code, 9500001)
	st.Expect(t, data.ObjectType[0].Name, "Hе имеет значения")
	st.Expect(t, len(data.ObjectType[0].ObjectSubType), 0)
	st.Expect(t, data.ObjectType[1].Code, 800001)
	st.Expect(t, data.ObjectType[1].Name, "Аэропорт")
	st.Expect(t, len(data.ObjectType[1].ObjectSubType), 0)
	st.Expect(t, data.ObjectType[2].Code, 900001)
	st.Expect(t, data.ObjectType[2].Name, "Достопримечательности")
	st.Expect(t, len(data.ObjectType[2].ObjectSubType), 3)
	st.Expect(t, data.ObjectType[2].ObjectSubType[0].Code, 900005)
	st.Expect(t, data.ObjectType[2].ObjectSubType[0].Name, "Аквапарки, дельфинарии")
	st.Expect(t, data.ObjectType[2].ObjectSubType[1].Code, 900010)
	st.Expect(t, data.ObjectType[2].ObjectSubType[1].Name, "Выставки и концертные залы")
	st.Expect(t, data.ObjectType[2].ObjectSubType[2].Code, 900012)
	st.Expect(t, data.ObjectType[2].ObjectSubType[2].Name, "Государственные учреждения")
}

func TestObjectSubTypeRequest_Error(t *testing.T) {
	testRequest("objectsubtype_error_example.xml", true)
	defer gockOff()

	_, err := acApi.ObjectSubTypeRequest(0)

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}

