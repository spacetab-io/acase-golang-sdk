package test

import (
	"context"
	"testing"

	"github.com/nbio/st"
)

func TestObjectRequest_Ok(t *testing.T) {
	testRequest("object_response_example.xml", false)
	defer gockOff()

	data, err := acApi.ObjectRequest(context.Background(), 0, 0, 0)
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, data.Action.Name, "LISTMETROSTYLE")
	st.Expect(t, data.Action.Parameters.ObjectTypeCode, "800104")
	st.Expect(t, data.Action.Parameters.ObjectSubTypeCode, "")
	st.Expect(t, data.Action.Parameters.CityCode, "2")
	st.Expect(t, len(data.Object), 2)
	st.Expect(t, data.Object[0].Code, 800229)
	st.Expect(t, data.Object[0].Name, "Авиамоторная")
	st.Expect(t, data.Object[0].ObjectSubType.Code, 1100025)
	st.Expect(t, data.Object[0].ObjectSubType.Name, "Калининская линия")
}

func TestObjectRequest_Error(t *testing.T) {
	testRequest("object_error_example.xml", true)
	defer gockOff()

	_, err := acApi.ObjectRequest(context.Background(), 0, 0, 0)

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}
