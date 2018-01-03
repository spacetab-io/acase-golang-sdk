package test

import (
	"testing"
	"gopkg.in/h2non/gock.v1"
	"github.com/nbio/st"
)

func TestObjectRequest_Ok(t *testing.T) {
	testRequest("object_response_example.xml")
	defer gock.Off()

	data, err := acApi.ObjectRequest(nil,nil,nil)
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, len(data.Action), 1)
	st.Expect(t, data.Action[0].Name, "LISTMETROSTYLE")
	st.Expect(t, data.Action[0].Parameters.ObjectTypeCode, 800104)
	st.Expect(t, data.Action[0].Parameters.ObjectSubTypeCode, 0)
	st.Expect(t, data.Action[0].Parameters.CityCode, 2)
	st.Expect(t, len(data.Object), 2)
	st.Expect(t, data.Object[0].Code, 800229)
	st.Expect(t, data.Object[0].Name, "Авиамоторная")
	st.Expect(t, data.Object[0].ObjectSubType.Code, 1100025)
	st.Expect(t, data.Object[0].ObjectSubType.Name, "Калининская линия")
}

func TestObjectRequest_Error(t *testing.T) {
	testRequest("object_error_example.xml")
	defer gock.Off()

	_, err := acApi.ObjectRequest(nil,nil,nil)

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}

