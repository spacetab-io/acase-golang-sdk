package test

import (
	"testing"

	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
)

func TestCitizenshipListRequest_Ok(t *testing.T)  {
	testRequest("csresp_example.xml")
	defer gock.Off()

	data, err := acApi.CitizenshipListRequest()
	er := getCustomErrorType()
	st.Expect(t, err, er)
	st.Expect(t, len(data.Citizenship), 128)
}

func TestCitizenshipListRequest_Error(t *testing.T) {
	testRequest("cserror_example.xml")
	defer gock.Off()

	_, err := acApi.CitizenshipListRequest()

	st.Expect(t, err.Code, "9998")
	st.Expect(t, err.Message, "Доступ запрещен !")
}
