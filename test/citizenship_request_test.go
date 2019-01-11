package test

import (
	"context"
	"testing"

	"github.com/nbio/st"
)

func TestCitizenshipListRequest_Ok(t *testing.T) {
	testRequest("citizenshiplist_response_example.xml", false)
	defer gockOff()

	data, err := acApi.CitizenshipListRequest(context.Background())
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, len(data.Citizenship), 128)
}

func TestCitizenshipListRequest_Error(t *testing.T) {
	testRequest("citizenshiplist_error_example.xml", true)
	defer gockOff()

	_, err := acApi.CitizenshipListRequest(context.Background())

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}
